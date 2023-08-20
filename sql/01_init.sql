CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    birthday DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE families (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    birthday DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO users (id, name, email, birthday) VALUES
    (1, 'A-san', 'j@hoge.jp', '1997-01-01'),
    (2, 'B-san', 'i@hoge.jp', '1998-01-02'),
    (3, 'C-san', 'h@hoge.jp', '1999-01-03'),
    (4, 'D-san', 'g@hoge.jp', '1991-01-04'),
    (5, 'E-san', 'f@hoge.jp', '1995-01-05'),
    (6, 'F-san', 'e@hoge.jp', '1996-01-06'),
    (7, 'G-san', 'd@hoge.jp', '1992-01-07'),
    (8, 'H-san', 'c@hoge.jp', '1993-01-08'),
    (9, 'I-san', 'b@hoge.jp', '1994-01-09'),
    (10, 'J-san', 'a@hoge.jp', '1990-01-10');

SELECT pg_catalog.setval('users_id_seq', 10);

INSERT INTO families (id, user_id, name, birthday) VALUES
    (1, 1, 'Aa-san', '2000-01-01'),
    (2, 1, 'Ab-san', '2000-01-02'),
    (3, 3, 'Ca-san', '2000-01-03'),
    (4, 3, 'Cb-san', '2000-01-04'),
    (5, 3, 'Cc-san', '2000-01-05'),
    (6, 6, 'Fa-san', '2000-01-06'),
    (7, 7, 'Ga-san', '2000-01-07'),
    (8, 8, 'Ha-san', '2000-01-08'),
    (9, 9, 'Ia-san', '2000-01-09'),
    (10, 9, 'Ib-san', '2000-01-10');

SELECT pg_catalog.setval('families_id_seq', 10);
