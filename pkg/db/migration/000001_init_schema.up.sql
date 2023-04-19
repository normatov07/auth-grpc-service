CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "role_id" bigint NOT NULL DEFAULT 1,
  "email" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT 'now()',
  "active" boolean NOT NULL DEFAULT TRUE,
  "login_at" timestamptz NOT NULL DEFAULT 'now()',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);
