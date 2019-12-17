--# 创建 frienddata 表
CREATE TABLE `frienddata` (
   `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
   `code` text NOT NULL COMMENT '邀请码',
   `fail` int NOT NULL COMMENT '失败次数',
   `creaator` text NOT NULL COMMENT '创建者',
   `is_delete` tinyint(1) NOT NULL COMMENT '删除否',
   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   PRIMARY KEY (`id`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8
