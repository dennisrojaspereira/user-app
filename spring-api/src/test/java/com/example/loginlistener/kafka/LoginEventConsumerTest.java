package com.example.loginlistener.kafka;

import com.example.loginlistener.model.LoginEvent;
import com.example.loginlistener.service.LoginEventService;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;

class LoginEventConsumerTest {
    @Test
    void consume_callsService() {
        LoginEventService service = Mockito.mock(LoginEventService.class);
        LoginEventConsumer consumer = new LoginEventConsumer();
        consumer.service = service;
        LoginEvent event = new LoginEvent("1", "test@example.com", "2025-08-24T12:00:00Z");
        consumer.consume(event);
        Mockito.verify(service).handleLoginEvent(event);
    }
}
