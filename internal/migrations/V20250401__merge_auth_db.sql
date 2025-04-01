BEGIN;

CREATE TABLE public.user
(
    id       BIGSERIAL NOT NULL,
    name     TEXT,
    email    TEXT      NOT NULL,
    password TEXT      NOT NULL,
    CONSTRAINT user_id_pk PRIMARY KEY (id),
    CONSTRAINT user_ukey UNIQUE (name, email)
);

ALTER TABLE IF EXISTS public.user_skill
    ADD CONSTRAINT user_fk FOREIGN KEY (id_user)
        REFERENCES public.user (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;

CREATE TABLE token
(
    id      BIGSERIAL NOT NULL
        CONSTRAINT token_pk PRIMARY KEY,
    token   TEXT      NOT NULL,
    user_id BIGSERIAL NOT NULL
        CONSTRAINT token_uk UNIQUE
        CONSTRAINT token_users_id_fk REFERENCES public.user
);


CREATE TABLE IF NOT EXISTS public.token
(
    id      BIGSERIAL NOT NULL,
    token   TEXT      NOT NULL,
    user_id bigint    NOT NULL,
    CONSTRAINT token_pk PRIMARY KEY (id),
    CONSTRAINT token_uk UNIQUE (user_id)
);

END;