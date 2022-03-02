CREATE TABLE `shippo_temp_trade_20220108` (
	`id` int(5) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`trade_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
	`trade_type` int(2) NOT NULL COMMENT '订单类型(0淘宝1支付宝2微信)',
	`trade_amount` int(11) NOT NULL COMMENT '订单金额',
	`amount_status` int(2) NOT NULL COMMENT '订单状态(0正常1没有到账2到账后返还)',
	`user_qq` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '订单归属用户的QQ',
	`user_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '订单归属用户的手机号',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci
COMMENT='临时订单表';