package controllers

import (
	"DataCerProject/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	//、解析请求数据
	var user models.User
	err :=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析错误，请重试！")
		return
	}
	//、保存用户数据到数据库
	_,err =user.SaveUser()
	//、返回前端结果
	if err != nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试！")
		return
	}
	r.TplName = "login.html"
}
