DO
$$
BEGIN
   IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'todos') THEN
      PERFORM dblink_exec('dbname=postgres', 'CREATE DATABASE todos');
   END IF;
END
$$;

\connect todos;

CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    created INT NOT NULL,
    due INT
);
