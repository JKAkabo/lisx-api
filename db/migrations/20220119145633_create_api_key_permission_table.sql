-- migrate:up
create table api_key_permission (
    api_key_id int references "api_key",
    resource varchar(100) not null ,
    can_create bool not null ,
    can_read bool not null ,
    can_update bool not null ,
    can_delete bool not null,
    primary key (api_key_id, resource)
);

-- migrate:down
drop table api_key_permission;
