BEGIN;

CREATE TABLE IF NOT EXISTS public.accounts(
    id            INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    external_id   uuid                                         NOT NULL default gen_random_uuid(), 
    name          text                                         NOT NULL,
    cpf           text                                         NOT NULL,
    secret        text                                         NOT NULL,
    balance       bigint                                       NOT NULL,
    created_at    timestamp with time zone                     NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT account_cpf_uk UNIQUE (cpf)
);

COMMIT;
