-- tabfile 文件表
CREATE TABLE `tabfile` (
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `file_sha1` char(40)      NOT NULL DEFAULT '' COMMENT '文件hash值',
    `file_name` varchar(256)  NOT NULL DEFAULT '' COMMENT '文件名',
    `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
    `file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储路径',
    `create_at` datetime               DEFAULT NOW() COMMENT '创建日期',
    `update_at` datetime               DEFAULT NOW() on update current_timestamp COMMENT '更新时间',
    `status`    int(11) NOT NULL DEFAULT '0' COMMENT '状态（可用/禁用/已删除）',
    `ext1`      int(11) DEFAULT '0' COMMENT '备用1',
    `ext2`      text COMMENT '备用2',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_file_hash` (`file_sha1`),
    KEY         `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件表';

-- tabuser 用户表
CREATE TABLE `tabuser` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
    `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT '用户密码',
    `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone` varchar(128) NOT NULL DEFAULT '' COMMENT '手机号码',
    `email_validated` tinyint(1) DEFAULT 0 COMMENT '邮箱是否验证',
    `phone_validated` tinyint(1) DEFAULT 0 COMMENT '手机号码是否验证',
    `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
    `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间',
    `profile` text COMMENT '用户属性',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT '账户状态（启用/禁用/锁定/标记删除等）',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_phone` (`phone`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- tabToken token表
CREATE TABLE `tabtoken` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
    `user_token` char(40) NOT NULL DEFAULT '' COMMENT 'token',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'token表';