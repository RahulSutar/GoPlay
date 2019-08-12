--
-- PostgreSQL database dump
--

-- Dumped from database version 11.4 (Ubuntu 11.4-1.pgdg18.04+1)
-- Dumped by pg_dump version 11.4 (Ubuntu 11.4-1.pgdg18.04+1)

-- Started on 2019-08-13 01:11:52 IST

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

--
-- TOC entry 197 (class 1259 OID 16400)
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    book_name text,
    author text,
    book_type text,
    book_description text,
    number_of_copies bigint,
    availability_status text
);


ALTER TABLE public.books OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 16398)
-- Name: books_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_id_seq OWNER TO postgres;

--
-- TOC entry 2932 (class 0 OID 0)
-- Dependencies: 196
-- Name: books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;


--
-- TOC entry 2800 (class 2604 OID 16403)
-- Name: books id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);


--
-- TOC entry 2926 (class 0 OID 16400)
-- Dependencies: 197
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.books (id, created_at, updated_at, deleted_at, book_name, author, book_type, book_description, number_of_copies, availability_status) FROM stdin;
41	2019-08-01 17:55:58.842814+05:30	2019-08-13 00:53:33.440276+05:30	\N	HowToCodeInGO	TestAuthor	Technical	Technical book for learning go lang	2	available
\.


--
-- TOC entry 2933 (class 0 OID 0)
-- Dependencies: 196
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.books_id_seq', 41, true);


--
-- TOC entry 2802 (class 2606 OID 16408)
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- TOC entry 2803 (class 1259 OID 16409)
-- Name: idx_books_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_books_deleted_at ON public.books USING btree (deleted_at);


-- Completed on 2019-08-13 01:11:52 IST

--
-- PostgreSQL database dump complete
--

