create table super_send_conn_info
(
    id       integer    not null
        constraint super_send_conn_info_pk
            primary key autoincrement,
    address  TEXT       not null,
    is_ssl   integer(1) not null,
    username text       not null,
    password text
);

