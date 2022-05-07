package models

import (
	"github.com/TruthHun/BookStack/oauth"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var ModelIAM = new(IAM)

type IAM struct {
	oauth.IAMUser
}

func (u *IAM) TableName() string {
	// db table name
	return "iam"
}

//iam用户的登录流程是这样的
//1、获取iam的用户信息，用iam的用户id查询member_id是否大于0，大于0则表示已绑定了用户信息，直接登录
//2、未绑定用户，先把iam数据入库，然后再跳转绑定页面

//根据iamid获取用户的iam数据。这里可以查询用户是否绑定了或者数据是否在库中存在
func (this *IAM) GetUserByIAMId(id int, cols ...string) (user IAM, err error) {
	qs := orm.NewOrm().QueryTable("md_iam").Filter("id", id)
	if len(cols) > 0 {
		err = qs.One(&user, cols...)
	} else {
		err = qs.One(&user)
	}
	return
}

//绑定用户
func (this *IAM) Bind(iamId, memberId interface{}) (err error) {
	beego.Debug("绑定iam账号", iamId, memberId)
	_, err = orm.NewOrm().QueryTable("md_iam").Filter("id", iamId).Update(orm.Params{"member_id": memberId})
	return
}
