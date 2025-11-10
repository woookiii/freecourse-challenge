package aws

import (
	"fmt"
	"geosqlwithcdn/config"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Aws struct {
	session *session.Session
	S3      *s3.S3
	cfg     *config.Config

	Bucket string
}

func NewAws(cfg *config.Config) *Aws {
	a := &Aws{cfg: cfg}
	var err error

	if a.session, err = session.NewSession(&aws.Config{
		Region:      aws.String(cfg.Aws.Region),
		Credentials: credentials.NewStaticCredentials(cfg.Aws.Key, cfg.Aws.SecretKey, ""),
	}); err != nil {
		panic(err)
	} else {
		a.Bucket = cfg.Aws.Bucket
		a.S3 = s3.New(a.session)
	}
	return a
}

func (a *Aws) PutFileToS3(key string, extension string, file *os.File) error {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(a.Bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(fmt.Sprintf("%s/%s", "image", extension)),
		ACL:         aws.String("public-read"),
	}
	_, err := a.S3.PutObject(input)

	return err
}
