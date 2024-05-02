CREATE TABLE IF NOT EXISTS user (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) NOT NULL,
  password VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS question (
  id SERIAL PRIMARY KEY,
  text TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS answer (
  id SERIAL PRIMARY KEY,
  question_id INT NOT NULL,
  text TEXT NOT NULL,
  FOREIGN KEY (question_id) REFERENCES question(id)
);
