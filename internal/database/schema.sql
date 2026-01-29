-- at the time of creating this sql schema, I was still learning postgres, maybe something would be wrong

create table api_keys {
    id uuid primary key default gen_random_uuid(),
    appName varchar(100) not null,
    keyHash varchar(255) not null,
    keyPrefix varchar(8) not null,
    isActive boolean default true,

    rateLimit int default 100,

    plainID uuid references plans(id),

    quotaLimit int default 8000,
    quotaUsed int default 0,
    quotaType varchar(20) default "free"
    quotaResetAt timestamptz,

    createdAt timestamptz default now(),
    updatedAt timestamptz default now(),
    deletedAt timestamptz
}

-- note: IMO, we should use timestamptz because we have to sync the time of sending mail to customers all at the same time
-- if the service need a specific time zone, I think we will redesign later

create table plans {
    id uuid primary key default gen_random_uuid(),
    name varchar(50) not null, -- free, premium and business
    maxQuota int not null,
    rateLimit int not null,
    price decimal(10,2)
    createdAt timestamptz default now()
}

create table api_key_usage {
    id uuid primary key default gen_random_uuid(),
    apiKeyID references api_keys(id) on delete cascade,
    -- on delete cascade: whenever a record in table B is deleted, any corresponding records in table A will also be deleted automatically
    -- I just want to make sure I will understand it in the future when reread it
    usageDate date not null,
    sentCount int default 0,
    failedAt int default 0,

    unique(apiKeyID, usageDate) -- for each key, only 1 row per day
}


create table emails {
    id uuid primary key default gen_random_uuid(),
    apiKeyID uuid references api_keys(id),
    idempotencyKey varchar(255) unique, -- though success or fail, it still counts as a single request

    fromMail varchar(255) not null,
    fromName varchar(100),
    toMail varchar(255) not null,
    toName varchar(100),,
    subject varchar(500) not null,

    templateID uuid references emails_templates(id),
    templateVars jsonb, -- for filtering usage by tags
    bodyHTML text,

    status varchar(200) default 'pending', -- pending, queued, sent, failed, cancelled

    createdAt timestamptz default now(),
    scheduledAt timestamptz default now(),
    sentAt timestamptz
}

create table emails_templates {
    id uuid primary key default gen_random_uuid(),
    emailID uuid references emails(id) on delete cascade,

    attemptNum int not null,
    status varchar(20) not null,
    reponseCode varchar(50), -- from smtp (e.g. 550, 421)
    errMsg text,

    providerUser varchar(50), -- extended feature, soon

    createdAt timestamptz default now()
}

-- please take note that these are logs for client side, server will not store it
create table emails_logs {
    id uuid primary key default gen_random_uuid(),
    emailLogID uuid references emails(id),
    status varchar(20) not null,
    msg text,
    metadata jsonb,
    createdAt timestamptz default now()
}

create table smpt_configs {
    id uuid primary key default gen_random_uuid(),
    name            varchar(100) not null,
    host            varchar(255) not null,
    port            int not null,
    username        varchar(255),
    passwordEnc    TEXT,  -- encrypted password
    useTls         boolean default true,
    isDefault      boolean default false,
    createdAt      timestamptz default now()
}

create index idx_usage_lookup on api_key_usage(apiKeyID, usageDate)
create index idx_emails_status_created on emails(status, createdAt)

create index idx_emails_status ON emails(status)
create index idx_emails_api_key ON emails(api_key_id)
create index idx_emails_idempotency ON emails(idempotency_key)
create index idx_email_logs_email ON email_logs(email_id)
