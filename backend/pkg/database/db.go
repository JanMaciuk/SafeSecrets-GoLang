package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	g *gorm.DB
}

func NewDB(path string) (*DB, error) {
	g, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := g.AutoMigrate(&Secret{}); err != nil {
		return nil, err
	}

	return &DB{
		g: g,
	}, nil
}

func (d *DB) GetSecretByID(id int) (*Secret, error) {
	var res Secret
	err := d.g.First(&res, id).Error
	return &res, err
}

func (d *DB) InsertSecret(secret *Secret) error {
	return d.g.Create(secret).Error
}

func (d *DB) UpdateSecret(secret *Secret) error {
	return d.g.Model(&Secret{ID: secret.ID}).Updates(secret).Error
}

func (d *DB) RemoveSecret(secret *Secret) error {
	return d.g.Delete(secret).Error
}
