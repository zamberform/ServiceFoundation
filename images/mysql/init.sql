SET FOREIGN_KEY_CHECKS=0;
-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `version` varchar(255) DEFAULT NULL,
  `build_code` int(10) DEFAULT '0',
  `url` varchar(255) DEFAULT NULL,
  `update_status` int(10) DEFAULT '0',
  `platform_type` int(10) DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NOW(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NOW(),
  `deleted_at` datetime DEFAULT NULL,
  `activated_at` datetime DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `uuid` varchar(50) DEFAULT NULL,
  `pass` varchar(100) DEFAULT NULL,
  `role` int(11) DEFAULT '0',
  `status` int(11) DEFAULT '0',
  `platform` int(11) DEFAULT '0',
  `avatar_url` varchar(500) DEFAULT NULL,
  `sex` tinyint(1) NOT NULL DEFAULT '0',
  `location` varchar(200) DEFAULT NULL,
  `introduce` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
