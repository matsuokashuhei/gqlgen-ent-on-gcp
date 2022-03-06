CREATE TABLE IF NOT EXISTS classes (
    `id` bigint(20) AUTO_INCREMENT NOT NULL,
    `name` varchar(255) NOT NULL,
    `tuition` bigint(20) NOT NULL,
    `start_date` date NOT NULL,
    `end_date` date,
    `schedule_id` bigint(20) NOT NULL,
    `instructor_id` bigint(20) NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    CONSTRAINT `fk_schedule_id_in_schedules` FOREIGN KEY (`schedule_id`) REFERENCES `schedules`(`id`),
    CONSTRAINT `fk_instructor_id_in_instructors` FOREIGN KEY (`instructor_id`) REFERENCES `instructors`(`id`)
);
