-- Crear la base de datos solo si no existe
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'series_tracker') THEN
        CREATE DATABASE series_tracker;
    END IF;
END $$;

-- Conectar a la base de datos 'series_tracker'
\c series_tracker;

-- Crear el usuario 'usuario' solo si no existe
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'usuario') THEN
        CREATE USER usuario WITH PASSWORD 'secret';
    END IF;
END $$;

-- Otorgar privilegios al usuario sobre la base de datos
GRANT ALL PRIVILEGES ON DATABASE series_tracker TO usuario;

-- Crear la tabla 'series'
CREATE TABLE IF NOT EXISTS series (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL UNIQUE,
    status VARCHAR(50) NOT NULL,
    last_episode_watched INT NOT NULL,
    total_episodes INT NOT NULL,
    ranking INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)
