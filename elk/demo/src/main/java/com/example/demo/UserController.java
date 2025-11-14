package com.example.demo;

import com.example.demo.dto.UserCreateRequestDto;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.web.bind.annotation.*;


@RestController
@RequestMapping("users")
public class UserController {

    private UserDocumentRepository userDocumentRepository;

    public UserController(UserDocumentRepository userDocumentRepository) {
        this.userDocumentRepository = userDocumentRepository;
    }

    @PostMapping
    public UserDocument createUser(@RequestBody UserCreateRequestDto userCreateRequestDto) {
        UserDocument user = new UserDocument(
                userCreateRequestDto.getId(),
                userCreateRequestDto.getName(),
                userCreateRequestDto.getAge(),
                userCreateRequestDto.getActive()
        );
        return userDocumentRepository.save(user);

    }



}
