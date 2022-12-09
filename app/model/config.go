package model

const (
	CONFIG_CONFIGS_PREFIX = "configs_"
	CONFIG_MODULE_FORUM   = "forum"
)

// 网站设置
const (
	CONFIG_FORUM_SITE_LOGO        = "logo"
	CONFIG_FORUM_SITE_FAVICON     = "favicon"
	CONFIG_FORUM_SITE_NAME        = "site_name"
	CONFIG_FORUM_SITE_DESCRIPTION = "site_description"
	CONFIG_FORUM_SITE_DOMAIN      = "site_domain"
	CONFIG_FORUM_SITE_SLOGAN      = "site_slogan"
)

// 邮箱设置

const (
	CONFIG_EMAIL_HOST          = "email_host"
	CONFIG_EMAIL_USERNAME      = "email_user"
	CONFIG_EMAIL_PASSWORD      = "email_password"
	CONFIG_EMAIL_SEND_USERNAME = "email_send_name"
)

const (
	// forum.posts_is_need_audit 发帖是否需要后台审核
	CONFIG_POSTS_IS_NEED_AUDIT = "posts_is_need_audit"

	// forum.posts_every_day_max 发帖每天最大数量
	CONFIG_POSTS_EVERY_DAY_MAX = "posts_every_day_max"

	// forum.posts_character_max 发帖字符最大数量
	CONFIG_POSTS_CHARACTER_MAX = "posts_character_max"

	// forum.posts_can_update_time 发帖可以修改的时间 ,如10min，则10min内可以修改
	CONFIG_POSTS_CAN_UPDATE_TIME = "posts_can_update_time"

	// forum.posts_can_update_reply_amount 发帖可以修改的回复数量,如2，则回复数量小于2可以修改

	CONFIG_POSTS_CAN_UPDATE_REPLY_AMOUNT = "posts_can_update_reply_amount"

	// forum.reply_is_need_audit 回帖是否需要后台审核
	CONFIG_REPLY_IS_NEED_AUDIT = "reply_is_need_audit"

	// forum.reply_every_day_max 回帖每天最大数量
	CONFIG_REPLY_EVERY_DAY_MAX = "reply_every_day_max"

	// forum.reply_character_max 回帖每个字符最大数量//
	CONFIG_REPLY_CHARACTER_MAX = "reply_character_max"

	// forum.token 赠送
	// forum.token_register_give 注册赠送token数量
	CONFIG_TOKEN_REGISTER_GIVE = "token_register_give"

	// forum.token_login_give 每日登录赠送token数量
	CONFIG_TOKEN_LOGIN_GIVE = "token_login_give"

	// forum.token 扣除
	// forum.token_establish_posts_deduct 发帖扣除token数量
	CONFIG_TOKEN_ESTABLISH_POSTS_DEDUCT = "token_establish_posts_deduct"

	// forum.token_update_posts_deduct 修改帖子扣除token数量
	CONFIG_TOKEN_UPDATE_POSTS_DEDUCT = "token_update_posts_deduct"

	// forum.token_establish_reply_deduct 回复扣除token数量
	CONFIG_TOKEN_ESTABLISH_REPLY_DEDUCT = "token_establish_reply_deduct"

	// forum.token_thanks_posts_deduct 感谢帖子扣除token数量
	CONFIG_TOKEN_THANKS_POSTS_DEDUCT = "token_thanks_posts_deduct"

	// forum.token_thanks_reply_deduct 感谢回复扣除token数量
	CONFIG_TOKEN_THANKS_REPLY_DEDUCT = "token_thanks_reply_deduct"
)

// 注册相关
const (
	CONFIG_REGISTER_SEND_EMAIL_DIFF_HOUR = "register_send_email_diff_hour"
	CONFIG_REGISTER_DEFAULT_AVATAR       = "register_default_avatar"
)
