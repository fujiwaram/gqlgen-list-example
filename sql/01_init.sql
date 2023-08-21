CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    birthday DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE friends (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id),
    friend_id INTEGER NOT NULL REFERENCES users (id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX friends_unique_idx ON friends (user_id, friend_id);

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

INSERT INTO friends (user_id, friend_id) VALUES
    (1,2),(1,3),(1,4),(1,5),(1,10),
    (2,1),
    (3,1),
    (4,1),
    (5,1),
    (6,7),(6,8),
    (7,6),
    (8,6);
