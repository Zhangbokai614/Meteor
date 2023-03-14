package service

import (
	"errors"

	"github.com/Zhangbokai614/go-template/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	errRouterNotExists = errors.New("Router not exists")
)

func CreateRole(r *model.APIRole) error {
	return dbConn.Transaction(func(tx *gorm.DB) error {
		role := &model.Role{
			Name: r.Name,
		}

		if err := tx.Model(role).Create(role).Error; err != nil {
			tx.Rollback()

			return err
		}

		pids := r.PerimssionsList
		if len(pids) == 0 {
			return nil
		}

		rps := []*model.RolePermissions{}
		for _, v := range pids {
			rps = append(rps, &model.RolePermissions{
				RID: role.ID,
				PID: v,
			})
		}

		if err := tx.Create(rps).Error; err != nil {
			tx.Rollback()

			return err
		}

		return nil
	})
}

func RoleNameIsExists(name string) bool {
	role := &model.Role{
		Name: name,
	}

	if r := dbConn.Model(role).Select("name").Where(role).Find(&role).RowsAffected; r > 0 {
		return true
	}

	return false
}

func RoleIDIsExists(rid uint) bool {
	role := &model.Role{
		ID: rid,
	}

	if r := dbConn.Model(role).Select("id").Where(role).Find(&role).RowsAffected; r > 0 {
		return true
	}

	return false
}

func QueryRole() (roles []*model.Role, err error) {
	if err = dbConn.Model(model.Role{}).Select("id, name").Scan(&roles).Error; err != nil {
		return nil, err
	}

	return
}

func QueryPermissions() (permissions []*model.Permissions, err error) {
	if err = dbConn.Model(model.Permissions{}).Select("id, router_permissions").Scan(&permissions).Error; err != nil {
		return nil, err
	}

	return
}

func DeleteRole(rid uint) (err error) {
	return dbConn.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(model.Role{}, "id = ?", rid).Error; err != nil {
			tx.Rollback()

			return err
		}

		if err = tx.Delete(model.RolePermissions{}, "r_id = ?", rid).Error; err != nil {
			tx.Rollback()

			return err
		}

		return nil
	})
}

func ModifyRolePermissions(r *model.APIRole) error {
	return dbConn.Transaction(func(tx *gorm.DB) error {
		role := &model.Role{
			Name: r.Name,
		}

		var id uint
		tx.Model(role).Select("id").Where(role).First(&id)

		pids := r.PerimssionsList

		if len(pids) == 0 {
			if err := tx.Delete(&model.RolePermissions{}, "r_id = ?", id).Error; err != nil {
				tx.Rollback()

				return err
			}
		} else {
			if err := tx.Delete(&model.RolePermissions{}, "r_id = ? and p_id not in ?", id, pids).Error; err != nil {
				tx.Rollback()

				return err
			}

			newPermissions := []*model.RolePermissions{}
			for _, v := range pids {
				newPermissions = append(newPermissions, &model.RolePermissions{
					RID: id,
					PID: v,
				})
			}

			if err := tx.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(newPermissions).Error; err != nil {
				tx.Rollback()

				return err
			}
		}

		return nil
	})
}

func ModifyUserRole(name string, rid uint) error {
	newRole := &model.User{RID: rid}
	if err := dbConn.Model(newRole).Where("name = ?", name).Update("r_id", newRole.RID).Error; err != nil {
		return err
	}

	return nil
}
