package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvaloidToken = errors.New("token is invalid")
	ErrExpiredToken  = errors.New("token has expired")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (pyl *Payload) Valid() error {
	if time.Now().After(pyl.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
