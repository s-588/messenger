# Messenger

This is project designed for learning purpose of microservice architecture, websockets, grpc, jwt and refresh tokens and Postgres.

Planned architecture:

- Auth service - handle authentication, authorization, registartion, token and permissions management. Keep user data in Postgres instance.
- User service - handle user profiles and user data. Keep user data in Postgres instance.
- Message service - handle message, chat and group management.

Potentially:

- Integrate RabbitMQ. Message service will push messages in it for other services.
- Add notificttion service - get messages from RabbitMQ and send notification to clients.
- Add search service - service that uses ElasticSearch and get messages from RabbitMQ.
