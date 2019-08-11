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

 Date: 14/05/2019 23:13:59
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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='etched server列表';

-- ----------------------------
-- Records of etcd_servers
-- ----------------------------
BEGIN;
INSERT INTO `etcd_servers` VALUES (1, 'v3', 'local', '127.0.0.1:2379', 'false', '', '', '', '', '', '本机测试');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
