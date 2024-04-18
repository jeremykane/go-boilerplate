-- Revert go-backend:1698972633_create_facility_types_table from pg

BEGIN;

DROP TABLE IF EXISTS facility_types;

COMMIT;
