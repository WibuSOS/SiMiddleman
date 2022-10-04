package models

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Nama     string `json:"nama" gorm:"not null;type:varchar(30)"`
	Role     string `json:"role" gorm:"not null;type:varchar(15);default:consumer"`
	NoHp     string `json:"noHp" gorm:"type:varchar(18)"`
	Email    string `json:"email" gorm:"unique;not null;type:varchar(30)"`
	Password string `json:"password" gorm:"not null;type:varchar(128)"`
	NoRek    string `json:"noRek" gorm:"type:varchar(18)"`
}

func (user *Users) ValidateUser() error {
	if user.Nama == "" || len(user.Nama) > 30 {
		return fmt.Errorf("nama tidak memenuhi syarat")
	}

	if (user.Role != "consumer" && user.Role != "admin") || len(user.Role) > 15 {
		return fmt.Errorf("role tidak memenuhi syarat")
	}

	regex := regexp.MustCompile(`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,7}$`)
	if matched := regex.MatchString(user.NoHp); user.NoHp == "" || !matched || len(user.NoHp) > 18 {
		return fmt.Errorf("no hp tidak memenuhi syarat")
	}

	regex = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if matched := regex.MatchString(user.Email); user.Email == "" || !matched || len(user.Email) > 30 {
		return fmt.Errorf("email tidak memenuhi syarat")
	}

<<<<<<< HEAD
	if user.Password == "" || (len(user.Password) > 18 && len(user.Password) >= 8) {
=======
	if user.Password == "" || len(user.Password) > 18 || len(user.Password) < 8 {
>>>>>>> 1734b93604cd615ec267b48b4dc01c1bdaa5efd0
		return fmt.Errorf("password tidak memenuhi syarat")
	}

	regex = regexp.MustCompile(`^[0-9]{4,18}$`)
	if matched := regex.MatchString(user.NoRek); user.NoRek == "" || !matched || len(user.NoRek) > 18 {
		return fmt.Errorf("no rek tidak memenuhi syarat")
	}

	return nil
}
