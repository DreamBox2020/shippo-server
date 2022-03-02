CREATE TABLE `shippo_picture` (
	`id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '相片物理路径',
	`uri` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '相片虚拟路径',
	`name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '相片名称',
	`mime` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '相片格式',
	`type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '相片类别',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;