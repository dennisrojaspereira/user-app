package com.example.loginlistener.application;

import com.example.loginlistener.domain.LoginEvent;
import com.example.loginlistener.domain.LoginEventRepository;
import com.example.loginlistener.domain.EmailSender;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;

class LoginEventUseCaseTest {
    @Test
    void handle_savesEventAndSendsEmail() {
        LoginEventRepository repo = Mockito.mock(LoginEventRepository.class);
        EmailSender sender = Mockito.mock(EmailSender.class);
        LoginEventUseCase useCase = new LoginEventUseCase(repo, sender);
        LoginEvent event = new LoginEvent("1", "test@example.com", "2025-08-24T12:00:00Z");
        useCase.handle(event);
        Mockito.verify(repo).save(event);
        Mockito.verify(sender).send("test@example.com", "Login Notification", "User 1 logged in at 2025-08-24T12:00:00Z");
    }
}
