CREATE TABLE IF NOT EXISTS rooms (
    `id` bigint(20) AUTO_INCREMENT NOT NULL,
    `studio_id` bigint(20) NOT NULL,
    `name` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    CONSTRAINT `fk_studio_id_in_rooms` FOREIGN KEY (`studio_id`) REFERENCES `studios`(`id`)
);
