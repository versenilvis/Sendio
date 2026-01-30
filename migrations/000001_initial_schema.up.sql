-- NOTE: AI generated, I was too lazy
-- Initial schema for Sendio

-- Plans
create table plans (
    id              uuid primary key default gen_random_uuid(),
    name            varchar(50) not null unique,
    max_quota       int not null,
    rate_limit      int not null,
    price           decimal(10,2) default 0,
    created_at      timestamptz default now()
);

insert into plans (name, max_quota, rate_limit, price) values
    ('free', 8000, 50, 0),
    ('premium', 20000, 200, 19.99),
    ('enterprise', 100000, 1000, 39.99);

-- API Keys
create table api_keys (
    id              uuid primary key default gen_random_uuid(),
    app_name        varchar(100) not null,
    key_hash        varchar(255) not null,
    key_prefix      varchar(8) not null,
    is_active       boolean default true,
    rate_limit      int default 100,
    created_at      timestamptz default now(),
    updated_at      timestamptz default now(),
    deleted_at      timestamptz
);

-- API Key Usage
create table api_key_usage (
    id              uuid primary key default gen_random_uuid(),
    api_key_id      uuid not null references api_keys(id) on delete cascade,
    -- on delete cascade: whenever a record in table B is deleted, any corresponding records in table A will also be deleted automatically
    usage_date      date not null,
    sent_count      int default 0,
    failed_count    int default 0,
    unique(api_key_id, usage_date)
);

-- Subscriptions
-- only active and trialing can use higher quota
create type subscription_status as enum (
    'active',       -- users can send emails
    'past_due',     -- payment failed, notify user to update payment method, grace period
    'cancelled',
    'expired',
    'trialing'      -- free trial for premium or enterprise (e.g. or 7 days)
);

create table subscriptions (
    id              uuid primary key default gen_random_uuid(),
    api_key_id      uuid not null references api_keys(id) on delete cascade,
    plan_id         uuid not null references plans(id),
    status          subscription_status default 'active',
    billing_period  varchar(20) default 'monthly',
    current_period_start timestamptz not null,
    current_period_end   timestamptz not null,
    quota_limit     int not null,
    quota_used      int default 0,
    quota_reset_at  timestamptz not null,
    -- stripe and vietnamese payment gateway
    payment_provider       varchar(50) default 'system', -- stripe, vietqr, sepay, manual
    provider_sub_id        varchar(255),
    payment_metadata       jsonb default '{}',
    created_at      timestamptz default now(),
    updated_at      timestamptz default now(),
    cancelled_at    timestamptz
);

-- Payment Transactions
create table payment_transactions (
    id              uuid primary key default gen_random_uuid(),
    subscription_id uuid not null references subscriptions(id) on delete cascade,
    amount          decimal(10,2) not null,
    currency        varchar(10) default 'VND',
    status          varchar(20) default 'pending',
    provider        varchar(50) not null,
    external_id     varchar(255),
    metadata        jsonb default '{}',
    created_at      timestamptz default now(),
    confirmed_at    timestamptz
);

-- Email Templates
create table email_templates (
    id              uuid primary key default gen_random_uuid(),
    api_key_id      uuid references api_keys(id) on delete cascade,
    name            varchar(100) not null,
    description     text,
    subject_template varchar(500),
    body_html       text,
    body_text       text,
    variables       jsonb default '[]',
    is_active       boolean default true,
    created_at      timestamptz default now(),
    updated_at      timestamptz default now()
);

-- Email Status Enum
create type email_status as enum (
    'pending',
    'queued',
    'processing',
    'sent',
    'delivered',
    'failed',
    'bounced',
    'cancelled'
);

