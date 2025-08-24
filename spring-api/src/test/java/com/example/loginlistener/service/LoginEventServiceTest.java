package com.example.loginlistener.service;

import com.example.loginlistener.model.LoginEvent;
import com.example.loginlistener.repository.LoginEventRepository;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.SimpleMailMessage;

class LoginEventServiceTest {
    @Test
    void handleLoginEvent_savesEventAndSendsEmail() {
        LoginEventRepository repo = Mockito.mock(LoginEventRepository.class);
        JavaMailSender sender = Mockito.mock(JavaMailSender.class);
        LoginEventService service = new LoginEventService();
        service.repository = repo;
        service.mailSender = sender;
        LoginEvent event = new LoginEvent("1", "test@example.com", "2025-08-24T12:00:00Z");
        service.handleLoginEvent(event);
        Mockito.verify(repo).save(event);
        Mockito.verify(sender).send(Mockito.any(SimpleMailMessage.class));
    }
}
