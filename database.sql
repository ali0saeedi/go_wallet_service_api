CREATE DATABASE  IF NOT EXISTS `wallet_servic`;
USE `wallet_service`;

DROP TABLE IF EXISTS `wallets`;
CREATE TABLE `wallets` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `balance` bigint(20),
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

LOCK TABLES `wallets` WRITE;
INSERT INTO `wallets` VALUES (1,100),(2,200),(3,300);
UNLOCK TABLES;