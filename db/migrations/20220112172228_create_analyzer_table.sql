-- migrate:up
create table analyzer
(
    id              serial primary key,
    name            varchar(50) not null,
    protocol        varchar(10) not null,
    ip              varchar(15),
    port            int,
    server_mode     bool        not null,
    serial_port     varchar(50),
    baud_rate       int,
    parity          int,
    data_bits       int,
    stop_bits       int,
    start_delimiter varchar(50),
    end_delimiter   varchar(50),
    user_id         int references "user"
);

-- migrate:down
drop table analyzer;
