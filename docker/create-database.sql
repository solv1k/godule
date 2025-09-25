SELECT 'CREATE DATABASE go_api_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'go_api_db')\gexec