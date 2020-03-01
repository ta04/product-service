--
-- PostgreSQL database dump
--

-- Dumped from database version 11.7 (Ubuntu 11.7-1.pgdg18.04+1)
-- Dumped by pg_dump version 11.7 (Ubuntu 11.7-1.pgdg18.04+1)

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

SET default_with_oids = false;

CREATE DATABASE product;

\connect product;

--
-- Name: product; Type: TABLE; Schema: public; Owner: sleepingnext
--

CREATE TABLE public.product (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(1000) NOT NULL,
    price numeric NOT NULL,
    picture character varying(255) NOT NULL,
    status boolean NOT NULL
);


ALTER TABLE public.product OWNER TO sleepingnext;

--
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: sleepingnext
--

CREATE SEQUENCE public.product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_id_seq OWNER TO sleepingnext;

--
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sleepingnext
--

ALTER SEQUENCE public.product_id_seq OWNED BY public.product.id;


--
-- Name: product id; Type: DEFAULT; Schema: public; Owner: sleepingnext
--

ALTER TABLE ONLY public.product ALTER COLUMN id SET DEFAULT nextval('public.product_id_seq'::regclass);


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: sleepingnext
--

COPY public.product (id, name, description, price, picture, status) FROM stdin;
7	This is a test product	Test	123.450000	test product.jpg	t
8	This is a test product	Test	123.450000	test product.jpg	t
9	This is a test product	This is a test product description	123.450000	test.jpg	t
6	This is a test product	This is a test product description	123.450000	test.jpg	t
\.


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sleepingnext
--

SELECT pg_catalog.setval('public.product_id_seq', 9, true);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: sleepingnext
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

