CREATE TABLE program_kerjas
(
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    slug         VARCHAR(255) UNIQUE NOT NULL,
    name         VARCHAR(255)        NOT NULL,
    kemenbiro_id UUID                NOT NULL REFERENCES kemenbiros (id) ON DELETE SET NULL,
    description  VARCHAR(2000)
);