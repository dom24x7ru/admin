--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.14
-- Dumped by pg_dump version 10.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: goadmin_menu_myid_seq; Type: SEQUENCE; Schema: public; Owner: webadmin
--

CREATE SEQUENCE public.goadmin_menu_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_menu_myid_seq OWNER TO webadmin;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: goadmin_menu; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_menu (
    id integer DEFAULT nextval('public.goadmin_menu_myid_seq'::regclass) NOT NULL,
    parent_id integer DEFAULT 0 NOT NULL,
    type integer DEFAULT 0,
    "order" integer DEFAULT 0 NOT NULL,
    title character varying(50) NOT NULL,
    header character varying(100),
    plugin_name character varying(100) NOT NULL,
    icon character varying(50) NOT NULL,
    uri character varying(3000) NOT NULL,
    uuid character varying(100),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_menu OWNER TO webadmin;

--
-- Name: goadmin_operation_log_myid_seq; Type: SEQUENCE; Schema: public; Owner: webadmin
--

CREATE SEQUENCE public.goadmin_operation_log_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_operation_log_myid_seq OWNER TO webadmin;

--
-- Name: goadmin_operation_log; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_operation_log (
    id integer DEFAULT nextval('public.goadmin_operation_log_myid_seq'::regclass) NOT NULL,
    user_id integer NOT NULL,
    path character varying(255) NOT NULL,
    method character varying(10) NOT NULL,
    ip character varying(15) NOT NULL,
    input text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_operation_log OWNER TO webadmin;

--
-- Name: goadmin_site_myid_seq; Type: SEQUENCE; Schema: public; Owner: webadmin
--

CREATE SEQUENCE public.goadmin_site_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_site_myid_seq OWNER TO webadmin;

--
-- Name: goadmin_site; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_site (
    id integer DEFAULT nextval('public.goadmin_site_myid_seq'::regclass) NOT NULL,
    key character varying(100) NOT NULL,
    value text NOT NULL,
    type integer DEFAULT 0,
    description character varying(3000),
    state integer DEFAULT 0,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_site OWNER TO webadmin;

--
-- Name: goadmin_permissions_myid_seq; Type: SEQUENCE; Schema: public; Owner: webadmin
--

CREATE SEQUENCE public.goadmin_permissions_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_permissions_myid_seq OWNER TO webadmin;

--
-- Name: goadmin_permissions; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_permissions (
    id integer DEFAULT nextval('public.goadmin_permissions_myid_seq'::regclass) NOT NULL,
    name character varying(50) NOT NULL,
    slug character varying(50) NOT NULL,
    http_method character varying(255),
    http_path text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_permissions OWNER TO webadmin;

--
-- Name: goadmin_role_menu; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_role_menu (
    role_id integer NOT NULL,
    menu_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_role_menu OWNER TO webadmin;

--
-- Name: goadmin_role_permissions; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_role_permissions (
    role_id integer NOT NULL,
    permission_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_role_permissions OWNER TO webadmin;

--
-- Name: goadmin_role_users; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_role_users (
    role_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_role_users OWNER TO webadmin;

--
-- Name: goadmin_roles_myid_seq; Type: SEQUENCE; Schema: public; Owner: webadmin
--

CREATE SEQUENCE public.goadmin_roles_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_roles_myid_seq OWNER TO webadmin;

--
-- Name: goadmin_roles; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_roles (
    id integer DEFAULT nextval('public.goadmin_roles_myid_seq'::regclass) NOT NULL,
    name character varying NOT NULL,
    slug character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_roles OWNER TO webadmin;

--
-- Name: goadmin_session_myid_seq; Type: SEQUENCE; Schema: public; Owner: webadmin
--

CREATE SEQUENCE public.goadmin_session_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_session_myid_seq OWNER TO webadmin;

--
-- Name: goadmin_session; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_session (
    id integer DEFAULT nextval('public.goadmin_session_myid_seq'::regclass) NOT NULL,
    sid character varying(50) NOT NULL,
    "values" character varying(3000) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_session OWNER TO webadmin;

--
-- Name: goadmin_user_permissions; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_user_permissions (
    user_id integer NOT NULL,
    permission_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_user_permissions OWNER TO webadmin;

--
-- Name: goadmin_users_myid_seq; Type: SEQUENCE; Schema: public; Owner: webadmin
--

CREATE SEQUENCE public.goadmin_users_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_users_myid_seq OWNER TO webadmin;

--
-- Name: goadmin_users; Type: TABLE; Schema: public; Owner: webadmin
--

CREATE TABLE public.goadmin_users (
    id integer DEFAULT nextval('public.goadmin_users_myid_seq'::regclass) NOT NULL,
    username character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    avatar character varying(255),
    remember_token character varying(100),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_users OWNER TO webadmin;

--
-- Data for Name: goadmin_operation_log; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_operation_log (id, user_id, path, method, ip, input, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: goadmin_site; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_site (id, key, value, description, state, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: goadmin_permissions; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_permissions (id, name, slug, http_method, http_path, created_at, updated_at) FROM stdin;
1	All permission	*		*	2019-09-10 00:00:00	2019-09-10 00:00:00
2	Dashboard	dashboard	GET,PUT,POST,DELETE	/	2019-09-10 00:00:00	2019-09-10 00:00:00
\.


--
-- Data for Name: goadmin_role_menu; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_role_menu (role_id, menu_id, created_at, updated_at) FROM stdin;
1	1	2019-09-10 00:00:00	2019-09-10 00:00:00
1	7	2019-09-10 00:00:00	2019-09-10 00:00:00
2	7	2019-09-10 00:00:00	2019-09-10 00:00:00
\.


--
-- Data for Name: goadmin_role_permissions; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_role_permissions (role_id, permission_id, created_at, updated_at) FROM stdin;
1	1	2019-09-10 00:00:00	2019-09-10 00:00:00
1	2	2019-09-10 00:00:00	2019-09-10 00:00:00
2	2	2019-09-10 00:00:00	2019-09-10 00:00:00
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
\.


--
-- Data for Name: goadmin_role_users; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_role_users (role_id, user_id, created_at, updated_at) FROM stdin;
1	1	2019-09-10 00:00:00	2019-09-10 00:00:00
2	2	2019-09-10 00:00:00	2019-09-10 00:00:00
\.


--
-- Data for Name: goadmin_roles; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_roles (id, name, slug, created_at, updated_at) FROM stdin;
1	Administrator	administrator	2019-09-10 00:00:00	2019-09-10 00:00:00
2	Operator	operator	2019-09-10 00:00:00	2019-09-10 00:00:00
\.


--
-- Data for Name: goadmin_session; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_session (id, sid, "values", created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: goadmin_user_permissions; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_user_permissions (user_id, permission_id, created_at, updated_at) FROM stdin;
1	1	2019-09-10 00:00:00	2019-09-10 00:00:00
2	2	2019-09-10 00:00:00	2019-09-10 00:00:00
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
\.


--
-- Data for Name: goadmin_users; Type: TABLE DATA; Schema: public; Owner: webadmin
--

COPY public.goadmin_users (id, username, password, name, avatar, remember_token, created_at, updated_at) FROM stdin;
1	admin	$2a$10$OxWYJJGTP2gi00l2x06QuOWqw5VR47MQCJ0vNKnbMYfrutij10Hwe	admin		tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh	2019-09-10 00:00:00	2019-09-10 00:00:00
2	operator	$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.	Operator		\N	2019-09-10 00:00:00	2019-09-10 00:00:00
\.


--
-- Name: goadmin_menu_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: webadmin
--

SELECT pg_catalog.setval('public.goadmin_menu_myid_seq', 7, true);


--
-- Name: goadmin_operation_log_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: webadmin
--

SELECT pg_catalog.setval('public.goadmin_operation_log_myid_seq', 1, true);


--
-- Name: goadmin_permissions_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: webadmin
--

SELECT pg_catalog.setval('public.goadmin_permissions_myid_seq', 2, true);


--
-- Name: goadmin_roles_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: webadmin
--

SELECT pg_catalog.setval('public.goadmin_roles_myid_seq', 2, true);


--
-- Name: goadmin_site_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: webadmin
--

SELECT pg_catalog.setval('public.goadmin_site_myid_seq', 1, true);


--
-- Name: goadmin_session_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: webadmin
--

SELECT pg_catalog.setval('public.goadmin_session_myid_seq', 1, true);


--
-- Name: goadmin_users_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: webadmin
--

SELECT pg_catalog.setval('public.goadmin_users_myid_seq', 2, true);


--
-- Name: goadmin_menu goadmin_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: webadmin
--

ALTER TABLE ONLY public.goadmin_menu
    ADD CONSTRAINT goadmin_menu_pkey PRIMARY KEY (id);


--
-- Name: goadmin_operation_log goadmin_operation_log_pkey; Type: CONSTRAINT; Schema: public; Owner: webadmin
--

ALTER TABLE ONLY public.goadmin_operation_log
    ADD CONSTRAINT goadmin_operation_log_pkey PRIMARY KEY (id);


--
-- Name: goadmin_permissions goadmin_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: webadmin
--

ALTER TABLE ONLY public.goadmin_permissions
    ADD CONSTRAINT goadmin_permissions_pkey PRIMARY KEY (id);


--
-- Name: goadmin_roles goadmin_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: webadmin
--

ALTER TABLE ONLY public.goadmin_roles
    ADD CONSTRAINT goadmin_roles_pkey PRIMARY KEY (id);


--
-- Name: goadmin_site goadmin_site_pkey; Type: CONSTRAINT; Schema: public; Owner: webadmin
--

ALTER TABLE ONLY public.goadmin_site
    ADD CONSTRAINT goadmin_site_pkey PRIMARY KEY (id);


--
-- Name: goadmin_session goadmin_session_pkey; Type: CONSTRAINT; Schema: public; Owner: webadmin
--

ALTER TABLE ONLY public.goadmin_session
    ADD CONSTRAINT goadmin_session_pkey PRIMARY KEY (id);


--
-- Name: goadmin_users goadmin_users_pkey; Type: CONSTRAINT; Schema: public; Owner: webadmin
--

ALTER TABLE ONLY public.goadmin_users
    ADD CONSTRAINT goadmin_users_pkey PRIMARY KEY (id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: webadmin
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM webadmin;
GRANT ALL ON SCHEMA public TO webadmin;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

