CREATE DATABASE g2fa;

CREATE TABLE UserInfo (
        Id INT PRIMARY KEY NOT NULL,
        Name TEXT NOT NULL,
        Email TEXT NOT NULL,
        Password VARCHAR(255) NOT NULL,
        Otp_enabled BOOLEAN NOT NULL,
        Otp_verified BOOLEAN NOT NULL,
        Otp_secret TEXT NOT NULL,
        Otp_auth_url TEXT NOT NULL
    );