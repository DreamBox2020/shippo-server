CREATE TABLE `shippo_auth` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`group_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`group_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`parent` bigint NULL,
	`type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`intro` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`router_rule` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`req_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;