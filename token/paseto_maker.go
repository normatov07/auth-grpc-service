package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto      *paseto.V2
	symetricKey string
}

func NewPasetoMaker(symetricKey string) (Maker, error) {
	if len(symetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:      paseto.NewV2(),
		symetricKey: symetricKey,
	}
	return maker, nil
}

// CreateToken implements Maker
func (pst *PasetoMaker) CreateToken(email string, duration time.Duration) (string, error) {
	paylod, err := NewPayload(email, duration)
	if err != nil {
		fmt.Errorf("Error creating payload: %v", err)
	}
	return pst.paseto.Encrypt([]byte(pst.symetricKey), paylod, nil)
}

// CreateToken implements Maker
func (pst *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := pst.paseto.Decrypt(token, []byte(pst.symetricKey), payload, nil)
	if err != nil {
		return nil, fmt.Errorf("token is invaild")
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
