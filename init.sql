CREATE TABLE "Employee" (
  "nip" BIGINT PRIMARY KEY,
  "name" varchar(100),
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "deleted_at" timestamp DEFAULT NULL
);

CREATE TYPE activity_type_enum AS ENUM ('check-in', 'check-out', 'break-start', 'break-end');

CREATE TABLE "AttendanceLog" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "employee_nip" BIGINT,
  "activity_type" activity_type_enum NOT NULL,
  "notes" varchar(100),
  "timestamp" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE INDEX "idx_employee_timestamp" ON "AttendanceLog" (DATE("timestamp"));
CREATE INDEX "idx_employee_nip" ON "AttendanceLog" ("employee_nip");

ALTER TABLE "AttendanceLog" ADD FOREIGN KEY ("employee_nip") REFERENCES "Employee" ("nip");
