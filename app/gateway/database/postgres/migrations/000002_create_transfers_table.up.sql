BEGIN;

CREATE TABLE IF NOT EXISTS public.transfers
(
    id                  INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    external_id         uuid        NOT NULL default gen_random_uuid(),
    account_origin_id   INTEGER        NOT NULL,
    account_destiny_id  INTEGER        NOT NULL,
    amount              bigint      NOT NULL DEFAULT 0,
    created_at          timestamp with time zone                     NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_origin_id FOREIGN KEY(account_origin_id) REFERENCES accounts(id),
    CONSTRAINT fk_destiny_id FOREIGN KEY(account_destiny_id) REFERENCES accounts(id)
); 

COMMIT;
