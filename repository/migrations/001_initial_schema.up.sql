CREATE TABLE profiles (
    id char(27) PRIMARY KEY,
    username TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT,
    avatar TEXT
);

CREATE UNIQUE INDEX username_idx ON profiles (username);
