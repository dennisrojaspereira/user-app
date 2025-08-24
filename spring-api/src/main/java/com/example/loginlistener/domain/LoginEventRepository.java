package com.example.loginlistener.domain;

public interface LoginEventRepository {
    void save(LoginEvent event);
}
