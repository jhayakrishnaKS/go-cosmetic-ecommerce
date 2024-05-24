package roles

import (
	"ecommerce/src/constants"
	"ecommerce/src/daos"
	"ecommerce/src/database/models"
	"ecommerce/src/dtos"
	"ecommerce/src/utils/context"

	"github.com/google/uuid"
)

type Role struct {
	role daos.RolesDAO
}

func New() *Role {
	return &Role{
		role: daos.NewRole(),
	}
}

func (r *Role) roleFromRegisterReq(req *dtos.RoleReq) *models.Roles {
	return &models.Roles{
		ID:   uuid.New().String(),
		Name: req.Name,
	}
}

func (r *Role) RegisterRoles(ctx *context.Context, req *dtos.RoleReq) error {
	role := r.roleFromRegisterReq(req)
	if ok, _ := r.role.CheckRoleExist(ctx, role.Name); ok {
		return constants.ErrRoleTaken
	}
	return r.role.Create(ctx, role)
}
