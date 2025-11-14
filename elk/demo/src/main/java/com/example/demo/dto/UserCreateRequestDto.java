package com.example.demo.dto;

public class UserCreateRequestDto {
    private String id;
    private String name;
    private Long age;
    private Boolean isActive;

    public String getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public Long getAge() {
        return age;
    }

    public Boolean getActive() {
        return isActive;
    }
}
