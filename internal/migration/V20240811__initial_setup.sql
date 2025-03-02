-- https://www.red-gate.com/blog/database-devops/flyway-naming-patterns-matter
create table operations
(
    id              BIGSERIAL               not null
        constraint operations_id_pk
            primary key,
    user_id         BIGSERIAL               not null
        constraint operations_user_id_fk
            references users,
    category        TEXT    default 'Other' not null,
    amount_of_money INTEGER default 0       not null,
    is_income       BOOLEAN default false   not null
);

create table courses
(
    id                BIGSERIAL not null
        constraint courses_id_pk
            primary key,
    title             TEXT      not null unique,
    description       TEXT,
    short_description TEXT
);

create table lessons
(
    id            BIGSERIAL not null
        constraint lessons_id_pk
            primary key,
    text          TEXT      not null,
    link_to_video TEXT
);

create table questions
(
    id       BIGSERIAL not null
        constraint questions_id_pk
            primary key,
    question TEXT      not null unique
);

create table answers
(
    id     BIGSERIAL not null
        constraint answers_id_pk
            primary key,
    answer TEXT      not null unique
);

create table articles
(
    id          BIGSERIAL not null
        constraint articles_id_pk
            primary key,
    title       TEXT      not null unique,
    description TEXT,
    text        TEXT      not null
);

create table question_answers
(
    id          BIGSERIAL not null
        constraint question_answers_id_pk
            primary key,
    question_id BIGSERIAL not null
        constraint question_answers_qid_fk
            references questions,
    answer_id   BIGSERIAL not null
        constraint question_answers_aid_fk
            references answers,
    is_correct  BOOLEAN   not null
);

