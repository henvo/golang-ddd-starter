package postgres

import "github.com/henvo/golang-ddd-starter/internal/domain"

func ToDomainUser(u *User) *domain.User {
	return &domain.User{}
}

func FromDomainUser(u *domain.User) *User {
	return &User{}
}
