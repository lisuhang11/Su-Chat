-- 用户基本信息
CREATE TABLE IF NOT EXISTS `users` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_type` tinyint DEFAULT '0' COMMENT '用户类型 0用户, 1机器人',
    `user_id` varchar(32) NOT NULL COMMENT '用户id',
    `nickname` varchar(50) NOT NULL COMMENT '昵称',
    `password` varchar(50) NOT NULL COMMENT '密码',
    `user_portrait` varchar(200) DEFAULT NULL COMMENT '用户头像',
    `phone` varchar(50) DEFAULT NULL COMMENT '手机号',
    `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
    `created_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_userid` (`user_id`),
    UNIQUE KEY `uniq_phone` (`phone`),
    UNIQUE KEY `uniq_email` (`email`),
    KEY `idx_userid` (`user_type`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT = '用户-基本信息表';

-- 用户扩展信息
CREATE TABLE IF NOT EXISTS `userexts` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id` varchar(32) DEFAULT NULL COMMENT '用户id',
    `item_key` varchar(50) DEFAULT NULL COMMENT '参数key',
    `item_value` varchar(2000) DEFAULT NULL COMMENT '参数value',
    `item_type` tinyint DEFAULT '0' COMMENT '参数类型',
    `updated_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_item_key` (`user_id`,`item_key`),
    KEY `idx_item_key` (`item_key`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT = '用户-补充信息表';



CREATE TABLE IF NOT EXISTS `conversations` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id` varchar(32) DEFAULT NULL COMMENT '用户id',
    `target_id` varchar(32) DEFAULT NULL COMMENT '接收者id',
    `sub_channel` varchar(32) DEFAULT '',
    `channel_type` tinyint DEFAULT '0' COMMENT '会话类型 1单聊, 2群聊，3聊天室，4系统，5群公告，6广播',
    `latest_msg_id` varchar(20) DEFAULT NULL COMMENT '最新消息id',
    `latest_msg` mediumblob COMMENT '最新消息体',
    `latest_unread_msg_index` int DEFAULT '0' COMMENT '最新未读index',
    `latest_read_msg_index` int DEFAULT '0' COMMENT '最新已读消息index',
    `latest_read_msg_id` varchar(20) DEFAULT NULL COMMENT '最新已读消息id',
    `latest_read_msg_time` bigint DEFAULT '0' COMMENT '最新已读时间',
    `sort_time` bigint DEFAULT '0' COMMENT 'sort time',
    `is_deleted` tinyint DEFAULT '0' COMMENT '是否删除 0 未删除，1已删除',
    `is_top` tinyint DEFAULT '0' COMMENT '是否置顶 0未指定，1置顶',
    `top_updated_time` bigint DEFAULT '0' COMMENT '置顶更新时间',
    `undisturb_type` tinyint DEFAULT '0' COMMENT '免打扰类型：0:取消免打扰；1:普通会话免打扰；',
    `sync_time` bigint DEFAULT '0' COMMENT '同步消息位点',
    `unread_tag` tinyint DEFAULT '0' COMMENT '未读tag',
    `conver_exts` mediumblob,
    `app_key` varchar(20) DEFAULT NULL COMMENT '应用key',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_app_key_user_id_target_id` (`app_key`,`user_id`,`target_id`,`sub_channel`,`channel_type`),
    KEY `idx_sync_time` (`app_key`,`user_id`,`sync_time`),
    KEY `idx_update_time` (`app_key`,`user_id`,`sort_time`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT = '会话';