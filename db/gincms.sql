/*
 Navicat Premium Data Transfer

 Source Server         : 本机phpstudy的mysql7
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : gincms

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 22/08/2023 21:41:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for online_table
-- ----------------------------
DROP TABLE IF EXISTS `online_table`;
CREATE TABLE `online_table`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'id',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表名',
  `comments` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表描述',
  `form_layout` tinyint(4) NULL DEFAULT NULL COMMENT '表单布局',
  `tree` tinyint(4) NULL DEFAULT NULL COMMENT '是否树  0：否   1：是',
  `tree_pid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '树父id',
  `tree_label` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '树展示列',
  `table_type` tinyint(4) NULL DEFAULT NULL COMMENT '表类型  0：单表',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '是否更新  0：否   1：是',
  `version` int(11) NULL DEFAULT NULL COMMENT '版本号',
  `deleted` tinyint(4) NULL DEFAULT NULL COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updater` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Online表单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of online_table
-- ----------------------------

-- ----------------------------
-- Table structure for online_table_column
-- ----------------------------
DROP TABLE IF EXISTS `online_table_column`;
CREATE TABLE `online_table_column`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字段名称',
  `comments` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字段描述',
  `length` int(11) NOT NULL COMMENT '字段长度',
  `point_length` int(11) NOT NULL COMMENT '小数点',
  `default_value` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '默认值',
  `column_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段类型',
  `column_pk` tinyint(4) NULL DEFAULT NULL COMMENT '字段主键 0：否  1：是',
  `column_null` tinyint(4) NULL DEFAULT NULL COMMENT '字段为空 0：否  1：是',
  `form_item` tinyint(4) NULL DEFAULT NULL COMMENT '表单项 0：否  1：是',
  `form_required` tinyint(4) NULL DEFAULT NULL COMMENT '表单必填 0：否  1：是',
  `form_input` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表单控件',
  `form_default` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表单控件默认值',
  `form_dict` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表单字典',
  `grid_item` tinyint(4) NULL DEFAULT NULL COMMENT '列表项 0：否  1：是',
  `grid_sort` tinyint(4) NULL DEFAULT NULL COMMENT '列表排序 0：否  1：是',
  `query_item` tinyint(4) NULL DEFAULT NULL COMMENT '查询项 0：否  1：是',
  `query_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '查询方式',
  `query_input` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '查询控件',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `table_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Online表单字段' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of online_table_column
-- ----------------------------

-- ----------------------------
-- Table structure for sys_attachment
-- ----------------------------
DROP TABLE IF EXISTS `sys_attachment`;
CREATE TABLE `sys_attachment`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附件名称',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附件地址',
  `size` bigint(20) NOT NULL DEFAULT 0 COMMENT '附件大小',
  `size_tip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '附件大小KB M显示',
  `platform` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '存储平台',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `version` int(11) NOT NULL DEFAULT 0 COMMENT '版本号',
  `deleted` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `updater` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '附件管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_attachment
-- ----------------------------
INSERT INTO `sys_attachment` VALUES (1, '16917279797514075421.png', 'upload/2023-08-11/16917279797514075421.png', 9758, '9.53K', 'local', 0, 0, 1, 1000, '2023-08-11 12:26:19', 1000, '2023-08-11 12:52:25');
INSERT INTO `sys_attachment` VALUES (2, '16917282224898124277.png', 'upload/2023-08-11/16917282224898124277.png', 9758, '9.53K', 'local', 0, 0, 0, 1000, '2023-08-11 12:30:23', 1000, '2023-08-11 12:30:23');
INSERT INTO `sys_attachment` VALUES (3, '16923728245227842104.png', 'upload/2023-08-18/16923728245227842104.png', 9758, '9.53K', 'local', 0, 0, 0, 1000, '2023-08-18 23:33:45', 1000, '2023-08-18 23:33:45');
INSERT INTO `sys_attachment` VALUES (4, '16923730925051878229.png', 'upload/2023-08-18/16923730925051878229.png', 9758, '9.53K', 'local', 0, 0, 0, 1000, '2023-08-18 23:38:12', 1000, '2023-08-18 23:38:12');
INSERT INTO `sys_attachment` VALUES (5, '16923734501815091733.png', 'upload/2023-08-18/16923734501815091733.png', 9758, '9.53K', 'local', 0, 0, 0, 1000, '2023-08-18 23:44:10', 1000, '2023-08-18 23:44:10');

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `dict_type_id` bigint(20) NOT NULL COMMENT '字典类型ID',
  `dict_label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典标签',
  `dict_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典值',
  `label_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签样式',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `version` int(11) NOT NULL DEFAULT 0 COMMENT '版本号',
  `deleted` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `updater` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 1, '停用', '0', 'danger', '', 1, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (2, 1, '正常', '1', 'primary', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (3, 2, '男', '0', 'primary', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (4, 2, '女', '1', 'success', '', 1, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (5, 2, '未知', '2', 'warning', '', 2, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-08 22:11:57');
INSERT INTO `sys_dict_data` VALUES (6, 3, '正常', '1', 'primary', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (7, 3, '停用', '0', 'danger', '', 1, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (8, 4, '全部数据', '0', '', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (9, 4, '本机构及子机构数据', '1', '', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (10, 4, '本机构数据', '2', '', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (11, 4, '本人数据', '3', '', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (12, 4, '自定义数据', '4', '', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (13, 5, '禁用', '0', 'danger', '', 1, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (14, 5, '启用', '1', 'primary', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (15, 6, '失败', '0', 'danger', '', 1, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (16, 6, '成功', '1', 'primary', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_data` VALUES (17, 7, '登录成功', '0', 'primary', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (18, 7, '退出成功', '1', 'warning', '', 1, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (19, 7, '验证码错误', '2', 'danger', '', 2, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (20, 7, '账号密码错误', '3', 'danger', '', 3, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (21, 8, '否', '0', 'primary', '', 1, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (22, 8, '是', '1', 'danger', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (23, 9, '是', '1', 'danger', '', 1, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (24, 9, '否', '0', 'primary', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (25, 10, '其它', '0', 'info', '', 10, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (26, 10, '查询', '1', 'primary', '', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (27, 10, '新增', '2', 'success', '', 1, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (28, 10, '修改', '3', 'warning', '', 2, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (29, 10, '删除', '4', 'danger', '', 3, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (30, 10, '导出', '5', 'info', '', 4, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');
INSERT INTO `sys_dict_data` VALUES (31, 10, '导入', '6', 'info', '', 5, 10000, 0, 0, 10000, '2023-08-01 14:38:34', 10000, '2023-08-01 14:38:34');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典类型',
  `dict_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典名称',
  `dict_source` tinyint(4) NOT NULL DEFAULT 0 COMMENT '来源  0：字典数据  1：动态SQL',
  `dict_sql` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '动态SQL',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `version` int(11) NOT NULL DEFAULT 0 COMMENT '版本号',
  `deleted` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `updater` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, 'post_status', '状态', 0, '', '岗位管理', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 0, '2023-08-08 16:34:36');
INSERT INTO `sys_dict_type` VALUES (2, 'user_gender', '性别', 0, '', '用户管理', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_type` VALUES (3, 'user_status', '状态', 0, '', '用户管理', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 0, '2023-08-08 16:35:12');
INSERT INTO `sys_dict_type` VALUES (4, 'role_data_scope', '数据范围', 0, '', '角色管理', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 0, '2023-08-08 16:35:12');
INSERT INTO `sys_dict_type` VALUES (5, 'enable_disable', '状态', 0, '', '功能状态：启用 | 禁用 ', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 0, '2023-08-08 16:35:12');
INSERT INTO `sys_dict_type` VALUES (6, 'success_fail', '状态', 0, '', '操作状态：成功 | 失败', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_type` VALUES (7, 'login_operation', '操作信息', 0, '', '登录管理', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_type` VALUES (8, 'params_type', '系统参数', 0, '', '参数管理', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_type` VALUES (9, 'user_super_admin', '用户是否是超管', 0, '', '用户是否是超管', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_dict_type` VALUES (10, 'log_operate_type', '操作类型', 0, '', '操作日志', 0, 10000, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');

-- ----------------------------
-- Table structure for sys_log_login
-- ----------------------------
DROP TABLE IF EXISTS `sys_log_login`;
CREATE TABLE `sys_log_login`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录IP',
  `address` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录地点',
  `user_agent` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'User Agent',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '登录状态  0：失败   1：成功',
  `operation` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作信息   0：登录成功   1：退出成功  2：验证码错误  3：账号密码错误',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '登录日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_log_login
-- ----------------------------

-- ----------------------------
-- Table structure for sys_log_operate
-- ----------------------------
DROP TABLE IF EXISTS `sys_log_operate`;
CREATE TABLE `sys_log_operate`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `module` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块名',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作名',
  `req_uri` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求URI',
  `req_method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求方法',
  `req_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求参数',
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作IP',
  `address` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录地点',
  `user_agent` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'User Agent',
  `operate_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '操作类型',
  `duration` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '执行时长',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '操作状态  0：失败   1：成功',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `real_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作人',
  `result_msg` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '返回消息',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_log_operate
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` bigint(20) NOT NULL DEFAULT 0 COMMENT '上级ID，一级菜单为0',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单URL',
  `authority` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '授权标识(多个用逗号分隔，如：sys:menu:list,sys:menu:save)',
  `type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '类型   0：菜单   1：按钮   2：接口',
  `open_style` tinyint(4) NOT NULL DEFAULT 0 COMMENT '打开方式   0：内部   1：外部',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `version` int(11) NOT NULL DEFAULT 0 COMMENT '版本号',
  `deleted` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `updater` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_pid`(`pid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 50 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, 0, '系统设置', '', '', 0, 0, 'icon-setting', 1, 0, 0, 10000, '2023-08-01 14:38:32', 10000, '2023-08-01 14:38:32');
INSERT INTO `sys_menu` VALUES (2, 1, '菜单管理', 'sys/menu/index', '', 0, 0, 'icon-menu', 0, 0, 0, 10000, '2023-08-01 14:38:32', 10000, '2023-08-01 14:38:32');
INSERT INTO `sys_menu` VALUES (3, 2, '查看', '', 'get:/admin_api/sys/menu_list|', 1, 0, '', 0, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:27:11');
INSERT INTO `sys_menu` VALUES (4, 2, '新增', '', 'post:/admin_api/sys/menu|', 1, 0, '', 1, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:28:00');
INSERT INTO `sys_menu` VALUES (5, 2, '修改', '', 'put:/admin_api/sys/menu|,get:/admin_api/sys/menu|', 1, 0, '', 2, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:32:27');
INSERT INTO `sys_menu` VALUES (6, 2, '删除', '', 'delete:/admin_api/sys/menu|', 1, 0, '', 3, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:32:57');
INSERT INTO `sys_menu` VALUES (7, 1, '数据字典', 'sys/dict/type', '', 0, 0, 'icon-insertrowabove', 1, 0, 0, 10000, '2023-08-01 14:38:32', 10000, '2023-08-01 14:38:32');
INSERT INTO `sys_menu` VALUES (8, 7, '查询', '', 'get:/admin_api/sys/dict_type_page|,get:/admin_api/sys/dict_data_page|', 1, 0, '', 0, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:44:59');
INSERT INTO `sys_menu` VALUES (9, 7, '新增', '', 'post:/admin_api/sys/dict_type|', 1, 0, '', 2, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:51:19');
INSERT INTO `sys_menu` VALUES (10, 7, '修改', '', 'put:/admin_api/sys/dict_type|,get:/admin_api/sys/dict_type|,all:/admin_api/sys/dict_data|', 1, 0, '', 1, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:54:45');
INSERT INTO `sys_menu` VALUES (11, 7, '删除', '', 'delete:/admin_api/sys/dict_type|', 1, 0, '', 3, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:51:29');
INSERT INTO `sys_menu` VALUES (12, 0, '权限管理', '', '', 0, 0, 'icon-safetycertificate', 0, 0, 0, 10000, '2023-08-01 14:38:32', 10000, '2023-08-01 14:38:32');
INSERT INTO `sys_menu` VALUES (13, 12, '岗位管理', 'sys/post/index', '', 0, 0, 'icon-solution', 2, 0, 0, 10000, '2023-08-01 14:38:32', 10000, '2023-08-01 14:38:32');
INSERT INTO `sys_menu` VALUES (14, 13, '查询', '', 'get:/admin_api/sys/post_page|', 1, 0, '', 0, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-16 22:26:58');
INSERT INTO `sys_menu` VALUES (15, 13, '新增', '', 'post:/admin_api/sys/post|', 1, 0, '', 1, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-16 22:27:49');
INSERT INTO `sys_menu` VALUES (16, 13, '修改', '', 'put:/admin_api/sys/post|,get:/admin_api/sys/post|', 1, 0, '', 2, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-16 22:28:28');
INSERT INTO `sys_menu` VALUES (17, 13, '删除', '', 'delete:/admin_api/sys/post|', 1, 0, '', 3, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-16 22:28:54');
INSERT INTO `sys_menu` VALUES (18, 12, '机构管理', 'sys/org/index', '', 0, 0, 'icon-cluster', 1, 0, 0, 10000, '2023-08-01 14:38:32', 10000, '2023-08-01 14:38:32');
INSERT INTO `sys_menu` VALUES (19, 18, '查询', '', 'get:/admin_api/sys/org_list|', 1, 0, '', 0, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-16 22:23:04');
INSERT INTO `sys_menu` VALUES (20, 18, '新增', '', 'post:/admin_api/sys/org|', 1, 0, '', 1, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-16 22:24:08');
INSERT INTO `sys_menu` VALUES (21, 18, '修改', '', 'put:/admin_api/sys/org|get:/admin_api/sys/org|', 1, 0, '', 2, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-16 22:25:40');
INSERT INTO `sys_menu` VALUES (22, 18, '删除', '', 'delete:/admin_api/sys/org|', 1, 0, '', 3, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-16 22:26:08');
INSERT INTO `sys_menu` VALUES (23, 12, '角色管理', 'sys/role/index', '', 0, 0, 'icon-team', 3, 0, 0, 10000, '2023-08-01 14:38:32', 10000, '2023-08-01 14:38:32');
INSERT INTO `sys_menu` VALUES (24, 23, '查询', '', 'get:/admin_api/sys/role_page|', 1, 0, '', 0, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:13:31');
INSERT INTO `sys_menu` VALUES (25, 23, '新增', '', 'post:/admin_api/sys/role|,get:/admin_api/sys/role_menu|', 1, 0, '', 1, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:21:01');
INSERT INTO `sys_menu` VALUES (26, 23, '修改', '', 'get:/admin_api/sys/user_page|,get:/admin_api/sys/role_user_page|,post:/admin_api/sys/role_user|,delete:/admin_apisys/role_user|,put:/admin_api/sys/role|,get:/admin_api/sys/role|,get:/admin_api/sys/role_menu|', 1, 0, '', 2, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 21:50:03');
INSERT INTO `sys_menu` VALUES (27, 23, '删除', '', 'delete:/admin_api/sys/role|', 1, 0, '', 3, 0, 0, 10000, '2023-08-01 14:38:32', 1000, '2023-08-20 20:23:37');
INSERT INTO `sys_menu` VALUES (28, 12, '用户管理', 'sys/user/index', '', 0, 0, 'icon-user', 0, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_menu` VALUES (29, 28, '查询', '', 'get:/admin_api/sys/user_page|', 1, 0, '', 0, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_menu` VALUES (30, 28, '新增', '', 'post:/admin_api/sys/user|', 1, 0, '', 1, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-16 22:21:44');
INSERT INTO `sys_menu` VALUES (31, 28, '修改', '', 'put:/admin_api/sys/user|,get:/admin_api/sys/user|,get:/admin_api/sys/role_list|,get:/admin_api/sys/post_list|', 1, 0, '', 2, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-20 21:54:41');
INSERT INTO `sys_menu` VALUES (32, 28, '删除', '', 'delete:/admin_api/sys/user|', 1, 0, '', 3, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-20 21:17:54');
INSERT INTO `sys_menu` VALUES (33, 0, '应用管理', '', '', 0, 0, 'icon-appstore', 2, 0, 1, 10000, '2023-08-01 14:38:33', 1000, '2023-08-08 15:40:35');
INSERT INTO `sys_menu` VALUES (34, 1, '附件记录', 'sys/attachment/index', '', 0, 0, 'icon-folder-fill', 3, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_menu` VALUES (35, 34, '查看', '', 'get:/admin_api/sys/attachment_page|', 1, 0, '', 0, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-20 20:58:19');
INSERT INTO `sys_menu` VALUES (37, 34, '删除', '', 'delete:/admin_api/sys/attachment|', 1, 0, '', 1, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-20 20:58:40');
INSERT INTO `sys_menu` VALUES (38, 0, '日志管理', '', '', 0, 0, 'icon-filedone', 3, 0, 0, 10000, '2023-08-01 14:38:33', 10000, '2023-08-01 14:38:33');
INSERT INTO `sys_menu` VALUES (39, 38, '登录日志', 'sys/log/login', 'get:/admin_api/sys/log_login_page|', 0, 0, 'icon-solution', 0, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-20 20:25:31');
INSERT INTO `sys_menu` VALUES (40, 28, '导入', '', '', 1, 0, '', 5, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-16 22:21:57');
INSERT INTO `sys_menu` VALUES (41, 28, '导出', '', '', 1, 0, '', 6, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-16 22:22:02');
INSERT INTO `sys_menu` VALUES (42, 1, '参数管理', 'sys/params/index', 'get:/admin_api/sys/params_page|,all:/admin_api/sys/params|', 0, 0, 'icon-filedone', 2, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-20 20:56:30');
INSERT INTO `sys_menu` VALUES (43, 1, '接口文档', '{{apiUrl}}/doc.html', '', 0, 1, 'icon-file-text-fill', 10, 0, 1, 10000, '2023-08-01 14:38:33', 1000, '2023-08-11 20:41:12');
INSERT INTO `sys_menu` VALUES (44, 0, '在线开发', '', '', 0, 0, 'icon-cloud', 2, 0, 1, 10000, '2023-08-01 14:38:33', 1000, '2023-08-08 15:40:52');
INSERT INTO `sys_menu` VALUES (45, 44, 'Online表单开发', 'online/table/index', 'online:table:all', 0, 0, 'icon-table', 0, 0, 1, 10000, '2023-08-01 14:38:33', 1000, '2023-08-08 15:40:50');
INSERT INTO `sys_menu` VALUES (46, 38, '操作日志', 'sys/log/operate', 'get:/admin_api/sys/operate_log_page|', 0, 0, 'icon-file-text', 1, 0, 0, 10000, '2023-08-01 14:38:33', 1000, '2023-08-20 20:26:14');
INSERT INTO `sys_menu` VALUES (47, 1, '文件管理', 'sys/filemanage/index', '', 0, 0, 'icon-folder-fill', 4, 0, 0, 1000, '2023-08-11 20:21:59', 1000, '2023-08-11 20:23:52');
INSERT INTO `sys_menu` VALUES (48, 0, '查询', '', 'get:/admin_api/sys/fileManage_dir_list|,get:/admin_api/sys/fileManage_dirFile_list|', 2, 0, '', 0, 0, 0, 1000, '2023-08-20 21:04:25', 1000, '2023-08-20 21:04:25');
INSERT INTO `sys_menu` VALUES (49, 34, '下载', '', 'post:/admin_api/down_file|', 1, 0, '', 0, 0, 0, 1000, '2023-08-20 21:08:07', 1000, '2023-08-20 21:08:45');

-- ----------------------------
-- Table structure for sys_org
-- ----------------------------
DROP TABLE IF EXISTS `sys_org`;
CREATE TABLE `sys_org`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` bigint(20) NOT NULL DEFAULT 0 COMMENT '上级ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '机构名称',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `version` int(11) NOT NULL DEFAULT 0 COMMENT '版本号',
  `deleted` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `updater` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_pid`(`pid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '机构管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_org
-- ----------------------------
INSERT INTO `sys_org` VALUES (1, 0, '技术部', 0, 0, 0, 0, 1, '2023-08-04 09:14:28', 0, '2023-08-04 09:14:31');
INSERT INTO `sys_org` VALUES (2, 0, '市场部', 0, 0, 0, 0, 1, '2023-08-04 09:14:56', 0, '2023-08-04 09:14:57');
INSERT INTO `sys_org` VALUES (3, 1, '开发组', 0, 0, 0, 0, 1, '2023-08-04 09:15:08', 0, '2023-08-04 09:15:11');
INSERT INTO `sys_org` VALUES (4, 2, '华南市场', 4, 0, 0, 0, 0, '2023-08-04 17:09:45', 0, '2023-08-04 19:47:08');
INSERT INTO `sys_org` VALUES (5, 0, '人事部', 8, 0, 0, 0, 0, '2023-08-04 17:28:13', 0, '2023-08-04 17:28:13');
INSERT INTO `sys_org` VALUES (6, 2, '东北市场部', 10, 0, 0, 0, 0, '2023-08-05 22:23:46', 0, '2023-08-05 22:23:46');

-- ----------------------------
-- Table structure for sys_params
-- ----------------------------
DROP TABLE IF EXISTS `sys_params`;
CREATE TABLE `sys_params`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `param_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '参数名称',
  `param_type` tinyint(4) NOT NULL COMMENT '系统参数   0：否   1：是',
  `param_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '参数键',
  `param_value` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '参数值',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `version` int(11) NOT NULL DEFAULT 0 COMMENT '版本号',
  `deleted` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `updater` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_param_key`(`param_key`) USING BTREE COMMENT '参数键值唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '参数管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_params
-- ----------------------------
INSERT INTO `sys_params` VALUES (1, '用户登录-验证码开关', 1, 'LOGIN_CAPTCHA', 'false', '是否开启验证码（true：开启，false：关闭）', 10000, 0, 0, 10000, '2023-08-01 14:38:34', 1000, '2023-08-10 14:03:44');
INSERT INTO `sys_params` VALUES (2, '用户进入系统必要的权限', 1, 'needAuthButNeedAllow', 'get:/admin_api/sys/user_info|,get:/admin_api/sys/menu_authority|,get:/sys/menu_nav|,get:/admin_api/sys/dict_type_all|,get:/admin_api/sys/menu_nav||,post:/admin_api/sys/auth_logout|,post:/admin_api/sys/file_upload|', '', 0, 0, 0, 1000, '2023-08-16 21:14:21', 1000, '2023-08-16 21:14:21');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `post_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位编码',
  `post_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位名称',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态  0：停用   1：正常',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `version` int(11) NOT NULL DEFAULT 0 COMMENT '版本号',
  `deleted` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `updater` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '岗位管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'dev', '技术开发人员', 3, 1, 0, 0, 0, 0, '2023-08-02 15:15:31', 1000, '2023-08-06 21:16:10');
INSERT INTO `sys_post` VALUES (2, 'test', '测试岗位11888', 2, 1, 0, 0, 0, 0, '2023-08-02 21:30:24', 1000, '2023-08-06 21:16:10');
INSERT INTO `sys_post` VALUES (4, 'product', '产品经理', 5, 1, 0, 0, 0, 0, '2023-08-02 22:02:48', 0, '2023-08-02 22:34:18');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint(4) NOT NULL DEFAULT 0 COMMENT '数据范围  0：全部数据  1：本机构及子机构数据  2：本机构数据  3：本人数据  4：自定义数据',
  `org_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '机构ID',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `version` int(11) NOT NULL DEFAULT 0 COMMENT '版本号',
  `deleted` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `updater` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_org_id`(`org_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '开发部', '开发部角色', 0, 1, 0, 0, 0, 0, '2023-08-04 20:09:22', 1000, '2023-08-16 00:21:33');
INSERT INTO `sys_role` VALUES (2, '人事部', '人事部角色', 0, 5, 0, 0, 0, 0, '2023-08-04 20:35:49', 1000, '2023-08-05 23:33:12');
INSERT INTO `sys_role` VALUES (4, '市场部', '市场部角色11', 0, 0, 0, 0, 0, 1000, '2023-08-06 12:38:24', 1000, '2023-08-07 19:53:40');
INSERT INTO `sys_role` VALUES (5, '超级管理员', '超级管理员尽量不要删了1', 0, 0, 0, 0, 0, 1000, '2023-08-06 18:47:12', 1000, '2023-08-12 17:35:32');

-- ----------------------------
-- Table structure for sys_role_data_scope
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_data_scope`;
CREATE TABLE `sys_role_data_scope`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `role_id` bigint(20) NULL DEFAULT NULL COMMENT '角色ID',
  `org_id` bigint(20) NULL DEFAULT NULL COMMENT '机构ID',
  `version` int(11) NULL DEFAULT NULL COMMENT '版本号',
  `deleted` tinyint(4) NULL DEFAULT NULL COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updater` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_role_id`(`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色数据权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_data_scope
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_m2m_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_m2m_menu`;
CREATE TABLE `sys_role_m2m_menu`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
  INDEX `idx_role_id`(`role_id`) USING BTREE,
  INDEX `idx_menu_id`(`menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色菜单关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_m2m_menu
-- ----------------------------
INSERT INTO `sys_role_m2m_menu` VALUES (1, 12);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 13);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 14);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 15);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 16);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 17);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 18);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 19);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 20);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 21);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 22);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 23);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 24);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 25);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 26);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 27);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 28);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 29);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 30);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 31);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 32);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 40);
INSERT INTO `sys_role_m2m_menu` VALUES (1, 41);
INSERT INTO `sys_role_m2m_menu` VALUES (4, 12);
INSERT INTO `sys_role_m2m_menu` VALUES (4, 28);
INSERT INTO `sys_role_m2m_menu` VALUES (4, 29);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 1);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 2);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 3);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 4);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 5);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 6);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 7);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 8);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 9);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 10);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 11);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 12);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 13);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 14);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 15);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 16);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 17);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 18);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 19);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 20);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 21);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 22);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 23);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 24);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 25);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 26);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 27);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 28);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 29);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 30);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 31);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 32);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 34);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 35);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 37);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 38);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 39);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 40);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 41);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 42);
INSERT INTO `sys_role_m2m_menu` VALUES (5, 46);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `real_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(4) NOT NULL DEFAULT 0 COMMENT '性别   0：男   1：女   2：未知',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `org_id` bigint(20) NOT NULL COMMENT '机构ID',
  `super_admin` tinyint(4) NOT NULL DEFAULT 0 COMMENT '超级管理员   0：否   1：是',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态  0：停用   1：正常',
  `tenant_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `version` int(11) NOT NULL DEFAULT 0 COMMENT '版本号',
  `deleted` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标识  0：正常   1：已删除',
  `creator` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `updater` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_username`(`username`) USING BTREE COMMENT '用户名唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 10014 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1000, 'admin', '$2a$14$Jjmaeq0MRH8P1DWsYE.ow.wOpXdZM5pa5oS6iNVtJIXoAwOq24cEi', 'admin', 'http://localhost:8066/upload/2023-08-19/16923769534472185970.jpg', 0, 'babamu@126.com', '13612345678', 3, 1, 1, 10000, 0, 0, 10000, '2023-08-01 14:38:32', 0, '2023-08-05 22:36:58');
INSERT INTO `sys_user` VALUES (10008, 'admin999', '', '张三', '', 1, '', '17773935066', 5, 0, 1, 0, 0, 0, 0, '2023-08-05 09:33:09', 0, '2023-08-05 22:30:51');
INSERT INTO `sys_user` VALUES (10009, '223565', '', '2323232332', '', 1, '', '23234234234', 1, 0, 1, 0, 0, 0, 0, '2023-08-05 10:12:24', 0, '2023-08-05 21:24:53');
INSERT INTO `sys_user` VALUES (10010, 'jack', '', '22332', '', 1, '', '11231232132', 5, 0, 1, 0, 0, 0, 0, '2023-08-05 11:21:29', 1000, '2023-08-06 21:39:54');
INSERT INTO `sys_user` VALUES (10012, 'tom522', '', 'tom9', '', 1, '', '17752523322', 5, 0, 1, 0, 0, 0, 0, '2023-08-05 11:31:44', 1000, '2023-08-06 21:38:45');
INSERT INTO `sys_user` VALUES (10013, 'xiao', '$2a$14$8AsNSMkQykpa.6yl4OeZ4uE7vHfE8W5SzSRanHEEq38YmgiWB16BC', 'ssdfsdfsdf', '', 0, '', '17773935555', 5, 0, 1, 0, 0, 0, 0, '2023-08-05 12:16:20', 1000, '2023-08-15 00:01:19');

-- ----------------------------
-- Table structure for sys_user_m2m_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_m2m_post`;
CREATE TABLE `sys_user_m2m_post`  (
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '岗位ID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE,
  INDEX `idx_user_id`(`user_id`) USING BTREE,
  INDEX `idx_post_id`(`post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户岗位关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_m2m_post
-- ----------------------------
INSERT INTO `sys_user_m2m_post` VALUES (1000, 1);
INSERT INTO `sys_user_m2m_post` VALUES (1000, 2);
INSERT INTO `sys_user_m2m_post` VALUES (1000, 4);
INSERT INTO `sys_user_m2m_post` VALUES (10008, 4);
INSERT INTO `sys_user_m2m_post` VALUES (10009, 1);
INSERT INTO `sys_user_m2m_post` VALUES (10009, 2);
INSERT INTO `sys_user_m2m_post` VALUES (10009, 4);
INSERT INTO `sys_user_m2m_post` VALUES (10010, 1);
INSERT INTO `sys_user_m2m_post` VALUES (10010, 2);
INSERT INTO `sys_user_m2m_post` VALUES (10010, 4);
INSERT INTO `sys_user_m2m_post` VALUES (10012, 1);
INSERT INTO `sys_user_m2m_post` VALUES (10012, 2);
INSERT INTO `sys_user_m2m_post` VALUES (10012, 4);
INSERT INTO `sys_user_m2m_post` VALUES (10013, 1);

-- ----------------------------
-- Table structure for sys_user_m2m_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_m2m_role`;
CREATE TABLE `sys_user_m2m_role`  (
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '角色ID',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `role_id`(`role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户角色关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_m2m_role
-- ----------------------------
INSERT INTO `sys_user_m2m_role` VALUES (10013, 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'tom');

SET FOREIGN_KEY_CHECKS = 1;
