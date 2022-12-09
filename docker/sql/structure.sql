/*
 Navicat Premium Data Transfer

 Source Server         : docker 本地
 Source Server Type    : MySQL
 Source Server Version : 100422 (10.4.22-MariaDB-1:10.4.22+maria~focal)
 Source Host           : localhost:3306
 Source Schema         : forum

 Target Server Type    : MySQL
 Target Server Version : 100422 (10.4.22-MariaDB-1:10.4.22+maria~focal)
 File Encoding         : 65001

 Date: 06/12/2022 23:17:07
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for forum_association
-- ----------------------------
DROP TABLE IF EXISTS `forum_association`;
CREATE TABLE `forum_association`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`         int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`        varchar(45) NOT NULL COMMENT '用户名',
    `target_id`       int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被感谢｜屏蔽|收藏 主题id|回复id',
    `additional_id`   int(11) unsigned NOT NULL DEFAULT 0 COMMENT '附加id，当target_id为回复id时，additional_id为主题id',
    `target_user_id`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被感谢｜屏蔽|收藏 用户id',
    `target_username` varchar(45) NOT NULL COMMENT '被感谢用户名',
    `type`            char(15)    NOT NULL DEFAULT '' COMMENT '类型 感谢主题: thanks_post,感谢回复: thanks_reply,屏蔽主题: shield_post,屏蔽回复: shield_reply,收藏主题:collect_post,收藏节点: collect_node,关注用户： follow_user,屏蔽用户: shield_user',
    `created_at`      datetime             DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY               `idx_user_id` (`user_id`) USING BTREE,
    KEY               `idx_target_user_id` (`target_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='论坛 感谢｜屏蔽|收藏| 关注    主题｜回复 |节点| 用户 关联表';

-- ----------------------------
-- Table structure for forum_balance_change_log
-- ----------------------------
DROP TABLE IF EXISTS `forum_balance_change_log`;
CREATE TABLE `forum_balance_change_log`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`    varchar(45)  NOT NULL COMMENT '用户名',
    `type`        char(50)     NOT NULL DEFAULT '' COMMENT '每日登录奖励:login, 每日活跃度奖励: activity, 感谢主题: thanks_posts,感谢回复: thanks_relpy,创建主题: create_posts,创建回复: create_reply,初始奖励: register',
    `amount`      int(11) NOT NULL DEFAULT 0 COMMENT '金额',
    `before`      int(11) unsigned NOT NULL DEFAULT 0 COMMENT '变动前余额',
    `after`       int(11) unsigned NOT NULL DEFAULT 0 COMMENT '变动后余额',
    `relation_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联主题id或关联回复id',
    `remark`      varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
    `created_at`  datetime              DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='论坛余额变动表';

-- ----------------------------
-- Table structure for forum_email_records
-- ----------------------------
DROP TABLE IF EXISTS `forum_email_records`;
CREATE TABLE `forum_email_records`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`   varchar(45)  NOT NULL COMMENT '用户名',
    `type`       varchar(50)  NOT NULL COMMENT '类型',
    `email`      varchar(255) NOT NULL COMMENT '邮箱',
    `title`      varchar(255) NOT NULL COMMENT '标题',
    `content`    text         NOT NULL COMMENT '内容',
    `error`      text         NOT NULL COMMENT '错误信息',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `update_at`  datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='邮件发送历史表';

-- ----------------------------
-- Table structure for forum_follow_or_shield_user_relation
-- ----------------------------
DROP TABLE IF EXISTS `forum_follow_or_shield_user_relation`;
CREATE TABLE `forum_follow_or_shield_user_relation`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`         int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`        varchar(45) NOT NULL COMMENT '用户名',
    `target_user_id`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被关注｜屏蔽用户id',
    `target_username` varchar(45) NOT NULL COMMENT '被关注｜屏蔽用户名',
    `type`            char(10)    NOT NULL DEFAULT '' COMMENT '类型 关注: follow,屏蔽: shield',
    `created_at`      datetime             DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY               `idx_user_id` (`user_id`) USING BTREE,
    KEY               `idx_target_user_id` (`target_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='论坛 关注|屏蔽  用户关联表';

-- ----------------------------
-- Table structure for forum_messages
-- ----------------------------
DROP TABLE IF EXISTS `forum_messages`;
CREATE TABLE `forum_messages`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`          int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`         varchar(45) NOT NULL COMMENT '用户名',
    `replied_user_id`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被回复用户id,用户a向用户b回复，用户b为 被回复用户id',
    `replied_username` varchar(45) NOT NULL COMMENT '被回复用户名',
    `post_id`          int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联主题id',
    `reply_id`         int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联回复id',
    `type`             varchar(50) NOT NULL DEFAULT '' COMMENT '消息类型',
    `is_read`          tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否已读，否: 0,是: 1',
    `created_at`       datetime             DEFAULT NULL COMMENT '创建时间',
    `deleted_at`       datetime             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY                `idx_user_id` (`user_id`) USING BTREE,
    KEY                `idx_replied_user_id` (`replied_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='论坛消息表';

-- ----------------------------
-- Table structure for forum_node_categories
-- ----------------------------
DROP TABLE IF EXISTS `forum_node_categories`;
CREATE TABLE `forum_node_categories`
(
    `id`                  int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`                varchar(45) NOT NULL COMMENT '分类名称',
    `parent_id`           int(11) unsigned NOT NULL DEFAULT 0 COMMENT '父节点id',
    `is_index_navigation` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否首页导航显示',
    `sort`                int(11) NOT NULL DEFAULT 0 COMMENT '显示顺序越小越靠前',
    `created_at`          datetime DEFAULT NULL COMMENT '创建时间',
    `deleted_at`          datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='论坛节点分类表';

-- ----------------------------
-- Table structure for forum_nodes
-- ----------------------------
DROP TABLE IF EXISTS `forum_nodes`;
CREATE TABLE `forum_nodes`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`             varchar(45) NOT NULL COMMENT '节点名称',
    `keyword`          varchar(45) NOT NULL COMMENT '节点关键词',
    `description`      text         DEFAULT NULL COMMENT '节点描述',
    `detail`           text         DEFAULT NULL COMMENT '节点详情',
    `img`              varchar(255) DEFAULT NULL COMMENT '节点图片',
    `parent_id`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '父节点id',
    `category_id`      int(11) unsigned NOT NULL DEFAULT 0 COMMENT '节点分类id',
    `is_index`         tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否首页显示',
    `is_virtual`       tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是虚拟节点，如今日最热，全部等节点，不是真实的节点',
    `is_disabled_edit` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用编辑和删除,1是 0否',
    `sort`             int(11) NOT NULL DEFAULT 0 COMMENT '显示顺序越小越靠前',
    `created_at`       datetime     DEFAULT NULL COMMENT '创建时间',
    `deleted_at`       datetime     DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=237 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='论坛节点表';

-- ----------------------------
-- Table structure for forum_posts
-- ----------------------------
DROP TABLE IF EXISTS `forum_posts`;
CREATE TABLE `forum_posts`
(
    `id`                  int(11) unsigned NOT NULL AUTO_INCREMENT,
    `node_id`             int(11) unsigned NOT NULL DEFAULT 0 COMMENT '节点id',
    `user_id`             int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`            varchar(45)  NOT NULL COMMENT '用户名',
    `title`               varchar(255) NOT NULL COMMENT '标题',
    `content_type`        char(20) DEFAULT 'general' COMMENT 'quill为使用quill编辑器，markdown为使用markdown编辑器,general 为普通文本',
    `content`             longtext DEFAULT NULL COMMENT '内容',
    `html_content`        longtext DEFAULT NULL COMMENT 'html内容',
    `top_end_time`        datetime DEFAULT NULL COMMENT '置顶截止时间,为空说明没有置顶',
    `character_amount`    int(11) unsigned NOT NULL DEFAULT 0 COMMENT '字符长度',
    `visit_amount`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '访问次数',
    `collection_amount`   int(11) unsigned NOT NULL DEFAULT 0 COMMENT '收藏次数',
    `reply_amount`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '回复次数',
    `thanks_amount`       int(11) unsigned NOT NULL DEFAULT 0 COMMENT '感谢次数',
    `shielded_amount`     int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被屏蔽次数',
    `status`              tinyint(1) DEFAULT 0 COMMENT '状态：0 未审核 1 已审核',
    `weight`              int(11) NOT NULL DEFAULT 0 COMMENT '权重',
    `reply_last_user_id`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '最后回复用户id',
    `reply_last_username` varchar(45)  NOT NULL COMMENT '最后回复用户名',
    `last_change_time`    datetime DEFAULT NULL COMMENT '主题最后变动时间',
    `created_at`          datetime DEFAULT NULL COMMENT '主题创建时间',
    `updated_at`          datetime DEFAULT NULL COMMENT '主题更新时间',
    `deleted_at`          datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY                   `idx_node_id` (`node_id`) USING BTREE,
    KEY                   `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='论坛主题表';

-- ----------------------------
-- Table structure for forum_prompts
-- ----------------------------
DROP TABLE IF EXISTS `forum_prompts`;
CREATE TABLE `forum_prompts`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `position`    varchar(50)  NOT NULL COMMENT '提示语位置',
    `content`     text         NOT NULL COMMENT '提示语内容',
    `description` varchar(255) NOT NULL COMMENT '简介',
    `is_disabled` tinyint(1) DEFAULT 0 COMMENT '状态：0 正常 1禁用',
    `created_at`  datetime DEFAULT NULL COMMENT '创建时间',
    `update_at`   datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='提示语内容';

-- ----------------------------
-- Table structure for forum_replies
-- ----------------------------
DROP TABLE IF EXISTS `forum_replies`;
CREATE TABLE `forum_replies`
(
    `id`                int(11) unsigned NOT NULL AUTO_INCREMENT,
    `posts_id`          int(11) unsigned NOT NULL DEFAULT 0 COMMENT '主题id',
    `user_id`           int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`          varchar(45)  NOT NULL COMMENT '用户名',
    `relation_user_ids` varchar(255) NOT NULL DEFAULT '' COMMENT '涉及用户ids，多个以逗号分隔',
    `content`           longtext              DEFAULT NULL COMMENT '内容',
    `character_amount`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '字符长度',
    `thanks_amount`     int(11) unsigned NOT NULL DEFAULT 0 COMMENT '感谢次数',
    `shielded_amount`   int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被屏蔽次数',
    `status`            tinyint(1) DEFAULT 0 COMMENT '状态：0 未审核 1 已审核',
    `created_at`        datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at`        datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at`        datetime              DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY                 `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='论坛回复表';

-- ----------------------------
-- Table structure for forum_sensitive_words
-- ----------------------------
DROP TABLE IF EXISTS `forum_sensitive_words`;
CREATE TABLE `forum_sensitive_words`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `type`       varchar(20)  NOT NULL COMMENT '类型',
    `word`       varchar(255) NOT NULL COMMENT '敏感词',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `update_at`  datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_word` (`word`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='敏感词表';

-- ----------------------------
-- Table structure for forum_user_posts_histories
-- ----------------------------
DROP TABLE IF EXISTS `forum_user_posts_histories`;
CREATE TABLE `forum_user_posts_histories`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`   varchar(45) NOT NULL COMMENT '用户名',
    `posts_id`   int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联主题id',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `update_at`  datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户访问主题历史表 ';

-- ----------------------------
-- Table structure for forum_users
-- ----------------------------
DROP TABLE IF EXISTS `forum_users`;
CREATE TABLE `forum_users`
(
    `id`                     int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`               varchar(45)  NOT NULL COMMENT '用户名',
    `email`                  varchar(50)  NOT NULL DEFAULT '' COMMENT 'email',
    `description`            varchar(255) NOT NULL DEFAULT '' COMMENT '简介',
    `password`               char(32)     NOT NULL COMMENT 'MD5密码',
    `avatar`                 varchar(200)          DEFAULT NULL COMMENT '头像地址',
    `status`                 int(11) DEFAULT 0 COMMENT '二进制位,0111 由低到高分别代表 禁止登录，禁止发帖，禁止回复，尚未激活',
    `posts_amount`           int(11) unsigned NOT NULL DEFAULT 0 COMMENT '创建主题次数',
    `reply_amount`           int(11) unsigned NOT NULL DEFAULT 0 COMMENT '回复次数',
    `shielded_amount`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被屏蔽次数',
    `follow_by_other_amount` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被关注次数',
    `today_activity`         int(11) unsigned NOT NULL DEFAULT 0 COMMENT '今日活跃度',
    `balance`                bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '余额',
    `site`                   varchar(200)          DEFAULT NULL COMMENT '个人站点',
    `company`                varchar(200)          DEFAULT NULL COMMENT '所在公司',
    `job`                    varchar(200)          DEFAULT NULL COMMENT '工作职位',
    `location`               varchar(200)          DEFAULT NULL COMMENT '所在地',
    `signature`              varchar(200)          DEFAULT NULL COMMENT '个人签名',
    `introduction`           varchar(500)          DEFAULT NULL COMMENT '个人简介',
    `remark`                 varchar(500)          DEFAULT NULL COMMENT '备注',
    `last_login_ip`          varchar(50)           DEFAULT NULL COMMENT '最后登陆IP',
    `last_login_time`        datetime              DEFAULT NULL COMMENT '最后登陆时间',
    `created_at`             datetime              DEFAULT NULL COMMENT '注册时间',
    `updated_at`             datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at`             datetime              DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_username` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='论坛用户表';

-- ----------------------------
-- Table structure for ga_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `ga_admin_log`;
CREATE TABLE `ga_admin_log`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
    `path`             varchar(255) NOT NULL DEFAULT '' COMMENT '请求路径',
    `method`           varchar(10)  NOT NULL DEFAULT '' COMMENT '请求方法',
    `path_name`        varchar(255) NOT NULL DEFAULT '' COMMENT '请求路径名称',
    `params`           text                  DEFAULT NULL COMMENT '请求参数',
    `response`         longtext              DEFAULT NULL COMMENT '响应结果',
    `created_at`       datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at`       datetime              DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for ga_admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `ga_admin_menu`;
CREATE TABLE `ga_admin_menu`
(
    `id`                   int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`                 varchar(20)  NOT NULL COMMENT '菜单名称',
    `path`                 varchar(100) NOT NULL DEFAULT '' COMMENT '前端路由地址，可以是外链',
    `parent_id`            int(11) unsigned NOT NULL DEFAULT 0 COMMENT '父id',
    `identification`       varchar(40)  NOT NULL DEFAULT '' COMMENT '后端权限标识符',
    `method`               varchar(10)  NOT NULL DEFAULT '' COMMENT '请求方法',
    `front_component_path` varchar(255)          DEFAULT NULL COMMENT '前端组件路径',
    `icon`                 varchar(100)          DEFAULT '#' COMMENT '菜单图标',
    `sort`                 tinyint(4) NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
    `status`               varchar(10)           DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `created_at`           datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at`           datetime              DEFAULT NULL COMMENT '更新时间',
    `type`                 varchar(12)  NOT NULL COMMENT '菜单类型',
    `link_type`            varchar(12)  NOT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=93 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='后台菜单表';

-- ----------------------------
-- Table structure for ga_administrator
-- ----------------------------
DROP TABLE IF EXISTS `ga_administrator`;
CREATE TABLE `ga_administrator`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`        varchar(45) NOT NULL COMMENT '用户名',
    `password`        char(32)    NOT NULL COMMENT 'MD5密码',
    `nickname`        varchar(45)  DEFAULT NULL COMMENT '昵称',
    `avatar`          varchar(200) DEFAULT NULL COMMENT '头像地址',
    `status`          varchar(10)  DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `remark`          varchar(500) DEFAULT NULL COMMENT '备注',
    `last_login_ip`   varchar(50)  DEFAULT NULL COMMENT '最后登陆IP',
    `last_login_date` datetime     DEFAULT NULL COMMENT '最后登陆时间',
    `created_at`      datetime     DEFAULT NULL COMMENT '注册时间',
    `updated_at`      datetime     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='管理员表';

-- ----------------------------
-- Table structure for ga_administrator_role
-- ----------------------------
DROP TABLE IF EXISTS `ga_administrator_role`;
CREATE TABLE `ga_administrator_role`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
    `role_id`          int(11) unsigned NOT NULL COMMENT '角色id',
    `created_at`       datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at`       datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='管理员角色关联表';

-- ----------------------------
-- Table structure for ga_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `ga_casbin_rule`;
CREATE TABLE `ga_casbin_rule`
(
    `ptype` varchar(10)  DEFAULT NULL,
    `v0`    varchar(256) DEFAULT NULL,
    `v1`    varchar(256) DEFAULT NULL,
    `v2`    varchar(256) DEFAULT NULL,
    `v3`    varchar(256) DEFAULT NULL,
    `v4`    varchar(256) DEFAULT NULL,
    `v5`    varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for ga_config
-- ----------------------------
DROP TABLE IF EXISTS `ga_config`;
CREATE TABLE `ga_config`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `module`     varchar(255) NOT NULL DEFAULT '' COMMENT '所属模块',
    `key`        varchar(255) NOT NULL DEFAULT '' COMMENT '键值',
    `value`      text                  DEFAULT NULL COMMENT '值',
    `created_at` datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime              DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `module_key_idx` (`module`,`key`)
) ENGINE=InnoDB AUTO_INCREMENT=1324 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='配置表';

-- ----------------------------
-- Table structure for ga_login_log
-- ----------------------------
DROP TABLE IF EXISTS `ga_login_log`;
CREATE TABLE `ga_login_log`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
    `ip`               varchar(30)  NOT NULL DEFAULT '' COMMENT 'ip地址',
    `browser`          varchar(10)  NOT NULL DEFAULT '' COMMENT '浏览器',
    `os`               varchar(255) NOT NULL DEFAULT '' COMMENT '操作系统',
    `status`           varchar(10)           DEFAULT 'normal' COMMENT '状态 success 成功 fail 失败',
    `created_at`       datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at`       datetime              DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for ga_role
-- ----------------------------
DROP TABLE IF EXISTS `ga_role`;
CREATE TABLE `ga_role`
(
    `id`             int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`           varchar(45) NOT NULL COMMENT '角色名',
    `identification` varchar(20) NOT NULL COMMENT '角色标识符',
    `sort`           tinyint(4) NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
    `status`         varchar(10) DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `created_at`     datetime    DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime    DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色表';

-- ----------------------------
-- Table structure for ga_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `ga_role_menu`;
CREATE TABLE `ga_role_menu`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `menu_id`    int(11) unsigned NOT NULL COMMENT '管理员id',
    `role_id`    int(11) unsigned NOT NULL COMMENT '角色id',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=941 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色和菜单权限关联表';

SET
FOREIGN_KEY_CHECKS = 1;
