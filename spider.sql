CREATE TABLE IF NOT EXISTS `spider`.`houses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `town` VARCHAR(255) NOT NULL COMMENT '城市',
  `area` VARCHAR(255) NOT NULL COMMENT '城区',
  `room` VARCHAR(255) NOT NULL COMMENT '房型',
  `totalPrice` VARCHAR(255) NOT NULL COMMENT '价格',
  `price` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL COMMENT '小区',
  `options` VARCHAR(255) NOT NULL COMMENT '城区',
  `createTime` INT(64) NOT NULL,
  `updateTime` INT(64) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_bin
COMMENT = '房源信息';