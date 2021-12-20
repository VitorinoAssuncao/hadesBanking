BEGIN;

CREATE TABLE IF NOT EXISTS public.transfers(
    id                          VARCHAR(36) COLLATE    pg_catalog."default"    NOT NULL,
    account_origin_id           VARCHAR(36) COLLATE    pg_catalog."default"     NOT NULL,
    account_destiny_id          VARCHAR(36) COLLATE    pg_catalog."default"     NOT NULL,
    balance                     bigint                                          NOT NULL,
    created_at    timestamp with time zone                     NOT NULL,
    CONSTRAINT transfers_pkey PRIMARY KEY (id)
);

COMMIT;
