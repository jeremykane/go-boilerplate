-- Deploy go-backend:1698972633_create_facility_types_table to pg

BEGIN;

CREATE TABLE IF NOT EXISTS facility_types (
    id bigserial NOT NULL,
    uuid varchar(64) NOT NULL,
    name varchar NOT NULL,
    bahasa_name varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT facility_types_pkey PRIMARY KEY (id),
    CONSTRAINT facility_types_un UNIQUE (uuid)
);


COMMIT;
