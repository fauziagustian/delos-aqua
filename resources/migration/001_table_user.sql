CREATE TABLE public.users (
    id bigint NOT NULL,
    name text,
    email text,
    password text
);

ALTER TABLE public.users OWNER TO aulianabil;

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO aulianabil;

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;