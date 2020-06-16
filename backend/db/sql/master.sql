SET foreign_key_checks=0;

DROP TABLE IF EXISTS `articles`;

CREATE TABLE `articles` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(100) NOT NULL,
    `body` LONGTEXT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    INDEX `created_at_idx` (`created_at`),
    INDEX `updated_at_idx` (`updated_at`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4;

SET foreign_key_checks=1;