/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3306
 Source Schema         : crm

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 24/01/2023 20:00:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for contract
-- ----------------------------
DROP TABLE IF EXISTS `contract`;
CREATE TABLE `contract` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(200) DEFAULT NULL COMMENT '合同名称',
  `amount` decimal(10,2) DEFAULT NULL COMMENT '金额',
  `begin_time` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '开始时间',
  `over_time` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '结束时间',
  `remarks` varchar(80) DEFAULT NULL COMMENT '备注',
  `cid` bigint DEFAULT NULL COMMENT '客户编号',
  `productlist` json DEFAULT NULL COMMENT '产品编号和数量',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `creator` bigint DEFAULT NULL COMMENT '创建人',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `updated` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `source` char(15) DEFAULT NULL COMMENT '来源',
  `phone` char(12) DEFAULT NULL COMMENT '手机号',
  `email` char(20) DEFAULT NULL COMMENT '邮箱',
  `industry` char(30) DEFAULT NULL COMMENT '所在行业',
  `level` char(10) DEFAULT NULL COMMENT '级别',
  `remarks` varchar(200) NOT NULL COMMENT '备注',
  `region` char(80) DEFAULT NULL COMMENT '地区',
  `address` varchar(255) DEFAULT NULL COMMENT '详细地址',
  `status` tinyint DEFAULT NULL COMMENT '成交状态',
  `creator` bigint DEFAULT NULL COMMENT '创建人',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `updated` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for mail_config
-- ----------------------------
DROP TABLE IF EXISTS `mail_config`;
CREATE TABLE `mail_config` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `stmp` char(50) DEFAULT NULL COMMENT 'ip地址或域名',
  `port` int DEFAULT NULL COMMENT '端口号',
  `auth_code` char(80) DEFAULT NULL COMMENT '授权码',
  `email` char(80) NOT NULL COMMENT '邮箱账号',
  `status` tinyint DEFAULT NULL COMMENT '服务状态',
  `creator` bigint DEFAULT NULL COMMENT '创建人',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `updated` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for notice
-- ----------------------------
DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `content` varchar(200) DEFAULT NULL COMMENT '通知内容',
  `status` tinyint DEFAULT NULL COMMENT '状态，1-已读，2-未读',
  `creator` bigint DEFAULT NULL COMMENT '创建者',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `updated` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `type` tinyint(1) DEFAULT NULL COMMENT '类型',
  `unit` char(5) DEFAULT NULL COMMENT '单位',
  `code` varchar(80) DEFAULT NULL COMMENT '编码',
  `price` decimal(10,2) DEFAULT NULL COMMENT '价格',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态，1-上架，2-下架',
  `creator` bigint DEFAULT NULL COMMENT '创建人',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `updated` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for subscribe
-- ----------------------------
DROP TABLE IF EXISTS `subscribe`;
CREATE TABLE `subscribe` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `uid` bigint DEFAULT NULL COMMENT '用户编号',
  `version` tinyint DEFAULT NULL COMMENT '版本，1-免费版，2-专业版，3-高级版',
  `expired` bigint DEFAULT NULL COMMENT '到期时间',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `updated` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `email` char(30) DEFAULT NULL COMMENT '邮箱',
  `password` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '密码',
  `name` char(10) DEFAULT NULL COMMENT '名称',
  `status` tinyint DEFAULT '1' COMMENT '状态，1-正常，2-注销',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `updated` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb3;

SET FOREIGN_KEY_CHECKS = 1;
