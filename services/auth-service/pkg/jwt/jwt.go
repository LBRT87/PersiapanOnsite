package jwt

import "honnef.co/go/tools/analysis/facts/generated"

type Manager struct {
	secret []byte
}

type Claims struct {
	UserID uint
	Role string
	jwt.RegisteredClaims 
}

func (m *Manager) generate (userID uint,role string)