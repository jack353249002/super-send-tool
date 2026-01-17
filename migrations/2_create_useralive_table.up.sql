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
      "last_ping_time" integer,
      "send_email_accounts" TEXT,
      "message_id" INTEGER,
      "smtp_ids" TEXT
);