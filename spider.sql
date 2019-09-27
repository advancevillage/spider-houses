-- MySQL Script generated by MySQL Workbench
-- Wed Sep 25 19:46:12 2019
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema spider
-- -----------------------------------------------------
-- 房价爬虫
DROP SCHEMA IF EXISTS `spider` ;

-- -----------------------------------------------------
-- Schema spider
--
-- 房价爬虫
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `spider` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin ;
USE `spider` ;

-- -----------------------------------------------------
-- Table `spider`.`houses`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `spider`.`houses` ;

CREATE TABLE IF NOT EXISTS `spider`.`houses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `city` VARCHAR(255) NOT NULL COMMENT '城市',
  `title` VARCHAR(255) NULL,
  `pattern` VARCHAR(255) NOT NULL COMMENT '格局',
  `address` VARCHAR(255) NOT NULL COMMENT '小区',
  `options` VARCHAR(255) NOT NULL COMMENT '城区',
  `totalPrice` VARCHAR(255) NOT NULL COMMENT '价格',
  `price` VARCHAR(255) NOT NULL,
  `createTime` INT(64) NOT NULL,
  `updateTime` INT(64) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_bin
COMMENT = '房源信息';


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
