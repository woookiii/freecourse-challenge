package com.example.elastic_search_practice_api;

import com.example.elastic_search_practice_api.domain.ProductDocument;
import org.springframework.data.elasticsearch.repository.ElasticsearchRepository;

public interface ProductDocumentRepository extends ElasticsearchRepository<ProductDocument, String> {
}