CREATE TABLE "users" (
    "username" VARCHAR(255) PRIMARY KEY,
    "password" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "token" VARCHAR(255)
);