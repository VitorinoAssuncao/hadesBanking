BEGIN;

CREATE TABLE IF NOT EXISTS public.transfers
(
    id                  INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    external_id         uuid        NOT NULL default gen_random_uuid(),
    account_origin_id   uuid        NOT NULL,
    account_destiny_id  uuid        NOT NULL,
    amount              bigint      NOT NULL DEFAULT 0,
    created_at          timestamp with time zone                     NOT NULL DEFAULT CURRENT_TIMESTAMP
); 

COMMIT;
