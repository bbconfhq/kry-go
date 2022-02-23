package entity

import (
	"gorm.io/gorm"
)

/*
CREATE TABLE IF NOT EXISTS `user`
(
    `id`      BIGINT       NOT NULL AUTO_INCREMENT,
    `name`    VARCHAR(30)  NOT NULL,
    `pw`      VARCHAR(512) NOT NULL,
    `nick`    VARCHAR(30)  NOT NULL,
    `bio`     VARCHAR(512) NOT NULL,
    `email`   VARCHAR(191) NOT NULL,
    `created` DATETIME(6)  NOT NULL DEFAULT NOW(),
    PRIMARY KEY (`id`),
    UNIQUE (`email`),
    UNIQUE (`name`),
    UNIQUE (`nick`)
) ENGINE = InnoDB;
*/

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	AvatarUrl string `gorm:"unique;not null"`
	Bio       string `gorm:"unique;not null"`
	Email     string `gorm:"unique"`
}
