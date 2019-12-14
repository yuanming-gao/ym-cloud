CREATE TABLE `token_table` (
	`id` CHAR(30) NOT NULL COMMENT 'uuid',
	`user_id` INT(10) NOT NULL COMMENT '对应的用户id',
	`expiration_time` BIGINT(10) NOT NULL COMMENT '过期时间戳10位',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
