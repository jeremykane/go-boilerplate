-- Revert go-backend:1698984344_create_location_types_table from pg

BEGIN;

DROP TABLE IF EXISTS location_types;

COMMIT;
