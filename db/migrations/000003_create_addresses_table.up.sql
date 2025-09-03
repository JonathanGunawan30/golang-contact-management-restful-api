CREATE TABLE "addresses" (
    "id" SERIAL PRIMARY KEY,
    "street" VARCHAR(255),
    "city" VARCHAR(100),
    "province" VARCHAR(100),
    "country" VARCHAR(100) NOT NULL,
    "postal_code" VARCHAR(10) NOT NULL,
    "contact_id" INT NOT NULL,
    CONSTRAINT fk_addresses_contact
        FOREIGN KEY("contact_id")
            REFERENCES "contacts"("id")
            ON DELETE CASCADE
);

CREATE INDEX idx_addresses_contact_id ON "addresses"("contact_id");