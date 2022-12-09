package model

type EmailType string

const (
	EMAIL_TYPE_REGISTER        EmailType = "register"
	EMAIL_TYPE_FORGET_PASSWORD EmailType = "forget_password"
	EMAIL_TYPE_CHANGE_EMAIL    EmailType = "change_email"
)
