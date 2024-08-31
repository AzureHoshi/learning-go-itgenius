package models

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// passwordMatch ฟังก์ชันสำหรับตรวจสอบรหัสผ่าน
func (u *User) PasswordMatches(plaintext string) (bool, error) {
	log.Println("u.Password:", u.Password)
	log.Println("plaintext:", plaintext)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plaintext))
	if err != nil {
		log.Println("error comparing password", err)
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// รหัสผ่านไม่ตรงกัน
			return false, nil
		default:
			// มีข้อผิดพลาดอื่น ๆ
			return false, err

		}
	}

	return true, nil
}
