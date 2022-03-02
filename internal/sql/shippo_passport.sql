CREATE TABLE `shippo_passport` (
	`id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` timestamp NULL,
	`token` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	`user_id` bigint(20) UNSIGNED NULL DEFAULT 0,
	`ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`ua` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
	`client` int(2) NULL DEFAULT 0,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;