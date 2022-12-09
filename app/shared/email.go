package shared

import (
	"context"
	"errors"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"net/smtp"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/frame/g"
)

var Email = email{}

type email struct {
}

// SendAsync 异步发送邮件
func (e *email) SendAsync(ctx context.Context, in *model.SendEmailInput) {
	go e.SendSync(ctx, in)
}

// SendSync 同步发送邮件
func (e *email) SendSync(ctx context.Context, in *model.SendEmailInput) error {

	configs, err := Config.Gets(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_EMAIL_HOST, model.CONFIG_EMAIL_USERNAME, model.CONFIG_EMAIL_PASSWORD, model.CONFIG_EMAIL_SEND_USERNAME)
	if err != nil {
		return err
	}
	from := configs[model.CONFIG_EMAIL_USERNAME].String()
	password := configs[model.CONFIG_EMAIL_PASSWORD].String()
	host := configs[model.CONFIG_EMAIL_HOST].String()
	sendUserName := configs[model.CONFIG_EMAIL_SEND_USERNAME].String() //发送邮件的人名称

	g.Dump(&model.SendOriginEmailInput{
		From:         from,
		SendUserName: sendUserName,
		Password:     password,
		Host:         host,
		To:           in.To,
		Subject:      in.Subject,
		Body:         in.Body,
		IsHtml:       in.IsHtml,
	})
	err = e.SendOrigin(ctx, &model.SendOriginEmailInput{
		From:         from,
		SendUserName: sendUserName,
		Password:     password,
		Host:         host,
		To:           in.To,
		Subject:      in.Subject,
		Body:         in.Body,
		IsHtml:       in.IsHtml,
	})

	errMessage := ""
	if err != nil {
		errMessage = err.Error()
	}
	_, err1 := dao.EmailRecords.Ctx(ctx).Insert(entity.EmailRecords{
		UserId:   in.UserId,
		Username: in.Username,
		Type:     gconv.String(in.Type),
		Email:    in.To,
		Title:    in.Subject,
		Content:  in.Body,
		Error:    errMessage,
	})
	if err1 != nil {
		return err1
	}
	return err
}

/**
 * 发送邮件
 * @param from 发送人
 * @param sendUserName 发送人名称
 * @param password 发送人密码 不是邮箱密码,需要登陆你的邮箱，在设置，账号，启用IMAP/SMTP服务，会发送一段身份验证符号给你，用这个登陆
 * @param host 发送人邮箱服务器地址
 * @param to 接收人,多个以;号隔开
 * @param subject 主题
 * @param body 内容
 * @param isHtml 是否是html格式
 */
// This is using the authentication method PLAIN. Unfortunately smtp-mail.outlook.com does not support this authentication method: 具体见：https://stackoverflow.com/questions/57783841/how-to-send-email-using-outlooks-smtp-servers

func (e *email) SendOrigin(ctx context.Context, in *model.SendOriginEmailInput) error {

	hp := strings.Split(in.Host, ":")

	auth := smtp.PlainAuth("", in.From, in.Password, hp[0])

	// 根据host判断是否是outlook 邮箱，如果是的话修改授权方式
	if gstr.Trim(in.Host) == "smtp.office365.com:587" {
		auth = LoginAuth(in.From, in.Password)
	}

	var content_type string

	if in.IsHtml {
		content_type = "Content-Type: text/html; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain; charset=UTF-8"
	}

	msg := []byte("To: " + in.To + "\r\nFrom: " + in.SendUserName + "<" + in.From + ">" + "\r\nSubject: " + in.Subject + "\r\n" + content_type + "\r\n\r\n" + in.Body)

	targets := strings.Split(in.To, ";")

	err := smtp.SendMail(in.Host, auth, in.From, targets, msg)
	return err
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}
