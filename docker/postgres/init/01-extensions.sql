ALTER USER postgres WITH PASSWORD 'postgres';

GRANT ALL PRIVILEGES ON DATABASE assessment TO postgres;

-- Create uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

