CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid PRIMARY KEY,
    "first_name" varchar(255),
    "last_name" varchar(255),
    "patronymic" varchar(255),
    "phone" varchar(255),
    "email" varchar(255),
    is_enabled bool default false,
    photo_url varchar(255),
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz
);

CREATE UNIQUE INDEX users_phone_unique_idx on users (phone);
CREATE UNIQUE INDEX users_email_unique_idx on users (email);