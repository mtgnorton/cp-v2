package shared

//func TestEmailsendOriginByQQ(t *testing.T) {
//	gtest.C(t, func(t *gtest.T) {
//		user := "851426308@qq.com"
//		password := "mrnjlfbtvhobbdcd"
//		host := "smtp.qq.com:587"
//		to := "15726204663@163.com"
//
//		subject := "TestEmailsendOriginByQQ"
//
//		body := `
//		<!DOCTYPE html>
//		<html lang="en">
//		<head>
//			<meta charset="iso-8859-15">
//			<title>MMOGA POWER</title>
//		</head>
//		<body>
//			GO 发送邮件，官方连包都帮我们写好了，真是贴心啊！！！
//		</body>
//		</html>`
//
//		sendUserName := "GOLANG SEND MAIL" //发送邮件的人名称
//		fmt.Println("send email")
//		err := Email.sendOrigin(user, sendUserName, password, host, to, subject, body, true)
//		t.AssertNil(err)
//	})
//}
//func TestEmail_SendByQQAsync(t *testing.T) {
//	gtest.C(t, func(t *gtest.T) {
//		Email.SendByQQAsync("TestEmail_SendByQQAsync", "TestEmail_SendByQQAsync", "15726204663@163.com", false)
//	})
//
//}
//
//func TestEmail_SendByQQSync(t *testing.T) {
//	gtest.C(t, func(t *gtest.T) {
//		Email.SendByQQSync("TestEmail_SendByQQAsync", "TestEmail_SendByQQAsync", "15726204663@163.com", false)
//	})
//}
