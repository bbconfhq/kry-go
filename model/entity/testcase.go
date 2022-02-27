package entity

import (
	"gorm.io/gorm"
)

/*
CREATE TABLE `testcase`
(
    `id`         BIGINT  NOT NULL AUTO_INCREMENT,
    `visible`    BOOLEAN NOT NULL,
    `input`      TEXT    NOT NULL,
    `output`     TEXT    NOT NULL,
    `problem_id` BIGINT  NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`problem_id`) REFERENCES `problem` (`id`)
) ENGINE = InnoDB;
*/

type Testcase struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Visible   bool   `gorm:"not null"`
	Input     string `gorm:"type:text;not null"`
	Output    string `gorm:"type:text;not null"`
	ProblemID uint
}
