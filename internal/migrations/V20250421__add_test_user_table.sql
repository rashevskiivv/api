create table test_user
(
    id_test             bigserial
        constraint test_user_test_id_fk
            references test,
    id_user             bigserial
        constraint test_user_user_id_fk
            references public."user" (id),
    score               int     not null,
    number_of_questions integer not null,
    constraint test_user_pkey
        primary key (id_test, id_user)
);

