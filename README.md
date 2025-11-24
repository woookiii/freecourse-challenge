study 디렉토리는 주로 golang, elasticsearch에 대한 공부기록을 담고 있습니다.

project에 대해서 설명해보겠습니다.

golang으로만 구성된 네 개의 서버가 있습니다.

세개의 서버는 worker 노드로 각각 redis, elasticsearch, replicadb에 대한 연결을 가지고 있습니다.

나머지 한개의 서버는 웹서버로 api엔드포인트를 가지고 있습니다.

POST 형태로 회원가입 요청이 오게 되면 이 웹서버는 이메일 중복확인 후에, uuid와 created_time을 생성, 패스워드를 암호화하여 master db에 저장합니다.

db에 저장이 성공하면, 사용자에게 200 ok response를 주는 것과 별개로 새로운 고루틴(스레드와 비슷한 개념이지만 os런타임이 아닌 go자체런타임)을 통해 동시에 
db에 insert 성공한 객체를 byte로 marshalling(serializing)하여 서버가 시작할 때, 만들어진 카프카 프로듀서 클라이언트를 통해 메세지를 
confluent kafka cloud에 발행합니다. 

이렇게 하게 되면, 하나의 고루틴을 통해 카프카에 메세지를 발행하고 카프카로부터 메시지가 제대로 발행됬는지 확인(ack)하고,
사용자에게 200 ok response를 주는 것보다 빠르게 사용자에게 응답값을 보여줄 수 있고,

동일하게 하나의 고루틴을 사용하여 사용자에게 200 ok response를 주고 나서 카프카에 메세지를 발행하여 
캐싱과 서칭, replica db에 대한 master db와의 동기화를 하는 것보다 빠르게 메시지를 발행, 
동기화를 빠르게 수행할 수 있다고 생각했습니다.

나머지 세개의 서버는 서버가 시작되면, 동일하게 카프카 컨슈머 클라이언트를 통해 발행될 토픽에 대해 별개의 고루틴에서 무한 루프를 돌며 리슨을 하고 있습니다.
메시지가 감지되면, 메시지의 토픽을 확인하여 토픽에 맞게 메시지를 unmarshalling(deserializing)하여 
각자 연결되어있는 replicadb, redis, elastic search(cloud)에 저장하도록 했습니다.

redis의 경우, 전체객체를 다시 byte로 marshalling해서 string으로 캐스팅을 해서 저장을 하기 보다는, redis 자료구조중 하나인 
hset을 이용하여 필드별로 저장을 하여, 나중에 캐싱된 값을 수정하거나, get할 때, 원하는 필드값만 읽어오는 것이 용이하도록 했습니다. 

kafka의 경우 도커 데스크탑환경에서 구동시에 콘테이너가 여러번 꺼지는 문제가 있었고,
엘라스틱 서치와 키바나의 경우 너무 무거워서 다른 콘테이너들과 함께 구동이 힘들었던 관계로 
무료체험으로 제공되는 confluent, elasticsearch 클라우드 환경을 학습하여 사용해보는 것으로 문제를 해결했습니다

db의 경우 uuid와 관련된 문제때문에 postgres를 sql/database 표준 데이터베이스 라이브러리에 드라이버를 통해서 사용했고,
kafka 클라이언트의 경우 shopify에서 만들어져서, 지금은 IBM에서 관리가 되고 있는 sarama를 학습하여 사용했고,
redis 클라이언트의 경우 go-redis보다 벤치마크 성능이 빠른 reuidis를 학습하여 사용했습니다.
elastic search 클라이언트의 경우 go-elasticsearch를 사용했고 typedClient(Spring data JPA와 같이 쿼리작성을 도와줌)를 학습하여 사용했습니다

