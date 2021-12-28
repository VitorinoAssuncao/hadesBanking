BEGIN;

CREATE TABLE IF NOT EXISTS public.accounts(
    id            VARCHAR(36)                                  NOT NULL,
    name          text                                         NOT NULL,
    cpf           text                                         NOT NULL,
    secret        text                                         NOT NULL,
    balance       bigint                                       NOT NULL,
    created_at    timestamp with time zone                     NOT NULL,
    CONSTRAINT accounts_pkey PRIMARY KEY (id)
);

COMMIT;
