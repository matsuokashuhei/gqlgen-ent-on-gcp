CREATE TABLE IF NOT EXISTS schedules (
    `id` bigint(20) AUTO_INCREMENT NOT NULL,
    `day_of_week` int NOT NULL,
    `start_time` varchar(8) NOT NULL,
    `end_time` varchar(8) NOT NULL,
    `room_id` bigint(20) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    CONSTRAINT `fk_room_id_in_rooms` FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`)
);
