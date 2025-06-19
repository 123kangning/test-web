CREATE TABLE `book`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`        varchar(128) NOT NULL COMMENT '书籍名称',
    `author`       varchar(128) NOT NULL COMMENT '作者',
    `price`        int          NOT NULL DEFAULT '0' COMMENT '价格',
    `publish_date` datetime              DEFAULT NULL COMMENT '出版日期',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='书籍表'

-- 用户表
CREATE TABLE `user`
(
    `id`              bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`            varchar(64)  NOT NULL COMMENT '用户名',
    `nickname`        varchar(64)  NOT NULL COMMENT '昵称',
    `email`           varchar(128) NOT NULL COMMENT '邮箱',
    `password`        varchar(128) NOT NULL COMMENT '密码',
    `create_time`     datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `last_login_time` datetime              DEFAULT NULL COMMENT '最后登录时间',
    `status`          tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态：1-正常，0-禁用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表'

-- 用户与书籍关联表（多对多关系）
CREATE TABLE `user_book`
(
    `user_id`       bigint unsigned NOT NULL COMMENT '用户ID',
    `book_id`       bigint unsigned NOT NULL COMMENT '书籍ID',
    `relation_type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '关系类型：1-已购买，2-想读，3-已读完',
    `create_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`user_id`, `book_id`),
    KEY             `user_book_book_id_fk` (`book_id`),
    CONSTRAINT `user_book_book_id_fk` FOREIGN KEY (`book_id`) REFERENCES books (`id`),
    CONSTRAINT `user_book_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES users (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户书籍关联表'


CREATE TABLE IF NOT EXISTS `alarm_channels`
(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
    `project_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '项目ID',
    `name` varchar(255) NOT NULL,
    `groups` text NOT NULL COMMENT '渠道信息',
    `mod` int(10) NOT NULL DEFAULT 0 COMMENT '外部可见',
    `is_deleted` int(10) NOT NULL DEFAULT 0 COMMENT '是否被删除',
    `created_by` varchar(50) NOT NULL,
    `create_time` datetime DEFAULT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    primary key (`id`),
    UNIQUE KEY `uni_tid_pid_name` (`tenant_id`,`project_id`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `alarm_rules` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `module_id` bigint(20) NOT NULL COMMENT '模块id',
    `template_id` bigint(20) NOT NULL COMMENT '模板id',
    `tenant_id` bigint(20) NOT NULL COMMENT '租户ID',
    `project_id` bigint(20) NOT NULL COMMENT '项目ID',
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `tags` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户添加的标签',
    `sys_tags` json NOT NULL COMMENT '系统添加的标签',
    `config` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '规则',
    `level` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '级别',
    `state` varchar(25) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'going_online' COMMENT '状态',
    `runtime_config` json NOT NULL COMMENT '运行配置',
    `run_state` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `run_last_time` datetime DEFAULT NULL,
    `run_next_time` datetime DEFAULT NULL,
    `run_failure_reason` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `gen_from` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT 'openapi' COMMENT '创建来源，openapi生成的规则不可变',
    `type` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '类型',
    `is_deleted` int(10) NOT NULL DEFAULT '0' COMMENT '是否被删除',
    `created_by` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `alarm_rule_channels`(
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `rule_id` bigint(20) NOT NULL COMMENT '规则id',
    `channel_id` bigint(20) NOT NULL COMMENT '渠道id',
    `create_time` datetime DEFAULT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
