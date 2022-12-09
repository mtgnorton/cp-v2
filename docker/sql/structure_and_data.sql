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

 Date: 06/12/2022 23:12:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for forum_association
-- ----------------------------
DROP TABLE IF EXISTS `forum_association`;
CREATE TABLE `forum_association`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`         int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`        varchar(45)      NOT NULL COMMENT '用户名',
    `target_id`       int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被感谢｜屏蔽|收藏 主题id|回复id',
    `additional_id`   int(11) unsigned NOT NULL DEFAULT 0 COMMENT '附加id，当target_id为回复id时，additional_id为主题id',
    `target_user_id`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被感谢｜屏蔽|收藏 用户id',
    `target_username` varchar(45)      NOT NULL COMMENT '被感谢用户名',
    `type`            char(15)         NOT NULL DEFAULT '' COMMENT '类型 感谢主题: thanks_post,感谢回复: thanks_reply,屏蔽主题: shield_post,屏蔽回复: shield_reply,收藏主题:collect_post,收藏节点: collect_node,关注用户： follow_user,屏蔽用户: shield_user',
    `created_at`      datetime                  DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_target_user_id` (`target_user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛 感谢｜屏蔽|收藏| 关注    主题｜回复 |节点| 用户 关联表';

-- ----------------------------
-- Records of forum_association
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for forum_balance_change_log
-- ----------------------------
DROP TABLE IF EXISTS `forum_balance_change_log`;
CREATE TABLE `forum_balance_change_log`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`    varchar(45)      NOT NULL COMMENT '用户名',
    `type`        char(50)         NOT NULL DEFAULT '' COMMENT '每日登录奖励:login, 每日活跃度奖励: activity, 感谢主题: thanks_posts,感谢回复: thanks_relpy,创建主题: create_posts,创建回复: create_reply,初始奖励: register',
    `amount`      int(11)          NOT NULL DEFAULT 0 COMMENT '金额',
    `before`      int(11) unsigned NOT NULL DEFAULT 0 COMMENT '变动前余额',
    `after`       int(11) unsigned NOT NULL DEFAULT 0 COMMENT '变动后余额',
    `relation_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联主题id或关联回复id',
    `remark`      varchar(255)     NOT NULL DEFAULT '' COMMENT '备注',
    `created_at`  datetime                  DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛余额变动表';

