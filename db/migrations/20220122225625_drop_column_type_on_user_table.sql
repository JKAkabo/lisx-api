-- migrate:up
alter table "user" drop column type;

-- migrate:down
alter table "user" add column type varchar(10) default 'USER';
