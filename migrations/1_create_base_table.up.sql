create table super_send_conn_info
(
    id                     integer           not null
        constraint super_send_conn_info_pk
            primary key autoincrement,
    address                TEXT              not null,
    is_ssl                 integer(1)        not null,
    username               text              not null,
    password               text,
    online                 integer(1),
    token                  text,
    refresh_token_interval integer default 0 not null,
    conn_last_login_time   integer default 0 not null
);