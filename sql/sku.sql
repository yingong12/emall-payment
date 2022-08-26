DROP TABLE IF EXISTS `t_sku`;
CREATE TABLE `t_sku` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sku` char(10) NOT NULL DEFAULT '' COMMENT '商品sku',
  `name` varchar(80) NOT NULL DEFAULT ''  COMMENT '名称',
  `storage`  int  NOT NULL DEFAULT 0 COMMENT '库存',
  `price`    int NOT NULL DEFAULT 0 COMMENT '标价',
  `purchase_price`  int NOT NULL DEFAULT 0 COMMENT '进价',
  `created_at`  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`sku`),
  KEY (`price`,`storage`),
  KEY (`purchase_price`,`storage`),
  KEY (`created_at`,`storage`),
  KEY (`updated_at`,`storage`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8  COMMENT='sku表';