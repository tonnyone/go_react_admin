--
-- PostgreSQL database dump
--

-- Dumped from database version 15.13 (Debian 15.13-0+deb12u1)
-- Dumped by pg_dump version 17.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

ALTER TABLE ONLY public.sys_user DROP CONSTRAINT users_pkey;
ALTER TABLE ONLY public.sys_user_role DROP CONSTRAINT user_roles_pkey;
ALTER TABLE ONLY public.sys_user DROP CONSTRAINT uni_users_phone;
ALTER TABLE ONLY public.sys_user DROP CONSTRAINT uni_users_email;
ALTER TABLE ONLY public.sys_role DROP CONSTRAINT uni_sys_role_name;
ALTER TABLE ONLY public.test DROP CONSTRAINT test_pkey;
ALTER TABLE ONLY public.sys_role DROP CONSTRAINT sys_role_pkey;
DROP TABLE public.test;
DROP TABLE public.sys_user_role;
DROP TABLE public.sys_user;
DROP TABLE public.sys_role;
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: sys_role; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sys_role (
    id character(22) NOT NULL,
    name character varying(64),
    describe character varying(256),
    type character varying(32) DEFAULT ''::character varying,
    created_at bigint,
    updated_at bigint,
    disabled boolean DEFAULT false
);


--
-- Name: sys_user; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sys_user (
    id character(22) NOT NULL,
    username character varying(64) NOT NULL,
    email character varying(128),
    phone character varying(64),
    password character varying(128) NOT NULL,
    department character varying(128),
    disabled boolean DEFAULT false,
    deleted boolean DEFAULT false,
    created_at bigint,
    updated_at bigint
);


--
-- Name: sys_user_role; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sys_user_role (
    user_id character(22) NOT NULL,
    role_id character(22) NOT NULL,
    created_at bigint,
    created_by character(36) NOT NULL
);


--
-- Name: test; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.test (
    id integer NOT NULL,
    create_time date,
    name character varying(255)
);


--
-- Name: test_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.test ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.test_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: sys_role; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.sys_role (id, name, describe, type, created_at, updated_at, disabled) FROM stdin;
MUmt8KTCEtRkTw84HpckbQ	普通用户	normal	custom	1756916648166	1756916648166	f
\.


--
-- Data for Name: sys_user; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.sys_user (id, username, email, phone, password, department, disabled, deleted, created_at, updated_at) FROM stdin;
3ko3nN62Y8NTx9jhZikcY7	无用户名	pgtest@example.com		87d7edd5f2568736ddb6b824041fa88c	sdaf	f	f	1756223070663	1756223070663
4DBSg76iuTGtiBQ4Wpt8vC	无用户名		18888888888	3f1f84d2aae285bb1b3ebc47ec576f72	asdfasdf	f	f	1756224142420	1756224142420
\.


--
-- Data for Name: sys_user_role; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.sys_user_role (user_id, role_id, created_at, created_by) FROM stdin;
3ko3nN62Y8NTx9jhZikcY7	MUmt8KTCEtRkTw84HpckbQ	1756917194791	                                    
\.


--
-- Data for Name: test; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.test (id, create_time, name) FROM stdin;
\.


--
-- Name: test_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.test_id_seq', 1, false);


--
-- Name: sys_role sys_role_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sys_role
    ADD CONSTRAINT sys_role_pkey PRIMARY KEY (id);


--
-- Name: test test_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.test
    ADD CONSTRAINT test_pkey PRIMARY KEY (id);


--
-- Name: sys_role uni_sys_role_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sys_role
    ADD CONSTRAINT uni_sys_role_name UNIQUE (name);


--
-- Name: sys_user uni_users_email; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sys_user
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- Name: sys_user uni_users_phone; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sys_user
    ADD CONSTRAINT uni_users_phone UNIQUE (phone);


--
-- Name: sys_user_role user_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sys_user_role
    ADD CONSTRAINT user_roles_pkey PRIMARY KEY (user_id, role_id);


--
-- Name: sys_user users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sys_user
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

