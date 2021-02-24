package database

import "gorm.io/gorm"

// User user
type User struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"name"`
}

// Create create user
func (u *User) Create() (err error) {
	var t User
	if err = engine.Model(new(User)).Where(
		User{Name: u.Name},
	).First(&t).Error; err == gorm.ErrRecordNotFound {
		return engine.Create(u).Error
	} else if err != nil {
		return
	}
	if err = engine.Model(new(User)).Where(User{Name: u.Name}).Updates(u).Error; err != nil {
		return
	}
	if err = engine.Model(new(User)).Where(User{Name: u.Name}).First(&u).Error; err != nil {
		return
	}
	return
}

// Find find all the questions
func (u *User) Find(options Options) (users []User, err error) {
	users = make([]User, 0)
	if options.Limit == 0 {
		options.Limit = pageSize
	}
	err = engine.Debug().Limit(options.Limit).Where(&User{Name: u.Name}).Offset(options.Offset).Find(&users).Error
	return
}
