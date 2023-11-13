CREATE INDEX ON "results" ("test_id");

CREATE UNIQUE INDEX ON "results" ("devices_id", "test_id");

ALTER TABLE "results" ADD FOREIGN KEY ("test_id") REFERENCES "tests" ("id");

