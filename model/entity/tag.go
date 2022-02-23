package entity

import (
	"gorm.io/gorm"
)

/*
CREATE TABLE `tag`
(
    `id`   BIGINT       NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(191) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;
*/

type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(191);unique;not null"`
}
