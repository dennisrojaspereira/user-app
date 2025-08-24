package com.example.loginlistener.web;

import com.example.loginlistener.application.LoginEventUseCase;
import com.example.loginlistener.domain.LoginEvent;
import com.example.loginlistener.infrastructure.MongoLoginEventRepository;
import com.example.loginlistener.infrastructure.SmtpEmailSender;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class LoginEventListener {
    private final LoginEventUseCase useCase;

    @Autowired
    public LoginEventListener(MongoLoginEventRepository repository, SmtpEmailSender emailSender) {
        this.useCase = new LoginEventUseCase(repository, emailSender);
    }

    @KafkaListener(topics = "user-login", groupId = "login-listener-group")
    public void consume(LoginEvent event) {
        useCase.handle(event);
    }
}
