package models

import "time"

// User описывает зарегистрированного пользователя.
type User struct {
	ID        int
	Email     string
	Password  string
	CreatedAt time.Time
}

// Form описывает форму, созданную пользователем.
type Form struct {
	ID        string
	UserID    int
	Name      string
	Email     string
	CreatedAt time.Time
}

// Submission описывает отправленную внешнюю форму.
type Submission struct {
	ID        int
	FormID    string
	Data      string
	CreatedAt time.Time
}

