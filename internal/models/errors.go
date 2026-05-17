package models

import "errors"

var (
	ErrAccountNotFound = errors.New("không tìm thấy tài khoản")
	ErrUsernameExists  = errors.New("tên đăng nhập đã tồn tại")
	ErrInvalidPassword = errors.New("mật khẩu không đúng")
)
