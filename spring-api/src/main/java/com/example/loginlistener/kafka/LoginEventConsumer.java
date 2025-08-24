package com.example.loginlistener.kafka;

import com.example.loginlistener.model.LoginEvent;
import com.example.loginlistener.service.LoginEventService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class LoginEventConsumer {
    @Autowired
    private LoginEventService service;

    @KafkaListener(topics = "user-login", groupId = "login-listener-group")
    public void consume(LoginEvent event) {
        service.handleLoginEvent(event);
    }
}
