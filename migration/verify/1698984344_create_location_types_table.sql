-- Verify go-backend:1698984344_create_location_types_table on pg

BEGIN;

SELECT id, uuid, "name", created_at, updated_at
FROM location_types WHERE TRUE;

ROLLBACK;
