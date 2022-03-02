CREATE TABLE `shippo_role` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`role_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`auth_group` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`auth_api` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`auth_page` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;