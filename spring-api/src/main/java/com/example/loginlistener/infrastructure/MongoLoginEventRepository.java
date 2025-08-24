package com.example.loginlistener.infrastructure;

import com.example.loginlistener.domain.LoginEvent;
import com.example.loginlistener.domain.LoginEventRepository;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface MongoLoginEventRepository extends MongoRepository<LoginEvent, String>, LoginEventRepository {
    default void save(LoginEvent event) {
        save((LoginEvent) event);
    }
}
