CREATE TABLE IF NOT EXISTS levels (
    id SERIAL PRIMARY KEY,
    map jsonb NOT NULL
);

CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    hitpoints INT NOT NULL,

    start_x INT NOT NULL,
    start_y INT NOT NULL,

    curr_x INT NOT NULL,
    curr_y INT NOT NULL,

    level_id INT NOT NULL,

    CONSTRAINT fk_level
        FOREIGN KEY(level_id)
            REFERENCES levels(id)
            ON DELETE CASCADE
);