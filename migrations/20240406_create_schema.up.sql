CREATE TABLE authors (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL
);

CREATE TABLE poems (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    author_id uuid REFERENCES authors(id),
    text character varying NOT NULL,
    title character varying,
    date timestamp without time zone
);