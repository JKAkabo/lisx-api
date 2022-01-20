-- migrate:up
alter table "user" add column is_admin bool not null default false;

-- migrate:down
alter table "user" drop column is_admin;