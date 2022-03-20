ALTER TABLE `instructors` DROP COLUMN `kana` AFTER `name`, ADD COLUMN `syllabic_characters` varchar(255) NOT NULL;
