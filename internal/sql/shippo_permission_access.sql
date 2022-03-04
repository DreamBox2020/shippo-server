CREATE TABLE `shippo_permission_access` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`access_rule` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`access_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;