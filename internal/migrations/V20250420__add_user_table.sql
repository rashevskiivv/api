create table public."user"
(
    id        BIGSERIAL not null
        constraint users_id_pk
            primary key,
    name      TEXT,
    email     TEXT      not null,
    interests TEXT[],
    constraint users_name_email_pk
        unique (email, name)
);