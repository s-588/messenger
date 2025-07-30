CREATE TABLE users(
    user_id UUID PRIMARY KEY,
    username varchar(255) UNIQUE,
    password varchar(60) NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255),
    email varchar(255),
    avatar varchar(255), 
    registered_at timestamp
);

CREATE TABLE conversations(
    conversation_id UUID PRIMARY KEY,
    name varchar(255) NOT NULL,
    creation_date timestamp
);

CREATE TABLE participants(
    participant_id UUID PRIMARY KEY,
    conversation_id UUID,
    user_id UUID,
    FOREIGN KEY(conversation_id) REFERENCES conversations(conversation_id) ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users(user_id)ON DELETE CASCADE
);

CREATE TABLE messages(
    message_id UUID PRIMARY KEY,
    conversation_id UUID,
    sender_id UUID,
    sended_at timestamp,
    message_body text,
    FOREIGN KEY (conversation_id) REFERENCES conversations(conversation_id)ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE attachments(
    attachment_id UUID PRIMARY KEY,
    message_id UUID,
    file_name varchar(255),
    file_type varchar(255),
    file_size int,
    file_path varchar(255),
    FOREIGN KEY (message_id) REFERENCES messages(message_id) ON DELETE CASCADE
);
