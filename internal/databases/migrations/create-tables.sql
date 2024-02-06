DROP TABLE IF EXISTS album;

CREATE TABLE album (
    id       SERIAL PRIMARY KEY,
    title    VARCHAR(128) NOT NULL,
    artist   VARCHAR(255) NOT NULL,
    price    FLOAT(4) NOT NULL
);

INSERT INTO album
    (title, artist, price)
VALUES
    ('Coleccion Suprema', 'Los Prisioneros', 6.99),
    ('Lo Mejor de Vilma Palma', 'Charly Garcia', 8.80),
    ('La Ley MTV Unplugged', 'La Ley', 7.30),
    ('Garcia 87/93', 'Charly Garcia', 8.50),
    ('Cancion Animal', 'Soda Estereo', 8.00),
    ('Signos', 'Soda Estereo', 9.00),
    ('Sueños Liquidos', 'Soda Estereo', 9.00),
    ('Mil Siluetas', 'La Union', 7.00);