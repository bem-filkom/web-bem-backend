CREATE TABLE program_kerja_penanggung_jawabs
(
    program_kerja_id UUID     NOT NULL REFERENCES program_kerjas (id) ON DELETE CASCADE,
    nim              CHAR(15) NOT NULL REFERENCES bem_members (nim) ON DELETE CASCADE,
    PRIMARY KEY (program_kerja_id, nim)
)