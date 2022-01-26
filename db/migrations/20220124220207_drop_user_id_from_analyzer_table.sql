-- migrate:up
alter table analyzer drop column user_id;

-- migrate:down
alter table analyzer add column user_id int references "user";
