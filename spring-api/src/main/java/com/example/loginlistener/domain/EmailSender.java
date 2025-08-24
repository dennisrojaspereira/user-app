package com.example.loginlistener.domain;

public interface EmailSender {
    void send(String to, String subject, String body);
}
