BEGIN;

CREATE TABLE IF NOT EXISTS public.transfers
(
    id                  integer     NOT NULL,
    external_id         uuid        NOT NULL,
    account_origin_id   uuid        NOT NULL,
    account_destiny_id  uuid        NOT NULL,
    ammount             bigint      NOT NULL DEFAULT 0,
    created_at          timestamp with time zone    NOT NULL,
    
    CONSTRAINT transfers_pkey PRIMARY KEY (id)
);

COMMIT;
