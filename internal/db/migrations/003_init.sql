-- Create universities table
CREATE TABLE IF NOT EXISTS universities (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    city_id INTEGER NOT NULL REFERENCES city(id) ON DELETE RESTRICT,
    address TEXT,
    UNIQUE(name, city_id)
);

-- Load universities data
CREATE TEMP TABLE tmp_university (
    name TEXT,
    region TEXT,
    city TEXT,
    address TEXT
);
COPY tmp_university(name, region, city, address)
FROM '/docker-entrypoint-initdb.d/university.csv'
WITH (FORMAT csv, DELIMITER ';', HEADER true, ENCODING 'UTF8');

INSERT INTO universities(name, city_id, address)
SELECT
    u.name,
    c.id,
    u.address
FROM tmp_university u
JOIN city c ON u.city ILIKE ('%' || c.name || '%')
ON CONFLICT (name, city_id) DO NOTHING;

DROP TABLE tmp_university;

-- Create teams table
CREATE TABLE IF NOT EXISTS teams (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    city INTEGER NOT NULL REFERENCES city(id) ON DELETE RESTRICT,
    university INTEGER NOT NULL REFERENCES universities(id) ON DELETE RESTRICT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS participants (
    id BIGSERIAL PRIMARY KEY,
    team_id BIGINT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    middle_name TEXT,
    telegram TEXT NOT NULL,
    phone TEXT NOT NULL,
    course SMALLINT NOT NULL CHECK (course >= 1 AND course <= 6),
    email TEXT NOT NULL,
    shirt_size TEXT NOT NULL CHECK (shirt_size IN ('XS', 'S', 'M', 'L', 'XL', 'XXL')),
    is_captain BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS uniq_captain_per_team
ON participants(team_id)
WHERE is_captain = TRUE;

CREATE TABLE IF NOT EXISTS consents (
    id BIGSERIAL PRIMARY KEY,
    team_id BIGINT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    pdn_captain BOOLEAN NOT NULL,
    pdn_team BOOLEAN NOT NULL,
    rules_ack BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS emails_log (
    id BIGSERIAL PRIMARY KEY,
    team_id BIGINT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    email TEXT NOT NULL,
    subject TEXT NOT NULL,
    status TEXT NOT NULL,
    error TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);