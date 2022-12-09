package model

//  from, sendUserName, password, host, to, subject, body string, isHtml bool
type SendOriginEmailInput struct {
	From, SendUserName, Password, Host, To, Subject, Body string
	IsHtml                                                bool
}

//  subject, body, to string, isHtml bool
type SendEmailInput struct {
	Subject, Body, To string
	IsHtml            bool
	Type              EmailType
	Username          string //可选
	UserId            uint   //可选
}
