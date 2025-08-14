-- Create mapping of regions to cities
CREATE TABLE IF NOT EXISTS region_city (
    id SERIAL PRIMARY KEY,
    region_id INTEGER NOT NULL REFERENCES region(id) ON DELETE RESTRICT,
    city_id INTEGER NOT NULL REFERENCES city(id) ON DELETE RESTRICT,
    UNIQUE(region_id, city_id)
);

-- Load raw mappings into temp table
CREATE TEMP TABLE tmp_location (
    region TEXT,
    city TEXT
);
COPY tmp_location(region, city)
FROM '/docker-entrypoint-initdb.d/tables/city.csv'
WITH (FORMAT csv, DELIMITER ';', HEADER true, ENCODING 'UTF8');

-- Insert distinct region-city associations
INSERT INTO region_city(region_id, city_id)
SELECT r.id, c.id
FROM tmp_location t
JOIN region r ON t.region = r.name
JOIN city c ON t.city = c.name
ON CONFLICT (region_id, city_id) DO NOTHING;

DROP TABLE tmp_location;
