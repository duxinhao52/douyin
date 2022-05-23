CREATE DATABASE IF NOT EXISTS douyin;
USE douyin;

DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
                             `user_id` bigint(20) NOT NULL,
                             `user_name` varchar(32) NOT NULL,
                             `password` varchar(32) NOT NULL,
                             `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                             `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `status` int(11) DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `video_info`;
CREATE TABLE `video_info` (
                              `video_id` bigint(20) NOT NULL,
                              `user_id` bigint(20) NOT NULL,
                              `title` varchar(100) DEFAULT NULL,
                              `url` varchar(200) NOT NULL,
                              `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                              `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `status` int(11) DEFAULT '1',
                              `cover_url` varchar(200) DEFAULT 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `like_info`;
CREATE TABLE `like_info` (
                             `user_id` bigint(20) NOT NULL,
                             `video_id` bigint(20) NOT NULL,
                             `status` int(11) NOT NULL,
                             `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                             `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `comment_info`;
CREATE TABLE `comment_info` (
                                `comment_id` bigint(20) NOT NULL,
                                `user_id` bigint(20) NOT NULL,
                                `video_id` bigint(20) NOT NULL,
                                `status` int(11) NOT NULL,
                                `content` text,
                                `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `follow_info`;
CREATE TABLE `follow_info` (
                               `user_id` bigint(20) NOT NULL,
                               `follow_user_id` bigint(20) NOT NULL,
                               `status` int(11) NOT NULL,
                               `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                               `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;