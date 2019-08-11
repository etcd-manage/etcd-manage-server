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

 Date: 11/08/2019 21:10:37
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='etched server列表';

-- ----------------------------
-- Records of etcd_servers
-- ----------------------------
BEGIN;
INSERT INTO `etcd_servers` VALUES (1, 'v3', 'local', '127.0.0.1:2379', '', 'false', '', '', '', '', '', '本机测试');
INSERT INTO `etcd_servers` VALUES (3, 'v3', '测试', '127.0.0.1:2379', '', 'false', '', '', '', '', '', '备注一下');
INSERT INTO `etcd_servers` VALUES (4, 'v2', 'v2测试', '127.0.0.1:2379', '', 'false', '', '', '', '', '', 'v2备注');
COMMIT;

-- ----------------------------
-- Table structure for role_etcd_servers
-- ----------------------------
DROP TABLE IF EXISTS `role_etcd_servers`;
CREATE TABLE `role_etcd_servers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `etcd_server_id` int(11) NOT NULL DEFAULT '0' COMMENT 'etcd服务id',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0读 1写 2读写',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限表';

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

SET FOREIGN_KEY_CHECKS = 1;
