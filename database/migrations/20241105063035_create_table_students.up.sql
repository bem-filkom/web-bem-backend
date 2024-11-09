CREATE TABLE students
(
    nim           CHAR(15) PRIMARY KEY REFERENCES users (id) ON DELETE CASCADE,
    program_studi VARCHAR(255) NOT NULL,
    fakultas      VARCHAR(255) NOT NULL
);
