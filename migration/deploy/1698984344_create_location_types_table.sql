-- Deploy go-backend:1698984344_create_location_types_table to pg

BEGIN;

CREATE TABLE IF NOT EXISTS location_types (
    id BIGSERIAL NOT NULL,
    uuid varchar(64) NOT NULL,
    "name" varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT location_types_pkey PRIMARY KEY (id),
    CONSTRAINT location_types_un UNIQUE (uuid)
);


COMMIT;
