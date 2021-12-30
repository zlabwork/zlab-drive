-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- 主机： 127.0.0.1
-- 生成日期： 2021-12-30 01:57:36
-- 服务器版本： 10.2.7-MariaDB
-- PHP 版本： 8.0.9

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
-- 表的结构 `zd_favorite`
--

CREATE TABLE `zd_favorite` (
  `id` bigint(20) NOT NULL,
  `uid` bigint(20) NOT NULL DEFAULT 0,
  `fid` bigint(20) NOT NULL DEFAULT 0,
  `ctime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `zd_files`
--

CREATE TABLE `zd_files` (
  `id` bigint(20) NOT NULL,
  `uid` bigint(20) NOT NULL DEFAULT 0 COMMENT 'userId',
  `uuid` varchar(36) NOT NULL DEFAULT '' COMMENT 'UUID',
  `mime` varchar(32) NOT NULL DEFAULT '' COMMENT 'MimeType',
  `size` int(11) NOT NULL DEFAULT 0 COMMENT '文件大小',
  `hash` varchar(40) NOT NULL DEFAULT '' COMMENT '文件哈希',
  `parent` bigint(20) NOT NULL DEFAULT 0 COMMENT '父目录',
  `path` varchar(1000) NOT NULL DEFAULT '' COMMENT '路径',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '文件名',
  `attr` varchar(120) DEFAULT '' COMMENT '文件属性',
  `f_ctime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '文件创建时间',
  `f_mtime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '文件修改时间',
  `ctime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '创建时间',
  `mtime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `zd_recent`
--

CREATE TABLE `zd_recent` (
  `id` bigint(20) NOT NULL,
  `uid` bigint(20) NOT NULL DEFAULT 0,
  `fid` bigint(20) NOT NULL DEFAULT 0,
  `ctime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `zd_trash`
--

CREATE TABLE `zd_trash` (
  `id` bigint(20) NOT NULL,
  `fid` bigint(20) NOT NULL DEFAULT 0 COMMENT 'FileId',
  `uid` bigint(20) NOT NULL DEFAULT 0 COMMENT 'userId',
  `uuid` varchar(36) NOT NULL DEFAULT '' COMMENT 'UUID',
  `mime` varchar(32) NOT NULL DEFAULT '' COMMENT 'MimeType',
  `size` int(11) NOT NULL DEFAULT 0 COMMENT '文件大小',
  `hash` varchar(40) NOT NULL DEFAULT '' COMMENT '文件哈希',
  `parent` bigint(20) NOT NULL DEFAULT 0 COMMENT '历史父目录',
  `path` varchar(1000) NOT NULL DEFAULT '' COMMENT '当前路径',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '文件名',
  `attr` varchar(120) DEFAULT '' COMMENT '文件属性',
  `f_ctime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '文件创建时间',
  `f_mtime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '文件修改时间',
  `ctime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '创建时间',
  `mtime` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转储表的索引
--

--
-- 表的索引 `zd_favorite`
--
ALTER TABLE `zd_favorite`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ut` (`uid`,`ctime`);

--
-- 表的索引 `zd_files`
--
ALTER TABLE `zd_files`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uuid` (`uuid`) USING BTREE,
  ADD KEY `pa` (`parent`);

--
-- 表的索引 `zd_recent`
--
ALTER TABLE `zd_recent`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ut` (`uid`,`ctime`);

--
-- 表的索引 `zd_trash`
--
ALTER TABLE `zd_trash`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uid` (`uid`,`mtime`) USING BTREE;

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `zd_favorite`
--
ALTER TABLE `zd_favorite`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `zd_files`
--
ALTER TABLE `zd_files`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `zd_recent`
--
ALTER TABLE `zd_recent`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `zd_trash`
--
ALTER TABLE `zd_trash`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
