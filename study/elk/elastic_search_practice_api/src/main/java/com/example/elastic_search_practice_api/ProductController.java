package com.example.elastic_search_practice_api;


import java.util.List;

import com.example.elastic_search_practice_api.domain.Product;
import com.example.elastic_search_practice_api.domain.ProductDocument;
import com.example.elastic_search_practice_api.dto.CreateProductRequestDto;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("products")
public class ProductController {

    private final ProductService productService;

    public ProductController(ProductService productService) {
        this.productService = productService;
    }

    @GetMapping()
    public ResponseEntity<List<Product>> getProducts(@RequestParam(defaultValue = "1") int page, @RequestParam(defaultValue = "10") int size) {
        List<Product> products = productService.getProducts(page, size);
        return ResponseEntity.ok(products);
    }


    @PostMapping()
    public ResponseEntity<Product> createProduct(@RequestBody CreateProductRequestDto createProductRequestDto) {
        Product product = productService.createProduct(createProductRequestDto);
        return ResponseEntity.ok(product);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteProduct(@PathVariable Long id) {
        productService.deleteProduct(id);
        return ResponseEntity.noContent().build();
    }

    @GetMapping("/suggestions")
    public ResponseEntity<?> getSuggestions(@RequestParam String query) {
        var suggestions = productService.getSuggestions(query);
        return ResponseEntity.ok(suggestions);
    }

    @GetMapping("/search")
    public ResponseEntity<?> searchProducts(
            @RequestParam String query,
            @RequestParam(required = false) String category,
            @RequestParam(defaultValue = "0") double minPrice,
            @RequestParam(defaultValue = "10000000") double maxPrice,
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "5") int size
    ) {
        var products = productService.searchProducts(
                query,
                category,
                minPrice,
                maxPrice,
                page,
                size
        );
        return ResponseEntity.ok(products);
    }
}
