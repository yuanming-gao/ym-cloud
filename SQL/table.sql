CREATE TABLE `user_table` (
	`id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '用户id',
	`user_phone` CHAR(11) NOT NULL COMMENT '用户手机号',
	`user_name` CHAR(20) NOT NULL DEFAULT '未添加用户昵称' COMMENT '用户昵称',
	`user_password` VARCHAR(50) NOT NULL COMMENT '密码',
	`user_position` VARCHAR(50) NOT NULL COMMENT '用户填写的位置',
	`create_time` BIGINT(10) NOT NULL COMMENT '注册日期',
	`update_time` BIGINT(10) NOT NULL COMMENT '更新日期',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_login_record_table` (
	`id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '每条记录的id',
	`login_phone` CHAR(11) NOT NULL COMMENT '登录的手机号',
	`login_ip` VARCHAR(15) NOT NULL COMMENT '登录的ip地址',
	`login_time` BIGINT(10) NOT NULL COMMENT '登录时间',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `file_table` (
	`id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '文件id',
	`file_name` VARCHAR(15) NOT NULL COMMENT '文件名',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


