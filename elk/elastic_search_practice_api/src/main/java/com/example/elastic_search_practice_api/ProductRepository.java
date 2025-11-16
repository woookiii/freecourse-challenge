package com.example.elastic_search_practice_api;

import com.example.elastic_search_practice_api.domain.Product;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ProductRepository extends JpaRepository<Product, Long> {

}
