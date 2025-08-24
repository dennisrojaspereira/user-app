package com.example.loginlistener.application;

import com.example.loginlistener.domain.LoginEvent;
import com.example.loginlistener.domain.LoginEventRepository;
import com.example.loginlistener.domain.EmailSender;

public class LoginEventUseCase {
    private final LoginEventRepository repository;
    private final EmailSender emailSender;

    public LoginEventUseCase(LoginEventRepository repository, EmailSender emailSender) {
        this.repository = repository;
        this.emailSender = emailSender;
    }

    public void handle(LoginEvent event) {
        repository.save(event);
        emailSender.send(event.getEmail(), "Login Notification", "User " + event.getUserId() + " logged in at " + event.getTimestamp());
    }
}
