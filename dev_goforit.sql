/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 80012
Source Host           : localhost:3306
Source Database       : dev_goforit

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2020-10-16 08:45:07
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `go_logs`
-- ----------------------------
DROP TABLE IF EXISTS `go_logs`;
CREATE TABLE `go_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `log_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '日志id',
  `create_time` int(11) DEFAULT '0' COMMENT '创建时间',
  `user_id` int(11) DEFAULT '0' COMMENT '用户id',
  `type_id` int(11) DEFAULT '0' COMMENT '类型id',
  `log_text` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '日志内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of go_logs
-- ----------------------------

-- ----------------------------
-- Table structure for `go_log_type`
-- ----------------------------
DROP TABLE IF EXISTS `go_log_type`;
CREATE TABLE `go_log_type` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL COMMENT '日志名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of go_log_type
-- ----------------------------
INSERT INTO `go_log_type` VALUES ('1', '登录');
INSERT INTO `go_log_type` VALUES ('2', '退出登录');

-- ----------------------------
-- Table structure for `go_memos`
-- ----------------------------
DROP TABLE IF EXISTS `go_memos`;
CREATE TABLE `go_memos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `content` varchar(255) DEFAULT NULL COMMENT '内容',
  `create_time` int(11) DEFAULT '0' COMMENT '创建时间',
  `user_id` int(11) DEFAULT '0' COMMENT '用户id',
  `classify_id` int(11) DEFAULT '0' COMMENT '分类id',
  `update_time` int(11) DEFAULT '0' COMMENT '更新时间',
  `delete_time` int(11) DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of go_memos
-- ----------------------------

-- ----------------------------
-- Table structure for `go_memo_classify`
-- ----------------------------
DROP TABLE IF EXISTS `go_memo_classify`;
CREATE TABLE `go_memo_classify` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(255) DEFAULT NULL COMMENT '名称',
  `create_time` int(11) DEFAULT '0' COMMENT '10',
  `user_id` int(11) DEFAULT '0' COMMENT '用户id',
  `delete_time` int(11) DEFAULT '0' COMMENT '删除时间',
  `pid` int(11) DEFAULT '0' COMMENT '父级id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of go_memo_classify
-- ----------------------------

-- ----------------------------
-- Table structure for `go_tokens`
-- ----------------------------
DROP TABLE IF EXISTS `go_tokens`;
CREATE TABLE `go_tokens` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `token` varchar(255) DEFAULT NULL COMMENT 'token',
  `expire` int(11) DEFAULT '0' COMMENT '到期时间',
  `user_id` int(11) DEFAULT '0' COMMENT '用户ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of go_tokens
-- ----------------------------

-- ----------------------------
-- Table structure for `go_users`
-- ----------------------------
DROP TABLE IF EXISTS `go_users`;
CREATE TABLE `go_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(50) DEFAULT NULL COMMENT '用户名',
  `password` varchar(250) DEFAULT NULL COMMENT '密码',
  `create_time` int(11) DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of go_users
-- ----------------------------
INSERT INTO `go_users` VALUES ('1', 'user', '4297f44b13955235245b2497399d7a93', '0');
INSERT INTO `go_users` VALUES ('2', 'user1', '4297f44b13955235245b2497399d7a93', '0');
