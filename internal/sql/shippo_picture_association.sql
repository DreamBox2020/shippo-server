CREATE TABLE `shippo_picture_association` (
	`id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`picture_id` bigint(20) NOT NULL COMMENT '相片序号',
	`album_id` bigint(20) NOT NULL COMMENT '相簿序号',
	`dynamic_id` bigint(20) NOT NULL COMMENT '动态序号',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;