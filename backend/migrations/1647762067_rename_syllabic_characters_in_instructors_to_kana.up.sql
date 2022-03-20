ALTER TABLE `instructors` DROP COLUMN `syllabic_characters`, ADD COLUMN `kana` varchar(255) NOT NULL AFTER `name`;
