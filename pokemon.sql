/*
 Navicat Premium Data Transfer

 Source Server         : mysqlserver
 Source Server Type    : MySQL
 Source Server Version : 50651
 Source Host           : 172.18.0.1:3307
 Source Schema         : pokemon

 Target Server Type    : MySQL
 Target Server Version : 50651
 File Encoding         : 65001

 Date: 31/12/2021 20:34:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for my_pokemons
-- ----------------------------
DROP TABLE IF EXISTS `my_pokemons`;
CREATE TABLE `my_pokemons` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `real_pokemon_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `counter` int(11) DEFAULT '1',
  `image_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

SET FOREIGN_KEY_CHECKS = 1;
