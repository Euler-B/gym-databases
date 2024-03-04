CREATE TABLE album (
    id       SERIAL PRIMARY KEY,
    title    VARCHAR(128) NOT NULL,
    artist   VARCHAR(255) NOT NULL,
    price    FLOAT(4) NOT NULL,
    currency VARCHAR(3) NOT NULL
);

INSERT INTO album
    (title, artist, price, currency)
VALUES
    ('Coleccion Suprema', 'Los Prisioneros', 6.99, 'USD'),
    ('Lo Mejor de Vilma Palma', 'Charly Garcia', 8.80, 'EUR'),
    ('La Ley MTV Unplugged', 'La Ley', 7.30, 'GBP'),
    ('Garcia 87/93', 'Charly Garcia', 866.15, 'CVE'),
    ('Cancion Animal', 'Soda Estereo', 1279.43, 'JPY'),
    ('Signos', 'Soda Estereo', 324.32, 'VES'),
    ('Sueños Liquidos', 'Soda Estereo', 2515.03, 'PKR'),
    ('Mil Siluetas', 'La Union', 2828.98, 'AMD');

    -- Leyenda de las monedas 
    -- USD = Dolar Estadounidense
    -- EUR = EURO
    -- GBP = Libra Esterlina
    -- JPY = Yen Japones
    -- VES = Bolivara Soberano
    -- PRK = Rupia Pakistani
    -- AMD = Dram Armenio
    -- CVE = Escudo Caboverdiano