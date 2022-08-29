-- sku table
DROP TABLE IF EXISTS `t_sku`;

CREATE TABLE `t_sku` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sku_id` char(14) NOT NULL DEFAULT '' COMMENT '商品sku',
  `name` varchar(80) NOT NULL DEFAULT '' COMMENT '名称',
  `storage` int NOT NULL DEFAULT 0 COMMENT '库存',
  `price` int NOT NULL DEFAULT 0 COMMENT '标价',
  `purchase_price` int NOT NULL DEFAULT 0 COMMENT '进价',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`sku`),
  KEY (`price`, `storage`),
  KEY (`purchase_price`, `storage`),
  KEY (`created_at`, `storage`),
  KEY (`updated_at`, `storage`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = 'sku表';

-- order table--
DROP TABLE IF EXISTS `t_order`;
CREATE TABLE `t_order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` char(14) NOT NULL DEFAULT '' COMMENT '订单id  ord_10位',
  `user_id` char(12) NOT NULL DEFAULT '' COMMENT '用户id',
  `sku_id` char(14) NOT NULL DEFAULT '' COMMENT '商品sku', 
  `details` text NOT NULL COMMENT '订单详情',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态 0-未付款 1-已付款 2-超时 3-退款中 4-已退款',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_id` (`order_id`),
  KEY `created_at` (`created_at`),
  KEY `updated_at` (`updated_at`),
  KEY `idx_user_id_state` (`user_id`,`state`),
  KEY `idx_sku_user_id_state` (`sku_id`,`user_id`,`state`)
  KEY `idx_state_created_at_sku_id` (`state`,`created_at`,`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单表'