CREATE TABLE IF NOT EXISTS studios (
    `id` bigint(20) AUTO_INCREMENT NOT NULL,
    `school_id` bigint(20) NOT NULL,
    `name` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    CONSTRAINT `fk_school_id_in_studios` FOREIGN KEY (`school_id`) REFERENCES `schools`(`id`)
);
