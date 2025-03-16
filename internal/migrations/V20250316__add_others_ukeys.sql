ALTER TABLE public.test
    ADD CONSTRAINT test_ukey
        UNIQUE (title, id_skill);

ALTER TABLE public.skill
    ADD CONSTRAINT skill_ukey
        UNIQUE (title);

ALTER TABLE public.question
    ADD CONSTRAINT question_ukey
        UNIQUE (question, id_test);

ALTER TABLE public.vacancy
    ADD CONSTRAINT vacancy_ukey
        UNIQUE (title, grade, date, description);
