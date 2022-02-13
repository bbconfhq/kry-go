package models

import (
	"gorm.io/gorm"
	"time"
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
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(191);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ProblemID uint
}
