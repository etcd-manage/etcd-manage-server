/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : 127.0.0.1:3306
 Source Schema         : etcd-manage

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 15/08/2019 20:27:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for etcd_servers
-- ----------------------------
DROP TABLE IF EXISTS `etcd_servers`;
CREATE TABLE `etcd_servers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `version` varchar(3) NOT NULL DEFAULT 'v3' COMMENT 'etcd版本',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT 'etcd服务名字',
  `address` varchar(600) NOT NULL COMMENT 'etcd地址列表',
  `prefix` varchar(100) NOT NULL DEFAULT '' COMMENT 'key前缀，建议不为空，防止大量key',
  `tls_enable` varchar(5) NOT NULL DEFAULT 'true' COMMENT '是否启用tls连接',
  `cert_file` text NOT NULL COMMENT '证书',
  `key_file` text NOT NULL COMMENT '证书',
  `ca_file` text NOT NULL COMMENT '证书',
  `username` varchar(60) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(60) NOT NULL DEFAULT '' COMMENT '密码',
  `desc` varchar(300) NOT NULL COMMENT '描述信息',
  `created_at` datetime DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='etched server列表';

-- ----------------------------
-- Records of etcd_servers
-- ----------------------------
BEGIN;
INSERT INTO `etcd_servers` VALUES (1, 'v3', 'local', '127.0.0.1:2379', '', 'false', '', '', '', '', '', '本机测试', NULL, NULL);
INSERT INTO `etcd_servers` VALUES (3, 'v3', '测试', '127.0.0.1:2379', '', 'false', '', '', '', '', '', '备注一下', NULL, NULL);
INSERT INTO `etcd_servers` VALUES (4, 'v2', 'v2测试', '127.0.0.1:2379', '', 'false', '', '', '', '', '', 'v2备注', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for role_etcd_servers
-- ----------------------------
DROP TABLE IF EXISTS `role_etcd_servers`;
CREATE TABLE `role_etcd_servers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `etcd_server_id` int(11) NOT NULL DEFAULT '0' COMMENT 'etcd服务id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0读 1写',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_role_id` (`role_id`),
  KEY `idx_etcd_server_id` (`etcd_server_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='角色权限表';

-- ----------------------------
-- Records of role_etcd_servers
-- ----------------------------
BEGIN;
INSERT INTO `role_etcd_servers` VALUES (1, 1, 1, 1, '2019-08-14 19:43:23', '2019-08-14 11:48:09');
COMMIT;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '角色名',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO `roles` VALUES (1, '高级管理员', '2019-08-14 19:43:44', '2019-08-14 19:43:48');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(60) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(300) NOT NULL DEFAULT '' COMMENT '邮箱',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'admin', '94d9484ada1e5639810258b515fa61f1', '', 1, '2019-08-12 20:19:15', '2019-08-14 11:43:56');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
