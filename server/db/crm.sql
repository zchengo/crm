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

 Date: 02/01/2023 10:55:52
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
-- Records of contract
-- ----------------------------
BEGIN;
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (64, '家居用品买卖合同', 22975.00, '2022-12-16', '2022-12-16', '', 64, '[{\"id\": 98, \"name\": \"电动车\", \"type\": 1, \"unit\": \"台\", \"count\": 10, \"price\": 1498, \"total\": 14980}, {\"id\": 99, \"name\": \"跑步机\", \"type\": 1, \"unit\": \"台\", \"count\": 5, \"price\": 1599, \"total\": 7995}]', 2, 29, 1671192364, 0);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (65, '车辆交易买卖合同', 29960.00, '2022-12-16', '2022-12-16', '无备注', 63, '[{\"id\": 98, \"name\": \"电动车\", \"type\": 1, \"unit\": \"台\", \"count\": 20, \"price\": 1498, \"total\": 29960}]', 1, 29, 1671192433, 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of customer
-- ----------------------------
BEGIN;
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (63, '河南中通有限公司', '搜索引擎', '13302308800', '1220080@163.com', '互联网', '普通客户', '宣传页分发附着交通工具宣传', '河南省,郑州市,金水区', '绿地之窗尚峰座', 1, 29, 1671192173, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (64, '可望企业股份有限公司', '电话咨询', '15133905680', '28033056@qq.com', '金融业', '重点客户', '现有客户引荐', '北京市,朝阳区', '黄厂路豆各庄2号', 1, 29, 1671192269, 0);
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of notice
-- ----------------------------
BEGIN;
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (37, '你登录了账号', 1, 29, 1672627977, 1672628013);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (38, '你登录了账号', 1, 29, 1672628053, 1672628062);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (39, '你登录了账号', 2, 29, 1672628056, 0);
INSERT INTO `notice` (`id`, `content`, `status`, `creator`, `created`, `updated`) VALUES (40, '你登录了账号', 2, 29, 1672628059, 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (98, '电动车', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km', 1, 29, 1671191995, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (99, '跑步机', 1, '台', '003', 1599.00, '跑步机家用智能可折叠免安装室内健身多功能走步机64CM宽大跑台', 2, 29, 1671192053, 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of subscribe
-- ----------------------------
BEGIN;
INSERT INTO `subscribe` (`id`, `uid`, `version`, `expired`, `created`, `updated`) VALUES (3, 29, 1, 0, 1671191625, 1671191817);
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
