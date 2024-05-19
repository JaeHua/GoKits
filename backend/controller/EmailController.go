package controller

import (
	"GinVue/common"
	"GinVue/model"
	"GinVue/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

// SendEmailValidate 发送邮件
func SendEmailValidate(em []string) (string, error) {
	e := email.NewEmail()
	e.From = fmt.Sprintf("GoKits <2624857134@qq.com>")
	e.To = em

	// 生成6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	t := time.Now().Format("2006-01-02 15:04:05")

	//设置文件发送的内容
	content := fmt.Sprintf(`
	<div style="font-family: Arial, sans-serif; background-color: #f5f5f5; padding: 20px; border-radius: 5px;">
	<div style="background-color: #ffffff; padding: 20px; border-radius: 5px; box-shadow: 0 0 10px rgba(0,0,0,0.1);">
	<div style="font-size: 18px; font-weight: bold; color: #333333; margin-bottom: 10px;">
	尊敬的 %s，您好！
	</div>
	<div style="padding: 8px 40px 8px 50px; background-color: #f2f2f2; border-radius: 5px;">
	<p style="color: #555555;">您于 %s 提交的邮箱验证，本次验证码为<u><strong style="font-size: 20px; color: #ff6600;">%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
	</div>
	<div style="margin-top: 10px;">
	<p style="color: #777777; font-size: 14px;">此邮箱为系统邮箱，请勿回复。</p>
	</div>
	</div>
	</div>
	`, em[0], t, vCode)
	e.Text = []byte(content)

	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2624857134@qq.com", "ggtwjomswvvheaci", "smtp.qq.com"))
	return vCode, err
}

// GetValidateCode 获取验证码
func GetValidateCode(ctx *gin.Context) {

	rd := common.GetRedis()

	// 获取目的邮箱
	//获取前端发来的验证码
	var mail = model.MailCode{}
	ctx.ShouldBind(&mail)
	em := []string{mail.Mail}
	//em := []string{ctx.PostForm("mail")}
	log.Println(em[0] + "hi")
	vCode, err := SendEmailValidate(em)
	if err != nil {
		log.Println(err)
		response.Response(ctx, http.StatusBadRequest, 400, nil, "验证码发送失败")
		return
	}

	// 验证码存入redis 并设置过期时间5分钟,注意key的唯一性
	parts := strings.Split(em[0], "@")
	err1 := rd.Db.Set("vCode"+parts[0], vCode, 300000000000).Err()

	if err1 != nil {
		log.Println(err.Error())
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "验证码存储失败")
		return
	}

	//发送成功
	response.Response(ctx, http.StatusOK, 200, gin.H{"vCode": vCode}, "验证码发送成功")

}

// ValidateEmailCode 验证邮箱
func ValidateEmailCode(ctx *gin.Context) {

	rd := common.GetRedis()
	//获取前端发来的验证码
	var em = model.MailCode{}
	ctx.ShouldBind(&em)

	vCode := em.VCode
	//log.Println("vCode", vCode)

	// 获取存储在redis中的验证码
	key := strings.Split("vCode"+em.Mail, "@")
	rdVCode := rd.Db.Get(key[0]).Val()
	//log.Println("key:rdvCode", key[0], rdVCode)

	if rdVCode != "" && vCode == rdVCode {
		response.Response(ctx, http.StatusOK, 200, nil, "验证成功")
		return
	} else {
		if rdVCode != "" && vCode != rdVCode {
			response.Response(ctx, http.StatusBadRequest, 400, nil, "验证码错误")
			return
		} else {
			response.Response(ctx, http.StatusBadRequest, 400, nil, "验证码失效")
			return
		}
	}
}
