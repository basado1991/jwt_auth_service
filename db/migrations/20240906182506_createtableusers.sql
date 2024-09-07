-- migrate:up
CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  refresh_token bytea
)

-- migrate:down
DROP TABLE users;
