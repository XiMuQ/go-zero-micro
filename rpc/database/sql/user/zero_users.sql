SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for zero_users
-- ----------------------------
DROP TABLE IF EXISTS `zero_users`;
CREATE TABLE `zero_users`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账号',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `gender` tinyint(1) NOT NULL DEFAULT 1 COMMENT '性别 1：未设置；2：男性；3：女性',
  `updated_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `updated_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `created_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `deleted_at` datetime(0) NOT NULL COMMENT '删除时间',
  `deleted_flag` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否删除 1：正常  2：已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `index_username`(`username`) USING BTREE COMMENT '用户名索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
