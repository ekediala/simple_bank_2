-- +goose Up
CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp DEFAULT 'now()',
  "balance" bigint NOT NULL DEFAULT 0,
  "currency" varchar(3) NOT NULL
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp DEFAULT 'now()',
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp DEFAULT 'now()',
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "from_entry_id" bigint NOT NULL,
  "to_entry_id" BIGINT NOT NULL
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("to_entry_id");

CREATE INDEX ON "transfers" ("from_entry_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'it can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'only positive integers allowed';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_entry_id") REFERENCES "entries" ("id") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_entry_id") REFERENCES "entries" ("id") ON DELETE CASCADE;


-- +goose Down
DROP TABLE "transfers";
DROP TABLE "entries";
DROP TABLE "accounts";