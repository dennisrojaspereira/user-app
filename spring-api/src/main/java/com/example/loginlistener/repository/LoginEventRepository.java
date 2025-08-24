package com.example.loginlistener.repository;

import com.example.loginlistener.model.LoginEvent;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface LoginEventRepository extends MongoRepository<LoginEvent, String> {
}
