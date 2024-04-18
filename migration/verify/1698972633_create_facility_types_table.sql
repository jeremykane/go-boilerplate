-- Verify go-backend:1698972633_create_facility_types_table on pg

BEGIN;

SELECT id, uuid, "name", bahasa_name, created_at, updated_at
FROM facility_types WHERE TRUE;


ROLLBACK;
