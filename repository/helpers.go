package repository

import (
	pg "github.com/clubo-app/protobuf/profile"
)

func (p Profile) ToGRPCProfile() *pg.Profile {
	return &pg.Profile{
		Id:        p.ID,
		Username:  p.Username,
		Firstname: p.Firstname,
		Lastname:  p.Lastname.String,
		Avatar:    p.Avatar.String,
	}
}
