CREATE TABLE `shippo_dynamic` (
	`id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`user_id` bigint(20) NOT NULL COMMENT '用户序号',
	`type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '动态类型',
	`content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '动态内容',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;