CREATE TABLE conversations(
    conversation_id UUID PRIMARY KEY,
    name varchar(255) NOT NULL,
    creation_date timestamp NOT NULL DEFAULT Now()
);

CREATE TABLE participants(
    participant_id UUID PRIMARY KEY,
    conversation_id UUID NOT NULL,
    user_id UUID NOT NULL,
    FOREIGN KEY(conversation_id) REFERENCES conversations(conversation_id) ON DELETE CASCADE
);

CREATE TABLE messages(
    message_id UUID PRIMARY KEY,
    conversation_id UUID NOT NULL,
    sender_id UUID NOT NULL,
    sent_at timestamp NOT NULL DEFAULT Now(),
    message_body text,
    FOREIGN KEY (conversation_id) REFERENCES conversations(conversation_id)ON DELETE CASCADE
);

CREATE TABLE attachments(
    attachment_id UUID PRIMARY KEY,
    message_id UUID NOT NULL,
    file_name varchar(255),
    file_type varchar(255),
    file_size int,
    file_url varchar(255),
    FOREIGN KEY (message_id) REFERENCES messages(message_id) ON DELETE CASCADE
);

CREATE INDEX participants_conversation_id_index ON participants(conversation_id);
CREATE INDEX participants_user_id_index ON participants(user_id);
CREATE INDEX messages_conversation_id_index ON messages(conversation_id);
CREATE INDEX messages_sender_id_index ON messages(sender_id);
CREATE INDEX attachments_message_id_index ON attachments(message_id);
