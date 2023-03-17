package repo

import (
	u "TEMPLATE_MICROSERVICE/genproto/user"
	
)

type UserStorageI interface {
	CreateUser(*u.UserRequest) (*u.UserResponse, error)
	GetUserById(*u.UserId) (*u.UserResponse, error)
	GetUsersAll(*u.UserListReq) (*u.Users, error)
	DeleteUser(*u.UserId)(*u.Users, error)
	UpdateUser(*u.UserUpdateReq)(*u.UserResponse, error)
	SearchUser(*u.UserSearch)(*u.Users, error)
}
