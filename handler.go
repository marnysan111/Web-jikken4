package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type Posts struct {
	Id       int
	Author   int
	Title    string
	Content  string
	Comments int
}

func UserGetAll() ([]Users, error) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("UserGetAll失敗: %w", err)
	}
	var users []Users
	err = db.Select("id", "name", "email", "password").Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("UserGetAll_Select失敗: %w", err)
	}
	return users, nil
}

func UserGetOne(id int) (Users, error) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	var user Users
	if err != nil {
		return Users{}, fmt.Errorf("UserGetOne失敗: %w", err)
	}
	err = db.First(&user, id).Error
	if err != nil {
		return Users{}, fmt.Errorf("UserGetOne_First失敗: %w", err)
	}
	return user, nil
}

func PostGetAll(author int) ([]Posts, error) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("PostGetAll失敗: %w", err)
	}
	var posts []Posts
	err = db.Where("author = ?", author).Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("PostGetAll_Select失敗: %w", err)
	}
	return posts, nil
}

func UserDelete(id int) error {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return fmt.Errorf("UserDelete失敗: %w", err)
	}
	var users Users
	err = db.Delete(&users, id).Error
	if err != nil {
		return fmt.Errorf("UserDelete_Delete失敗: %w", err)
	}
	return nil
}

func UserUpdate(id int, name string, email string) error {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return fmt.Errorf("UserUpdate失敗: %w", err)
	}
	var users Users
	err = db.First(&users, id).Error
	if err != nil {
		return fmt.Errorf("UserUpdate_First失敗: %w", err)
	}
	users.Name = name
	users.Email = email
	db.Save(&users)
	return nil
}
