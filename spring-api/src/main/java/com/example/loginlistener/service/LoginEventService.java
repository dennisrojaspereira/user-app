package com.example.loginlistener.service;

import com.example.loginlistener.model.LoginEvent;
import com.example.loginlistener.repository.LoginEventRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Service;

@Service
public class LoginEventService {
    @Autowired
    private LoginEventRepository repository;
    @Autowired
    private JavaMailSender mailSender;

    public void handleLoginEvent(LoginEvent event) {
        repository.save(event);
        SimpleMailMessage message = new SimpleMailMessage();
        message.setTo(event.getEmail());
        message.setSubject("Login Notification");
        message.setText("User " + event.getUserId() + " logged in at " + event.getTimestamp());
        mailSender.send(message);
    }
}
