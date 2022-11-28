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

 Date: 28/11/2022 15:18:33
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of contract
-- ----------------------------
BEGIN;
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (14, '家居用品买卖合同', 8170.00, '2022-11-27', '2022-11-27', '客户对价格很满意', 21, '[{\"id\": 51, \"name\": \"洗发露\", \"type\": 1, \"unit\": \"瓶\", \"count\": 20, \"price\": 200, \"total\": 4000}, {\"id\": 52, \"name\": \"运动鞋\", \"type\": 1, \"unit\": \"只\", \"count\": 30, \"price\": 139, \"total\": 4170}]', 1, 23, 1669550577, 1669553025);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (15, '家居用品买卖合同', 8170.00, '2022-11-27', '2022-11-27', '客户对价格很满意', 21, '[{\"id\": 51, \"name\": \"洗发露\", \"type\": 1, \"unit\": \"瓶\", \"count\": 20, \"price\": 200, \"total\": 4000}, {\"id\": 52, \"name\": \"运动鞋\", \"type\": 1, \"unit\": \"只\", \"count\": 30, \"price\": 139, \"total\": 4170}]', 1, 23, 1669550577, 1669553025);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (16, '家居用品买卖合同', 8170.00, '2022-11-27', '2022-11-27', '客户对价格很满意', 21, '[{\"id\": 51, \"name\": \"洗发露\", \"type\": 1, \"unit\": \"瓶\", \"count\": 20, \"price\": 200, \"total\": 4000}, {\"id\": 52, \"name\": \"运动鞋\", \"type\": 1, \"unit\": \"只\", \"count\": 30, \"price\": 139, \"total\": 4170}]', 1, 23, 1669550577, 1669553025);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (18, '家居用品买卖合同', 8170.00, '2022-11-27', '2022-11-27', '客户对价格很满意', 21, '[{\"id\": 51, \"name\": \"洗发露\", \"type\": 1, \"unit\": \"瓶\", \"count\": 20, \"price\": 200, \"total\": 4000}, {\"id\": 52, \"name\": \"运动鞋\", \"type\": 1, \"unit\": \"只\", \"count\": 30, \"price\": 139, \"total\": 4170}]', 1, 23, 1669550577, 1669553025);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (19, '家居用品买卖合同', 8170.00, '2022-11-27', '2022-11-27', '客户对价格很满意', 21, '[{\"id\": 51, \"name\": \"洗发露\", \"type\": 1, \"unit\": \"瓶\", \"count\": 20, \"price\": 200, \"total\": 4000}, {\"id\": 52, \"name\": \"运动鞋\", \"type\": 1, \"unit\": \"只\", \"count\": 30, \"price\": 139, \"total\": 4170}]', 1, 23, 1669550577, 1669553025);
INSERT INTO `contract` (`id`, `name`, `amount`, `begin_time`, `over_time`, `remarks`, `cid`, `productlist`, `status`, `creator`, `created`, `updated`) VALUES (20, '家居用品买卖合同', 8170.00, '2022-11-27', '2022-11-27', '客户对价格很满意', 21, '[{\"id\": 51, \"name\": \"洗发露\", \"type\": 1, \"unit\": \"瓶\", \"count\": 20, \"price\": 200, \"total\": 4000}, {\"id\": 52, \"name\": \"运动鞋\", \"type\": 1, \"unit\": \"只\", \"count\": 30, \"price\": 139, \"total\": 4170}]', 1, 23, 1669550577, 1669553025);
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of customer
-- ----------------------------
BEGIN;
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (17, '河南正商泰盛裕网电有限公司 ', '电话咨询', '13308322331', '12637237@qq.com', '金融业', '普通客户', '无备注', '河南省,郑州市,上街区', '淮阳路南段23号院', 2, 19, 1669181566, 1669181797);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (18, '海航科技股份有限公司 ', '线上注册', '18905850015', '1737283@163.com', '房地产', '重点客户', '无备注', '北京市,朝阳区', '黄厂路豆各庄3号', 1, 19, 1669182133, 0);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (19, '可望企业股份有限公司', '搜索引擎', '12043480098', '9843834@gmail.com', '物流运输', '普通客户', '无备注', '河南省,郑州市,金水区', '绿地之窗尚峰座', 1, 19, 1669182378, 1669182392);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (20, '广北嵊州天燃气集团有限公司', '电话咨询', '13500205580', '123456@163.com', '金融业', '重点客户', '交流大会上收集的名片', '广东省,茂名市,电白区', '坡心镇谭莲工业区市民大道', 1, 23, 1669547870, 1669548361);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (21, '河南郑州石油集团有限公司', '搜索引擎', '12300561200', 'shiyou@qq.com', '文化传媒', '非优先客户', '', '河南省,郑州市,金水区', '绿地之窗尚峰座', 1, 23, 1669548325, 1669548341);
INSERT INTO `customer` (`id`, `name`, `source`, `phone`, `email`, `industry`, `level`, `remarks`, `region`, `address`, `status`, `creator`, `created`, `updated`) VALUES (22, '可望企业股份有限公司', '线上注册', '15890800054', '200300@163.com', '物流运输', '普通客户', '现有客户引荐', '北京市,朝阳区', '黄厂路豆各庄2号', 1, 23, 1669548489, 0);
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (44, '运动鞋', 1, '只', '024', 139.00, '男鞋跑步鞋2021新款冬季潮流男士厚底旅游透气保暖网面休闲运动鞋黑色板鞋防水皮面慢跑鞋子男加厚 黑白皮面-【店长推荐】 42码 ', 1, 19, 1669178384, 1669178428);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (45, '洗衣液', 1, '瓶', '012', 89.90, '全效洁净除菌洗衣液（洁雅百合香）超值套装16斤（2KG*4瓶）家庭装 去渍去污无残留 ', 2, 19, 1669178484, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (46, '洗发露', 1, '瓶', '013', 200.00, '洗发水清爽去油型1KG（持久去屑清洁止痒清爽柔润）柠檬香 男士女士通用 新老包装随机发货 ', 1, 19, 1669178546, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (47, '电动车', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km ', 2, 19, 1669178630, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (48, '跑步机', 1, '台', '005', 1599.00, '生态伙伴跑步机家用智能可折叠免安装室内健身多功能走步机64CM宽大跑台（支持HUAWEI HiLink） ', 1, 19, 1669178665, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (50, '洗衣液', 1, '瓶', '013', 98.80, '全效洁净除菌洗衣液（洁雅百合香）超值套装16斤（2KG*4瓶）家庭装 去渍去污无残留', 1, 23, 1669535316, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (51, '洗发露', 1, '瓶', '005', 200.00, '洗发水清爽去油型1KG（持久去屑清洁止痒清爽柔润）柠檬香 男士女士通用 新老包装随机发货', 1, 23, 1669535360, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (52, '运动鞋', 1, '只', '008', 139.00, '男鞋跑步鞋2021新款冬季潮流男士厚底旅游透气保暖网面休闲运动鞋黑色板鞋防水皮面慢跑鞋子男加厚 黑白皮面-【店长推荐】 42码', 1, 23, 1669535416, 0);
INSERT INTO `product` (`id`, `name`, `type`, `unit`, `code`, `price`, `description`, `status`, `creator`, `created`, `updated`) VALUES (53, '工装', 1, '只', '023', 89.00, '芯耐磨工作服套装男士纯棉电焊防烫建筑工地汽修春秋工装军绿户外劳保服 军绿A款 【精品581】套装 2XL 175码（120-140斤）', 1, 23, 1669535475, 0);
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
  `version` tinyint DEFAULT NULL COMMENT '系统版本，1-免费版，2-个人版',
  `expired` bigint DEFAULT NULL COMMENT '订阅到期时间',
  `status` tinyint DEFAULT '1' COMMENT '状态，1-正常，2-注销',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `updated` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `email`, `password`, `name`, `version`, `expired`, `status`, `created`, `updated`) VALUES (23, '12345@qq.com', '$2a$10$61Tabrfk/s3TKuzErG4cx.o9VOtaLhmeb3K4x5lkaN2dHns004dNe', '', 2, 1672126771, 1, 1669534759, 1669534807);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
