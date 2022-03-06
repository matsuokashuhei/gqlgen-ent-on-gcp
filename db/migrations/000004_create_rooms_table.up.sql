CREATE TABLE IF NOT EXISTS rooms (
    `id` bigint(20) AUTO_INCREMENT NOT NULL,
    `name` varchar(255) NOT NULL,
    `capacity` int NOT NULL,
    `studio_id` bigint(20) NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    CONSTRAINT `fk_studio_id_in_rooms` FOREIGN KEY (`studio_id`) REFERENCES `studios`(`id`)
);
