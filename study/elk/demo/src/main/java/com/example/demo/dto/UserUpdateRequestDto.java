package com.example.demo.dto;

public class UserUpdateRequestDto {
    private String name;
    private Long age;
    private Boolean isActive;

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
