# Spring Boot microservice for user login events

This service listens to Kafka for user login events, sends an email notification, and stores the event in MongoDB.

## Features
- Kafka consumer for login events
- Email notification on login
- MongoDB persistence for login events
- Unit and integration tests

## How to run
1. Configure Kafka, MongoDB, and SMTP settings in `application.yml`.
2. Build and run the service:
   ```bash
   ./mvnw spring-boot:run
   ```

## How to test
```bash
./mvnw test
```

## Environment variables
- `KAFKA_BOOTSTRAP_SERVERS`
- `MONGODB_URI`
- `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASS`

## Endpoints
- No REST endpoints; service runs as a Kafka consumer.

## Project structure
- `src/main/java` - main source code
- `src/test/java` - tests

---

See Makefile for build and test commands.
