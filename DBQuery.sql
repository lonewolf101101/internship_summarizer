CREATE TABLE model (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE users (
    uuid UUID PRIMARY KEY,
    auth_type VARCHAR(50) NOT NULL,
    password_hash TEXT,
    name VARCHAR(100),
    phone_number VARCHAR(20),
    email VARCHAR(100) UNIQUE NOT NULL,
    google_id VARCHAR(100) UNIQUE,
    profile_picture TEXT,
    is_verified BOOLEAN DEFAULT FALSE,
    last_login TIMESTAMPTZ,
    self_deleted_at TIMESTAMPTZ,
) INHERITS (model);

CREATE TABLE roles (
    rid SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
) INHERITS (model);

CREATE TABLE user_roles (
    urid SERIAL PRIMARY KEY, 
    uuid UUID NOT NULL,
    rid INT NOT NULL,  -- Change to INT to match the type of rid in roles
    name TEXT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (uuid) REFERENCES users(uuid) ON DELETE CASCADE,
    CONSTRAINT fk_role FOREIGN KEY (rid) REFERENCES roles(rid) ON DELETE CASCADE
) INHERITS (model);

CREATE TABLE sum_history (
    sid SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    contents TEXT,
    summary TEXT,
    CONSTRAINT fk_user FOREIGN KEY (uuid) REFERENCES users(uuid) ON DELETE CASCADE
) INHERITS (model);
