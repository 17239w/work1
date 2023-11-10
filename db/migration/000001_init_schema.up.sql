CREATE TABLE "devices" (
  "id" bigserial PRIMARY KEY,
  "device_name" varchar NOT NULL,
  "device_manufacturer" varchar NOT NULL,
  "device_origin" varchar NOT NULL,
  "production_date" timestamptz NOT NULL DEFAULT 'now()',
  "testing_date" timestamptz NOT NULL DEFAULT 'now()',
  "device_model" varchar NOT NULL
);

CREATE TABLE "tests" (
  "id" bigserial PRIMARY KEY,
  "test_name" varchar NOT NULL,
  "devices_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "gear" bigint NOT NULL,
  "percentage" bigint NOT NULL,
  "lower_limit" bigint NOT NULL,
  "upper_limit" bigint NOT NULL,
  "test_data" bigint NOT NULL
);

CREATE TABLE "results" (
  "id" bigserial PRIMARY KEY,
  "test_id" bigint NOT NULL,
  "devices_id" bigint NOT NULL,
  "voltage" bigint NOT NULL,
  "point_number" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "temperature" bigint NOT NULL,
  "humidity" bigint NOT NULL
);

CREATE INDEX ON "tests" ("devices_id");

CREATE INDEX ON "results" ("devices_id");

COMMENT ON COLUMN "tests"."percentage" IS 'must be positive';

COMMENT ON COLUMN "tests"."lower_limit" IS 'must be positive';

COMMENT ON COLUMN "tests"."upper_limit" IS 'must be positive';

COMMENT ON COLUMN "tests"."test_data" IS 'must be positive';

ALTER TABLE "tests" ADD FOREIGN KEY ("devices_id") REFERENCES "devices" ("id");

ALTER TABLE "results" ADD FOREIGN KEY ("devices_id") REFERENCES "devices" ("id");
