-- MySQL dump 10.13  Distrib 8.0.20, for Linux (x86_64)
--
-- Host: localhost    Database: test
-- ------------------------------------------------------
-- Server version	8.0.20


--
-- Current Database: `test`
--

CREATE DATABASE  `test` ;

USE `test`;

--
-- Table structure for table `app_user`
--

DROP TABLE IF EXISTS `app_user`;

CREATE TABLE `app_user` (
                            `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
                            `nick_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '昵称',
                            `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '密码',
                            `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '电子邮箱',
                            `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '手机号',
                            `gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '性别',
                            `birthday` date DEFAULT NULL COMMENT '生日',
                            `created_at` int DEFAULT NULL COMMENT '创建时间',
                            `updated_at` int DEFAULT NULL COMMENT '更新时间',
                            `deleted_at` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '删除时间',
                            `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1:有效 0:无效',
                            PRIMARY KEY (`id`),
                            KEY `nickName` (`nick_name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';


--
-- Dumping data for table `app_user`
--

LOCK TABLES `app_user` WRITE;
INSERT INTO `app_user` VALUES (4,'张三','fcea920f7412b5da7be0cf42b8c93759','123@qq.com','1234567',0,'1991-03-24',1613719714,1613719714,NULL,1),(5,'李四','0a1c6944cb66d02ccefac35620ce2e51','123@qq.com','123456798',0,'1991-03-24',1613721320,1613721320,NULL,1),(7,'刘二','b06462fbe31a4183e8589fbaccb5bdd5','123@qq.com','17600113419',0,'1991-03-24',1613721643,1613721643,NULL,1);
UNLOCK TABLES;
