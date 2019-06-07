package auth

import (
	"errors"
	"server/middleware/language"
	"server/models/database"
	"server/pkg/codes"
	"server/pkg/gdb"
)

func searchUser(userId uint) (user database.User, err error) {
	if err := gdb.Instance().Where("id = ?", userId).Find(&user).Error; err == nil {
		if len(user.Email) <= 0 {
			return user, errors.New(language.GetMsg(codes.ERROR_USER_NOT_SIGNIN))
		}
		return user, nil
	}

	return user, errors.New(language.GetMsg(codes.ERROR_USER_NOT_FOUND))
}

func searchVipUser(userId uint) (user database.User, err error) {
	if err := gdb.Instance().Where("id = ?", userId).Find(&user).Error; err == nil {
		if len(user.Email) <= 0 {
			return user, errors.New(language.GetMsg(codes.ERROR_USER_NOT_SIGNIN))
		} else if user.Status <= 1 {
			return user, errors.New(language.GetMsg(codes.ERROR_USER_NOT_VIP))
		}
		return user, nil
	}

	return user, errors.New(language.GetMsg(codes.ERROR_USER_NOT_FOUND))
}
