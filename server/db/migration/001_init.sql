-- +goose up
CREATE TABLE IF NOT EXISTS "users" (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL
);

-- +goose down
DROP TABLE IF EXISTS "users";
