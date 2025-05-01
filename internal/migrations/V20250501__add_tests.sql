--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+1)
-- Dumped by pg_dump version 16.8

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: skill; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.skill (id, title) FROM stdin;
1	git
2	linux
\.


--
-- Data for Name: test; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.test (id, title, description, duration, id_skill) FROM stdin;
1	Тест на знание git	Git - негласный стандарт среди VCS. Этот тест по проверке базовых знаний.	10	1
2	Тест на знание Linux	Linux часто используется как серверная ОС. Знания по работе с UNIX-подобными системами являются важными для любого специалиста в IT.	8	2
\.


--
-- Data for Name: question; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.question (id, question, id_test) FROM stdin;
1	Чем отличается git pull от git fetch?	1
2	Чем отличается git merge от git rebase?	1
3	Что такое репозиторий?	1
4	Какой командой можно отправить подтвержденные и включенные изменения на сервер?	1
5	Что делает команда git commit?	1
6	Как инициализировать репозиторий git локально?	1
7	git cherry-pick - это...	1
8	Как в консоли переключить ветку?	1
9	Эта команда создает локальную ветку с указанным названием. Что это за команда?	1
10	Что означает статус файла untracked в выводе команды git status?	1
11	Какой командой можно узнать текущую директорию?	2
12	Что делает команда ls?	2
13	Команда для создания новой директории?	2
14	Как посмотреть содержимое текстового файла?	2
15	Какой командой можно удалить файл?	2
16	Что делает команда chmod?	2
17	Какой командой посмотреть запущенные процессы?	2
18	Какая команда используется для установки пакетов в Debian/Ubuntu?	2
19	Что означает `..` в командной строке Linux?	2
20	Как завершить процесс по PID?	2
\.


--
-- Data for Name: answer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.answer (id, answer, id_question, is_right) FROM stdin;
1	git fetch скачивает изменения и не меняет состояния файлов, когда git pull выполняет скачивание и изменяет состояния файлов	1	t
2	git fetch - это связка git pull и git merge, когда git pull только скачивает изменения с удаленного репозитория	1	f
3	git pull и git fetch делают одно и то же	1	f
4	git merge от git rebase ничем не отличается	2	f
5	git merge производит слияние, когда git rebase копирует коммитыgit merge производит слияние, когда git rebase копирует коммиты	2	t
6	git rebase производит слияние, когда git merge копирует коммиты	2	f
7	git merge удаляет старые коммиты, когда git rebase создает новый удаленный репозиторий	2	f
8	Это место хранения исходного кода проекта и его истории изменений	3	t
9	Это сервер, где хранятся только готовые версии программ	3	f
10	Это система контроля версий	3	f
11	Это локальная директория	3	f
12	git push	4	t
13	git merge	4	f
14	git add	4	f
15	git commit	4	f
16	Отправляет изменения на сервер	5	f
17	Начинает отслеживать указанные файлы	5	f
18	Сохраняет зафиксированные изменения в локальной истории	5	t
19	Удаляет старые коммиты и оставляет только последний	5	f
20	git init	6	t
21	git start	6	f
22	git create	6	f
23	git add	6	f
24	Слияние всех изменений из другой ветки	7	f
25	Применение одного или нескольких конкретных коммитов из другой ветки	7	t
26	Удаление одного или нескольких конретных коммитов из текущей локальной ветки	7	f
27	Копирование последнего коммита из главной (стандартной) ветки	7	f
28	git checkout <branch-name>	8	t
29	git branch <branch-name>	8	f
30	git move <branch-name>	8	f
31	git switch <branch-name>	8	f
32	git branch <branch-name>	9	t
33	git switch -c <branch-name>	9	t
34	git switch <branch-name>	9	f
35	git new <branch-name>	9	f
36	Файл не отслеживается Git'ом и не был добавлен в индекс	10	t
37	Файл уже закоммичен	10	f
38	Файл удалён из репозитория	10	f
39	Файл отслеживается, но ещё не был закоммичен	10	f
40	pwd	11	t
41	ls	11	f
42	cd	11	f
43	dir	11	f
44	Создает директорию	12	f
45	Переходит в директорию	12	f
46	Показывает список файлов	12	t
47	Удаляет файлы	12	f
48	mkdir	13	t
49	rmdir	13	f
50	touch	13	f
51	newdir	13	f
52	cat	14	t
53	less	14	t
54	open	14	f
55	cd	14	f
56	rm	15	t
57	del	15	f
58	rmdir	15	f
59	erase	15	f
60	Меняет права доступа к файлам	16	t
61	Меняет владельца	16	f
62	Меняет имя файла	16	f
63	Перемещает файл	16	f
64	jobs	17	f
65	ps	17	t
66	top	17	t
67	run	17	f
68	yum	18	f
69	apt install	18	t
70	rpm install	18	f
71	dnf	18	f
72	Текущая директория	19	f
73	Домашняя директория	19	f
74	Родительская директория	19	t
75	Корневая директория	19	f
76	kill	20	t
77	end	20	f
78	terminate	20	f
79	stop	20	f
\.

--
-- Name: answer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.answer_id_seq', 79, true);


--
-- Name: question_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.question_id_seq', 20, true);


--
-- Name: skill_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.skill_id_seq', 2, true);


--
-- Name: test_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.test_id_seq', 2, true);


--
-- Name: test_user_id_test_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.test_user_id_test_seq', 1, false);


--
-- Name: test_user_id_user_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.test_user_id_user_seq', 1, false);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 1, false);


--
-- Name: vacancy_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vacancy_id_seq', 1, false);


--
-- PostgreSQL database dump complete
--

