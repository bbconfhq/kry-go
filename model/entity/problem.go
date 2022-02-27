package entity

import (
	"gorm.io/gorm"
)

/*
CREATE TABLE `problem`
(
    `id`           BIGINT        NOT NULL AUTO_INCREMENT,
    `title`        VARCHAR(191)  NOT NULL,
    `content`      TEXT          NOT NULL,
    `created`      DATETIME(6)   NOT NULL,
    `modified`     DATETIME(6)   NOT NULL,
    `input`        TEXT          NOT NULL,
    `output`       TEXT          NOT NULL,
    `note`         TEXT          NOT NULL,
    `time_limit`   DECIMAL(7, 5) NOT NULL,
    `memory_limit` INTEGER       NOT NULL,
    `submit_count` INTEGER       NOT NULL,
    `accept_count` INTEGER       NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`title`)
) ENGINE = InnoDB;

// TODO: Alter table for initial auto_increment value
ALTER TABLE `problem`
    AUTO_INCREMENT = 1000;
*/

type Problem struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	Title       string  `gorm:"type:varchar(191);unique;not null"`
	Content     string  `gorm:"type:text;not null"`
	Note        string  `gorm:"type:text;not null"`
	TimeLimit   float64 `gorm:"type:decimal(7,5);not null"`
	MemoryLimit uint    `gorm:"not null"`
	SubmitCount uint    `gorm:"default:0;not null"`
	AcceptCount uint    `gorm:"default:0;not null"`
	Testcases   []Testcase
	Tags        []Tag `gorm:"many2many:problem_tags;"`
}
