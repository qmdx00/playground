# 版本号机制业务实现 Demo

## 表结构定义

```SQL
CREATE TABLE `account` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户id',
  `balance` decimal(18,2) NOT NULL COMMENT '当前余额，精确到分',
  `version` int NOT NULL DEFAULT '0' COMMENT '版本号',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '账户状态（1：正常、2：冻结）',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_uid` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `account_flow` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `flow_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '流水号',
  `account_id` bigint NOT NULL COMMENT '关联的账户id',
  `amount` decimal(18,2) NOT NULL COMMENT '变动金额（正：进账，负：出账）',
  `balance_before` decimal(18,2) NOT NULL COMMENT '变动前余额',
  `balance_after` decimal(18,2) NOT NULL COMMENT '变动后余额',
  `type` tinyint NOT NULL COMMENT '流水类型（1：充值、2：消费、3：退款、4：提现）',
  `biz_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务单号',
  `version_seq` int NOT NULL COMMENT '关联账户的版本号（用于追溯）',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_fno` (`flow_no`),
  KEY `idx_aid_cat` (`account_id`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```