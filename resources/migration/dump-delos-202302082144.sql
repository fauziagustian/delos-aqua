--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9
-- Dumped by pg_dump version 13.6

-- Started on 2023-02-08 21:44:27

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
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 2860 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 205 (class 1259 OID 33448)
-- Name: farms; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.farms (
    farm_id bigint NOT NULL,
    name text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.farms OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 33446)
-- Name: farms_farm_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.farms_farm_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.farms_farm_id_seq OWNER TO postgres;

--
-- TOC entry 2861 (class 0 OID 0)
-- Dependencies: 204
-- Name: farms_farm_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.farms_farm_id_seq OWNED BY public.farms.farm_id;


--
-- TOC entry 207 (class 1259 OID 33470)
-- Name: ponds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ponds (
    pond_id bigint NOT NULL,
    name text,
    farm_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.ponds OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 33468)
-- Name: ponds_pond_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ponds_pond_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ponds_pond_id_seq OWNER TO postgres;

--
-- TOC entry 2862 (class 0 OID 0)
-- Dependencies: 206
-- Name: ponds_pond_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ponds_pond_id_seq OWNED BY public.ponds.pond_id;


--
-- TOC entry 209 (class 1259 OID 33492)
-- Name: user_agents; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_agents (
    user_agent_id bigint NOT NULL,
    method_url text,
    user_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.user_agents OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 33490)
-- Name: user_agents_user_agent_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_agents_user_agent_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_agents_user_agent_id_seq OWNER TO postgres;

--
-- TOC entry 2863 (class 0 OID 0)
-- Dependencies: 208
-- Name: user_agents_user_agent_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_agents_user_agent_id_seq OWNED BY public.user_agents.user_agent_id;


--
-- TOC entry 203 (class 1259 OID 33415)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    user_id bigint NOT NULL,
    name text,
    email text,
    password text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 33413)
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_user_id_seq OWNER TO postgres;

--
-- TOC entry 2864 (class 0 OID 0)
-- Dependencies: 202
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- TOC entry 2710 (class 2604 OID 33451)
-- Name: farms farm_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.farms ALTER COLUMN farm_id SET DEFAULT nextval('public.farms_farm_id_seq'::regclass);


--
-- TOC entry 2711 (class 2604 OID 33473)
-- Name: ponds pond_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ponds ALTER COLUMN pond_id SET DEFAULT nextval('public.ponds_pond_id_seq'::regclass);


--
-- TOC entry 2712 (class 2604 OID 33495)
-- Name: user_agents user_agent_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_agents ALTER COLUMN user_agent_id SET DEFAULT nextval('public.user_agents_user_agent_id_seq'::regclass);


--
-- TOC entry 2709 (class 2604 OID 33418)
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- TOC entry 2850 (class 0 OID 33448)
-- Dependencies: 205
-- Data for Name: farms; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.farms (farm_id, name, created_at, updated_at, deleted_at) FROM stdin;
3	farm-3	2023-02-07 15:57:29.748602+07	2023-02-07 15:57:29.748602+07	\N
4	farm-4	2023-02-08 01:22:23.133407+07	2023-02-08 01:22:23.133407+07	\N
1	farm-1-edit	2023-02-07 15:50:27.573548+07	2023-02-08 01:42:22.676884+07	2023-02-08 02:10:09.068998+07
2	farm-2	2023-02-07 15:57:25.122106+07	2023-02-07 15:57:25.122106+07	2023-02-08 02:13:13.330461+07
5	farm-1	2023-02-08 02:14:15.830624+07	2023-02-08 02:14:15.830624+07	\N
\.


--
-- TOC entry 2852 (class 0 OID 33470)
-- Dependencies: 207
-- Data for Name: ponds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ponds (pond_id, name, farm_id, created_at, updated_at, deleted_at) FROM stdin;
2	pond2-B	4	2023-02-08 03:38:01.089697+07	2023-02-08 03:38:01.089697+07	\N
1	pond2-A	5	2023-02-08 03:34:54.780134+07	2023-02-08 11:39:58.771841+07	2023-02-08 12:09:08.918273+07
3	pond1-A	5	2023-02-08 13:30:28.236598+07	2023-02-08 13:30:28.236598+07	\N
\.


--
-- TOC entry 2854 (class 0 OID 33492)
-- Dependencies: 209
-- Data for Name: user_agents; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_agents (user_agent_id, method_url, user_id, created_at, updated_at, deleted_at) FROM stdin;
3	GET /pond	1	2023-02-08 13:37:52.572785+07	2023-02-08 13:37:52.572785+07	\N
4	GET /pond	1	2023-02-08 13:38:05.839535+07	2023-02-08 13:38:05.839535+07	\N
1	GET /farm	1	2023-02-08 12:53:53.309684+07	2023-02-08 12:53:53.309684+07	\N
2	GET /farm	1	2023-02-08 12:55:48.024606+07	2023-02-08 12:55:48.024606+07	\N
\.


--
-- TOC entry 2848 (class 0 OID 33415)
-- Dependencies: 203
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (user_id, name, email, password, created_at, updated_at, deleted_at) FROM stdin;
1	fauzi	fauzi@delos.com	$2a$04$MUHIM8UD633vRPpmhCltTu6MYITLQwgEWmocjO.IQ2wO/CTGV21Ei	2023-02-06 23:32:20.478433+07	2023-02-06 23:32:20.478433+07	\N
\.


--
-- TOC entry 2865 (class 0 OID 0)
-- Dependencies: 204
-- Name: farms_farm_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.farms_farm_id_seq', 5, true);


--
-- TOC entry 2866 (class 0 OID 0)
-- Dependencies: 206
-- Name: ponds_pond_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ponds_pond_id_seq', 3, true);


--
-- TOC entry 2867 (class 0 OID 0)
-- Dependencies: 208
-- Name: user_agents_user_agent_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_agents_user_agent_id_seq', 4, true);


--
-- TOC entry 2868 (class 0 OID 0)
-- Dependencies: 202
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_user_id_seq', 1, true);


--
-- TOC entry 2716 (class 2606 OID 33456)
-- Name: farms farms_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.farms
    ADD CONSTRAINT farms_pkey PRIMARY KEY (farm_id);


--
-- TOC entry 2718 (class 2606 OID 33478)
-- Name: ponds ponds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ponds
    ADD CONSTRAINT ponds_pkey PRIMARY KEY (pond_id);


--
-- TOC entry 2720 (class 2606 OID 33500)
-- Name: user_agents user_agents_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_agents
    ADD CONSTRAINT user_agents_pkey PRIMARY KEY (user_agent_id);


--
-- TOC entry 2714 (class 2606 OID 33423)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


-- Completed on 2023-02-08 21:44:28

--
-- PostgreSQL database dump complete
--

