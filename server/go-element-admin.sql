/*
 Navicat Premium Data Transfer

 Source Server         : docker
 Source Server Type    : MySQL
 Source Server Version : 80017
 Source Host           : 127.0.0.1:3306
 Source Schema         : vue-element-admin

 Target Server Type    : MySQL
 Target Server Version : 80017
 File Encoding         : 65001

 Date: 07/11/2020 20:30:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '1', '/users', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/user', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/user/:user_id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/users', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/user/pwd/:user_id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/user/logs', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/user/logs', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/menus', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/menu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/menu/:menu_id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/menus', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/paths', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/path', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/path/:path_id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/paths', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/roles', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/role', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/role/:role_id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/roles', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '3', '/users', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '3', '/user/logs', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '3', '/menus', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '3', '/paths', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '3', '/roles', 'GET', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
  `menu_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '上级ID',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '类型',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图标',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '路由名称',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '路由组件',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '路由地址',
  `redirect` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '跳转地址',
  `permission` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '隐藏',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (7, 0, '系统管理', 2, 'M', 'component', 'System', 'Layout', '/system', '/system/user', '', 0, '2020-11-01 14:15:23', '2020-05-24 18:51:21');
INSERT INTO `menus` VALUES (8, 7, '菜单管理', 4, 'C', 'list', 'Menu', '/system/menu', '/system/menu', '', 'system:menu', 0, '2020-10-25 01:57:16', '2020-05-24 18:52:26');
INSERT INTO `menus` VALUES (9, 7, '角色管理', 0, 'C', 'peoples', 'Role', '/system/role', '/system/role', '', 'system:role', 0, NULL, '2020-05-24 18:53:31');
INSERT INTO `menus` VALUES (10, 7, '用户管理', 5, 'C', 'user', 'User', '/system/user', '/system/user', '', 'system:user', 0, '2020-10-25 01:57:07', '2020-05-24 18:54:26');
INSERT INTO `menus` VALUES (11, 7, '登录日志', 0, 'C', 'log', 'Log', '/system/log', '/system/log', '', 'system:log', 0, NULL, '2020-05-24 18:55:20');
INSERT INTO `menus` VALUES (12, 0, 'Demo', 0, 'M', 'example', 'Demo', 'Layout', '/demo', '', NULL, 0, '2020-05-30 15:00:24', '2020-05-25 20:31:21');
INSERT INTO `menus` VALUES (13, 12, 'Demo-1', 0, 'M', 'example', 'Demo1', '/demo/demo-1', '/demo/demo-1', '', NULL, 0, '2020-05-30 15:00:55', '2020-05-25 20:32:13');
INSERT INTO `menus` VALUES (14, 12, 'Demo-2', 0, 'M', 'example', 'Demo2', '/demo/demo-2', '/demo/demo-2', '', NULL, 0, NULL, '2020-05-25 20:32:54');
INSERT INTO `menus` VALUES (15, 13, 'Demo-1-1', 0, 'C', 'example', 'Demo11', '/demo/demo-1/demo-1-1', '/demo/demo-1/demo-1-1', '', NULL, 0, '2020-05-25 20:33:47', '2020-05-25 20:33:29');
INSERT INTO `menus` VALUES (16, 14, 'Demo-2-1', 0, 'C', 'example', 'Demo21', '/demo/demo-2/demo-2-1', '/demo/demo-2/demo-2-1', '', '', 0, '2020-10-24 16:06:37', '2020-05-25 20:34:43');
INSERT INTO `menus` VALUES (19, 7, '接口管理', 3, 'C', 'network', 'Interface', '/system/path', '/system/path', '', 'system:path', 0, '2020-11-01 14:26:53', '2020-10-25 01:56:50');
INSERT INTO `menus` VALUES (20, 21, '接口文档', 0, 'C', 'swagger', 'Swagger', 'Layout', 'http://localhost:8001/swagger/index.html', '', '', 0, '2020-11-01 14:14:11', '2020-10-25 14:42:21');
INSERT INTO `menus` VALUES (21, 0, '系统工具', 1, 'M', 'system', 'Settings', 'Layout', '/settings', '', '', 0, '2020-11-01 14:15:29', '2020-11-01 14:13:52');
INSERT INTO `menus` VALUES (22, 21, '表单生成器', 0, 'C', 'build', 'FormGen', 'Layout', 'https://mrhj.gitee.io/form-generator/', '', '', 0, NULL, '2020-11-01 14:22:32');

-- ----------------------------
-- Table structure for paths
-- ----------------------------
DROP TABLE IF EXISTS `paths`;
CREATE TABLE `paths`  (
  `path_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父ID',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '路径',
  `method` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求方法',
  `type` char(5) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '接口类型',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`path_id`) USING BTREE,
  INDEX `idx_path`(`path`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '接口表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of paths
-- ----------------------------
INSERT INTO `paths` VALUES (1, 0, '用户管理', '', '', 'M', NULL, '2020-10-25 02:21:28');
INSERT INTO `paths` VALUES (4, 1, '用户列表', '/users', 'GET', 'J', NULL, '2020-10-25 14:26:39');
INSERT INTO `paths` VALUES (6, 1, '添加用户', '/user', 'POST', 'J', NULL, '2020-10-25 14:22:04');
INSERT INTO `paths` VALUES (7, 1, '修改用户', '/user/:user_id', 'PUT', 'J', NULL, '2020-10-25 14:22:23');
INSERT INTO `paths` VALUES (8, 1, '删除用户', '/users', 'DELETE', 'J', NULL, '2020-10-25 14:22:45');
INSERT INTO `paths` VALUES (9, 1, '修改用户密码', '/user/pwd/:user_id', 'PUT', 'J', NULL, '2020-10-25 14:23:48');
INSERT INTO `paths` VALUES (10, 1, '登录日志列表', '/user/logs', 'GET', 'J', NULL, '2020-10-25 14:25:12');
INSERT INTO `paths` VALUES (11, 1, '删除登录日志', '/user/logs', 'DELETE', 'J', NULL, '2020-10-25 14:25:37');
INSERT INTO `paths` VALUES (12, 0, '菜单管理', '', '', 'M', NULL, '2020-10-25 14:26:43');
INSERT INTO `paths` VALUES (13, 12, '菜单列表', '/menus', 'GET', 'J', NULL, '2020-10-25 14:27:18');
INSERT INTO `paths` VALUES (14, 12, '添加菜单', '/menu', 'POST', 'J', NULL, '2020-10-25 14:27:34');
INSERT INTO `paths` VALUES (15, 12, '修改菜单', '/menu/:menu_id', 'PUT', 'J', NULL, '2020-10-25 14:27:54');
INSERT INTO `paths` VALUES (16, 12, '删除菜单', '/menus', 'DELETE', 'J', NULL, '2020-10-25 14:28:10');
INSERT INTO `paths` VALUES (17, 0, '接口管理', '', '', 'M', NULL, '2020-10-25 14:28:43');
INSERT INTO `paths` VALUES (18, 17, '接口列表', '/paths', 'GET', 'J', NULL, '2020-10-25 14:28:59');
INSERT INTO `paths` VALUES (19, 17, '添加接口', '/path', 'POST', 'J', '2020-10-25 14:29:45', '2020-10-25 14:29:22');
INSERT INTO `paths` VALUES (20, 17, '修改接口', '/path/:path_id', 'PUT', 'J', NULL, '2020-10-25 14:31:12');
INSERT INTO `paths` VALUES (21, 17, '删除接口', '/paths', 'DELETE', 'J', NULL, '2020-10-25 14:31:26');
INSERT INTO `paths` VALUES (22, 0, '角色管理', '', '', 'M', NULL, '2020-10-25 14:31:48');
INSERT INTO `paths` VALUES (23, 22, '角色列表', '/roles', 'GET', 'J', NULL, '2020-10-25 14:32:10');
INSERT INTO `paths` VALUES (24, 22, '添加角色', '/role', 'POST', 'J', NULL, '2020-10-25 14:33:45');
INSERT INTO `paths` VALUES (25, 22, '修改角色', '/role/:role_id', 'PUT', 'J', NULL, '2020-10-25 14:34:06');
INSERT INTO `paths` VALUES (26, 22, '删除角色', '/roles', 'DELETE', 'J', NULL, '2020-10-25 14:34:29');

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
  `role_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态',
  `path_ids` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '接口ID',
  `menu_ids` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '菜单ID',
  `buttons` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '权限标识',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`role_id`, `role_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (1, '管理员', '拥有全部权限', 1, '[12,15,16,14,13,22,23,24,25,26,17,18,19,20,21,1,4,6,7,8,9,10,11]', '[7,10,8,19,9,11,21,20,22,12,13,15,14,16]', '[{\"btns\":[\"system:user:query\",\"system:user:add\",\"system:user:edit\",\"system:user:del\"],\"menu_id\":10},{\"btns\":[\"system:menu:del\",\"system:menu:edit\",\"system:menu:add\",\"system:menu:query\"],\"menu_id\":8},{\"btns\":[\"system:role:query\",\"system:role:add\",\"system:role:del\",\"system:role:edit\"],\"menu_id\":9},{\"btns\":[\"system:log:query\",\"system:log:del\"],\"menu_id\":11},{\"btns\":[\"system:path:query\",\"system:path:add\",\"system:path:edit\",\"system:path:del\"],\"menu_id\":19}]', '2020-11-01 14:22:56', '2020-05-16 21:14:50');
INSERT INTO `roles` VALUES (3, '观察者', '', 1, '[13,22,23,17,18,1,4,10,12]', '[7,10,8,19,9,11,12,13,15,14,16]', '[{\"btns\":[\"system:user:query\"],\"menu_id\":10},{\"btns\":[\"system:menu:query\"],\"menu_id\":8},{\"btns\":[\"system:log:query\"],\"menu_id\":11},{\"btns\":[\"system:role:query\"],\"menu_id\":9},{\"btns\":[\"system:path:query\",\"system:path:add\",\"system:path:edit\"],\"menu_id\":19}]', '2020-11-01 16:26:40', '2020-05-22 18:50:23');

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles`  (
  `user_role_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`user_role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO `user_roles` VALUES (9, 3, 3, '2020-10-18 16:25:51');
INSERT INTO `user_roles` VALUES (10, 1, 1, '2020-10-18 16:26:14');
INSERT INTO `user_roles` VALUES (24, 3, 11, '2020-10-24 06:52:29');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '帐号',
  `password` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `idx_user_name`(`user_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'bigfool', 'efd2708635b70c935fb9015e5fb28fe5', 1, '2020-11-01 13:57:08', '2020-05-16 21:15:43');
INSERT INTO `users` VALUES (3, 'usenav', '492a4d82c53b26c6e87be11a6dae3e70', 1, '2020-10-18 17:01:47', '2020-05-23 19:43:54');
INSERT INTO `users` VALUES (11, 'www', 'e10adc3949ba59abbe56e057f20f883e', 0, '2020-10-24 14:52:47', '2020-10-24 14:52:29');

SET FOREIGN_KEY_CHECKS = 1;
