package com.example.loginlistener.domain;

public class LoginEvent {
    private String userId;
    private String email;
    private String timestamp;

    public LoginEvent(String userId, String email, String timestamp) {
        this.userId = userId;
        this.email = email;
        this.timestamp = timestamp;
    }
    public String getUserId() { return userId; }
    public String getEmail() { return email; }
    public String getTimestamp() { return timestamp; }
}
