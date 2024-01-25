DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        CREATE TYPE public."user_role" AS ENUM (
            'admin',
            'moderator',
            'user'
        );
    END IF;
END $$;