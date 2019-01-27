# int(10): 这个10表示的是数据显示的长度为10位。
# int(M) [undesigned] [zerofill]，加上zerofill后则会对于不满足指定的显示位的数据会在其前面加上0

# mysql> create table t (t int(3) zerofill);
# Query OK, 0 rows affected (0.00 sec)

# mysql> insert into t set t = 10;
# Query OK, 1 row affected (0.00 sec)

# mysql> select * from t;
# +——+
# | t |
# +——+
# | 010 |
# +——+
# 1 row in set (0.11 sec)

CREATE TABLE `users` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uid` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户uid',
  `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '用户名称',
  `phone` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '用户电话',
  `create_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY `uk_uid` (`uid`) USING BTREE COMMENT '唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
