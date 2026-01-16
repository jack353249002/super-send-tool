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

CREATE TABLE check_user_alive (
     "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
     "username" varchar(50) NOT NULL,
     "password" text NOT NULL,
     "day_login_first_time" integer NOT NULL,
     "send_id" integer,
     "send_email_action_timeout" INT,
     "salt" text NOT NULL,
     "last_send_email_time" integer NOT NULL DEFAULT 0,
     "super_send_conn_info_id" integer,
     "position" TEXT,
     "last_ping_time" integer
);