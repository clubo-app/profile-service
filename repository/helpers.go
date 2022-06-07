package repository

import (
	"fmt"
	"net/url"

	pg "github.com/clubo-app/protobuf/profile"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	g "github.com/golang-migrate/migrate/v4/source/github"
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

const version = 1

func validateSchema(url url.URL) error {
	url.Scheme = "pgx"
	url2 := fmt.Sprintf("%v%v", url.String(), "?sslmode=disable")
	g := g.Github{}
	d, err := g.Open("github://clubo-app/profile-service/repository/migrations")
	if err != nil {
		return err
	}
	defer d.Close()

	m, err := migrate.NewWithSourceInstance("github", d, url2)

	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	defer m.Close()
	return nil
}
