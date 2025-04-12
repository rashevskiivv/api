BEGIN;


CREATE TABLE IF NOT EXISTS public.answer
(
    id          bigserial NOT NULL,
    answer      text      NOT NULL,
    id_question bigint    NOT NULL,
    is_right    boolean   NOT NULL,
    CONSTRAINT answer_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.question
(
    id       bigserial NOT NULL,
    question text      NOT NULL,
    id_test  bigint    NOT NULL,
    CONSTRAINT question_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.skill
(
    id    bigserial NOT NULL,
    title text      NOT NULL,
    CONSTRAINT skill_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.skill_vacancy
(
    id_vacancy bigint NOT NULL,
    id_skill   bigint NOT NULL,
    CONSTRAINT skill_vacancy_pkey PRIMARY KEY (id_vacancy, id_skill)
);

CREATE TABLE IF NOT EXISTS public.test
(
    id          bigserial NOT NULL,
    title       text      NOT NULL,
    description text,
    duration    int,
    id_skill    bigint,
    CONSTRAINT test_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.user_skill
(
    id_user           bigint  NOT NULL,
    id_skill          bigint  NOT NULL,
    proficiency_level integer NOT NULL,
    CONSTRAINT user_skill_pkey PRIMARY KEY (id_user, id_skill)
);

CREATE TABLE IF NOT EXISTS public.vacancy
(
    id          bigserial NOT NULL,
    title       text      NOT NULL,
    grade       text,
    date        timestamp without time zone,
    description text,
    CONSTRAINT vacancy_pkey PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.answer
    ADD CONSTRAINT question_fk FOREIGN KEY (id_question)
        REFERENCES public.question (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID;


ALTER TABLE IF EXISTS public.question
    ADD CONSTRAINT test_fk FOREIGN KEY (id_test)
        REFERENCES public.test (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID;


ALTER TABLE IF EXISTS public.skill_vacancy
    ADD CONSTRAINT skill_fk FOREIGN KEY (id_skill)
        REFERENCES public.skill (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.skill_vacancy
    ADD CONSTRAINT vacancy_fk FOREIGN KEY (id_vacancy)
        REFERENCES public.vacancy (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.test
    ADD CONSTRAINT skill_fk FOREIGN KEY (id_skill)
        REFERENCES public.skill (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID;


ALTER TABLE IF EXISTS public.user_skill
    ADD CONSTRAINT skill_fk FOREIGN KEY (id_skill)
        REFERENCES public.skill (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;

END;