-- migrate:up
alter table analyzer add column format varchar(10) not null default '';
alter table analyzer add column format_spec text not null default '';

-- migrate:down
alter table analyzer drop column format;
alter table analyzer drop column format_spec;
