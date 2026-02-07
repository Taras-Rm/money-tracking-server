package hasher

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	cost int64
}

func NewHasher(cost int64) *Hasher {
	return &Hasher{
		cost,
	}
}

func (h *Hasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), int(h.cost))
	return string(bytes), err
}

func (h *Hasher) VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
