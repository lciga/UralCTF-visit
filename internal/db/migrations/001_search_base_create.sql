-- Create regions table
CREATE TABLE IF NOT EXISTS region (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Create cities table
CREATE TABLE IF NOT EXISTS city (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);


-- загружаем данные во временную таблицу и вставляем без дублирования
-- Load raw locations into temp table
CREATE TEMP TABLE tmp_location (
    region TEXT,
    name TEXT
);
COPY tmp_location(region, name)
FROM '/docker-entrypoint-initdb.d/city.csv'
WITH (FORMAT csv, DELIMITER ';', HEADER true, ENCODING 'UTF8');

-- Insert distinct regions
INSERT INTO region(name)
SELECT DISTINCT region FROM tmp_location
ON CONFLICT (name) DO NOTHING;

-- Insert distinct cities
INSERT INTO city(name)
SELECT DISTINCT name FROM tmp_location
ON CONFLICT (name) DO NOTHING;

DROP TABLE tmp_location;