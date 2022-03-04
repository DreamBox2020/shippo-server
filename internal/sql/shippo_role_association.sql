CREATE TABLE `shippo_role_association` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`role_id` bigint NOT NULL COMMENT '角色序号',
	`policy_id` bigint NOT NULL COMMENT '权限策略序号',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;