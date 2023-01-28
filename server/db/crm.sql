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

 Date: 28/01/2023 19:19:16
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
) ENGINE=InnoDB AUTO_INCREMENT=20026 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of contract
-- ----------------------------
BEGIN;
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20001, '电动车交易1', 89880.00, '2023-01-28', '2023-01-30', '无备注', 1, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20002, '电动车交易2', 89880.00, '2023-01-28', '2023-01-30', '无备注', 2, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20003, '电动车交易3', 89880.00, '2023-01-28', '2023-01-30', '无备注', 3, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20004, '电动车交易4', 89880.00, '2023-01-28', '2023-01-30', '无备注', 4, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 1674901109);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20005, '电动车交易5', 89880.00, '2023-01-28', '2023-01-30', '无备注', 5, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20006, '电动车交易6', 89880.00, '2023-01-28', '2023-01-30', '无备注', 6, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20007, '电动车交易7', 89880.00, '2023-01-28', '2023-01-30', '无备注', 7, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20008, '电动车交易8', 89880.00, '2023-01-28', '2023-01-30', '无备注', 8, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20009, '电动车交易9', 89880.00, '2023-01-28', '2023-01-30', '无备注', 9, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 1674901130);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20010, '电动车交易10', 89880.00, '2023-01-28', '2023-01-30', '无备注', 10, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20011, '电动车交易11', 89880.00, '2023-01-28', '2023-01-30', '无备注', 11, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20012, '电动车交易12', 89880.00, '2023-01-28', '2023-01-30', '无备注', 12, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20013, '电动车交易13', 89880.00, '2023-01-28', '2023-01-30', '无备注', 13, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20014, '电动车交易14', 89880.00, '2023-01-28', '2023-01-30', '无备注', 14, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20015, '电动车交易15', 89880.00, '2023-01-28', '2023-01-30', '无备注', 15, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20016, '电动车交易16', 89880.00, '2023-01-28', '2023-01-30', '无备注', 16, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20017, '电动车交易17', 89880.00, '2023-01-28', '2023-01-30', '无备注', 17, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20018, '电动车交易18', 89880.00, '2023-01-28', '2023-01-30', '无备注', 18, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20019, '电动车交易19', 89880.00, '2023-01-28', '2023-01-30', '无备注', 19, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20020, '电动车交易20', 89880.00, '2023-01-28', '2023-01-30', '无备注', 20, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20021, '电动车交易21', 89880.00, '2023-01-28', '2023-01-30', '无备注', 21, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20022, '电动车交易22', 89880.00, '2023-01-28', '2023-01-30', '无备注', 22, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20023, '电动车交易23', 89880.00, '2023-01-28', '2023-01-30', '无备注', 23, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 2, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20024, '电动车交易24', 89880.00, '2023-01-28', '2023-01-30', '无备注', 24, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20025, '电动车交易25', 89880.00, '2023-01-28', '2023-01-30', '无备注', 25, '[{\"id\": 1, \"name\": \"电动车1\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 2, \"name\": \"电动车2\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}, {\"id\": 3, \"name\": \"电动车3\", \"type\": 1, \"unit\": \"台\", \"count\": 30, \"price\": 1498, \"total\": 44940}]', 1, 29, 1674900672, 0);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of customer
-- ----------------------------
BEGIN;
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (1, '北京文化有限公司1', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (2, '北京文化有限公司2', '线上注册', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 1674899808);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (3, '北京文化有限公司3', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (4, '北京文化有限公司4', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '互联网', '普通客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 1674900143);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (5, '北京文化有限公司5', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (6, '北京文化有限公司6', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 2, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (7, '北京文化有限公司7', '搜索引擎', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 2, 29, 1674899545, 1674900036);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (8, '北京文化有限公司8', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '互联网', '非优先客户', '', '北京市,朝阳区', '望京东园615号', 2, 29, 1674899545, 1674899824);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (9, '北京文化有限公司9', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (10, '北京文化有限公司10', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '文化传媒', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 1674899838);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (11, '北京文化有限公司11', '促销', '13050803360', 'bjwenhua@gmail.com', '互联网', '非优先客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 1674900100);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (12, '北京文化有限公司12', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 2, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (13, '北京文化有限公司13', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (14, '北京文化有限公司14', '邮件咨询', '13050803360', 'bjwenhua@gmail.com', '文化传媒', '非优先客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 1674899962);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (15, '北京文化有限公司15', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 2, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (16, '北京文化有限公司16', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (17, '北京文化有限公司17', '转介绍', '13050803360', 'bjwenhua@gmail.com', '文化传媒', '非优先客户', '', '北京市,朝阳区', '望京东园615号', 2, 29, 1674899545, 1674900109);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (18, '北京文化有限公司18', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 2, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (19, '北京文化有限公司19', '线上注册', '13050803360', 'bjwenhua@gmail.com', '政府', '普通客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 1674900124);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (20, '北京文化有限公司20', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 2, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (21, '北京文化有限公司21', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '物流运输', '普通客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 1674899902);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (22, '北京文化有限公司22', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (23, '北京文化有限公司23', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 2, 29, 1674899545, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (24, '北京文化有限公司24', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '房地产', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 1674899914);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (25, '北京文化有限公司25', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 0);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of mail_config
-- ----------------------------
BEGIN;
INSERT INTO `mail_config` (`id`, `stmp`, `port`, `auth_code`, `email`, `status`, `creator`, `created`, `updated`) VALUES (11, 'smtp.qq.com', 465, 'zrzxsebacrpfdaeg', '200300666@qq.com', 2, 29, 1674901189, 1674901237);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=154 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of notice
-- ----------------------------
BEGIN;
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (133, '你订阅了专业版', 2, 0, 1674803772, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (142, '你登录了账号', 2, 29, 1674819114, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (143, '你登录了账号', 2, 29, 1674821501, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (144, '你登录了账号', 2, 29, 1674821522, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (145, '你登录了账号', 2, 29, 1674821570, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (146, '你登录了账号', 2, 29, 1674821584, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (147, '你登录了账号', 2, 29, 1674821635, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (148, '你登录了账号', 2, 29, 1674821669, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (149, '你登录了账号', 2, 29, 1674872716, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (150, '你登录了账号', 2, 29, 1674872767, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (151, '你登录了账号', 2, 29, 1674874711, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (152, '你登录了账号', 2, 29, 1674899094, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (153, '你登录了账号', 2, 29, 1674904112, 0);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (1, '电动车1', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km', 1, 29, 1671191995, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (2, '电动车2', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km', 1, 29, 1671191995, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (3, '电动车3', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km', 2, 29, 1671191995, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (4, '电动车4', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km', 2, 29, 1671191995, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (5, '电动车5', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km', 1, 29, 1671191995, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (6, '电动车6', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km', 1, 29, 1671191995, 0);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of subscribe
-- ----------------------------
BEGIN;
INSERT INTO `subscribe` (`id`, `uid`, `version`, `expired`, `created`, `updated`) VALUES (4, 29, 2, 1701895772, 1674803772, 0);
COMMIT;

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

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `email`, `password`, `name`, `status`, `created`, `updated`) VALUES (29, '1655064994@qq.com', '$2a$10$62yO.fxSfNlstacxZfTtdO2uuR9YKG6hykuVTBIMc06CEJ3BWW/Ny', '', 1, 1671191625, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
