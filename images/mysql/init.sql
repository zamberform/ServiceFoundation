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
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NOW(),
  `platform_id` int(11) DEFAULT NULL,
  FOREIGN KEY (`platform_id`) 
  REFERENCES `platform` (`id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- Add datas into app
-- ----------------------------

-- ----------------------------
-- Table structure for platform
-- ----------------------------
DROP TABLE IF EXISTS `platform`;
CREATE TABLE `platform` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type_code` int(10) DEFAULT '0',
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
  `last_login_at` datetime DEFAULT NOW(),
  `name` varchar(250) DEFAULT NULL,
  `uuid` varchar(250) DEFAULT NULL,
  `email` varchar(250) DEFAULT NULL,
  `pass` varchar(250) DEFAULT NULL,
  `platform_id` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT '0',
  `avatar_url` varchar(500) DEFAULT NULL,
  `sex` tinyint(1) NOT NULL DEFAULT '0',
  `introduce` varchar(500) DEFAULT NULL,
  FOREIGN KEY (`platform_id`) 
  REFERENCES `platform` (`id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- Table structure for user_package_log
-- ----------------------------
DROP TABLE IF EXISTS `user_package_log`;
CREATE TABLE `user_package_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NOW(),
  `user_id` int(11) DEFAULT '0',
  `package_id` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- Table structure for package
-- ----------------------------
DROP TABLE IF EXISTS `package`;
CREATE TABLE `package` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(250) DEFAULT NULL,
  `price` int(11) DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NOW(),
  `show_start_at` datetime DEFAULT NULL,
  `show_end_at` datetime DEFAULT NULL,
  `enable_start_at` datetime DEFAULT NULL,
  `enable_end_at` datetime DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(250) DEFAULT NULL,
  `status` int(11) DEFAULT '0',
  `content_desc` varchar(250) DEFAULT NULL,
  `comment_flg` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NOW(),
  `tag_id` int(11) DEFAULT NULL,
  FOREIGN KEY (`tag_id`) 
  REFERENCES `tag` (`id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `status` int(11) DEFAULT '0',
  `comment` varchar(250) DEFAULT NULL,
  `created_at` datetime DEFAULT NOW(),
  `article_id` int(11) DEFAULT NULL,
  FOREIGN KEY (`article_id`) 
  REFERENCES `article` (`id`),
  `user_id` int(11) DEFAULT NULL,
  FOREIGN KEY (`user_id`) 
  REFERENCES `user` (`id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(250) DEFAULT NULL,
  `color` varchar(250) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NOW(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(250) DEFAULT NULL,
  `pass` varchar(250) DEFAULT NULL,
  `email` varchar(250) DEFAULT NULL,
  `role` int(11) DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NOW(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- add self admin user
-- ----------------------------

