DO $$ BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'beaver-admin') THEN
        CREATE DATABASE "beaver-admin";
    END IF;
END $$;