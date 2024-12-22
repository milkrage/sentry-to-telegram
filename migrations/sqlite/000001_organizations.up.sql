CREATE TABLE organizations(
    slug            TEXT    PRIMARY KEY,
    installation_id TEXT    UNIQUE,
    refresh_token   TEXT
);