-- Emails
create table emails (
    id              uuid primary key default gen_random_uuid(),
    api_key_id      uuid not null references api_keys(id),
    idempotency_key varchar(255) unique, -- though success or fail, it still counts as a single request
    from_email      varchar(255) not null,
    from_name       varchar(100),
    to_email        varchar(255) not null,
    to_name         varchar(100),
    subject         varchar(500) not null,
    template_id     uuid references email_templates(id),
    template_vars   jsonb,
    body_html       text,
    body_text       text,
    tags            jsonb default '[]',
    metadata        jsonb default '{}',
    status          email_status default 'pending',
    retry_count     int default 0,
    max_retries     int default 3,
    created_at      timestamptz default now(),
    scheduled_at    timestamptz,
    queued_at       timestamptz,
    sent_at         timestamptz,
    delivered_at    timestamptz
);

-- SMTP Configs
create table smtp_configs (
    id              uuid primary key default gen_random_uuid(),
    name            varchar(100) not null,
    host            varchar(255) not null,
    port            int not null,
    username        varchar(255),
    password_enc    text,
    use_tls         boolean default true,
    use_ssl         boolean default false,
    is_default      boolean default false,
    is_active       boolean default true,
    last_used_at    timestamptz,
    failure_count   int default 0,
    created_at      timestamptz default now(),
    updated_at      timestamptz default now()
);

-- Email Delivery Attempts
create table email_delivery_attempts (
    id              uuid primary key default gen_random_uuid(),
    email_id        uuid not null references emails(id) on delete cascade,
    attempt_num     int not null,
    status          varchar(20) not null,       -- success, failed, timeout
    smtp_code       varchar(10),                -- e.g., 250, 550, 421
    smtp_message    text,
    error_message   text,
    smtp_config_id  uuid references smtp_configs(id),
    duration_ms     int,
    created_at      timestamptz default now()
);

-- Email Events
create type email_event_type as enum (
    'created',
    'queued',
    'sent',
    'delivered',
    'opened',
    'clicked',
    'bounced',
    'complained',
    'unsubscribed',
    'failed'
);

create table email_events (
    id              uuid primary key default gen_random_uuid(),
    email_id        uuid not null references emails(id) on delete cascade,
    event_type      email_event_type not null,
    description     text,
    metadata        jsonb default '{}',
    created_at      timestamptz default now()
);

-- Indexes
create index idx_api_keys_prefix on api_keys(key_prefix);
create index idx_api_keys_active on api_keys(is_active) where is_active = true;
create index idx_api_key_usage_lookup on api_key_usage(api_key_id, usage_date);
create index idx_subscriptions_api_key on subscriptions(api_key_id);
create index idx_subscriptions_status on subscriptions(status) where status = 'active';
create index idx_subscriptions_provider_id on subscriptions(provider_sub_id) where provider_sub_id is not null;
create index idx_emails_api_key on emails(api_key_id);
create index idx_emails_status on emails(status);
create index idx_emails_status_created on emails(status, created_at);
create index idx_emails_idempotency on emails(idempotency_key) where idempotency_key is not null;
create index idx_emails_scheduled on emails(scheduled_at) where scheduled_at is not null and status = 'pending';
create index idx_emails_to_email on emails(to_email);
create index idx_delivery_attempts_email on email_delivery_attempts(email_id);
create index idx_email_events_email on email_events(email_id);
create index idx_email_events_type on email_events(event_type);
create index idx_templates_api_key on email_templates(api_key_id);
create index idx_transactions_sub on payment_transactions(subscription_id);
create index idx_transactions_external on payment_transactions(external_id) where external_id is not null;
create index idx_transactions_status on payment_transactions(status);

-- Functions & Triggers
-- auto-update updated_at timestamp
create or replace function update_updated_at_column()
returns trigger as $$
begin
    new.updated_at = now();
    return new;
end;
$$ language 'plpgsql';

create trigger update_api_keys_updated_at
    before update on api_keys
    for each row
    execute function update_updated_at_column();

create trigger update_email_templates_updated_at
    before update on email_templates
    for each row
    execute function update_updated_at_column();

create trigger update_smtp_configs_updated_at
    before update on smtp_configs
    for each row
    execute function update_updated_at_column();

create trigger update_subscriptions_updated_at
    before update on subscriptions
    for each row
    execute function update_updated_at_column();
