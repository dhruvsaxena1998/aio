package middlewares

import "gorm.io/gorm"

type Middleware struct {
	db *gorm.DB
}

func NewMiddleware(db *gorm.DB) *Middleware {
	return &Middleware{
		db: db,
	}
}
