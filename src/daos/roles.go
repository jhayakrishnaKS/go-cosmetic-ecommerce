package daos

import (
	"ecommerce/src/database/models"
	"ecommerce/src/utils/context"
	"log"
)

type RolesDAO interface {
	Create(ctx *context.Context, role *models.Roles) error
	CheckRoleExist(ctx*context.Context,name string)(bool,error)
}

type Roles struct {
}

func NewRole() RolesDAO {
	return &Roles{}
}

func (r *Roles) Create(ctx *context.Context, role *models.Roles) error {
	err := ctx.DB.Table("roles").Create(role).Error
	if err != nil {
		log.Println("Unable to create Role. Err:", err)
		return err
	}
	return nil
}

func(r *Roles)CheckRoleExist(ctx*context.Context,name string)(bool,error){
	var cnt int 
	err:=ctx.DB.Table("roles").Select("count(*)").Where("name=?",name).Scan(&cnt).Error
	if err != nil {
		log.Println("Unable to create role.Err:", err)
		return false,err
	}
	return cnt>0,nil
}