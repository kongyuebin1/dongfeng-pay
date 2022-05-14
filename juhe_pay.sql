-- MySQL dump 10.13  Distrib 5.6.51, for Linux (x86_64)
--
-- Host: localhost    Database: juhe_pay
-- ------------------------------------------------------
-- Server version	5.6.51

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `account_history_info`
--

DROP TABLE IF EXISTS `account_history_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account_history_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `account_uid` varchar(100) NOT NULL COMMENT '账号uid',
  `account_name` varchar(100) NOT NULL COMMENT '账户名称',
  `type` varchar(20) NOT NULL DEFAULT '' COMMENT '减款，加款',
  `amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '操作对应金额对应的金额',
  `balance` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '操作后的当前余额',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8 COMMENT='账户账户资金动向表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_history_info`
--

LOCK TABLES `account_history_info` WRITE;
/*!40000 ALTER TABLE `account_history_info` DISABLE KEYS */;
INSERT INTO `account_history_info` VALUES (24,'8888c9kit6bimggos5kk0c8g','天天','plus_amount',50000.000,50000.000,'2022-04-27 11:54:28','2022-04-27 11:54:28');
/*!40000 ALTER TABLE `account_history_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `account_info`
--

DROP TABLE IF EXISTS `account_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `status` varchar(20) NOT NULL DEFAULT 'active' COMMENT '状态',
  `account_uid` varchar(100) NOT NULL COMMENT '账户uid，对应为merchant_uid或者agent_uid',
  `account_name` varchar(100) NOT NULL COMMENT '账户名称，对应的是merchant_name或者agent_name',
  `balance` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '账户余额',
  `settle_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '已经结算了的金额',
  `loan_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '押款金额',
  `wait_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '待结算资金',
  `freeze_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '账户冻结金额',
  `payfor_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '账户代付中金额',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account_uid` (`account_uid`),
  UNIQUE KEY `account_name` (`account_name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='账户记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_info`
--

LOCK TABLES `account_info` WRITE;
/*!40000 ALTER TABLE `account_info` DISABLE KEYS */;
INSERT INTO `account_info` VALUES (1,'active','8888c254gk8isf001cqrj6og','测试账号',6100.000,6100.000,0.000,0.000,0.000,0.000,'2022-04-23 10:38:38','2021-04-29 13:46:57'),(3,'active','8888c9kit6bimggos5kk0c8g','天天',50000.000,50000.000,0.000,0.000,0.000,0.000,'2022-04-27 11:54:28','2022-04-27 11:52:57');
/*!40000 ALTER TABLE `account_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `agent_info`
--

DROP TABLE IF EXISTS `agent_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `agent_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `status` varchar(20) NOT NULL DEFAULT 'active' COMMENT '代理状态状态',
  `agent_name` varchar(100) NOT NULL COMMENT '代理名称',
  `agent_password` varchar(50) NOT NULL COMMENT '代理登录密码',
  `pay_password` varchar(50) NOT NULL COMMENT '支付密码',
  `agent_uid` varchar(100) NOT NULL COMMENT '代理编号',
  `agent_phone` varchar(15) NOT NULL COMMENT '代理手机号',
  `agent_remark` text COMMENT '备注',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `agent_name` (`agent_name`),
  UNIQUE KEY `agent_uid` (`agent_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='代理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `agent_info`
--

LOCK TABLES `agent_info` WRITE;
/*!40000 ALTER TABLE `agent_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `agent_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bank_card_info`
--

DROP TABLE IF EXISTS `bank_card_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bank_card_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `uid` varchar(100) NOT NULL COMMENT '唯一标识',
  `user_name` varchar(100) NOT NULL COMMENT '用户名称',
  `bank_name` varchar(100) NOT NULL COMMENT '银行名称',
  `bank_code` varchar(30) NOT NULL COMMENT '银行编码',
  `bank_account_type` varchar(20) NOT NULL COMMENT '银行账号类型',
  `account_name` varchar(50) NOT NULL COMMENT '银行账户名称',
  `bank_no` varchar(50) NOT NULL COMMENT '银行账号',
  `identify_card` varchar(100) NOT NULL COMMENT '证件类型',
  `certificate_no` varchar(100) NOT NULL COMMENT '证件号码',
  `phone_no` varchar(50) NOT NULL COMMENT '手机号码',
  `bank_address` varchar(200) NOT NULL COMMENT '银行地址',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='银行卡表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bank_card_info`
--

LOCK TABLES `bank_card_info` WRITE;
/*!40000 ALTER TABLE `bank_card_info` DISABLE KEYS */;
INSERT INTO `bank_card_info` VALUES (5,'3333c9kirsbimggos5kk0c5g','8888c254gk8isf001cqrj6og','11','11','private','11','11','identify-card','11','11','11','2022-04-27 11:50:09','2022-04-27 11:50:09');
/*!40000 ALTER TABLE `bank_card_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `legend_any_money`
--

DROP TABLE IF EXISTS `legend_any_money`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `legend_any_money` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `template_name` varchar(50) NOT NULL DEFAULT 'OK' COMMENT '模板名称',
  `game_money_name` varchar(30) DEFAULT NULL COMMENT '游戏币名称，默认是元宝，也可以是钻石、点券',
  `game_money_scale` int(11) NOT NULL DEFAULT '100' COMMENT '游戏币比例，默认是1：100',
  `limit_low` double(20,2) NOT NULL DEFAULT '10.00' COMMENT '最低值金额，默认是10元',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='充值任意金额类型';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `legend_any_money`
--

LOCK TABLES `legend_any_money` WRITE;
/*!40000 ALTER TABLE `legend_any_money` DISABLE KEYS */;
INSERT INTO `legend_any_money` VALUES (13,'技术测试模板','钻石代付',1000,10.00,'2021-05-16 15:03:39','2021-05-16 15:03:39');
/*!40000 ALTER TABLE `legend_any_money` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `legend_area`
--

DROP TABLE IF EXISTS `legend_area`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `legend_area` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `area_name` varchar(150) NOT NULL DEFAULT 'OK' COMMENT '分区名称',
  `uid` varchar(50) NOT NULL COMMENT '分区id',
  `group_name` varchar(150) NOT NULL COMMENT '分组id',
  `notify_url` varchar(1024) DEFAULT NULL COMMENT '通知地址',
  `attach_params` varchar(1024) DEFAULT NULL COMMENT '通知参数',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `template_name` varchar(120) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`) USING BTREE,
  UNIQUE KEY `area_name` (`area_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=200016 DEFAULT CHARSET=utf8 COMMENT='分区列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `legend_area`
--

LOCK TABLES `legend_area` WRITE;
/*!40000 ALTER TABLE `legend_area` DISABLE KEYS */;
INSERT INTO `legend_area` VALUES (200006,'琼琼2','c2htuhoisf00uunvkbb0','分组12','代付','大幅度发','2021-05-18 23:37:43','2021-05-18 23:37:43','技术测试模板'),(200007,'琼琼3','c2htum8isf00uunvkbbg','分组12','的','地方','2021-05-18 23:38:01','2021-05-18 23:38:01','技术测试模板'),(200008,'琼琼4','c2htuugisf00uunvkbc0','分组12','的','2','2021-05-18 23:38:34','2021-05-18 23:38:34','技术测试模板'),(200009,'琼琼5','c2htv0gisf00uunvkbcg','分组12','6','6','2021-05-18 23:38:42','2021-05-18 23:38:42','技术测试模板'),(200010,'琼琼7','c2htv28isf00uunvkbd0','分组12','7','7','2021-05-18 23:38:49','2021-05-18 23:38:49','技术测试模板'),(200011,'琼琼8','c2htv6gisf00uunvkbdg','分组1','8','8','2021-05-18 23:39:06','2021-05-18 23:39:06','技术测试模板'),(200012,'琼琼9','c2htvigisf00uunvkbe0','分组1','9','9','2021-05-18 23:39:54','2021-05-18 23:39:54','技术测试模板'),(200013,'琼琼11','c2htvloisf00uunvkbeg','分组1','11','11','2021-05-18 23:40:07','2021-05-18 23:40:07','技术测试模板'),(200015,'琼琼33','c2htvpgisf00uunvkbfg','分组12','33','33','2021-05-19 11:33:55','2021-05-18 23:40:22','技术测试模板');
/*!40000 ALTER TABLE `legend_area` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `legend_fix_money`
--

DROP TABLE IF EXISTS `legend_fix_money`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `legend_fix_money` (
  `uid` varchar(32) NOT NULL COMMENT '唯一id',
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `template_name` varchar(50) NOT NULL DEFAULT 'OK' COMMENT '模板名称',
  `price` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '售价，默认是0',
  `goods_name` varchar(120) DEFAULT NULL COMMENT '商品名称',
  `goods_no` varchar(60) DEFAULT NULL COMMENT '商品编号',
  `buy_times` int(11) NOT NULL DEFAULT '1' COMMENT '该商品可够次数，默认为1',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='充值固定金额类型';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `legend_fix_money`
--

LOCK TABLES `legend_fix_money` WRITE;
/*!40000 ALTER TABLE `legend_fix_money` DISABLE KEYS */;
INSERT INTO `legend_fix_money` VALUES ('',15,'',1.00,'1','1',1,'2021-05-16 14:18:06','2021-05-16 14:18:06'),('',16,'',1.00,'1','1',1,'2021-05-16 14:18:59','2021-05-16 14:18:59'),('',17,'',1.00,'1','1',1,'2021-05-16 14:21:46','2021-05-16 14:21:46');
/*!40000 ALTER TABLE `legend_fix_money` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `legend_fix_present`
--

DROP TABLE IF EXISTS `legend_fix_present`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `legend_fix_present` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `template_name` varchar(50) NOT NULL DEFAULT 'OK' COMMENT '模板名称',
  `money` int(11) NOT NULL DEFAULT '0' COMMENT '金额，默认是0',
  `present_money` int(11) DEFAULT NULL COMMENT '赠送金额',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `uid` varchar(32) NOT NULL DEFAULT '唯一id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='固定金额赠送';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `legend_fix_present`
--

LOCK TABLES `legend_fix_present` WRITE;
/*!40000 ALTER TABLE `legend_fix_present` DISABLE KEYS */;
INSERT INTO `legend_fix_present` VALUES (8,'技术测试模板',3,3,'2021-05-16 15:03:39','2021-05-16 14:55:34','F5NNP3OC41EKBORPYBQV'),(14,'技术测试模板',5,5,'2021-05-16 15:03:39','2021-05-16 15:03:13','3WBHP7Q6T421WNKNDW4M'),(15,'技术测试模板',4,4,'2021-05-16 15:03:39','2021-05-16 15:03:23','PV29NV5PGV32KSI847CV');
/*!40000 ALTER TABLE `legend_fix_present` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `legend_group`
--

DROP TABLE IF EXISTS `legend_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `legend_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `group_name` varchar(50) NOT NULL DEFAULT 'OK' COMMENT '分组名称',
  `uid` varchar(50) NOT NULL COMMENT '分组id',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `group_name` (`group_name`) USING BTREE,
  UNIQUE KEY `group_uid` (`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 COMMENT='分组列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `legend_group`
--

LOCK TABLES `legend_group` WRITE;
/*!40000 ALTER TABLE `legend_group` DISABLE KEYS */;
INSERT INTO `legend_group` VALUES (8,'分组5','c2h7h6oisf00jf91l220','2021-05-17 22:07:23','2021-05-17 22:07:23'),(9,'分组6','c2h7h7oisf00jf91l22g','2021-05-17 22:07:27','2021-05-17 22:07:27'),(10,'分组7','c2h7h8gisf00jf91l230','2021-05-17 22:07:30','2021-05-17 22:07:30'),(11,'分组8','c2h7h98isf00jf91l23g','2021-05-17 22:07:33','2021-05-17 22:07:33'),(12,'分组9','c2h7ha8isf00jf91l240','2021-05-17 22:07:37','2021-05-17 22:07:37'),(13,'分组10','c2h7hboisf00jf91l24g','2021-05-17 22:07:43','2021-05-17 22:07:43'),(14,'分组11','c2h7hcgisf00jf91l250','2021-05-17 22:07:46','2021-05-17 22:07:46'),(16,'分组1','c2h7i0oisf00jf91l260','2021-05-17 22:09:07','2021-05-17 22:09:07'),(18,'分组12','c2h7i70isf00jf91l270','2021-05-17 22:09:32','2021-05-17 22:09:32');
/*!40000 ALTER TABLE `legend_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `legend_scale_present`
--

DROP TABLE IF EXISTS `legend_scale_present`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `legend_scale_present` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `template_name` varchar(50) NOT NULL DEFAULT 'OK' COMMENT '模板名称',
  `money` int(11) NOT NULL DEFAULT '0' COMMENT '金额，默认是0',
  `present_scale` decimal(20,3) DEFAULT NULL COMMENT '赠送比例',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `uid` varchar(32) NOT NULL DEFAULT '唯一id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='按百分比赠送';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `legend_scale_present`
--

LOCK TABLES `legend_scale_present` WRITE;
/*!40000 ALTER TABLE `legend_scale_present` DISABLE KEYS */;
INSERT INTO `legend_scale_present` VALUES (1,'技术测试',7,7.000,'2021-05-14 23:34:39','2021-05-14 23:34:39','XEYZVFES4YUIU2MSEXMR'),(2,'技术测试',8,8.000,'2021-05-14 23:34:39','2021-05-14 23:34:39','DBYAVA464EUEPE97T2M9'),(4,'技术测试',5,5.000,'2021-05-14 23:41:04','2021-05-14 23:41:04','9UXT5BO1KMEMCPYLNGPL'),(5,'技术测试',6,6.000,'2021-05-14 23:41:04','2021-05-14 23:41:04','471NH5XKPRIO4B2T5O6D'),(8,'技术测试模板',6,6.000,'2021-05-16 15:03:39','2021-05-16 15:03:39','87JOQSRQ1K9BE5NHMVJO'),(9,'技术测试模板',7,7.000,'2021-05-16 15:03:39','2021-05-16 15:03:39','8SYB9VBM1CPDW8RXSD88');
/*!40000 ALTER TABLE `legend_scale_present` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `legend_scale_template`
--

DROP TABLE IF EXISTS `legend_scale_template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `legend_scale_template` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `merchant_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '商户uid',
  `template_name` varchar(50) NOT NULL DEFAULT 'OK' COMMENT '模板名称',
  `user_uid` varchar(50) NOT NULL DEFAULT 'role' COMMENT '用户标识',
  `user_warn` varchar(240) DEFAULT NULL COMMENT '用户标识提醒',
  `money_type` varchar(32) NOT NULL DEFAULT 'any' COMMENT '金额类型，any-任意金额，fix-固定金额',
  `present_type` varchar(32) NOT NULL DEFAULT 'close' COMMENT '赠送方式，close-关闭，fix-固定金额的赠送，scale-按按百分比赠送',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `templete_name` (`template_name`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8 COMMENT='传奇比例模板';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `legend_scale_template`
--

LOCK TABLES `legend_scale_template` WRITE;
/*!40000 ALTER TABLE `legend_scale_template` DISABLE KEYS */;
INSERT INTO `legend_scale_template` VALUES (53,'8888c254gk8isf001cqrj6og','技术测试模板','我是狗','的反间谍法','radio-any-money','scale-present','2021-05-16 15:03:39','2021-05-16 14:23:02');
/*!40000 ALTER TABLE `legend_scale_template` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu_info`
--

DROP TABLE IF EXISTS `menu_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `menu_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `menu_order` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '一级菜单的排名顺序',
  `menu_uid` varchar(40) NOT NULL COMMENT '一级菜单的唯一标识',
  `first_menu` varchar(50) NOT NULL COMMENT '一级菜单名称，字符不能超过50',
  `second_menu` text COMMENT '二级菜单名称，每个之间用|隔开',
  `creater` varchar(20) NOT NULL COMMENT '创建者的id',
  `status` varchar(10) NOT NULL DEFAULT 'active' COMMENT '菜单的状态情况，默认是active',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_first_menu` (`first_menu`),
  UNIQUE KEY `u_menu_uid` (`menu_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='存放左侧栏的菜单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu_info`
--

LOCK TABLES `menu_info` WRITE;
/*!40000 ALTER TABLE `menu_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `menu_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `merchant_deploy_info`
--

DROP TABLE IF EXISTS `merchant_deploy_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `merchant_deploy_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `status` varchar(20) NOT NULL DEFAULT 'active' COMMENT '商户状态状态',
  `merchant_uid` varchar(100) NOT NULL COMMENT '商户uid',
  `pay_type` varchar(50) DEFAULT NULL COMMENT '支付配置',
  `single_road_uid` varchar(100) DEFAULT NULL COMMENT '单通道uid',
  `single_road_name` varchar(200) DEFAULT NULL COMMENT '单通道名称',
  `single_road_platform_rate` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '单通到平台净利率',
  `single_road_agent_rate` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '单通到代理净利率',
  `roll_road_code` varchar(100) DEFAULT NULL COMMENT '轮询通道编码',
  `roll_road_name` varchar(200) DEFAULT NULL COMMENT '轮询通道名称',
  `roll_road_platform_rate` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '轮询通道平台净利率',
  `roll_road_agent_rate` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '轮询通道代理净利率',
  `is_loan` varchar(10) NOT NULL DEFAULT 'NO' COMMENT '是否押款',
  `loan_rate` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '押款比例，默认是0',
  `loan_days` int(11) NOT NULL DEFAULT '0' COMMENT '押款的天数，默认0天',
  `unfreeze_hour` int(11) NOT NULL DEFAULT '0' COMMENT '每天解款的时间点，默认是凌晨',
  `wait_unfreeze_amount` decimal(20,3) DEFAULT NULL COMMENT '等待解款的金额',
  `loan_amount` decimal(20,3) DEFAULT NULL COMMENT '押款中的金额',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='商户通道配置；\r\n单通道给商户的汇率=single_road_platform_rate+single_road_agent_rate+basic_fee；\r\n轮询通道汇率=roll_road_platform_rate+roll_road_agent_rate+basic_fee；';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `merchant_deploy_info`
--

LOCK TABLES `merchant_deploy_info` WRITE;
/*!40000 ALTER TABLE `merchant_deploy_info` DISABLE KEYS */;
INSERT INTO `merchant_deploy_info` VALUES (3,'active','8888c254gk8isf001cqrj6og','WEIXIN_SCAN','','',0.000,0.000,'0914','轮询池1',1.000,1.000,'yes',0.500,0,21,0.000,0.000,'2021-11-08 21:56:40','2021-08-18 17:16:32');
/*!40000 ALTER TABLE `merchant_deploy_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `merchant_info`
--

DROP TABLE IF EXISTS `merchant_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `merchant_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `status` varchar(20) NOT NULL DEFAULT 'active' COMMENT '商户状态状态',
  `belong_agent_uid` varchar(100) DEFAULT NULL COMMENT '所属代理uid',
  `belong_agent_name` varchar(100) DEFAULT NULL COMMENT '所属代理名称',
  `merchant_name` varchar(100) NOT NULL DEFAULT '客户' COMMENT '商户名称',
  `merchant_uid` varchar(100) NOT NULL COMMENT '商户uid',
  `merchant_key` varchar(100) NOT NULL COMMENT '商户key',
  `merchant_secret` varchar(100) NOT NULL COMMENT '商户密钥',
  `login_account` varchar(100) NOT NULL COMMENT '登录账号',
  `login_password` varchar(100) NOT NULL COMMENT '登录密码',
  `auto_settle` varchar(10) NOT NULL DEFAULT 'YES' COMMENT '是否自动结算',
  `auto_pay_for` varchar(10) NOT NULL DEFAULT 'YES' COMMENT '是否自动代付',
  `white_ips` text COMMENT '配置ip白名单',
  `remark` text COMMENT '备注',
  `single_pay_for_road_uid` varchar(100) DEFAULT NULL COMMENT '单代付代付通道uid',
  `single_pay_for_road_name` varchar(200) DEFAULT NULL COMMENT '单代付通道名称',
  `roll_pay_for_road_code` varchar(100) DEFAULT NULL COMMENT '轮询代付通道编码',
  `roll_pay_for_road_name` varchar(200) DEFAULT NULL COMMENT '轮询代付通道名称',
  `payfor_fee` double DEFAULT NULL COMMENT '代付手续费',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `merchant_uid` (`merchant_uid`),
  UNIQUE KEY `merchant_name` (`merchant_name`),
  UNIQUE KEY `merchant_key` (`merchant_key`),
  UNIQUE KEY `merchant_secret` (`merchant_secret`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='商户支付配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `merchant_info`
--

LOCK TABLES `merchant_info` WRITE;
/*!40000 ALTER TABLE `merchant_info` DISABLE KEYS */;
INSERT INTO `merchant_info` VALUES (1,'active','123','的范德萨发的','测试账号','8888c254gk8isf001cqrj6og','kkkkc254gk8isf001cqrj6p0','ssssc254gk8isf001cqrj6pg','17343601111','E10ADC3949BA59ABBE56E057F20F883E','yes','yes','','这个是用来测试账号的','4444c4vdosgisf0020c06r0g','代丽宝','','',0,'2022-04-23 10:36:51','2021-04-29 13:46:57'),(3,'active','','','天天','8888c9kit6bimggos5kk0c8g','kkkkc9kit6bimggos5kk0c90','ssssc9kit6bimggos5kk0c9g','18888888888','E10ADC3949BA59ABBE56E057F20F883E','no','yes','','','','','','',0,'2022-04-27 11:52:57','2022-04-27 11:52:57');
/*!40000 ALTER TABLE `merchant_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `merchant_load_info`
--

DROP TABLE IF EXISTS `merchant_load_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `merchant_load_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `status` varchar(20) NOT NULL DEFAULT 'no' COMMENT 'no-没有结算，yes-结算',
  `merchant_uid` varchar(100) NOT NULL COMMENT '商户uid',
  `road_uid` varchar(50) NOT NULL COMMENT '通道uid',
  `load_date` varchar(50) NOT NULL COMMENT '押款日期，格式2019-10-10',
  `load_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '押款金额',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商户对应的每条通道的押款信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `merchant_load_info`
--

LOCK TABLES `merchant_load_info` WRITE;
/*!40000 ALTER TABLE `merchant_load_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `merchant_load_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `notify_info`
--

DROP TABLE IF EXISTS `notify_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notify_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `type` varchar(10) NOT NULL COMMENT '支付订单-order， 代付订单-payfor',
  `bank_order_id` varchar(50) NOT NULL COMMENT '系统订单id',
  `merchant_order_id` varchar(50) NOT NULL COMMENT '下游商户订单id',
  `status` varchar(20) NOT NULL DEFAULT 'wait' COMMENT '状态字段',
  `times` int(11) NOT NULL DEFAULT '0' COMMENT '回调次数',
  `url` text COMMENT '回调的url',
  `response` text COMMENT '回调返回的结果',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `merchant_order_id` (`merchant_order_id`),
  UNIQUE KEY `bank_order_id` (`bank_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='支付回调';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `notify_info`
--

LOCK TABLES `notify_info` WRITE;
/*!40000 ALTER TABLE `notify_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `notify_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_info`
--

DROP TABLE IF EXISTS `order_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `merchant_order_id` varchar(50) NOT NULL COMMENT '下游商户提交过来的订单id',
  `shop_name` varchar(100) NOT NULL COMMENT '商品名称',
  `order_period` varchar(3) NOT NULL DEFAULT '0' COMMENT '订单有效时间，小时制',
  `bank_order_id` varchar(50) NOT NULL COMMENT '平台自身的订单id',
  `bank_trans_id` varchar(50) NOT NULL COMMENT '上游返回的订单id',
  `order_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '订单提交金额',
  `show_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '展示在用户面前待支付的金额',
  `fact_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '实际支付金额',
  `roll_pool_code` varchar(50) DEFAULT NULL COMMENT '轮询产品编码',
  `roll_pool_name` varchar(100) DEFAULT NULL COMMENT '轮询产品名称',
  `road_uid` varchar(100) NOT NULL COMMENT '通道uid',
  `road_name` varchar(200) NOT NULL COMMENT '通道名称',
  `pay_product_code` varchar(100) NOT NULL COMMENT '支付产品编码',
  `pay_product_name` varchar(200) NOT NULL COMMENT '支付产品名称',
  `pay_type_code` varchar(50) NOT NULL COMMENT '支付类型编码',
  `pay_type_name` varchar(100) NOT NULL COMMENT '支付类型名称',
  `os_type` varchar(5) NOT NULL COMMENT '平台类型，苹果app-0， 安卓app-1，苹果H5-3，安卓H5-4，pc-5',
  `status` varchar(20) NOT NULL DEFAULT 'wait' COMMENT '等待支付-wait,支付成功-success, 支付失败-failed',
  `refund` varchar(5) NOT NULL DEFAULT 'no' COMMENT '退款-yes， 没有退款-no',
  `refund_time` varchar(100) DEFAULT NULL COMMENT '退款时间',
  `freeze` varchar(5) NOT NULL DEFAULT 'no' COMMENT '冻结-yes， 没有-no',
  `freeze_time` varchar(100) DEFAULT NULL COMMENT '冻结时间',
  `unfreeze` varchar(5) NOT NULL DEFAULT 'no' COMMENT '解冻-yes，没有-no',
  `unfreeze_time` varchar(100) DEFAULT NULL COMMENT '解冻时间',
  `return_url` text COMMENT '订单支付后，跳转的地址',
  `notify_url` text COMMENT '订单回调给下游的地址',
  `merchant_uid` varchar(100) NOT NULL COMMENT '商户uid，表示订单是哪个商户的',
  `merchant_name` varchar(200) NOT NULL COMMENT '商户名称',
  `agent_uid` varchar(100) DEFAULT NULL COMMENT '代理uid，表示该商户是谁的代理',
  `agent_name` varchar(200) DEFAULT NULL COMMENT '代理名称',
  `response` text COMMENT '上游返回的结果',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `merchant_order_id` (`merchant_order_id`),
  UNIQUE KEY `bank_order_id` (`bank_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_info`
--

LOCK TABLES `order_info` WRITE;
/*!40000 ALTER TABLE `order_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_profit_info`
--

DROP TABLE IF EXISTS `order_profit_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_profit_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `merchant_name` varchar(100) NOT NULL COMMENT '商户名称',
  `merchant_uid` varchar(50) NOT NULL COMMENT '商户uid',
  `agent_uid` varchar(100) DEFAULT NULL COMMENT '代理uid，表示该商户是谁的代理',
  `agent_name` varchar(200) DEFAULT NULL COMMENT '代理名称',
  `pay_product_code` varchar(100) NOT NULL COMMENT '支付产品编码',
  `pay_product_name` varchar(200) NOT NULL COMMENT '支付产品名称',
  `pay_type_code` varchar(50) NOT NULL COMMENT '支付类型编码',
  `pay_type_name` varchar(100) NOT NULL COMMENT '支付类型名称',
  `status` varchar(20) NOT NULL DEFAULT 'wait' COMMENT '等待支付-wait,支付成功-success, 支付失败-failed',
  `merchant_order_id` varchar(50) NOT NULL COMMENT '下游商户提交过来的订单id',
  `bank_order_id` varchar(50) NOT NULL COMMENT '平台自身的订单id',
  `bank_trans_id` varchar(50) NOT NULL COMMENT '上游返回的订单id',
  `order_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '订单提交金额',
  `show_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '展示在用户面前待支付的金额',
  `fact_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '实际支付金额',
  `user_in_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '商户入账金额',
  `all_profit` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '总的利润，包括上游，平台，代理',
  `supplier_rate` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '上游的汇率',
  `platform_rate` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '平台自己的手续费率',
  `agent_rate` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '代理的手续费率',
  `supplier_profit` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '上游的利润',
  `platform_profit` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '平台利润',
  `agent_profit` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '代理利润',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `merchant_order_id` (`merchant_order_id`),
  UNIQUE KEY `bank_order_id` (`bank_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单利润表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_profit_info`
--

LOCK TABLES `order_profit_info` WRITE;
/*!40000 ALTER TABLE `order_profit_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_profit_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_settle_info`
--

DROP TABLE IF EXISTS `order_settle_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_settle_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `pay_product_code` varchar(100) NOT NULL COMMENT '支付产品编码',
  `pay_product_name` varchar(200) NOT NULL COMMENT '支付产品名称',
  `pay_type_code` varchar(50) NOT NULL COMMENT '支付类型编码',
  `pay_type_name` varchar(100) NOT NULL COMMENT '支付类型名称',
  `merchant_uid` varchar(100) NOT NULL COMMENT '商户uid，表示订单是哪个商户的',
  `road_uid` varchar(50) NOT NULL COMMENT '通道uid',
  `merchant_name` varchar(200) NOT NULL COMMENT '商户名称',
  `merchant_order_id` varchar(50) NOT NULL COMMENT '下游商户提交过来的订单id',
  `bank_order_id` varchar(50) NOT NULL COMMENT '平台自身的订单id',
  `settle_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '结算金额',
  `is_allow_settle` varchar(10) NOT NULL DEFAULT 'yes' COMMENT '是否允许结算，允许-yes，不允许-no',
  `is_complete_settle` varchar(10) NOT NULL DEFAULT 'no' COMMENT '该笔订单是否结算完毕，没有结算-no，结算完毕-yes',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `merchant_order_id` (`merchant_order_id`),
  UNIQUE KEY `bank_order_id` (`bank_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单结算表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_settle_info`
--

LOCK TABLES `order_settle_info` WRITE;
/*!40000 ALTER TABLE `order_settle_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_settle_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payfor_info`
--

DROP TABLE IF EXISTS `payfor_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payfor_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `payfor_uid` varchar(100) NOT NULL COMMENT '代付唯一uid',
  `merchant_uid` varchar(100) NOT NULL COMMENT '发起代付的商户uid',
  `merchant_name` varchar(200) NOT NULL COMMENT '发起代付的商户名称',
  `merchant_order_id` varchar(50) DEFAULT NULL COMMENT '下游代付订单id',
  `bank_order_id` varchar(50) NOT NULL COMMENT '系统代付订单id',
  `bank_trans_id` varchar(50) NOT NULL COMMENT '上游返回的代付订单id',
  `road_uid` varchar(100) NOT NULL COMMENT '所用的代付通道uid',
  `road_name` varchar(200) NOT NULL COMMENT '所有通道的名称',
  `roll_pool_code` varchar(100) DEFAULT NULL COMMENT '所用轮询池编码',
  `roll_pool_name` varchar(200) DEFAULT NULL COMMENT '所用轮询池的名称',
  `payfor_fee` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '代付手续费',
  `payfor_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '代付到账金额',
  `payfor_total_amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '代付总金额',
  `bank_code` varchar(20) NOT NULL COMMENT '银行编码',
  `bank_name` varchar(100) NOT NULL COMMENT '银行名称',
  `bank_account_name` varchar(100) NOT NULL COMMENT '银行开户名称',
  `bank_account_no` varchar(50) NOT NULL COMMENT '银行开户账号',
  `bank_account_type` varchar(20) DEFAULT 'private' COMMENT '银行卡类型，对私-private，对公-public',
  `country` varchar(50) NOT NULL DEFAULT '中国' COMMENT '开户所属国家',
  `province` varchar(50) NOT NULL DEFAULT '' COMMENT '银行卡开户所属省',
  `city` varchar(50) NOT NULL DEFAULT '' COMMENT '银行卡开户所属城市',
  `ares` varchar(50) DEFAULT NULL COMMENT '所属地区',
  `bank_account_address` text COMMENT '银行开户具体街道',
  `phone_no` varchar(20) NOT NULL COMMENT '开户所用手机号',
  `give_type` varchar(50) NOT NULL DEFAULT 'payfor_road' COMMENT '下发类型，payfor_road-通道打款，payfor_hand-手动打款，payfor_refuse-拒绝打款',
  `type` varchar(36) NOT NULL DEFAULT 'auto' COMMENT '代付类型，self_api-系统发下， 管理员手动下发给商户-self_merchant，管理自己提现-self_help',
  `notify_url` text COMMENT '代付结果回调给下游的地址',
  `status` varchar(20) NOT NULL DEFAULT 'wait' COMMENT '审核-payfor_confirm,系统处理中-payfor_solving，银行处理中-payfor_banking，代付成功-success, 代付失败-failed',
  `is_send` varchar(10) NOT NULL DEFAULT 'no' COMMENT '未发送-no，已经发送-yes',
  `request_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '发起请求时间',
  `response_time` varchar(32) DEFAULT '0000-00-00 00:00:00' COMMENT '上游做出响应的时间',
  `response_content` text COMMENT '代付的最终结果',
  `remark` text COMMENT '代付备注',
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `payfor_uid` (`payfor_uid`),
  UNIQUE KEY `bank_order_id` (`bank_order_id`),
  UNIQUE KEY `merchant_order_id` (`merchant_order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='代付表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payfor_info`
--

LOCK TABLES `payfor_info` WRITE;
/*!40000 ALTER TABLE `payfor_info` DISABLE KEYS */;
INSERT INTO `payfor_info` VALUES (17,'ppppc9ro8qu7matk051l4cfg','8888c9kit6bimggos5kk0c8g','天天','c9ro8qu7matk051l4cg0','4444c9ro8qu7matk051l4cgg','','','','','',2.000,3333.000,3335.000,'11','11','11','11','private','','','','','11','11','','self_merchant','','payfor_confirm','no','2022-05-08 16:50:51','','','','2022-05-08 16:50:51','2022-05-08 16:50:51');
/*!40000 ALTER TABLE `payfor_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `power_info`
--

DROP TABLE IF EXISTS `power_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `power_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `first_menu_uid` varchar(40) NOT NULL COMMENT '一级菜单的唯一标识',
  `second_menu_uid` varchar(40) NOT NULL COMMENT '二级菜单的唯一标识',
  `second_menu` varchar(50) NOT NULL COMMENT '二级菜单的名称',
  `power_item` varchar(50) NOT NULL COMMENT '权限项的名称',
  `power_id` varchar(200) NOT NULL COMMENT '权限的ID',
  `creater` varchar(20) NOT NULL COMMENT '创建者的id',
  `status` varchar(10) NOT NULL DEFAULT 'active' COMMENT '菜单的状态情况，默认是active',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_power_id` (`power_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='存放控制页面的一些功能操作';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `power_info`
--

LOCK TABLES `power_info` WRITE;
/*!40000 ALTER TABLE `power_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `power_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `road_info`
--

DROP TABLE IF EXISTS `road_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `road_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `status` varchar(20) NOT NULL DEFAULT 'active通道状态',
  `road_name` varchar(100) NOT NULL COMMENT '通道名称',
  `road_uid` varchar(100) NOT NULL COMMENT '通道唯一id',
  `remark` varchar(100) DEFAULT NULL COMMENT '备注',
  `product_name` varchar(100) NOT NULL COMMENT '上游产品名称',
  `product_uid` varchar(100) NOT NULL COMMENT '上游产品编号',
  `pay_type` varchar(50) NOT NULL COMMENT '支付类型',
  `basic_fee` double NOT NULL COMMENT '基本汇率/成本汇率',
  `settle_fee` double NOT NULL COMMENT '代付手续费',
  `total_limit` double NOT NULL COMMENT '通道总额度',
  `today_limit` double NOT NULL COMMENT '每日最多额度',
  `single_min_limit` double NOT NULL COMMENT '单笔最小金额',
  `single_max_limit` double NOT NULL COMMENT '单笔最大金额',
  `star_hour` int(11) NOT NULL COMMENT '通道开始时间',
  `end_hour` int(11) NOT NULL COMMENT '通道结束时间',
  `params` text COMMENT '参数json格式',
  `today_income` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '当天的收入',
  `total_income` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '通道总收入',
  `today_profit` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '当天的收益',
  `total_profit` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '通道总收益',
  `balance` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '通道的余额',
  `request_all` int(11) DEFAULT '0' COMMENT '请求总次数',
  `request_success` int(11) DEFAULT '0' COMMENT '请求成功次数',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `road_name` (`road_name`),
  UNIQUE KEY `road_uid` (`road_uid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='通道数据表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `road_info`
--

LOCK TABLES `road_info` WRITE;
/*!40000 ALTER TABLE `road_info` DISABLE KEYS */;
INSERT INTO `road_info` VALUES (3,'active','代丽宝','4444c4vdosgisf0020c06r0g','代丽宝','代丽支付','DAILI','WEIXIN_SCAN',1,2,1000000,10000,1,1000,0,23,'{}',300.000,300.000,5.000,5.000,0.000,4,3,'2021-11-08 15:56:42','2021-09-13 13:06:58'),(4,'active','歌力思','4444c9hsa9bimggv91apihc0','','快付支付','KF','QQ_SYT',1,1,10000,10000,1,10000,0,23,'',0.000,0.000,0.000,0.000,0.000,0,0,'2022-04-23 09:21:41','2022-04-23 09:21:41');
/*!40000 ALTER TABLE `road_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `road_pool_info`
--

DROP TABLE IF EXISTS `road_pool_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `road_pool_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `status` varchar(20) NOT NULL DEFAULT 'active' COMMENT '通道池状态',
  `road_pool_name` varchar(100) NOT NULL COMMENT '通道池名称',
  `road_pool_code` varchar(100) NOT NULL COMMENT '通道池编号',
  `road_uid_pool` text COMMENT '通道池里面的通道uid',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `road_pool_name` (`road_pool_name`),
  UNIQUE KEY `road_pool_code` (`road_pool_code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='通道池';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `road_pool_info`
--

LOCK TABLES `road_pool_info` WRITE;
/*!40000 ALTER TABLE `road_pool_info` DISABLE KEYS */;
INSERT INTO `road_pool_info` VALUES (1,'active','轮询池1','0914','4444c4vdosgisf0020c06r0g','2021-09-15 13:07:44','2021-09-15 13:07:36');
/*!40000 ALTER TABLE `road_pool_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_info`
--

DROP TABLE IF EXISTS `role_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `role_name` varchar(100) NOT NULL COMMENT '角色名称',
  `role_uid` varchar(200) NOT NULL COMMENT '角色唯一标识号',
  `show_first_menu` text NOT NULL COMMENT '可以展示的一级菜单名',
  `show_first_uid` text NOT NULL COMMENT '可以展示的一级菜单uid',
  `show_second_menu` text NOT NULL COMMENT '可以展示的二级菜单名',
  `show_second_uid` text NOT NULL COMMENT '可以展示的二级菜单uid',
  `show_power` text NOT NULL COMMENT '可以展示的权限项名称',
  `show_power_uid` text NOT NULL COMMENT '可以展示的权限项uid',
  `remark` text NOT NULL COMMENT '角色描述',
  `creater` varchar(20) NOT NULL COMMENT '创建者的id',
  `status` varchar(10) NOT NULL DEFAULT 'active' COMMENT '菜单的状态情况，默认是active',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_power_name` (`role_name`),
  UNIQUE KEY `u_role_uid` (`role_uid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_info`
--

LOCK TABLES `role_info` WRITE;
/*!40000 ALTER TABLE `role_info` DISABLE KEYS */;
INSERT INTO `role_info` VALUES (1,'超级管理员','c62dtroisf0022dhe5ig','','','','','','','开发','10086','active','2021-11-05 15:36:15','2022-04-27 11:49:53'),(2,'555','c9hs6irimggv91apih60','','','','','','','','10086','active','2022-04-23 09:13:47','2022-04-27 11:49:53');
/*!40000 ALTER TABLE `role_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `second_menu_info`
--

DROP TABLE IF EXISTS `second_menu_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `second_menu_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `first_menu_order` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '一级菜单对应的顺序',
  `menu_order` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '二级菜单的排名顺序',
  `first_menu_uid` varchar(40) NOT NULL COMMENT '二级菜单的唯一标识',
  `first_menu` varchar(50) NOT NULL COMMENT '一级菜单名称，字符不能超过50',
  `second_menu_uid` varchar(40) NOT NULL COMMENT '二级菜单唯一标识',
  `second_menu` varchar(225) NOT NULL COMMENT '二级菜单名称',
  `second_router` varchar(200) NOT NULL COMMENT '二级菜单路由',
  `creater` varchar(20) NOT NULL COMMENT '创建者的id',
  `status` varchar(10) NOT NULL DEFAULT 'active' COMMENT '菜单的状态情况，默认是active',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_second_menu` (`second_menu`),
  UNIQUE KEY `u_second_menu_uid` (`second_menu_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='存放左侧栏的二级菜单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `second_menu_info`
--

LOCK TABLES `second_menu_info` WRITE;
/*!40000 ALTER TABLE `second_menu_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `second_menu_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_info`
--

DROP TABLE IF EXISTS `user_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `user_id` varchar(40) NOT NULL COMMENT '用户登录号',
  `passwd` varchar(40) NOT NULL COMMENT '用户登录密码',
  `nick` varchar(30) NOT NULL DEFAULT 'kity' COMMENT '用户昵称',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  `ip` varchar(30) NOT NULL DEFAULT '127.0.0.1' COMMENT '用户当前ip',
  `status` varchar(10) NOT NULL DEFAULT 'active' COMMENT '该用户的状态 active、unactive、delete',
  `role` varchar(100) NOT NULL DEFAULT 'nothing' COMMENT '管理者分配的角色',
  `role_name` varchar(200) NOT NULL DEFAULT '普通操作员' COMMENT '操作员分配的角色名称',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='管理员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_info`
--

LOCK TABLES `user_info` WRITE;
/*!40000 ALTER TABLE `user_info` DISABLE KEYS */;
INSERT INTO `user_info` VALUES (1,'10086','E10ADC3949BA59ABBE56E057F20F883E','admin',NULL,'::1','active','nothing','普通操作员','2021-04-21 13:30:02','2022-05-08 08:43:01');
/*!40000 ALTER TABLE `user_info` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-14 14:41:55
