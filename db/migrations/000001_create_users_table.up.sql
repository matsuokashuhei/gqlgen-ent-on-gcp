CREATE TABLE IF NOT EXISTS users (
    `id` bigint(20) AUTO_INCREMENT NOT NULL,
    `name` varchar(255) NOT NULL,
    `firebase_uid` varchar(255) NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);
