--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3 (Debian 16.3-1.pgdg120+1)
-- Dumped by pg_dump version 16.3 (Debian 16.3-1.pgdg120+1)

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
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: genres; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.genres (
    id integer NOT NULL,
    genre character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.genres OWNER TO admin;

--
-- Name: genres_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.genres_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.genres_id_seq OWNER TO admin;

--
-- Name: genres_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.genres_id_seq OWNED BY public.genres.id;


--
-- Name: movies; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    release_date date,
    runtime integer,
    mpaa_rating character varying(10),
    description text,
    image character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.movies OWNER TO admin;

--
-- Name: movies_genres; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.movies_genres (
    movie_id integer NOT NULL,
    genre_id integer NOT NULL
);


ALTER TABLE public.movies_genres OWNER TO admin;

--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.movies_id_seq OWNER TO admin;

--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO admin;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO admin;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: genres id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.genres ALTER COLUMN id SET DEFAULT nextval('public.genres_id_seq'::regclass);


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.genres (id, genre, created_at, updated_at) FROM stdin;
1	Action	2024-09-04 05:59:20.475888	2024-09-04 05:59:20.475888
2	Comedy	2024-09-04 05:59:20.475888	2024-09-04 05:59:20.475888
3	Drama	2024-09-04 05:59:20.475888	2024-09-04 05:59:20.475888
4	Horror	2024-09-04 05:59:20.475888	2024-09-04 05:59:20.475888
5	Sci-Fi	2024-09-04 05:59:20.475888	2024-09-04 05:59:20.475888
\.


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.movies (id, title, release_date, runtime, mpaa_rating, description, image, created_at, updated_at) FROM stdin;
1	Inception	2010-07-16	148	PG-13	A thief who steals corporate secrets through the use of dream-sharing technology.	inception.jpg	2024-09-04 05:56:54.011137	2024-09-04 05:56:54.011137
2	The Matrix	1999-03-31	136	R	A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.	matrix.jpg	2024-09-04 05:56:54.011137	2024-09-04 05:56:54.011137
3	Interstellar	2014-11-07	169	PG-13	A team of explorers travel through a wormhole in space in an attempt to ensure humanity's survival.	interstellar.jpg	2024-09-04 05:56:54.011137	2024-09-04 05:56:54.011137
4	The Dark Knight	2008-07-18	152	PG-13	When the menace known as the Joker emerges from his mysterious past, he wreaks havoc and chaos on the people of Gotham.	dark_knight.jpg	2024-09-04 05:56:54.011137	2024-09-04 05:56:54.011137
5	Fight Club	1999-10-15	139	R	An insomniac office worker and a devil-may-care soap maker form an underground fight club that evolves into much more.	fight_club.jpg	2024-09-04 05:56:54.011137	2024-09-04 05:56:54.011137
\.


--
-- Data for Name: movies_genres; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.movies_genres (movie_id, genre_id) FROM stdin;
1	1
1	5
2	1
2	5
3	3
3	5
4	1
4	3
5	3
5	1
8	1
8	2
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.users (id, first_name, last_name, email, password, created_at, updated_at) FROM stdin;
1	John	Doe	admin@example.com	$2y$10$2EZg0uvLLzU86E6eo6hMI.25K4uaNnxRNfs3.nMVxh.rCQNqcA4E.	2024-09-04 05:53:16.910928	2024-09-04 05:53:16.910928
\.


--
-- Name: genres_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.genres_id_seq', 1, false);


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.movies_id_seq', 8, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: genres genres_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (id);


--
-- Name: movies_genres movies_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT movies_genres_pkey PRIMARY KEY (movie_id, genre_id);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: movies_genres movies_genres_genre_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT movies_genres_genre_id_fkey FOREIGN KEY (genre_id) REFERENCES public.genres(id);


--
-- PostgreSQL database dump complete
--

