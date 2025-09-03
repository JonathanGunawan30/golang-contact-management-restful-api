CREATE TABLE "contacts" (
    "id" SERIAL PRIMARY KEY,
    "first_name" VARCHAR(100) NOT NULL,
    "last_name" VARCHAR(100),
    "email" VARCHAR(100),
    "phone" VARCHAR(20),
    "username" VARCHAR(255) NOT NULL,
    CONSTRAINT fk_contacts_user
        FOREIGN KEY("username")
            REFERENCES "users"("username")
);

CREATE INDEX idx_contacts_username ON "contacts"("username");