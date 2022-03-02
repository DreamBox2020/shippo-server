CREATE TABLE `shippo_user` (
	`id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '手机号',
	`email` varchar(17) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '邮箱',
	`nickname` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '昵称',
	`avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '头像',
	`exp` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '经验',
	`coin` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '硬币',
	`role` int(2) UNSIGNED NULL DEFAULT 0 COMMENT '角色(0普通用户1超级管理员2管理员3审核4禁用)',
	PRIMARY KEY (`id`),
	Unique KEY `idx_user_phone`(`phone`) USING BTREE,
	KEY `idx_user_deleted_at`(`deleted_at`) USING BTREE,
	Unique KEY `idx_user_nickname`(`nickname`) USING BTREE,
	Unique KEY `idx_user_email`(`email`) USING BTREE
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;