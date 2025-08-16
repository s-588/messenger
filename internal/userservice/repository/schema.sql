CREATE TABLE usersData(
    user_id UUID PRIMARY KEY,
    first_name varchar(255) NOT NULL,
    last_name varchar(255),
    email varchar(255) UNIQUE,
    avatar_url varchar(255)
);

CREATE INDEX users_email_index ON users (email);
