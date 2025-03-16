ALTER TABLE public.answer
    ADD CONSTRAINT answer_ukey
        UNIQUE (answer, id_question);