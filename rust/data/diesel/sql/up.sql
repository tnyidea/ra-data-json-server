CREATE TABLE addresses
(
    id           varchar default gen_random_uuid() not null primary key,
--     created_at   timestamp with time zone,
--     updated_at   timestamp with time zone,
--     deleted_at   timestamp with time zone,
    first_name   varchar not null,
    last_name    varchar not null,
    company_name varchar not null,
    address      varchar not null,
    city         varchar not null,
    county       varchar not null,
    state        varchar not null,
    zip          varchar not null,
    phone1       varchar not null,
    phone2       varchar not null,
    email        varchar not null,
    web          varchar not null
)