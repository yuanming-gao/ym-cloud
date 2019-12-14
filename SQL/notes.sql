CREATE TABLE `notes_table` (
	`id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '笔记id',
	`user_id` INT(10) NOT NULL COMMENT '作者id',
	`user_name` VARCHAR(100) NOT NULL COMMENT '作者名字',
	`title` VARCHAR(250) NOT NULL COMMENT '题目',
	`content` TEXT NOT NULL COMMENT '用户填写的位置',
	`tags` VARCHAR(250) NOT NULL DEFAULT '尚未添加标签' COMMENT '标签',
	`create_time` BIGINT(10) NOT NULL COMMENT '注册日期',
	`update_time` BIGINT(10) NOT NULL COMMENT '更新日期',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
