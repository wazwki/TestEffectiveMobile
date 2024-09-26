-- 000001_create_songs_table.up.sql
CREATE TABLE IF NOT EXISTS songs (
  id SERIAL PRIMARY KEY,
  group_name VARCHAR(255) NOT NULL,
  song_name VARCHAR(255) NOT NULL,
  release_date DATE,
  text TEXT,
  link TEXT
);