CREATE TABLE kemenbiros
(
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    abbreviation VARCHAR(15)  NOT NULL,
    name         VARCHAR(255) NOT NULL,
    description  VARCHAR(2000)
);

CREATE UNIQUE INDEX kemenbiros_abbreviation_key ON kemenbiros (UPPER(abbreviation));
