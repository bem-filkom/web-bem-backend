CREATE TABLE bem_members
(
    nim          CHAR(15) PRIMARY KEY REFERENCES students (nim) ON DELETE CASCADE,
    kemenbiro_id UUID         NOT NULL REFERENCES kemenbiros (id) ON DELETE CASCADE,
    position     VARCHAR(255) NOT NULL
);