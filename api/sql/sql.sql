DO $$ 
BEGIN 
   IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'devbook') THEN
      CREATE DATABASE devbook;
   END IF;
END $$;

\c devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(20) NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
