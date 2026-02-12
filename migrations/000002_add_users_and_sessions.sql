-- +goose Up
-- Users
create table users (
    id              uuid primary key default gen_random_uuid(),
    email           varchar(255) not null unique,
    name            varchar(100),
    avatar_url      text,
    is_active       boolean default true,
    created_at      timestamptz default now(),
    updated_at      timestamptz default now()
);

-- OAuth Accounts (1 user -> many providers)
create table oauth_accounts (
    id              uuid primary key default gen_random_uuid(),
    user_id         uuid not null references users(id) on delete cascade,
    provider        varchar(20) not null,   -- 'google', 'github'
    provider_id     varchar(255) not null,
    access_token    text,
    refresh_token   text,
    created_at      timestamptz default now(),
    updated_at      timestamptz default now(),

    unique(provider, provider_id)
);

-- Sessions (Server-side tracking)
create table sessions (
    id              uuid primary key default gen_random_uuid(),
    user_id         uuid not null references users(id) on delete cascade,
    token_hash      varchar(255) not null unique,
    expires_at      timestamptz not null,
    created_at      timestamptz default now()
);

-- Indexes
create index idx_users_email on users(email);
create index idx_oauth_accounts_user_id on oauth_accounts(user_id);
create index idx_sessions_token_hash on sessions(token_hash);
create index idx_sessions_expires_at on sessions(expires_at);

-- Triggers for updated_at
create trigger update_users_updated_at before update on users for each row execute function update_updated_at_column();
create trigger update_oauth_accounts_updated_at before update on oauth_accounts for each row execute function update_updated_at_column();

-- +goose Down
drop trigger if exists update_oauth_accounts_updated_at on oauth_accounts;
drop trigger if exists update_users_updated_at on users;
drop table if exists sessions;
drop table if exists oauth_accounts;
drop table if exists users;
