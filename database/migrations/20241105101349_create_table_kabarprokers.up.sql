CREATE TABLE kabar_prokers
(
    id               VARCHAR(255) PRIMARY KEY,
    program_kerja_id UUID         NOT NULL REFERENCES program_kerjas (id) ON DELETE SET NULL,
    title            VARCHAR(255) NOT NULL,
    content          TEXT,
    created_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX kabar_prokers_created_at_idx ON kabar_prokers (created_at);

CREATE TRIGGER update_kabar_prokers_updated_at
    BEFORE UPDATE
    ON kabar_prokers
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TABLE kabarproker_images
(
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    kabar_proker_id VARCHAR(255) NOT NULL REFERENCES kabar_prokers (id) ON DELETE CASCADE,
    url             TEXT         NOT NULL,
    description     VARCHAR(255)
);
