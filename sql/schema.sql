CREATE TABLE setting (
    title VARCHAR PRIMARY KEY,
    detail JSONB NOT NULL
);

-- Insert Default Settings
INSERT INTO setting (title, detail) VALUES ('general', '{"communityName": "Amigo", "description": "An elegant weapon from a more civilizd age.","defaultTheme":"default23"}');

CREATE TABLE "user" (
    id uuid PRIMARY KEY,
    user_number SERIAL NOT NULL,
    display_name VARCHAR
);

CREATE UNIQUE INDEX unique_display_name ON "user" (display_name);

CREATE TABLE preference (
    user_id uuid REFERENCES "user"(id) NOT NULL PRIMARY KEY,
    theme VARCHAR
);
