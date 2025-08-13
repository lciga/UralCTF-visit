CREATE TABLE IF NOT EXISTS teams (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    city TEXT NOT NULL,
    university TEXT NOT NULL,
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