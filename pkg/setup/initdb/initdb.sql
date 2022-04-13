\c db_nra;

CREATE TABLE component (
  id serial PRIMARY KEY,
  name VARCHAR,
  creation TIMESTAMP
);

CREATE TABLE release_manager (
  id serial PRIMARY KEY,
  name VARCHAR,
  version VARCHAR
);


INSERT INTO component (name, creation) VALUES ('Component', '2004-10-19 10:23:54');
INSERT INTO release_manager (name, version) VALUES ('Relman', '3.x');


