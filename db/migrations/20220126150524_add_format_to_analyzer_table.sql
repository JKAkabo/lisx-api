-- migrate:up
alter table analyzer add column format varchar(10);
alter table analyzer add column format_spec text;

-- migrate:down
alter table analyzer drop column format;
alter table analyzer drop column format_spec;
