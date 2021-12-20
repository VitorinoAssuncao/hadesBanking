BEGIN;

CREATE TABLE IF NOT EXISTS public.accounts(
    id            VARCHAR(36) COLLATE     pg_catalog."default" NOT NULL,
    name          text COLLATE            pg_catalog."default" NOT NULL,
    cpf           text COLLATE            pg_catalog."default" NOT NULL,
    secret        text COLLATE            pg_catalog."default" NOT NULL,
    balance       bigint                                       NOT NULL,
    created_at    timestamp with time zone                     NOT NULL,
    CONSTRAINT accounts_pkey PRIMARY KEY (id)
);

COMMIT;
