CREATE TABLE IF NOT EXISTS instructors (
    `id` bigint(20) AUTO_INCREMENT NOT NULL,
    `name` varchar(255) NOT NULL,
    `syllabic_characters` varchar(255),
    `biography` text,
    `phone_number` varchar(32),
    `email` varchar(255),
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);
