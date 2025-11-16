package com.example.elastic_search_practice_api;

import co.elastic.clients.elasticsearch._types.query_dsl.MultiMatchQuery;
import co.elastic.clients.elasticsearch._types.query_dsl.Query;
import co.elastic.clients.elasticsearch._types.query_dsl.TextQueryType;
import com.example.elastic_search_practice_api.domain.Product;
import com.example.elastic_search_practice_api.domain.ProductDocument;
import com.example.elastic_search_practice_api.dto.CreateProductRequestDto;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.elasticsearch.client.elc.NativeQuery;
import org.springframework.data.elasticsearch.core.ElasticsearchOperations;
import org.springframework.data.elasticsearch.core.SearchHits;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ProductService {

    private final ProductRepository productRepository;
    private final ProductDocumentRepository productDocumentRepository;
    private final ElasticsearchOperations elasticsearchOperations;


    public ProductService(ProductRepository productRepository, ProductDocumentRepository productDocumentRepository, ElasticsearchOperations elasticsearchOperations) {
        this.productRepository = productRepository;
        this.productDocumentRepository = productDocumentRepository;
        this.elasticsearchOperations = elasticsearchOperations;
    }

    public List<Product> getProducts(int page, int size) {
        Pageable pageable = PageRequest.of(page - 1, size);
        return productRepository.findAll(pageable).getContent();
    }

    public Product createProduct(CreateProductRequestDto createProductRequestDto) {
        Product product = new Product(
                createProductRequestDto.getName(),
                createProductRequestDto.getDescription(),
                createProductRequestDto.getPrice(),
                createProductRequestDto.getRating(),
                createProductRequestDto.getCategory()
        );

        Product savedProduct = productRepository.save(product);

        ProductDocument productDocument = new ProductDocument(
                savedProduct.getId().toString(),
                savedProduct.getName(),
                savedProduct.getDescription(),
                savedProduct.getPrice(),
                savedProduct.getRating(),
                savedProduct.getCategory()
        );

        productDocumentRepository.save(productDocument);

        return savedProduct;
    }

    public void deleteProduct(Long id) {
        productRepository.deleteById(id);
        productDocumentRepository.deleteById(id.toString());
    }

    public List<String> getSuggestions(String query) {
        Query multiMatchQuery = MultiMatchQuery.of(m -> m
                .query(query)
                .type(TextQueryType.BoolPrefix)
                .fields("name.auto_complete", "name.auto_complete._2gram", "name.auto_complete._3gram")
        )._toQuery();

        NativeQuery nativeQuery = NativeQuery.builder()
                .withQuery(multiMatchQuery)
                .withPageable(PageRequest.of(0, 5))
                .build();

        SearchHits<ProductDocument> searchHits = this.elasticsearchOperations.search(nativeQuery, ProductDocument.class);//GET /products/_search {...}
        return searchHits.getSearchHits().stream()
                .map(hit -> {
                    ProductDocument productDocument = hit.getContent();
                    return productDocument.getName();
                }).toList();
    }
}