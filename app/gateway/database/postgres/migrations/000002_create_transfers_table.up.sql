BEGIN;

CREATE TABLE IF NOT EXISTS public.transfers
(
    id                  SERIAL     NOT NULL,
    external_id         VARCHAR(36)        NOT NULL,
    account_origin_id   VARCHAR(36)        NOT NULL,
    account_destiny_id  VARCHAR(36)        NOT NULL,
    ammount             bigint      NOT NULL DEFAULT 0,
    created_at          timestamp with time zone    NOT NULL,
    
    CONSTRAINT transfers_pkey PRIMARY KEY (id)
);

COMMIT;
