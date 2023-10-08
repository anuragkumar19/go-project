-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    avatar TEXT UNIQUE DEFAULT 'https://api.dicebear.com/avatar.svg' NOT NULL,
    is_email_verified BOOLEAN DEFAULT FALSE NOT NULL,
    otp INT,
    otp_expiry TIMESTAMP,
    password TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
    updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE users;