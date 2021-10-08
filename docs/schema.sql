-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2021-09-10 20:45:28
-- 服务器版本： 5.7.21
-- PHP 版本： 7.3.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `zlab_drive`
--

-- --------------------------------------------------------

--
-- 表的结构 `zd_files`
--

CREATE TABLE `zd_files` (
  `id` bigint(20) NOT NULL,
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'UserId',
  `uuid` varchar(36) NOT NULL DEFAULT '' COMMENT 'UUID',
  `mime` varchar(32) NOT NULL DEFAULT '' COMMENT 'MimeType',
  `size` int(11) NOT NULL DEFAULT '0' COMMENT '文件大小',
  `hash` varchar(40) NOT NULL DEFAULT '' COMMENT '文件哈希',
  `parent` bigint(20) NOT NULL DEFAULT '0' COMMENT '父目录',
  `path` varchar(1000) NOT NULL DEFAULT '' COMMENT '路径',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '文件名',
  `attr` varchar(120) DEFAULT '' COMMENT '文件属性',
  `f_ctime` int(11) NOT NULL DEFAULT '0' COMMENT '文件创建时间',
  `f_mtime` int(11) NOT NULL DEFAULT '0' COMMENT '文件修改时间',
  `ctime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `mtime` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转储表的索引
--

--
-- 表的索引 `zd_files`
--
ALTER TABLE `zd_files`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uuid` (`uuid`) USING BTREE,
  ADD KEY `pa` (`parent`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `zd_files`
--
ALTER TABLE `zd_files`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
