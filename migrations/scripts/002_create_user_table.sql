CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    handle VARCHAR(10) UNIQUE NOT NULL,
    mobile_number CHAR(10) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    status VARCHAR(20) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT status_check CHECK (status IN ('active', 'inactive'))
);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_indexes
        WHERE schemaname = 'public' AND indexname = 'unique_email_not_deleted'
    ) THEN
        CREATE UNIQUE INDEX unique_email_not_deleted
        ON users (email) WHERE deleted_at is NULL;
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM pg_indexes
        WHERE schemaname = 'public' AND indexname = 'unique_mobile_number_not_deleted'
    ) THEN
        CREATE UNIQUE INDEX unique_mobile_number_not_deleted
        ON users (mobile_number) WHERE deleted_at is NULL;
    END IF;

    IF NOT EXISTS (
        SELECT 1
        FROM pg_trigger
        WHERE tgname = 'handle_user_updated_at'
        AND tgrelid = 'users'::regclass
    ) THEN
        CREATE TRIGGER handle_user_updated_at
        BEFORE UPDATE ON users
        FOR EACH ROW
        EXECUTE FUNCTION set_updated_at();
    END IF;
END
$$;
