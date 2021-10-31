package services

import "github.com/axiaoxin-com/pink-lady/models"

type UserRepo struct {
	model models.User //用户模型
	base  curd        //通用的增删改查
}

// func RegisterUserRepo() UserRepo {
// 	var repo UserRepo
// 	// repo.RepoBase.session = DB()
// 	return *repo
// }
