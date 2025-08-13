CREATE TABLE users(
    user_id UUID PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(60) NOT NULL,
    registered_at timestamp NOT NULL DEFAULT Now()
);