-- ----------------------------
-- Records of forum_balance_change_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for forum_email_records
-- ----------------------------
DROP TABLE IF EXISTS `forum_email_records`;
CREATE TABLE `forum_email_records`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`   varchar(45)      NOT NULL COMMENT '用户名',
    `type`       varchar(50)      NOT NULL COMMENT '类型',
    `email`      varchar(255)     NOT NULL COMMENT '邮箱',
    `title`      varchar(255)     NOT NULL COMMENT '标题',
    `content`    text             NOT NULL COMMENT '内容',
    `error`      text             NOT NULL COMMENT '错误信息',
    `created_at` datetime                  DEFAULT NULL COMMENT '创建时间',
    `update_at`  datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='邮件发送历史表';

-- ----------------------------
-- Records of forum_email_records
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for forum_follow_or_shield_user_relation
-- ----------------------------
DROP TABLE IF EXISTS `forum_follow_or_shield_user_relation`;
CREATE TABLE `forum_follow_or_shield_user_relation`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`         int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`        varchar(45)      NOT NULL COMMENT '用户名',
    `target_user_id`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被关注｜屏蔽用户id',
    `target_username` varchar(45)      NOT NULL COMMENT '被关注｜屏蔽用户名',
    `type`            char(10)         NOT NULL DEFAULT '' COMMENT '类型 关注: follow,屏蔽: shield',
    `created_at`      datetime                  DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_target_user_id` (`target_user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛 关注|屏蔽  用户关联表';

-- ----------------------------
-- Records of forum_follow_or_shield_user_relation
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for forum_messages
-- ----------------------------
DROP TABLE IF EXISTS `forum_messages`;
CREATE TABLE `forum_messages`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`          int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`         varchar(45)      NOT NULL COMMENT '用户名',
    `replied_user_id`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被回复用户id,用户a向用户b回复，用户b为 被回复用户id',
    `replied_username` varchar(45)      NOT NULL COMMENT '被回复用户名',
    `post_id`          int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联主题id',
    `reply_id`         int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联回复id',
    `type`             varchar(50)      NOT NULL DEFAULT '' COMMENT '消息类型',
    `is_read`          tinyint(1)       NOT NULL DEFAULT 0 COMMENT '是否已读，否: 0,是: 1',
    `created_at`       datetime                  DEFAULT NULL COMMENT '创建时间',
    `deleted_at`       datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_replied_user_id` (`replied_user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛消息表';

-- ----------------------------
-- Records of forum_messages
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for forum_node_categories
-- ----------------------------
DROP TABLE IF EXISTS `forum_node_categories`;
CREATE TABLE `forum_node_categories`
(
    `id`                  int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`                varchar(45)      NOT NULL COMMENT '分类名称',
    `parent_id`           int(11) unsigned NOT NULL DEFAULT 0 COMMENT '父节点id',
    `is_index_navigation` tinyint(1)       NOT NULL DEFAULT 1 COMMENT '是否首页导航显示',
    `sort`                int(11)          NOT NULL DEFAULT 0 COMMENT '显示顺序越小越靠前',
    `created_at`          datetime                  DEFAULT NULL COMMENT '创建时间',
    `deleted_at`          datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 14
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛节点分类表';

-- ----------------------------
-- Records of forum_node_categories
-- ----------------------------
BEGIN;
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (1, '分享与探索', 0, 1, 0, '2022-11-30 23:18:01', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (2, 'V2EX', 0, 1, 0, '2022-11-30 23:18:13', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (3, 'Apple', 0, 1, 0, '2022-11-30 23:18:16', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (4, '前端开发', 0, 1, 0, '2022-11-30 23:18:53', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (5, '编程语言', 0, 1, 0, '2022-11-30 23:19:11', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (6, '后端架构', 0, 1, 0, '2022-11-30 23:19:36', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (7, '机器学习', 0, 1, 0, '2022-11-30 23:20:03', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (8, 'iOS', 0, 1, 0, '2022-11-30 23:20:16', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (9, 'Geek', 0, 1, 0, '2022-11-30 23:20:22', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (10, '游戏', 0, 1, 0, '2022-11-30 23:23:09', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (11, '生活', 0, 1, 0, '2022-11-30 23:23:30', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (12, 'Internet', 0, 1, 0, '2022-11-30 23:23:59', NULL);
INSERT INTO `forum_node_categories` (`id`, `name`, `parent_id`, `is_index_navigation`, `sort`, `created_at`,
                                     `deleted_at`)
VALUES (13, '城市', 0, 1, 0, '2022-11-30 23:24:06', NULL);
COMMIT;

-- ----------------------------
-- Table structure for forum_nodes
-- ----------------------------
DROP TABLE IF EXISTS `forum_nodes`;
CREATE TABLE `forum_nodes`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`             varchar(45)      NOT NULL COMMENT '节点名称',
    `keyword`          varchar(45)      NOT NULL COMMENT '节点关键词',
    `description`      text                      DEFAULT NULL COMMENT '节点描述',
    `detail`           text                      DEFAULT NULL COMMENT '节点详情',
    `img`              varchar(255)              DEFAULT NULL COMMENT '节点图片',
    `parent_id`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '父节点id',
    `category_id`      int(11) unsigned NOT NULL DEFAULT 0 COMMENT '节点分类id',
    `is_index`         tinyint(1)       NOT NULL DEFAULT 0 COMMENT '是否首页显示',
    `is_virtual`       tinyint(1)       NOT NULL DEFAULT 0 COMMENT '是否是虚拟节点，如今日最热，全部等节点，不是真实的节点',
    `is_disabled_edit` tinyint(1)       NOT NULL DEFAULT 0 COMMENT '是否禁用编辑和删除,1是 0否',
    `sort`             int(11)          NOT NULL DEFAULT 0 COMMENT '显示顺序越小越靠前',
    `created_at`       datetime                  DEFAULT NULL COMMENT '创建时间',
    `deleted_at`       datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 237
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛节点表';

-- ----------------------------
-- Records of forum_nodes
-- ----------------------------
BEGIN;
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (1, '最热', 'today-hot', '', '', '/upload/backend//co38rnfd6a540q9ywj.png', 0, 0, 1, 1, 1, 0,
        '2022-11-04 12:54:27', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (2, '七日最热', 'day7-hot', '', '', '/upload/backend//co38srfxut1knxndzc.png', 0, 0, 1, 1, 1, 0,
        '2022-11-04 12:55:53', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (3, '全部', 'all', '', '', '', 0, 0, 1, 1, 1, 0, '2022-11-04 12:56:11', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (4, '关注', 'follow-user', '', '', '', 0, 0, 1, 1, 1, 0, '2022-11-04 12:57:30', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (5, '节点', 'follow-node', '', '', '', 0, 0, 1, 1, 1, 0, '2022-11-04 14:25:51', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (6, '分享与探索', '分享与探索', '', NULL, '', 0, 1, 0, 0, 0, 0, '2022-11-30 23:18:01', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (7, '问与答', 'qna', '一个更好的世界需要你持续地提出好问题。', NULL, '/upload/backend/node/copqbbex98o8fjxhao.png', 6, 1, 0, 0, 0, 0,
        '2022-11-30 23:18:03', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (8, '分享发现', 'share', '分享你看到的好玩的，有信息量的，欢迎从这里获取灵感。', NULL, '/upload/backend/node/copqbc4vtzbkdxlybw.png', 6, 1, 0,
        0, 0, 0, '2022-11-30 23:18:04', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (9, '分享创造', 'create', '欢迎你在这里发布自己的最新作品！', NULL, '/upload/backend/node/copqbcv4sppchocub7.png', 6, 1, 0, 0, 0, 0,
        '2022-11-30 23:18:06', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (10, '奇思妙想', 'ideas', '让你的创意在这里自由流动吧。', NULL, '/upload/backend/node/copqbdmhogc8z6bfm4.png', 6, 1, 0, 0, 0, 0,
        '2022-11-30 23:18:08', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (11, '分享邀请码', 'in', '这里分享各类新酷网站的邀请码。', NULL, '/upload/backend/node/copqbegs6wcoadysgv.png', 6, 1, 0, 0, 0, 0,
        '2022-11-30 23:18:09', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (12, '自言自语', 'autistic', '&nbsp;', NULL, '/upload/backend/node/copqbf52cwvkvnazch.png', 6, 1, 0, 0, 0, 0,
        '2022-11-30 23:18:11', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (13, '设计', 'design',
        '<strong>Beautiful</strong> <code>adj.</code> <em>Pleasing the senses or mind aesthetically.</em>', NULL,
        '/upload/backend/node/copqbflrvaqwnmckbb.png', 6, 1, 0, 0, 0, 0, '2022-11-30 23:18:12', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (14, 'Blog', 'blog', '', NULL, '', 6, 1, 0, 0, 0, 0, '2022-11-30 23:18:13', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (15, 'V2EX', 'V2EX', '', NULL, '', 0, 2, 0, 0, 0, 0, '2022-11-30 23:18:13', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (16, '反馈', 'feedback', '欢迎在这里提出你对 V2EX 的任何疑问和建议', NULL, '/upload/backend/node/copqbgkbuo4wlkjw9v.png', 15, 2, 0,
        0, 0, 0, '2022-11-30 23:18:14', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (17, 'Project Babel', 'babel', '', NULL, '/upload/backend/node/copqbh2lfnrcmsycyv.png', 15, 2, 0, 0, 0, 0,
        '2022-11-30 23:18:15', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (18, '使用指南', 'guide', '这里有关于使用 V2EX 的各种技巧和提示。', NULL, '/upload/backend/node/copqbhisx27ccpakt5.png', 15, 2, 0, 0,
        0, 0, '2022-11-30 23:18:16', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (19, 'Apple', 'Apple', '', '', '/upload/backend/node/copqcfxiuhooipztrw.png', 0, 3, 1, 0, 0, 0,
        '2022-11-30 23:18:16', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (20, 'macOS', 'macos', 'The world’s most advanced desktop operating system.', NULL,
        '/upload/backend/node/copqbi7s1z2glhvx7y.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:18', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (21, 'iPhone', 'iphone', 'Say hello to the future.', '', '/upload/backend/node/copqbirnftyg8gqpcu.png', 19, 3, 0,
        0, 0, 0, '2022-11-30 23:18:19', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (22, 'MacBook Pro', 'mbp',
        'State-of-the-art processors. All-new graphics. Breakthrough high-speed I/O. Three very big leaps forward.',
        NULL, '/upload/backend/node/copqbj9o71rkygctzq.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:20', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (23, 'iOS', 'ios', '', NULL, '', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:20', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (24, 'iPad', 'ipad', '<a href=\"/go/apple\">Apple</a> 公司设计的全新理念的基于 <a href=\"/go/ios\">iOS</a> 的平板电脑。', NULL,
        '/upload/backend/node/copqbkcpfzxct6jkje.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:22', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (25, 'MacBook', 'macbook', '', NULL, '/upload/backend/node/copqbkucce68jfirz6.png', 19, 3, 0, 0, 0, 0,
        '2022-11-30 23:18:23', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (26, ' WATCH', 'watch', '<a href=\"/go/apple\">Apple</a> 公司设计的智能手表产品', NULL,
        '/upload/backend/node/copqblk0jb9ccdtzo9.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:25', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (27, '配件', 'accessory', '关于各类新酷配件的购买和使用的讨论。', NULL, '/upload/backend/node/copqbm68ap88s4tfyo.png', 19, 3, 0, 0,
        0, 0, '2022-11-30 23:18:26', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (28, 'MacBook Air', 'mba', 'Apple 设计的最轻巧的笔记本电脑', NULL, '/upload/backend/node/copqbn1f6c00ihtxkw.png', 19, 3, 0,
        0, 0, 0, '2022-11-30 23:18:28', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (29, 'iMac', 'imac', 'All-in-one design.', NULL, '/upload/backend/node/copqbnpr707smt82u5.png', 19, 3, 0, 0, 0,
        0, '2022-11-30 23:18:30', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (30, 'Mac mini', 'macmini', '<a href=\"/go/apple\">Apple</a> 公司设计的世界上最轻巧的桌面主机。', NULL,
        '/upload/backend/node/copqbohbmzs0kba66g.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:31', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (31, 'Xcode', 'xcode', '', NULL, '/upload/backend/node/copqbpe7puwoqak65z.png', 19, 3, 0, 0, 0, 0,
        '2022-11-30 23:18:33', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (32, 'Mac Pro', 'macpro', 'Built for creativity on an epic scale.', NULL,
        '/upload/backend/node/copqbqu5dgzc4i5gpd.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:36', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (33, 'WWDC', 'wwdc', 'Apple Worldwide Developers Conference', NULL,
        '/upload/backend/node/copqbs9t3bfc4ylo4l.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:39', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (34, 'iPod', 'ipod', '<a href=\"/go/apple\">Apple</a> 公司设计的便携式媒体播放器。', NULL,
        '/upload/backend/node/copqbsy8atwgehcgkw.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:41', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (35, 'AirPods', 'airpods', '<a href=\"/go/apple\">Apple</a> 公司设计的无线耳机', NULL,
        '/upload/backend/node/copqbu90m6l44j6hvs.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:44', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (36, 'Mac Studio', 'macstudio', 'Mac Studio is an entirely new Mac desktop.', NULL,
        '/upload/backend/node/copqbuptr5ugwendtr.png', 19, 3, 0, 0, 0, 0, '2022-11-30 23:18:45', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (37, 'iWork', 'iwork', '', NULL, '/upload/backend/node/copqbvupsxb4xjxzb1.png', 19, 3, 0, 0, 0, 0,
        '2022-11-30 23:18:47', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (38, 'MobileMe', 'mobileme', '', NULL, '/upload/backend/node/copqbwcx9m0ogfuwuf.png', 19, 3, 0, 0, 0, 0,
        '2022-11-30 23:18:48', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (39, 'iLife', 'ilife', '', NULL, '/upload/backend/node/copqbx8os65s9ioo1w.png', 19, 3, 0, 0, 0, 0,
        '2022-11-30 23:18:50', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (40, 'HomeKit', 'homekit', '', NULL, '/upload/backend/node/copqbxobq2oww1sd6f.png', 19, 3, 0, 0, 0, 0,
        '2022-11-30 23:18:51', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (41, 'GarageBand', 'garageband', '', NULL, '/upload/backend/node/copqby21ihb4j18euk.png', 19, 3, 0, 0, 0, 0,
        '2022-11-30 23:18:52', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (42, 'iMovie', 'imovie', 'macOS / iOS 上简单好用的视频编辑软件', NULL, '/upload/backend/node/copqbyg0bk74c0v5gm.png', 19, 3,
        0, 0, 0, 0, '2022-11-30 23:18:53', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (43, '前端开发', '前端开发', '', NULL, '', 0, 4, 0, 0, 0, 0, '2022-11-30 23:18:53', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (44, '微信', 'wechat', '关于微信及微信小程序的讨论节点', NULL, '/upload/backend/node/copqbywwa76g46xjft.png', 43, 4, 0, 0, 0, 0,
        '2022-11-30 23:18:54', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (45, '前端开发', 'fe', '', NULL, '/upload/backend/node/copqbzb8egu8t4erg6.png', 43, 4, 0, 0, 0, 0,
        '2022-11-30 23:18:55', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (46, 'Chrome', 'chrome',
        'Google Chrome is a browser that combines a minimal design with sophisticated technology to make the web faster, safer, and easier.',
        NULL, '', 43, 4, 0, 0, 0, 0, '2022-11-30 23:18:56', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (47, 'Vue.js', 'vue', 'The Progressive JavaScript Framework', NULL,
        '/upload/backend/node/copqc09g7b6wccxrct.png', 43, 4, 0, 0, 0, 0, '2022-11-30 23:18:57', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (48, 'React', 'react', '一个用于构建用户界面的 JavaScript 库', NULL, '/upload/backend/node/copqc0obojwgdaifbb.png', 43, 4, 0,
        0, 0, 0, '2022-11-30 23:18:58', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (49, '浏览器', 'browsers', '', NULL, '/upload/backend/node/copqc12y1hxszkw8vc.png', 43, 4, 0, 0, 0, 0,
        '2022-11-30 23:18:59', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (50, 'CSS', 'css', 'Cascading Style Sheet，层叠样式表，网页外观设计的标准技术。', NULL,
        '/upload/backend/node/copqc1hov68obubwnv.png', 43, 4, 0, 0, 0, 0, '2022-11-30 23:19:00', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (51, 'Firefox', 'firefox', 'Mozilla Firefox is a free and open source web browser.', NULL,
        '/upload/backend/node/copqc2b89eqg82ur71.png', 43, 4, 0, 0, 0, 0, '2022-11-30 23:19:01', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (52, 'Flutter', 'flutter',
        '来自 <a href=\"/go/google\">Google</a> 的同时支持 <a href=\"/go/android\">Android</a> 和 <a href=\"/go/idev\">iOS</a> 的移动应用 UI 框架。',
        NULL, '/upload/backend/node/copqc2rny4copzzhtj.png', 43, 4, 0, 0, 0, 0, '2022-11-30 23:19:02', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (53, 'Edge', 'edge', '基于 Chromium 的全新 Microsoft Edge', NULL, '/upload/backend/node/copqc38bkxco7tb73a.png', 43,
        4, 0, 0, 0, 0, '2022-11-30 23:19:03', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (54, 'Angular', 'angular', 'Superheroic JavaScript MVW Framework', NULL,
        '/upload/backend/node/copqc3nq9m1cuggrvh.png', 43, 4, 0, 0, 0, 0, '2022-11-30 23:19:04', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (55, 'Electron', 'electron', 'Build cross platform desktop apps with JavaScript, HTML, and CSS.', NULL,
        '/upload/backend/node/copqc4tzcz88t3bbyp.png', 43, 4, 0, 0, 0, 0, '2022-11-30 23:19:07', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (56, 'Web Dev', 'webdev', '讨论 W3C 的各项标准的细节和实现。', NULL, '/upload/backend/node/copqc5a9heywqazkdx.png', 43, 4, 0,
        0, 0, 0, '2022-11-30 23:19:08', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (57, 'Ionic', 'ionic', 'The beautiful, open source front-end SDK for developing hybrid mobile apps with HTML5.',
        NULL, '/upload/backend/node/copqc5z1rzs0o2liwx.png', 43, 4, 0, 0, 0, 0, '2022-11-30 23:19:09', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (58, 'Next.js', 'nextjs', '', NULL, '', 43, 4, 0, 0, 0, 0, '2022-11-30 23:19:10', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (59, 'Vite', 'vite', '下一代前端开发与构建工具', NULL, '/upload/backend/node/copqc6m5w85sswwkr4.png', 43, 4, 0, 0, 0, 0,
        '2022-11-30 23:19:11', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (60, 'Nuxt.js', 'nuxtjs', '', NULL, '', 43, 4, 0, 0, 0, 0, '2022-11-30 23:19:11', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (61, '编程语言', '编程语言', '', NULL, '', 0, 5, 0, 0, 0, 0, '2022-11-30 23:19:11', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (62, 'Python', 'python', '这里讨论各种 Python 语言编程话题，也包括 Django，Tornado 等框架的讨论。这里是一个能够帮助你解决实际问题的地方。', NULL,
        '/upload/backend/node/copqc7ebs85srace7m.png', 61, 5, 0, 0, 0, 0, '2022-11-30 23:19:12', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (63, 'Java', 'java', 'Sun 公司发明，被广泛使用的一门编程语言。', NULL, '/upload/backend/node/copqc7yju1egnevxjj.png', 61, 5, 0, 0,
        0, 0, '2022-11-30 23:19:14', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (64, 'PHP', 'php', 'PHP 是一门被广泛使用的编程语言，尤其是在各类互联网站项目中。PHP 代码可以被很容易地嵌入到 HTML 中。', NULL,
        '/upload/backend/node/copqc9ldlskojg66qb.png', 61, 5, 0, 0, 0, 0, '2022-11-30 23:19:17', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (65, 'JavaScript', 'js',
        'JavaScript (sometimes abbreviated JS) is a prototype-based scripting language that is dynamic, weakly typed and has first-class functions.',
        NULL, '/upload/backend/node/copqcbn3luzsfjvhtp.png', 61, 5, 0, 0, 0, 0, '2022-11-30 23:19:22', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (66, 'Go 编程语言', 'go', 'Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。', NULL, '/upload/backend/node/copqcc49bx7k7rdqmx.png',
        61, 5, 0, 0, 0, 0, '2022-11-30 23:19:23', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (67, 'Node.js', 'nodejs',
        'Node.js is a platform built on <a href=\"http://code.google.com/p/v8/\" target=\"_blank\">Chrome\'s JavaScript runtime</a> for easily building fast, scalable network applications.',
        NULL, '/upload/backend/node/copqcck1dylk5ogxra.png', 61, 5, 0, 0, 0, 0, '2022-11-30 23:19:24', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (68, 'C++', 'cpp', '关于编程语言 C++ 的使用讨论', NULL, '/upload/backend/node/copqccyjjv4whydldv.png', 61, 5, 0, 0, 0, 0,
        '2022-11-30 23:19:25', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (69, 'HTML', 'html', '超文本标记语言 HyperText Markup Language', NULL, '/upload/backend/node/copqcdghu5k0syisbz.png',
        61, 5, 0, 0, 0, 0, '2022-11-30 23:19:26', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (70, 'Swift', 'swift', '来自 Apple 的类型安全的编程语言。', NULL, '/upload/backend/node/copqcduckkco0vbeua.png', 61, 5, 0, 0,
        0, 0, '2022-11-30 23:19:26', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (71, '.NET', 'dotnet', '', NULL, '/upload/backend/node/copqceajifv4duacv4.png', 61, 5, 0, 0, 0, 0,
        '2022-11-30 23:19:27', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (72, 'Rust', 'rust',
        'Rust is a systems programming language that runs blazingly fast, prevents almost all crashes, and eliminates data races.',
        NULL, '/upload/backend/node/copqceswh67cdpucku.png', 61, 5, 0, 0, 0, 0, '2022-11-30 23:19:29', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (73, 'Ruby on Rails', 'ror', 'Full stack web application framework.', NULL,
        '/upload/backend/node/copqcf8t7dls15y5ur.png', 61, 5, 0, 0, 0, 0, '2022-11-30 23:19:29', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (74, 'C#', 'csharp', '', NULL, '/upload/backend/node/copqcg56rapcpl9lmc.png', 61, 5, 0, 0, 0, 0,
        '2022-11-30 23:19:31', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (75, 'Ruby', 'ruby',
        'A dynamic, interpreted, open source programming language with a focus on simplicity and productivity.', NULL,
        '/upload/backend/node/copqcgm9r5fcfkt6ls.png', 61, 5, 0, 0, 0, 0, '2022-11-30 23:19:32', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (76, 'TypeScript', 'typescript', 'JavaScript that scales.', NULL, '/upload/backend/node/copqch0r7b9k0pinmf.png',
        61, 5, 0, 0, 0, 0, '2022-11-30 23:19:33', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (77, 'Kotlin', 'kotlin', '', NULL, '/upload/backend/node/copqchnin53cheiyvb.png', 61, 5, 0, 0, 0, 0,
        '2022-11-30 23:19:35', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (78, 'Lua', 'lua', 'Lua is a powerful, fast, lightweight, embeddable scripting language.', NULL,
        '/upload/backend/node/copqcibslpvckwlysc.png', 61, 5, 0, 0, 0, 0, '2022-11-30 23:19:36', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (79, '后端架构', '后端架构', '', NULL, '', 0, 6, 0, 0, 0, 0, '2022-11-30 23:19:36', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (80, '云计算', 'cloud', '关于云计算技术和平台的综合讨论区。', NULL, '/upload/backend/node/copqcj3zz9c8e6gx1a.png', 79, 6, 0, 0, 0, 0,
        '2022-11-30 23:19:38', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (81, '服务器', 'server', '关于服务器选择和使用的技术讨论。', NULL, '/upload/backend/node/copqcjrjroognvlrgi.png', 79, 6, 0, 0, 0, 0,
        '2022-11-30 23:19:39', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (82, 'DNS', 'dns', '互联网基础协议——DNS', NULL, '/upload/backend/node/copqck7x1fw8zz17lq.png', 79, 6, 0, 0, 0, 0,
        '2022-11-30 23:19:40', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (83, 'MySQL', 'mysql', '地球上最流行的关系数据库。被大量的互联网企业运用在网站项目中。具有非常丰富的生态系统。', NULL,
        '/upload/backend/node/copqckolpipsxpr9bn.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:41', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (84, 'NGINX', 'nginx',
        'A HTTP and mail proxy server licensed under a 2-clause BSD-like license. By Igor Sysoev.', NULL,
        '/upload/backend/node/copqclj453e0n772yn.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:43', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (85, 'Docker', 'docker', '基于 Linux Container 技术的轻量级虚拟化环境。', NULL, '/upload/backend/node/copqcm22zodsybyei9.png',
        79, 6, 0, 0, 0, 0, '2022-11-30 23:19:44', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (86, '数据库', 'db', '', NULL, '/upload/backend/node/copqcmht1js0eq4ze4.png', 79, 6, 0, 0, 0, 0,
        '2022-11-30 23:19:45', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (87, 'Ubuntu', 'ubuntu',
        'Super-fast, easy to use and free, the Ubuntu operating system powers millions of desktops, netbooks and servers around the world.',
        NULL, '/upload/backend/node/copqcmvv5tk8cqk3rd.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:46', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (88, 'Django', 'django',
        'Django is a high-level Python Web framework that encourages rapid development and clean, pragmatic design.',
        NULL, '/upload/backend/node/copqcnli8uowjkqdho.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:48', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (89, 'Amazon Web Services', 'aws',
        'Amazon Web Services (AWS) delivers a set of services that together form a reliable, scalable, and inexpensive computing platform “in the cloud”.',
        NULL, '/upload/backend/node/copqco1pa5nsscpumg.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:49', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (90, 'Kubernetes', 'k8s', '适用于大规模生产环境的容器编排管理平台。', NULL, '/upload/backend/node/copqcojjhck0060bdj.png', 79, 6, 0,
        0, 0, 0, '2022-11-30 23:19:50', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (91, 'MongoDB', 'mongodb',
        'MongoDB is a free and open-source cross-platform document-oriented database program.', NULL,
        '/upload/backend/node/copqcp4yax4gwaiya5.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:51', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (92, 'Redis', 'redis', 'Redis 是一个高性能的数据结构服务器。Redis 中的 key 可以支持多种不同的数据结构，包括：字符串，列表，集合，sort set 等等。', NULL,
        '/upload/backend/node/copqcpu6uybkby2bjo.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:53', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (93, 'DevOps', 'devops', '我们不仅写代码，我们同时也关心代码将如何运行。', NULL, '/upload/backend/node/copqcqrk7snkl3vmfk.png', 79, 6,
        0, 0, 0, 0, '2022-11-30 23:19:55', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (94, 'Flask', 'flask', 'Flask is a microframework for Python based on Werkzeug, Jinja 2 and good intentions.',
        NULL, '/upload/backend/node/copqcrh2mgtshker6p.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:56', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (95, 'Elasticsearch', 'elasticsearch', '一个基于 Lucene 构建的、Apache 协议开源的、分布式的、提供 REST API 的搜索引擎。', NULL,
        '/upload/backend/node/copqcrv9jffkawdil5.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:57', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (96, 'Tornado', 'tornado',
        'Tornado is an open source version of the scalable, non-blocking web server and tools that power FriendFeed.',
        NULL, '/upload/backend/node/copqcsd5hctcsgzxcl.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:19:58', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (97, 'API', 'api', '', NULL, '/upload/backend/node/copqcsvptrywu0lr5n.png', 79, 6, 0, 0, 0, 0,
        '2022-11-30 23:19:59', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (98, 'Cloudflare', 'cloudflare', 'Cloudflare 提供全球 CDN 和 Anti-DDoS 服务。', NULL,
        '/upload/backend/node/copqctm6z0z4vzlffy.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:20:01', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (99, 'LeanCloud', 'leancloud', '最好的移动应用开发一站式服务，让小团队也能做出大产品。', NULL,
        '/upload/backend/node/copqcu136nj4rc1cyn.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:20:02', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (100, 'Timescale', 'timescale',
        'An open-source time-series database fully compatible with Postgres for fast ingest and complex queries.', NULL,
        '/upload/backend/node/copqcut6i7lkurrily.png', 79, 6, 0, 0, 0, 0, '2022-11-30 23:20:03', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (101, '机器学习', '机器学习', '', NULL, '', 0, 7, 0, 0, 0, 0, '2022-11-30 23:20:03', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (102, '机器学习', 'ml', '机器学习是人工智能的一个分支', NULL, '/upload/backend/node/copqcvfyoei01ngwlu.png', 101, 7, 0, 0, 0, 0,
        '2022-11-30 23:20:05', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (103, '数学', 'math', '数学的乐趣', NULL, '/upload/backend/node/copqcvw32ddchzutsv.png', 101, 7, 0, 0, 0, 0,
        '2022-11-30 23:20:06', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (104, 'TensorFlow', 'tensorflow', 'An open-source software library for Machine Intelligence', NULL,
        '/upload/backend/node/copqcwtdmnc0pvquvx.png', 101, 7, 0, 0, 0, 0, '2022-11-30 23:20:08', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (105, '自然语言处理', 'nlp', '', NULL, '/upload/backend/node/copqcx7wbvuwso3eea.png', 101, 7, 0, 0, 0, 0,
        '2022-11-30 23:20:09', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (106, 'CUDA', 'cuda', '来自 NVIDIA 的并行运算框架', NULL, '/upload/backend/node/copqcy6zyjfsvn6xkb.png', 101, 7, 0, 0, 0,
        0, '2022-11-30 23:20:11', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (107, 'Torch', 'torch',
        'Torch is a scientific computing framework with wide support for machine learning algorithms that puts GPUs first. It is easy to use and efficient, thanks to an easy and fast scripting language, LuaJIT, and an underlying C/CUDA implementation.',
        NULL, '/upload/backend/node/copqcz4x1k4ogvp0yf.png', 101, 7, 0, 0, 0, 0, '2022-11-30 23:20:13', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (108, 'Keras', 'keras',
        'Keras is a high-level neural networks API, written in Python and capable of running on top of either TensorFlow or Theano.',
        NULL, '/upload/backend/node/copqczij4vuonvcpz3.png', 101, 7, 0, 0, 0, 0, '2022-11-30 23:20:14', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (109, 'Core ML', 'coreml', 'iOS 11 开发环境新功能——机器学习开发套件', NULL, '/upload/backend/node/copqd0i0yu2gp4a2gi.png', 101,
        7, 0, 0, 0, 0, '2022-11-30 23:20:16', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (110, 'iOS', 'iOS', '', NULL, '', 0, 8, 0, 0, 0, 0, '2022-11-30 23:20:16', '2022-11-30 23:25:26');
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (111, 'iDev', 'idev',
        'iOS 及 OS X 开发技术讨论区，iOS 是 <a href=\"/go/iphone\">iPhone</a> 及 <a href=\"/go/ipad\">iPad</a> 上运行的操作系统。', NULL,
        '/upload/backend/node/copqd1bguw4g1kblw1.png', 23, 8, 0, 0, 0, 0, '2022-11-30 23:20:18', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (112, 'iCode', 'icode', '', NULL, '', 23, 8, 0, 0, 0, 0, '2022-11-30 23:20:18', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (113, 'iMarketing', 'imarketing', '独立 iOS 开发者推广讨论', NULL, '/upload/backend/node/copqd20dx368pkc4hu.png', 23, 8,
        0, 0, 0, 0, '2022-11-30 23:20:19', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (114, 'iAd', 'iad', '', NULL, '/upload/backend/node/copqd2m9ubbkgmviwp.png', 23, 8, 0, 0, 0, 0,
        '2022-11-30 23:20:20', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (115, 'iTransfer', 'itransfer', '这里讨论 iOS App Transfer', NULL, '/upload/backend/node/copqd3gag320qzi18n.png', 23,
        8, 0, 0, 0, 0, '2022-11-30 23:20:22', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (116, 'Geek', 'Geek', '', NULL, '', 0, 9, 0, 0, 0, 0, '2022-11-30 23:20:22', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (117, '程序员', 'programmer', 'While code monkeys are not eating bananas, they\'re coding.', NULL,
        '/upload/backend/node/copqd3v6co8g8ixh56.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:20:23', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (118, '宽带症候群', 'bb', '网速很重要。比快更快。', NULL, '/upload/backend/node/copqd49xkqa8h46u4x.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:20:24', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (119, 'Android', 'android', '', NULL, '', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:18', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (120, 'Linux', 'linux',
        'Linux is a Unix-like computer operating system assembled under the model of free and open source software development and distribution.',
        NULL, '/upload/backend/node/copqel0ck934hina51.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:19', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (121, '外包', 'outsourcing', '', NULL, '/upload/backend/node/copqelxpn4mgwmkapn.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:21', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (122, '硬件', 'hardware', '硬件发烧友的讨论节点', NULL, '/upload/backend/node/copqemeipr08py080h.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:22', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (123, 'Windows', 'windows', 'Windows, not walls.', NULL, '/upload/backend/node/copqemy3wje00wialn.png', 116, 9,
        0, 0, 0, 0, '2022-11-30 23:22:23', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (124, 'Bitcoin', 'bitcoin', 'P2P digital currency', NULL, '/upload/backend/node/copqeneyrp2oj16wcc.png', 116, 9,
        0, 0, 0, 0, '2022-11-30 23:22:24', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (125, '汽车', 'car', '关于买车、开车及汽车文化的技术讨论', NULL, '/upload/backend/node/copqentrfh08hcofev.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:25', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (126, '路由器', 'router', '数据包流动的乐趣', NULL, '/upload/backend/node/copqeo8idqfkp0h3bx.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:26', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (127, '站长', 'webmaster', '', NULL, '/upload/backend/node/copqeopn769kbioqis.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:27', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (128, '编程', 'programming', 'Have fun and make money and art.', NULL,
        '/upload/backend/node/copqep3vf7bcsbales.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:28', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (129, 'GitHub', 'github', 'A global place for programmers.', NULL, '/upload/backend/node/copqepzgl2fkf5tpg6.png',
        116, 9, 0, 0, 0, 0, '2022-11-30 23:22:30', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (130, 'OpenWrt', 'openwrt', 'Wireless Freedom', NULL, '/upload/backend/node/copqeqhhtim0yhbtug.png', 116, 9, 0,
        0, 0, 0, '2022-11-30 23:22:31', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (131, 'Visual Studio Code', 'vscode', '来自 <a href=\"/go/microsoft\">Microsoft</a> 的开源代码编辑器，支持通过插件扩展功能。', NULL,
        '/upload/backend/node/copqeqwxx8hkeso7mj.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:32', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (132, 'Linode', 'linode', 'Deploy and Manage Linux Virtual Servers in the Linode Cloud.', NULL,
        '/upload/backend/node/copqeruocddktc1sng.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:34', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (133, '区块链', 'blockchain', '关于区块链技术的开发和应用的讨论。', NULL, '/upload/backend/node/copqesahh6lk5a17kj.png', 116, 9, 0,
        0, 0, 0, '2022-11-30 23:22:35', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (134, 'Markdown', 'markdown', 'Markdown is a text-to-HTML conversion tool for web writers.', NULL,
        '/upload/backend/node/copqesz611vcueruyp.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:36', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (135, '设计师', 'designer', '', NULL, '/upload/backend/node/copqetgi8rco6kxre5.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:37', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (136, 'Kindle', 'kindle', 'Amazon 公司设计生产的电子阅读设备', NULL, '/upload/backend/node/copqetvga8a0ijjz6a.png', 116, 9, 0,
        0, 0, 0, '2022-11-30 23:22:38', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (137, 'Raspberry Pi', 'pi',
        'The Raspberry Pi is a credit-card sized computer that plugs into your TV and a keyboard.', NULL, '', 116, 9, 0,
        0, 0, 0, '2022-11-30 23:22:39', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (138, '游戏开发', 'gamedev', '无论你正在制作自己的 indie game，或是在大公司参与 AAA 大作，这里是大家讨论技术和理念的地方。', NULL,
        '/upload/backend/node/copqeux77854mzrtqi.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:40', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (139, '字体排印', 'typography', '', NULL, '/upload/backend/node/copqevbiq0nk1v1nr5.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:41', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (140, 'Atom', 'atom', '来自 <a href=\"/go/github\">GitHub</a> 的开放源代码的编辑器', NULL,
        '/upload/backend/node/copqew0dnoawfdhjrj.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:43', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (141, '商业模式', 'business', '', NULL, '/upload/backend/node/copqewf3kypkxp4thp.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:44', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (142, 'SONY', 'sony', 'ソニー株式会社', NULL, '/upload/backend/node/copqewv8icnkdjqcop.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:45', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (143, 'LeetCode', 'leetcode', '关于 LeetCode 上遇到的各类问题的技术讨论', NULL, '/upload/backend/node/copqexalij3kxwxeqi.png',
        116, 9, 0, 0, 0, 0, '2022-11-30 23:22:46', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (144, 'Photoshop', 'photoshop',
        'Adobe Photoshop CS6 software delivers state-of-the-art imaging magic, new creative options, and blazingly fast performance. Photo editing and more.',
        NULL, '/upload/backend/node/copqexos6nswoszfoh.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:46', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (145, 'Amazon', 'amazon', 'Online shopping from the earth\'s biggest selection.', NULL,
        '/upload/backend/node/copqeyggvn2wksjh7q.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:48', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (146, 'Serverless', 'serverless', '', NULL, '/upload/backend/node/copqeyvp1cnshegprj.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:49', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (147, '电动汽车', 'ev', '', NULL, '/upload/backend/node/copqezl380m0euhilk.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:22:51', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (148, 'GitLab', 'gitlab', 'GitLab 的安装和使用经验分享', NULL, '/upload/backend/node/copqf0b2ji80a4fpeh.png', 116, 9, 0, 0,
        0, 0, '2022-11-30 23:22:52', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (149, 'RSS', 'rss', '一种古老而高效的信息获取方式。这个节点讨论 RSS 源的实现及各种阅读器。', NULL, '/upload/backend/node/copqf0vhqgeopdr6mb.png',
        116, 9, 0, 0, 0, 0, '2022-11-30 23:22:53', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (150, '云修电脑', 'ifix', '如果你的电脑或者电子设备遇到了什么奇怪问题，可以发到这里来让大家帮你出出主意。', NULL,
        '/upload/backend/node/copqf1bnwm8gqrfsdh.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:54', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (151, 'LEGO', 'lego', 'The toy building brick.', NULL, '/upload/backend/node/copqf251d174i0xpk6.png', 116, 9, 0,
        0, 0, 0, '2022-11-30 23:22:56', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (152, 'GitHub Copilot', 'copilot',
        '来自 GitHub 的基于 AI 的结对编程工具 <a href=\"https://copilot.github.com/\" target=\"_blank\">Copilot</a>', NULL,
        '/upload/backend/node/copqf2vwixn4cttebi.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:22:58', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (153, 'IPFS', 'ipfs', 'IPFS is the Distributed Web', NULL, '/upload/backend/node/copqf3n2d81svm5m0a.png', 116, 9,
        0, 0, 0, 0, '2022-11-30 23:22:59', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (154, '以太坊', 'ethereum', '以太坊是一个为去中心化应用程序而生的全球开源平台。', NULL, '/upload/backend/node/copqf4giafrstuxspn.png', 116,
        9, 0, 0, 0, 0, '2022-11-30 23:23:01', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (155, '加密货币', 'crypto', '', NULL, '/upload/backend/node/copqf4uwdtm8ecmrfm.png', 116, 9, 0, 0, 0, 0,
        '2022-11-30 23:23:02', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (156, 'DJI', 'dji', '大疆创新——全球无人机控制与航拍影像系统先驱', NULL, '/upload/backend/node/copqf5b2bo28u6cgbq.png', 116, 9, 0, 0,
        0, 0, '2022-11-30 23:23:03', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (157, 'Blender', 'blender', '开源的 3D 创作工具', NULL, '/upload/backend/node/copqf5qtre7ka0znab.png', 116, 9, 0, 0, 0,
        0, '2022-11-30 23:23:04', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (158, 'Logseq', 'logseq', '一个尊重隐私的，本地优先的个人知识库管理软件', NULL, '/upload/backend/node/copqf6edn2u0f5e0og.png', 116, 9,
        0, 0, 0, 0, '2022-11-30 23:23:05', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (159, 'MetaMask', 'metamask', '关于 MetaMask 的使用问题的讨论。', NULL, '/upload/backend/node/copqf72utny8sjiaot.png', 116,
        9, 0, 0, 0, 0, '2022-11-30 23:23:07', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (160, '奇绩创坛', 'miracleplus', '奇绩创坛为创业公司提供初创期融资。目标是帮助创业者迈出第一步，并帮助创业公司成长到足够优秀。', NULL,
        '/upload/backend/node/copqf7qcvs8gfm2dtj.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:23:08', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (161, 'GameDB', 'gamedb', '8-bit / 16-bit 老游戏的截图和视频的数据库，为了我们记忆中那些永恒的经典', NULL,
        '/upload/backend/node/copqf83444l4r2rhna.png', 116, 9, 0, 0, 0, 0, '2022-11-30 23:23:09', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (162, '游戏', '游戏', '', NULL, '', 0, 10, 0, 0, 0, 0, '2022-11-30 23:23:09', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (163, '游戏', 'games', 'Life is short, have more fun.', NULL, '/upload/backend/node/copqf8lelv80slybtw.png', 162,
        10, 0, 0, 0, 0, '2022-11-30 23:23:10', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (164, 'Steam', 'steam', '关于游戏平台 Steam 和掌机硬件 Steam Deck 的讨论节点。', NULL,
        '/upload/backend/node/copqf90v8ntsk5fefe.png', 162, 10, 0, 0, 0, 0, '2022-11-30 23:23:11', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (165, 'Nintendo Switch', 'switch',
        '任天堂 Switch（日语：ニンテンドースイッチ，英语：Nintendo Switch）是日本任天堂公司出品的电子游戏机，于 2017 年 3 月 3 日开始发售。', NULL,
        '/upload/backend/node/copqf9hftk1cchgtig.png', 162, 10, 0, 0, 0, 0, '2022-11-30 23:23:12', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (166, 'PlayStation 4', 'ps4', 'Greatness Awaits', NULL, '/upload/backend/node/copqf9xlyjfc35nah2.png', 162, 10,
        0, 0, 0, 0, '2022-11-30 23:23:13', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (167, 'Minecraft', 'minecraft',
        'Minecraft is a game about breaking and placing blocks. At first, people built structures to protect against nocturnal monsters, but as the game grew players worked together to create wonderful, imaginative things.',
        NULL, '/upload/backend/node/copqfada95tcczngds.png', 162, 10, 0, 0, 0, 0, '2022-11-30 23:23:14', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (168, 'iGame', 'igame', 'iOS 上有很多精彩游戏。', NULL, '/upload/backend/node/copqfau3w7jcuts5fe.png', 162, 10, 0, 0, 0,
        0, '2022-11-30 23:23:15', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (169, 'StarCraft 2', 'sc2', '', NULL, '/upload/backend/node/copqfb8bagtshvbowt.png', 162, 10, 0, 0, 0, 0,
        '2022-11-30 23:23:16', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (170, 'Battlefield 3', 'bf3', '', NULL, '/upload/backend/node/copqfbo9luvcz6pade.png', 162, 10, 0, 0, 0, 0,
        '2022-11-30 23:23:17', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (171, 'World of Warcraft', 'wow', '魔兽世界是由暴雪娱乐制作的一款大型多人在线角色扮演（MMORPG）游戏。', NULL,
        '/upload/backend/node/copqfcamqt7cifxgvh.png', 162, 10, 0, 0, 0, 0, '2022-11-30 23:23:18', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (172, '怀旧游戏', 'retro', '那些在记忆中最美的像素和 8-bit 旋律', NULL, '/upload/backend/node/copqfcqju420ilyvbf.png', 162, 10, 0,
        0, 0, 0, '2022-11-30 23:23:19', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (173, 'EVE', 'eve', 'The world\'s greatest virtual universe.', NULL,
        '/upload/backend/node/copqfd9i4hfcjysrik.png', 162, 10, 0, 0, 0, 0, '2022-11-30 23:23:20', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (174, '精灵宝可梦', 'pokemon', '', NULL, '/upload/backend/node/copqfe0rg1v4l6po1k.png', 162, 10, 0, 0, 0, 0,
        '2022-11-30 23:23:22', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (175, 'PlayStation 5', 'ps5', 'Play Has No Limits', NULL, '/upload/backend/node/copqfesg5l7cqiv4us.png', 162, 10,
        0, 0, 0, 0, '2022-11-30 23:23:24', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (176, '3DS', '3ds', '', NULL, '/upload/backend/node/copqffljo2d4537yde.png', 162, 10, 0, 0, 0, 0,
        '2022-11-30 23:23:25', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (177, 'Gran Turismo', 'gt', 'The real driving simulator.', NULL, '/upload/backend/node/copqffz97dtkvcjgbq.png',
        162, 10, 0, 0, 0, 0, '2022-11-30 23:23:26', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (178, 'Battlefield 4', 'bf4', 'EA and DICE\'s next-generation first-person-shooter.', NULL,
        '/upload/backend/node/copqfges2954gimplr.png', 162, 10, 0, 0, 0, 0, '2022-11-30 23:23:27', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (179, 'Wii U', 'wiiu', 'The Wii U is a video game console from Nintendo and the successor to the Wii.', NULL,
        '/upload/backend/node/copqfgsxy10gkgnufd.png', 162, 10, 0, 0, 0, 0, '2022-11-30 23:23:28', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (180, 'Battlefield V', 'bfv', '战地系列（Battlefield）在 2018 年末的最新一代，游戏设定在第二次世界大战。', NULL,
        '/upload/backend/node/copqfhnvc49cf3kvcv.png', 162, 10, 0, 0, 0, 0, '2022-11-30 23:23:30', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (181, '生活', '生活', '', NULL, '', 0, 11, 0, 0, 0, 0, '2022-11-30 23:23:30', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (182, '二手交易', 'all4all', '为自己的闲置物品找到更好的主人。', NULL, '/upload/backend/node/copqfi38lgy8mpn6zy.png', 181, 11, 0, 0,
        0, 0, '2022-11-30 23:23:31', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (183, '酷工作', 'jobs', '做有趣的有意义的事情。', '', '/upload/backend/node/copqfiiqxzs0biy9ob.png', 181, 11, 1, 0, 0, 0,
        '2022-11-30 23:23:32', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (184, '职场话题', 'career', '这里，我们聊聊那些工作中遇到的开心和不开心的事。', NULL, '/upload/backend/node/copqfj0mzqi8e0qmme.png', 181, 11,
        0, 0, 0, 0, '2022-11-30 23:23:33', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (185, '求职', 'cv', '欢迎在这里发布自己的求职简历。', NULL, '/upload/backend/node/copqfjg2ahy0lmwmeg.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:34', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (186, '天黑以后', 'afterdark', '白天和晚上的我，是不一样的。', NULL, '/upload/backend/node/copqfjwgntjkhcycxt.png', 181, 11, 0, 0,
        0, 0, '2022-11-30 23:23:35', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (187, '免费赠送', 'free', '&nbsp;', NULL, '/upload/backend/node/copqfkbsvtbsjtydb8.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:36', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (188, '音乐', 'music', 'Music is an art form whose medium is sound and silence.', NULL,
        '/upload/backend/node/copqfl3r5wzkygb8jf.png', 181, 11, 0, 0, 0, 0, '2022-11-30 23:23:37', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (189, '电影', 'movie', '用 90 分钟去体验另外一个世界。', NULL, '/upload/backend/node/copqflj8519ksguwp8.png', 181, 11, 0, 0, 0,
        0, '2022-11-30 23:23:38', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (190, '物物交换', 'exchange', '将自己不需要的闲置物品拿出来和大家交换吧。', NULL, '', 181, 11, 0, 0, 0, 0, '2022-11-30 23:23:39', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (191, '投资', 'invest', 'Can you make money with money?', NULL, '/upload/backend/node/copqfm9hv5rc8bxbte.png', 181,
        11, 0, 0, 0, 0, '2022-11-30 23:23:40', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (192, '团购', 'tuan', '', NULL, '', 181, 11, 0, 0, 0, 0, '2022-11-30 23:23:40', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (193, '剧集', 'tv', '', NULL, '/upload/backend/node/copqfn09v5w8w2g48y.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:41', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (194, '旅行', 'travel', '你会把上大学的学费用来环游世界么？', NULL, '/upload/backend/node/copqfnfdk1coe2fnzs.png', 181, 11, 0, 0, 0,
        0, '2022-11-30 23:23:42', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (195, '信用卡', 'creditcard', '', NULL, '', 181, 11, 0, 0, 0, 0, '2022-11-30 23:23:43', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (196, '美酒与美食', 'taste', '关于那些好喝和好吃的', NULL, '/upload/backend/node/copqfp0yvz5sdfwjcp.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:46', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (197, '阅读', 'reading', '少上网，多读书', NULL, '/upload/backend/node/copqfpsqt1g89ggap1.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:48', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (198, '摄影', 'photograph', '', NULL, '/upload/backend/node/copqfq89nybcxvo1fn.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:49', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (199, '宠物', 'pet', '', NULL, '/upload/backend/node/copqfqmidk2w7djgbm.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:49', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (200, 'Baby', 'baby', '', NULL, '', 181, 11, 0, 0, 0, 0, '2022-11-30 23:23:51', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (201, '绿茵场', 'soccer', 'Brazil 2014', NULL, '/upload/backend/node/copqfrljhww013bcj3.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:51', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (202, '咖啡', 'coffee',
        'Coffee is a brewed drink prepared from roasted coffee beans, the seeds of berries from certain Coffea species.',
        NULL, '/upload/backend/node/copqfsd9xra8zvsxtj.png', 181, 11, 0, 0, 0, 0, '2022-11-30 23:23:53', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (203, '日记', 'diary', '', NULL, '', 181, 11, 0, 0, 0, 0, '2022-11-30 23:23:54', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (204, '骑行', 'bike', '', NULL, '', 181, 11, 0, 0, 0, 0, '2022-11-30 23:23:55', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (205, '植物', 'plant', '', NULL, '/upload/backend/node/copqfto3wmpsqktchj.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:56', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (206, '蘑菇', 'mushroom', '', NULL, '/upload/backend/node/copqfuc6pwdksd0q4d.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:58', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (207, '行程控', 'mileage', '', NULL, '/upload/backend/node/copqfuulc9bkvbgisf.png', 181, 11, 0, 0, 0, 0,
        '2022-11-30 23:23:59', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (208, 'Internet', 'Internet', '', NULL, '', 0, 12, 0, 0, 0, 0, '2022-11-30 23:23:59', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (209, 'Google', 'google',
        'Google’s mission is to organize the world’s information and make it universally accessible and useful.', NULL,
        '/upload/backend/node/copqfvfttxewpviepg.png', 208, 12, 0, 0, 0, 0, '2022-11-30 23:24:00', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (210, 'YouTube', 'youtube', '世界上最大的视频网站，由前 PayPal 员工创建，后来被 Google 收购。一个你绝对不应该错过的信息源。', NULL,
        '/upload/backend/node/copqfw0393nsecl1ks.png', 208, 12, 0, 0, 0, 0, '2022-11-30 23:24:01', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (211, '哔哩哔哩', 'bilibili', '', NULL, '', 208, 12, 0, 0, 0, 0, '2022-11-30 23:24:02', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (212, 'Notion', 'notion', 'Notion 爱好者的社区', NULL, '/upload/backend/node/copqfxc8elk0idn7d5.png', 208, 12, 0, 0, 0,
        0, '2022-11-30 23:24:04', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (213, 'Reddit', 'reddit',
        'Reddit, stylized as reddit, is a social news and entertainment website where registered users submit content in the form of either a link or a text (\"self\") post.',
        NULL, '/upload/backend/node/copqfy2jiemolyj5px.png', 208, 12, 0, 0, 0, 0, '2022-11-30 23:24:06', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (214, '城市', '城市', '', NULL, '', 0, 13, 0, 0, 0, 0, '2022-11-30 23:24:06', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (215, '北京', 'beijing', '', NULL, '/upload/backend/node/copqfykxgkfku4ur2c.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:07', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (216, '上海', 'shanghai', '', NULL, '/upload/backend/node/copqfzfjutoggoahwz.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:09', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (217, '深圳', 'shenzhen', '', NULL, '/upload/backend/node/copqg04col2ow6phkv.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:10', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (218, '杭州', 'hangzhou', '', NULL, '/upload/backend/node/copqg0utfuzs8rbxcs.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:12', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (219, '成都', 'chengdu', '', NULL, '/upload/backend/node/copqg1cm7ocovb81pd.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:13', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (220, '广州', 'guangzhou', '', NULL, '/upload/backend/node/copqg1w86aa0a93xyz.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:14', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (221, '武汉', 'wuhan', '', NULL, '/upload/backend/node/copqg2sizp2w5mzwsy.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:16', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (222, '南京', 'nanjing', '', NULL, '/upload/backend/node/copqg3nnygoo2874l5.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:18', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (223, '西安', 'xian', '', NULL, '/upload/backend/node/copqg4jgdg9kpmqxl4.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:20', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (224, '重庆', 'chongqing', '', NULL, '/upload/backend/node/copqg514n2w0ttdiq8.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:21', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (225, '昆明', 'kunming', '机场代码 KMG • 电话区号 0871 • 人口 726 万', NULL, '/upload/backend/node/copqg60081ggcejxtw.png',
        214, 13, 0, 0, 0, 0, '2022-11-30 23:24:23', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (226, '苏州', 'suzhou', '', NULL, '/upload/backend/node/copqg6hmfhk0usxrre.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:24', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (227, '厦门', 'xiamen', '', NULL, '/upload/backend/node/copqg7gl64pswvdulr.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:26', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (228, '天津', 'tianjin', '', NULL, '/upload/backend/node/copqg85ry54oud4vxu.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:24:28', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (229, '青岛', 'qingdao', '', NULL, '/upload/backend/node/copqgn4uapzk4o4jny.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:25:00', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (230, 'San Francisco', 'sanfrancisco',
        'San Francisco, officially the City and County of San Francisco, is the leading financial and cultural center of Northern California and the San Francisco Bay Area.',
        NULL, '/upload/backend/node/copqgo1k2h6wj3gbba.png', 214, 13, 0, 0, 0, 0, '2022-11-30 23:25:02', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (231, 'New York', 'nyc', '', NULL, '/upload/backend/node/copqgp03v85ks7aemr.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:25:04', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (232, 'Los Angeles', 'la', '', NULL, '/upload/backend/node/copqgpk33xi0cw24jw.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:25:05', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (233, '东京', 'tokyo', '', NULL, '/upload/backend/node/copqgq9fv42078nhgz.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:25:07', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (234, '贵阳', 'guiyang', '', NULL, '/upload/backend/node/copqgrl1qhiwnwxhx6.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:25:10', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (235, 'Singapore', 'singapore', '', NULL, '/upload/backend/node/copqgs27r6gglgdzoy.png', 214, 13, 0, 0, 0, 0,
        '2022-11-30 23:25:11', NULL);
INSERT INTO `forum_nodes` (`id`, `name`, `keyword`, `description`, `detail`, `img`, `parent_id`, `category_id`,
                           `is_index`, `is_virtual`, `is_disabled_edit`, `sort`, `created_at`, `deleted_at`)
VALUES (236, 'Boston', 'boston', '', NULL, '', 214, 13, 0, 0, 0, 0, '2022-11-30 23:25:12', NULL);
COMMIT;

-- ----------------------------
-- Table structure for forum_posts
-- ----------------------------
DROP TABLE IF EXISTS `forum_posts`;
CREATE TABLE `forum_posts`
(
    `id`                  int(11) unsigned NOT NULL AUTO_INCREMENT,
    `node_id`             int(11) unsigned NOT NULL DEFAULT 0 COMMENT '节点id',
    `user_id`             int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`            varchar(45)      NOT NULL COMMENT '用户名',
    `title`               varchar(255)     NOT NULL COMMENT '标题',
    `content_type`        char(20)                  DEFAULT 'general' COMMENT 'quill为使用quill编辑器，markdown为使用markdown编辑器,general 为普通文本',
    `content`             longtext                  DEFAULT NULL COMMENT '内容',
    `html_content`        longtext                  DEFAULT NULL COMMENT 'html内容',
    `top_end_time`        datetime                  DEFAULT NULL COMMENT '置顶截止时间,为空说明没有置顶',
    `character_amount`    int(11) unsigned NOT NULL DEFAULT 0 COMMENT '字符长度',
    `visit_amount`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '访问次数',
    `collection_amount`   int(11) unsigned NOT NULL DEFAULT 0 COMMENT '收藏次数',
    `reply_amount`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '回复次数',
    `thanks_amount`       int(11) unsigned NOT NULL DEFAULT 0 COMMENT '感谢次数',
    `shielded_amount`     int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被屏蔽次数',
    `status`              tinyint(1)                DEFAULT 0 COMMENT '状态：0 未审核 1 已审核',
    `weight`              int(11)          NOT NULL DEFAULT 0 COMMENT '权重',
    `reply_last_user_id`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '最后回复用户id',
    `reply_last_username` varchar(45)      NOT NULL COMMENT '最后回复用户名',
    `last_change_time`    datetime                  DEFAULT NULL COMMENT '主题最后变动时间',
    `created_at`          datetime                  DEFAULT NULL COMMENT '主题创建时间',
    `updated_at`          datetime                  DEFAULT NULL COMMENT '主题更新时间',
    `deleted_at`          datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_node_id` (`node_id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛主题表';

-- ----------------------------
-- Records of forum_posts
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for forum_prompts
-- ----------------------------
DROP TABLE IF EXISTS `forum_prompts`;
CREATE TABLE `forum_prompts`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `position`    varchar(50)      NOT NULL COMMENT '提示语位置',
    `content`     text             NOT NULL COMMENT '提示语内容',
    `description` varchar(255)     NOT NULL COMMENT '简介',
    `is_disabled` tinyint(1) DEFAULT 0 COMMENT '状态：0 正常 1禁用',
    `created_at`  datetime   DEFAULT NULL COMMENT '创建时间',
    `update_at`   datetime   DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='提示语内容';

-- ----------------------------
-- Records of forum_prompts
-- ----------------------------
BEGIN;
INSERT INTO `forum_prompts` (`id`, `position`, `content`, `description`, `is_disabled`, `created_at`, `update_at`)
VALUES (1, 'post-new',
        '<p><span style=\"color: rgb(187, 187, 187);\">每日创建主题数量最大为10,主题的最大字符数量为20000,每次创建主题扣除10积分</span></p><p><br></p><p><span style=\"color: rgb(187, 187, 187);\">每当有人回复该主题时,op获得10积分,有人感谢主题时op获得10积分</span></p>',
        '', 0, '2022-11-30 13:46:32', '2022-11-30 17:47:39');
INSERT INTO `forum_prompts` (`id`, `position`, `content`, `description`, `is_disabled`, `created_at`, `update_at`)
VALUES (2, 'reply-new', '<p>每日最大回复数量为100,回复最大字符数量为500,每次回复扣除10积分,有人感谢回复时回复者获得10积分</p>', '', 0, '2022-11-30 17:42:15',
        '2022-11-30 17:47:09');
COMMIT;

-- ----------------------------
-- Table structure for forum_replies
-- ----------------------------
DROP TABLE IF EXISTS `forum_replies`;
CREATE TABLE `forum_replies`
(
    `id`                int(11) unsigned NOT NULL AUTO_INCREMENT,
    `posts_id`          int(11) unsigned NOT NULL DEFAULT 0 COMMENT '主题id',
    `user_id`           int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`          varchar(45)      NOT NULL COMMENT '用户名',
    `relation_user_ids` varchar(255)     NOT NULL DEFAULT '' COMMENT '涉及用户ids，多个以逗号分隔',
    `content`           longtext                  DEFAULT NULL COMMENT '内容',
    `character_amount`  int(11) unsigned NOT NULL DEFAULT 0 COMMENT '字符长度',
    `thanks_amount`     int(11) unsigned NOT NULL DEFAULT 0 COMMENT '感谢次数',
    `shielded_amount`   int(11) unsigned NOT NULL DEFAULT 0 COMMENT '被屏蔽次数',
    `status`            tinyint(1)                DEFAULT 0 COMMENT '状态：0 未审核 1 已审核',
    `created_at`        datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`        datetime                  DEFAULT NULL COMMENT '更新时间',
    `deleted_at`        datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛回复表';

-- ----------------------------
-- Records of forum_replies
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for forum_sensitive_words
-- ----------------------------
DROP TABLE IF EXISTS `forum_sensitive_words`;
CREATE TABLE `forum_sensitive_words`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `type`       varchar(20)      NOT NULL COMMENT '类型',
    `word`       varchar(255)     NOT NULL COMMENT '敏感词',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `update_at`  datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_word` (`word`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='敏感词表';

-- ----------------------------
-- Records of forum_sensitive_words
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for forum_user_posts_histories
-- ----------------------------
DROP TABLE IF EXISTS `forum_user_posts_histories`;
CREATE TABLE `forum_user_posts_histories`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`   varchar(45)      NOT NULL COMMENT '用户名',
    `posts_id`   int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联主题id',
    `created_at` datetime                  DEFAULT NULL COMMENT '创建时间',
    `update_at`  datetime                  DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='用户访问主题历史表 ';

-- ----------------------------
-- Records of forum_user_posts_histories
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for forum_users
-- ----------------------------
DROP TABLE IF EXISTS `forum_users`;
CREATE TABLE `forum_users`
(
    `id`                     int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `username`               varchar(45)         NOT NULL COMMENT '用户名',
    `email`                  varchar(50)         NOT NULL DEFAULT '' COMMENT 'email',
    `description`            varchar(255)        NOT NULL DEFAULT '' COMMENT '简介',
    `password`               char(32)            NOT NULL COMMENT 'MD5密码',
    `avatar`                 varchar(200)                 DEFAULT NULL COMMENT '头像地址',
    `status`                 int(11)                      DEFAULT 0 COMMENT '二进制位,0111 由低到高分别代表 禁止登录，禁止发帖，禁止回复，尚未激活',
    `posts_amount`           int(11) unsigned    NOT NULL DEFAULT 0 COMMENT '创建主题次数',
    `reply_amount`           int(11) unsigned    NOT NULL DEFAULT 0 COMMENT '回复次数',
    `shielded_amount`        int(11) unsigned    NOT NULL DEFAULT 0 COMMENT '被屏蔽次数',
    `follow_by_other_amount` int(11) unsigned    NOT NULL DEFAULT 0 COMMENT '被关注次数',
    `today_activity`         int(11) unsigned    NOT NULL DEFAULT 0 COMMENT '今日活跃度',
    `balance`                bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '余额',
    `site`                   varchar(200)                 DEFAULT NULL COMMENT '个人站点',
    `company`                varchar(200)                 DEFAULT NULL COMMENT '所在公司',
    `job`                    varchar(200)                 DEFAULT NULL COMMENT '工作职位',
    `location`               varchar(200)                 DEFAULT NULL COMMENT '所在地',
    `signature`              varchar(200)                 DEFAULT NULL COMMENT '个人签名',
    `introduction`           varchar(500)                 DEFAULT NULL COMMENT '个人简介',
    `remark`                 varchar(500)                 DEFAULT NULL COMMENT '备注',
    `last_login_ip`          varchar(50)                  DEFAULT NULL COMMENT '最后登陆IP',
    `last_login_time`        datetime                     DEFAULT NULL COMMENT '最后登陆时间',
    `created_at`             datetime                     DEFAULT NULL COMMENT '注册时间',
    `updated_at`             datetime                     DEFAULT NULL COMMENT '更新时间',
    `deleted_at`             datetime                     DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_username` (`username`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛用户表';

-- ----------------------------
-- Records of forum_users
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ga_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `ga_admin_log`;
CREATE TABLE `ga_admin_log`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
    `path`             varchar(255)     NOT NULL DEFAULT '' COMMENT '请求路径',
    `method`           varchar(10)      NOT NULL DEFAULT '' COMMENT '请求方法',
    `path_name`        varchar(255)     NOT NULL DEFAULT '' COMMENT '请求路径名称',
    `params`           text                      DEFAULT NULL COMMENT '请求参数',
    `response`         longtext                  DEFAULT NULL COMMENT '响应结果',
    `created_at`       datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`       datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- ----------------------------
-- Records of ga_admin_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ga_admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `ga_admin_menu`;
CREATE TABLE `ga_admin_menu`
(
    `id`                   int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`                 varchar(20)      NOT NULL COMMENT '菜单名称',
    `path`                 varchar(100)     NOT NULL DEFAULT '' COMMENT '前端路由地址，可以是外链',
    `parent_id`            int(11) unsigned NOT NULL DEFAULT 0 COMMENT '父id',
    `identification`       varchar(40)      NOT NULL DEFAULT '' COMMENT '后端权限标识符',
    `method`               varchar(10)      NOT NULL DEFAULT '' COMMENT '请求方法',
    `front_component_path` varchar(255)              DEFAULT NULL COMMENT '前端组件路径',
    `icon`                 varchar(100)              DEFAULT '#' COMMENT '菜单图标',
    `sort`                 tinyint(4)       NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
    `status`               varchar(10)               DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `created_at`           datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`           datetime                  DEFAULT NULL COMMENT '更新时间',
    `type`                 varchar(12)      NOT NULL COMMENT '菜单类型',
    `link_type`            varchar(12)      NOT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 93
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='后台菜单表';

-- ----------------------------
-- Records of ga_admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (1, '系统管理', 'system', 0, '', '', '', 'system', 10, 'normal', '2022-01-09 19:44:50', '2022-11-29 13:46:57',
        'directory', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (2, '管理员列表', 'administrator', 1, '/administrator-list', 'get', 'system/administrator/index', 'user', 0, 'normal',
        '2022-01-09 19:46:32', '2022-02-23 16:51:44', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (3, '角色列表', 'role', 1, '/role-list', 'get', 'system/role/index', 'peoples', 0, 'normal', '2022-01-09 19:47:59',
        '2022-02-23 17:02:24', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (4, '角色新增', '', 3, '/role-store', 'post', '', '', 0, 'normal', '2022-01-09 19:49:15', '2022-01-09 19:49:15',
        'operation', '');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (5, '角色更新', '', 3, '/role-update', 'put', '', '', 0, 'normal', '2022-01-09 19:49:36', '2022-01-09 19:49:36',
        'operation', '');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (6, '角色删除', '', 3, '/role-destroy', 'delete', '', '', 0, 'normal', '2022-01-09 19:49:46', '2022-01-09 19:49:46',
        'operation', '');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (7, '管理员新增', '', 2, '/administrator-store', 'post', '', '', 0, 'normal', '2022-01-09 19:50:30',
        '2022-02-21 15:47:38', 'operation', '');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (8, '管理员删除', '', 2, '/administrator-destroy', 'delete', '', '', 0, 'normal', '2022-01-09 19:54:44',
        '2022-01-09 19:54:44', 'operation', '');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (9, '管理员更新', '', 2, '/administrator-update', 'put', '', '', 0, 'normal', '2022-01-09 19:55:13',
        '2022-01-09 19:55:13', 'operation', '');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (10, '菜单列表', 'menu', 1, '/menu-list', 'get', 'system/menu/index', 'tree-table', 0, 'normal', NULL, NULL, 'link',
        'tab');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (11, '菜单新增', '', 10, '/menu-store', 'POST', '', '#', 0, 'normal', NULL, '2022-02-21 16:10:46', 'operation',
        'tab');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (12, '菜单更新', '', 10, '/menu-update', 'PUT', '', '#', 0, 'normal', NULL, '2022-02-21 16:10:54', 'operation',
        'tab');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (13, '菜单删除', '', 10, '/menu-destroy', 'DELETE', '', '#', 0, 'normal', NULL, '2022-02-23 16:55:40', 'operation',
        'tab');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (15, 'aaa', 'test', 14, 'fffff', 'GET', '', '', 1, 'normal', '2022-02-18 13:31:26', '2022-02-21 17:34:05',
        'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (18, '管理员信息', '', 2, '/administrator-info', 'GET', '', '', 1, 'normal', '2022-02-21 16:12:38',
        '2022-02-23 17:32:55', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (19, '角色信息', '', 3, '/role-info', 'GET', '', '', 1, 'normal', '2022-02-21 16:18:39', '2022-02-23 15:17:30',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (20, '菜单信息', '', 10, '/menu-info', 'GET', '', '', 1, 'normal', '2022-02-21 16:19:16', '2022-02-23 15:05:50',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (22, '配置管理', 'config', 1, '/config-list', 'GET', 'system/config/index', 'edit', 0, 'normal',
        '2022-03-19 17:40:06', '2022-04-27 18:39:20', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (23, '配置更新', '', 22, '/config-update', 'PUT', '', '', 1, 'normal', '2022-04-28 14:29:18', '2022-04-28 14:32:53',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (24, '操作日志', '/operationlog', 1, '/operation-log-list', 'GET', 'system/operationlog/index', 'documentation', 6,
        'normal', '2022-05-05 12:02:45', '2022-05-05 15:14:38', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (27, '用户钱包地址', 'binance_user_address', 26, '/binance-user-address-list', 'GET', 'binance/useraddress/index',
        'list', 1, 'normal', '2022-06-15 16:09:12', '2022-06-16 15:43:17', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (28, '归集列表', 'collect-list', 26, '/binance-collect-list', 'GET', 'binance/collect/index', 'list', 3, 'normal',
        '2022-06-16 14:53:57', '2022-06-16 15:42:56', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (29, '提现列表', 'bianance-withdraw', 26, '/binance-withdraw-list', 'GET', 'binance/withdraw/index', 'list', 4,
        'normal', '2022-06-16 16:01:32', '2022-06-16 16:01:32', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (30, '转账队列', 'binance-queue-task-list', 26, '/binance-queue-task-list', 'GET', 'binance/queue_task/index',
        'list', 5, 'normal', '2022-06-16 16:50:54', '2022-06-16 17:45:59', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (31, '通知列表', 'binance-notify-list', 26, '/binance-notify-list', 'GET', 'binance/notify/index', 'list', 6,
        'normal', '2022-06-16 17:45:52', '2022-06-16 17:45:52', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (32, '合约管理', 'binance-contract-list', 26, '/binance-contract-list', 'GET', 'binance/contract/index', 'education',
        0, 'normal', '2022-06-17 14:27:26', '2022-06-17 14:27:26', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (33, '新增合约', '', 32, '/binance-contract-store', 'POST', '', '', 0, 'normal', '2022-06-17 15:11:22',
        '2022-06-17 15:17:10', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (34, '更新合约', '', 32, '/binance-contract-update', 'PUT', '', '', 1, 'normal', '2022-06-17 15:11:35',
        '2022-06-17 15:50:17', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (35, '合约信息', '', 32, '/binance-contract-info', 'GET', '', '', 22, 'normal', '2022-06-17 15:24:22',
        '2022-06-17 15:24:22', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (36, '删除合约', '', 32, '/binance-contract-destroy', 'DELETE', '', '', 6, 'normal', '2022-06-17 15:50:11',
        '2022-06-17 15:50:11', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (37, '丢失区块列表', 'binance-lose-block-list', 26, '/binance-lose-block-list', 'GET', 'binance/lose_block/index',
        'list', 2, 'normal', '2022-06-17 16:28:59', '2022-06-17 16:28:59', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (38, '新增区块', '', 37, '/binance-lose-block-store', 'POST', '', '', 0, 'normal', '2022-06-17 16:29:19',
        '2022-06-17 16:29:19', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (39, '删除区块', '', 37, '/binance-lose-block-destroy', 'DELETE', '', '', 1, 'normal', '2022-06-17 16:30:10',
        '2022-06-17 16:31:23', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (40, '更新队列', '', 30, '/binance-queue-task-update', 'PUT', '', '', 3, 'normal', '2022-06-22 10:29:25',
        '2022-06-22 10:30:55', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (41, '删除队列任务', '', 30, '/binance-queue-task-destroy', 'DELETE', '', '', 5, 'normal', '2022-06-22 10:34:02',
        '2022-06-22 10:34:22', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (42, '更新归集', '', 28, '/binance-collect-update', 'PUT', '', '', 3, 'normal', '2022-06-22 11:06:28',
        '2022-06-22 11:06:28', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (43, '删除归集', '', 28, '/binance-collect-destroy', 'DELETE', '', '', 4, 'normal', '2022-06-22 11:06:48',
        '2022-06-22 11:06:48', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (44, '更新提现', '', 29, '/binance-withdraw-update', 'PUT', '', '', 2, 'normal', '2022-06-22 11:24:42',
        '2022-06-22 11:24:42', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (45, '删除提现', '', 29, '/binance-withdraw-destroy', 'DELETE', '', '', 4, 'normal', '2022-06-22 11:24:59',
        '2022-06-22 11:26:31', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (46, '更新通知', '', 31, '/binance-notify-update', 'PUT', '', '', 3, 'normal', '2022-06-22 11:36:41',
        '2022-06-22 11:36:41', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (47, '删除通知', '', 31, '/binance-notify-destroy', 'DELETE', '', '', 5, 'normal', '2022-06-22 11:37:01',
        '2022-06-22 11:37:01', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (48, '其它管理', 'forum', 0, '', 'GET', '', 'table', 9, 'normal', '2022-09-06 15:31:54', '2022-11-29 13:47:05',
        'directory', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (49, '节点管理', 'node', 85, '/node-list', 'GET', 'forum/node/index', 'example', 1, 'normal', '2022-09-06 15:35:40',
        '2022-11-28 15:01:32', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (50, '新增节点', '', 49, '/node-store', 'POST', '', '', 0, 'normal', '2022-09-06 15:36:43', '2022-09-06 15:36:43',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (51, '更新节点', '', 49, '/node-update', 'PUT', '', '', 1, 'normal', '2022-09-06 15:38:35', '2022-09-06 15:39:01',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (52, '删除节点', '', 49, '/node-destroy', 'DELETE', '', '', 2, 'normal', '2022-09-06 15:38:56',
        '2022-09-06 15:38:56', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (53, '节点信息', '', 49, '/node-info', 'GET', '', '', 0, 'normal', '2022-09-06 15:39:46', '2022-09-06 15:39:46',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (54, '上传节点图片', '', 49, '/node-upload-img', 'POST', '', '', 1, 'normal', '2022-10-29 16:27:00',
        '2022-10-29 16:27:00', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (55, '下载节点图片', '', 49, '/node-download-img', 'POST', '', '', 0, 'normal', '2022-10-29 21:41:13',
        '2022-10-29 21:49:13', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (57, '用户列表', 'users', 87, '/user-list', 'GET', 'forum/user/index', 'user', 0, 'normal', '2022-10-29 21:59:54',
        '2022-11-29 13:20:32', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (58, '新增用户', '', 57, '/user-store', 'POST', '', '', 0, 'normal', '2022-10-29 22:01:02', '2022-10-29 22:01:09',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (59, '更新用户', '', 57, '/user-update', 'PUT', '', '', 0, 'normal', '2022-10-29 22:01:56', '2022-10-29 22:01:56',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (60, '用户信息', '', 57, '/user-info', 'GET', '', '', 0, 'normal', '2022-10-29 22:02:12', '2022-10-29 22:02:12',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (61, '删除用户', '', 57, '/user-destroy', 'DELETE', '', '', 0, 'normal', '2022-10-29 22:02:35',
        '2022-10-29 22:02:35', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (62, '主题列表', 'posts', 86, '/post-list', 'GET', 'forum/post/index', 'post', 2, 'normal', '2022-10-30 12:35:38',
        '2022-11-29 13:53:24', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (63, '删除主题', '', 62, '/post-destroy', 'DELETE', '', '', 0, 'normal', '2022-10-30 12:36:17',
        '2022-10-30 13:48:30', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (64, '查看主题', '', 62, '/post-info', 'GET', '', '', 0, 'normal', '2022-10-30 12:36:40', '2022-10-30 12:36:40',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (65, '回复列表', 'reply', 86, '/reply-list', 'GET', 'forum/reply/index', 'message', 3, 'normal',
        '2022-10-30 13:47:31', '2022-11-28 15:03:39', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (66, '查看回复', '', 65, '/reply-info', 'GET', '', '', 0, 'normal', '2022-10-30 13:47:50', '2022-10-30 13:48:05',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (67, '删除回复', '', 65, '/reply-destroy', 'DELETE', '', '', 0, 'normal', '2022-10-30 13:48:22',
        '2022-10-30 13:48:22', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (68, '论坛设置', 'forum_config', 0, '/forum-config-list', 'GET', 'forum/config/index', 'component', 8, 'normal',
        '2022-10-30 21:00:41', '2022-11-29 13:47:11', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (69, '配置更新', '', 68, '/forum-config-update', 'PUT', '', '', 0, 'normal', '2022-10-30 21:01:47',
        '2022-10-30 21:01:47', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (70, '审核回复', '', 65, '/reply-audit', 'POST', '', '', 0, 'normal', '2022-10-30 22:14:22', '2022-10-30 22:14:22',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (71, '敏感词管理', 'sensitive-word', 48, '/sensitive-word-list', 'GET', 'forum/sensitive/index', 'dict', 5, 'normal',
        '2022-11-07 16:06:00', '2022-11-07 21:04:56', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (72, '保存', '', 71, '/sensitive-word-store', 'POST', '', '', 0, 'normal', '2022-11-07 16:07:05',
        '2022-11-07 16:07:05', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (73, '删除', '', 71, '/sensitive-word-destroy', 'DELETE', '', '', 0, 'normal', '2022-11-07 16:07:23',
        '2022-11-07 16:07:23', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (74, '提示语管理', 'prompt', 48, '/prompt-list', 'GET', 'forum/prompt/index', 'form', 7, 'normal',
        '2022-11-11 16:58:20', '2022-11-12 13:31:48', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (75, '新增', '', 74, '/prompt-store', 'POST', '', '', 0, 'normal', '2022-11-11 16:59:00', '2022-11-11 16:59:00',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (76, '删除', '', 74, '/prompt-destroy', 'DELETE', '', '', 0, 'normal', '2022-11-11 16:59:15',
        '2022-11-11 16:59:15', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (77, '更新', '', 74, '/prompt-update', 'PUT', '', '', 0, 'normal', '2022-11-11 16:59:35', '2022-11-11 16:59:35',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (78, '邮件发送记录', 'email-record', 48, '/email-record-list', 'GET', 'forum/email-record/index', 'email', 8, 'normal',
        '2022-11-12 13:30:41', '2022-11-12 13:31:58', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (79, '删除', '', 78, '/email-record-destroy', 'DELETE', '', '', 0, 'normal', '2022-11-12 13:31:01',
        '2022-11-12 13:31:01', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (80, '节点分类管理', 'node-category', 85, '/node-category-list', 'GET', 'forum/node-category/index', 'build', 0,
        'normal', '2022-11-24 18:52:43', '2022-11-28 15:01:55', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (81, '新增', '', 80, '/node-category-store', 'POST', '', '', 0, 'normal', '2022-11-24 18:53:23',
        '2022-11-24 18:53:23', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (82, '删除', '', 80, '/node-category-destroy', 'DELETE', '', '', 0, 'normal', '2022-11-24 18:53:59',
        '2022-11-24 18:53:59', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (83, '更新', '', 80, '/node-category-update', 'PUT', '', '', 0, 'normal', '2022-11-24 18:55:00',
        '2022-11-24 18:55:00', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (84, '详情', '', 80, '/node-category-info', 'GET', '', '', 0, 'normal', '2022-11-24 18:59:48',
        '2022-11-24 18:59:48', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (85, '节点管理', 'node-manage', 0, '', 'GET', '', 'cascader', 7, 'normal', '2022-11-28 15:01:11',
        '2022-11-29 13:47:19', 'directory', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (86, '主题和回复', 'post-reply', 0, '', 'GET', '', 'form', 6, 'normal', '2022-11-28 15:03:19', '2022-11-29 13:47:23',
        'directory', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (87, '用户管理', 'user-manage', 0, '', 'GET', '', 'peoples', 5, 'normal', '2022-11-29 13:19:43',
        '2022-11-29 13:47:37', 'directory', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (88, '财务记录', 'balance-log', 87, '/balance-log-list', 'GET', 'forum/balance-log/index', 'money', 0, 'normal',
        '2022-11-29 13:21:53', '2022-11-29 13:29:40', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (89, '删除', '', 88, '/balance-log-destroy', 'GET', '', '', 0, 'normal', '2022-11-29 13:23:37',
        '2022-11-29 13:23:37', 'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (90, '关联列表', 'association', 48, '/association-list', 'GET', 'forum/association/index', 'swagger', 5, 'normal',
        '2022-11-29 17:11:52', '2022-11-29 17:11:52', 'link', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (91, '管理', '', 62, '/post-update', 'PUT', '', '', 0, 'normal', '2022-12-06 15:29:51', '2022-12-06 15:29:51',
        'operation', 'internal');
INSERT INTO `ga_admin_menu` (`id`, `name`, `path`, `parent_id`, `identification`, `method`, `front_component_path`,
                             `icon`, `sort`, `status`, `created_at`, `updated_at`, `type`, `link_type`)
VALUES (92, '管理', '', 65, '/reply-update', 'PUT', '', '', 0, 'normal', '2022-12-06 16:31:17', '2022-12-06 16:31:17',
        'operation', 'internal');
COMMIT;

-- ----------------------------
-- Table structure for ga_administrator
-- ----------------------------
DROP TABLE IF EXISTS `ga_administrator`;
CREATE TABLE `ga_administrator`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`        varchar(45)      NOT NULL COMMENT '用户名',
    `password`        char(32)         NOT NULL COMMENT 'MD5密码',
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
) ENGINE = InnoDB
  AUTO_INCREMENT = 20
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='管理员表';

-- ----------------------------
-- Records of ga_administrator
-- ----------------------------
BEGIN;
INSERT INTO `ga_administrator` (`id`, `username`, `password`, `nickname`, `avatar`, `status`, `remark`, `last_login_ip`,
                                `last_login_date`, `created_at`, `updated_at`)
VALUES (1, 'admin', 'a66abb5684c45962d887564f08346e8d', 'admin', '/upload/backend/avatar/1/copgl6ijn8ootfxb4s',
        'normal', 'sssss', '', NULL, '2022-01-09 19:38:04', '2022-11-30 15:40:45');
COMMIT;

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
) ENGINE = InnoDB
  AUTO_INCREMENT = 34
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='管理员角色关联表';

-- ----------------------------
-- Records of ga_administrator_role
-- ----------------------------
BEGIN;
INSERT INTO `ga_administrator_role` (`id`, `administrator_id`, `role_id`, `created_at`, `updated_at`)
VALUES (4, 1, 1, '2022-02-19 19:33:49', '2022-02-19 19:33:49');
COMMIT;

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
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Records of ga_casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('g', 'admin', 'super', '', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/administrator-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/role-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/role-store', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/role-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/role-destroy', 'delete', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/administrator-store', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/administrator-destroy', 'delete', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/administrator-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/menu-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/menu-store', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/menu-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/menu-destroy', 'delete', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/administrator-info', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/role-info', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/menu-info', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/config-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/config-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/operation-log-list', 'get', '', '', '');

INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/node-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/node-store', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/node-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/node-destroy', 'delete', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/node-info', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/node-upload-img', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('p', 'super', '/node-download-img', 'post', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for ga_config
-- ----------------------------
DROP TABLE IF EXISTS `ga_config`;
CREATE TABLE `ga_config`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `module`     varchar(255)     NOT NULL DEFAULT '' COMMENT '所属模块',
    `key`        varchar(255)     NOT NULL DEFAULT '' COMMENT '键值',
    `value`      text                      DEFAULT NULL COMMENT '值',
    `created_at` datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `module_key_idx` (`module`, `key`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1324
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='配置表';

-- ----------------------------
-- Records of ga_config
-- ----------------------------
BEGIN;
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (228, 'forum', 'posts_character_max', '20000', '2022-09-13 16:43:51', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (229, 'forum', 'posts_every_day_max', '10', '2022-09-13 16:43:51', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (230, 'forum', 'token_establish_posts_deduct', '10', '2022-09-13 16:43:51', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (236, 'forum', 'token_register_give', '1000', '2022-09-14 16:52:32', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (237, 'forum', 'token_login_give', '20', '2022-09-14 16:52:32', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (243, 'forum', 'posts_is_need_audit', '0', '2022-10-30 21:22:54', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (248, 'forum', 'posts_can_update_time', '10', '2022-10-30 21:24:53', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (258, 'forum', 'posts_can_update_reply_amount', '2', '2022-10-30 21:25:01', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (262, 'forum', 'reply_is_need_audit', '0', '2022-10-30 21:25:22', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (263, 'forum', 'reply_character_max', '500', '2022-10-30 21:25:22', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (264, 'forum', 'reply_every_day_max', '100', '2022-10-30 21:25:22', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (271, 'forum', 'token_thanks_reply_deduct', '10', '2022-10-30 21:25:33', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (272, 'forum', 'token_update_posts_deduct', '10', '2022-10-30 21:25:33', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (273, 'forum', 'token_establish_reply_deduct', '10', '2022-10-30 21:25:33', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (277, 'forum', 'token_thanks_posts_deduct', '10', '2022-10-30 21:25:33', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (296, 'forum', 'logo', '/upload/backend/logo/coq94pe5l640dtg2ih.png?v0.36217617248247724', '2022-10-31 22:08:26',
        '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (311, 'forum', 'site_name', 'cp-v2', '2022-10-31 22:08:41', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (335, 'forum', 'site_domain', 'http://localhost:8201/', '2022-11-01 12:49:50', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (354, 'forum', 'email_password', '', '2022-11-01 13:31:08', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (365, 'forum', 'email_user', '', '2022-11-01 13:31:08', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (366, 'forum', 'email_host', '', '2022-11-01 13:31:08', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (397, 'forum', 'email_sendname', 'cp-v2', '2022-11-01 13:31:42', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (420, 'forum', 'email_send_name', 'cp-v2', '2022-11-01 13:33:25', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (676, 'forum', 'register_send_email_diff_hour', '6', '2022-11-12 20:16:57', '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (703, 'forum', 'register_default_avatar', '/upload/frontend/avatar/default.png', '2022-11-15 15:40:28',
        '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (781, 'forum', 'site_description', '该网站完全模仿v2ex,只为学习使用,v2ex网站地址为www.v2ex.com', '2022-11-24 17:20:59',
        '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (828, 'forum', 'site_slogan', '该网站完全模仿v2ex,只为学习使用\n<br>\nv2ex网站地址为www.v2ex.com', '2022-11-25 14:48:36',
        '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (978, 'forum', 'default_avatar', '/upload/frontend/avatar/default.png', '2022-11-28 23:19:11',
        '2022-12-01 15:08:37');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (984, 'backend', 'is_open_verify_captcha', '0', '2022-11-28 23:40:16', '2022-11-28 23:40:16');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`)
VALUES (1256, 'forum', 'favicon', '/upload/backend/favicon/coq9jr1iu1oos4emdx.png?v0.5803719275527459',
        '2022-12-01 14:21:06', '2022-12-01 15:08:37');
COMMIT;

-- ----------------------------
-- Table structure for ga_login_log
-- ----------------------------
DROP TABLE IF EXISTS `ga_login_log`;
CREATE TABLE `ga_login_log`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
    `ip`               varchar(30)      NOT NULL DEFAULT '' COMMENT 'ip地址',
    `browser`          varchar(10)      NOT NULL DEFAULT '' COMMENT '浏览器',
    `os`               varchar(255)     NOT NULL DEFAULT '' COMMENT '操作系统',
    `status`           varchar(10)               DEFAULT 'normal' COMMENT '状态 success 成功 fail 失败',
    `created_at`       datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`       datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- ----------------------------
-- Records of ga_login_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ga_role
-- ----------------------------
DROP TABLE IF EXISTS `ga_role`;
CREATE TABLE `ga_role`
(
    `id`             int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`           varchar(45)      NOT NULL COMMENT '角色名',
    `identification` varchar(20)      NOT NULL COMMENT '角色标识符',
    `sort`           tinyint(4)       NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
    `status`         varchar(10)               DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `created_at`     datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 12
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='角色表';

-- ----------------------------
-- Records of ga_role
-- ----------------------------
BEGIN;
INSERT INTO `ga_role` (`id`, `name`, `identification`, `sort`, `status`, `created_at`, `updated_at`)
VALUES (1, '超级管理员', 'super', 0, 'normal', '2022-01-09 22:32:23', '2022-10-29 21:49:21');
COMMIT;

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
) ENGINE = InnoDB
  AUTO_INCREMENT = 941
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='角色和菜单权限关联表';

-- ----------------------------
-- Records of ga_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (892, 1, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (893, 2, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (894, 7, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (895, 8, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (896, 9, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (897, 18, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (898, 3, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (899, 4, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (900, 5, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (901, 6, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (902, 19, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (903, 10, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (904, 11, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (905, 12, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (906, 13, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (907, 20, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (908, 22, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (909, 23, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (910, 24, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (911, 48, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (912, 49, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (913, 50, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (914, 53, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (915, 55, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (916, 51, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (917, 54, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (918, 52, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (919, 26, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (920, 32, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (921, 33, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (922, 34, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (923, 36, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (924, 35, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (925, 27, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (926, 37, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (927, 38, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (928, 39, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (929, 28, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (930, 42, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (931, 43, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (932, 29, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (933, 44, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (934, 45, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (935, 30, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (936, 40, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (937, 41, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (938, 31, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (939, 46, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
INSERT INTO `ga_role_menu` (`id`, `menu_id`, `role_id`, `created_at`, `updated_at`)
VALUES (940, 47, 1, '2022-10-29 21:49:22', '2022-10-29 21:49:22');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
