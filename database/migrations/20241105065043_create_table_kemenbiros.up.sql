CREATE TABLE kemenbiros
(
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    abbreviation VARCHAR(15) UNIQUE NOT NULL,
    name         VARCHAR(255)       NOT NULL
);

CREATE INDEX ON kemenbiros USING HASH (abbreviation);
