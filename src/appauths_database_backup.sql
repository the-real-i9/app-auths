--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3 (Ubuntu 16.3-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.3 (Ubuntu 16.3-1.pgdg22.04+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: auth_user; Type: TABLE; Schema: public; Owner: i9
--

CREATE TABLE public.auth_user (
    user_id integer NOT NULL,
    email character varying NOT NULL,
    username character varying NOT NULL,
    password character varying NOT NULL,
    totp_setup_key character varying,
    mfa_enabled boolean DEFAULT false NOT NULL,
    mfa_type character varying,
    CONSTRAINT mfa_type_options CHECK ((((mfa_type)::text = 'otp'::text) OR ((mfa_type)::text = 'totp'::text)))
);


ALTER TABLE public.auth_user OWNER TO i9;

--
-- Name: auth_user_user_id_seq; Type: SEQUENCE; Schema: public; Owner: i9
--

CREATE SEQUENCE public.auth_user_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.auth_user_user_id_seq OWNER TO i9;

--
-- Name: auth_user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: i9
--

ALTER SEQUENCE public.auth_user_user_id_seq OWNED BY public.auth_user.user_id;


--
-- Name: ongoing_auth; Type: TABLE; Schema: public; Owner: i9
--

CREATE TABLE public.ongoing_auth (
    k character varying(64) DEFAULT ''::character varying NOT NULL,
    v bytea NOT NULL,
    e bigint DEFAULT '0'::bigint NOT NULL
);


ALTER TABLE public.ongoing_auth OWNER TO i9;

--
-- Name: ongoing_process; Type: TABLE; Schema: public; Owner: i9
--

CREATE TABLE public.ongoing_process (
    k character varying(64) DEFAULT ''::character varying NOT NULL,
    v bytea NOT NULL,
    e bigint DEFAULT '0'::bigint NOT NULL
);


ALTER TABLE public.ongoing_process OWNER TO i9;

--
-- Name: auth_user user_id; Type: DEFAULT; Schema: public; Owner: i9
--

ALTER TABLE ONLY public.auth_user ALTER COLUMN user_id SET DEFAULT nextval('public.auth_user_user_id_seq'::regclass);


--
-- Data for Name: auth_user; Type: TABLE DATA; Schema: public; Owner: i9
--

COPY public.auth_user (user_id, email, username, password, totp_setup_key, mfa_enabled, mfa_type) FROM stdin;
\.


--
-- Data for Name: ongoing_auth; Type: TABLE DATA; Schema: public; Owner: i9
--

COPY public.ongoing_auth (k, v, e) FROM stdin;
\.


--
-- Data for Name: ongoing_process; Type: TABLE DATA; Schema: public; Owner: i9
--

COPY public.ongoing_process (k, v, e) FROM stdin;
\.


--
-- Name: auth_user_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: i9
--

SELECT pg_catalog.setval('public.auth_user_user_id_seq', 1, false);


--
-- Name: ongoing_process ongoing_process_pkey; Type: CONSTRAINT; Schema: public; Owner: i9
--

ALTER TABLE ONLY public.ongoing_process
    ADD CONSTRAINT ongoing_process_pkey PRIMARY KEY (k);


--
-- Name: e; Type: INDEX; Schema: public; Owner: i9
--

CREATE INDEX e ON public.ongoing_auth USING btree (e);


--
-- PostgreSQL database dump complete
--

