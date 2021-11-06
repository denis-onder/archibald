package builders

import (
	"crypto/rand"

	. "github.com/denis-onder/archibald/domain"
	"golang.org/x/crypto/argon2"
)

// Argon2 params
type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func hashPassword(password string) (string, error) {
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return string(hash), nil
}

func UserBuilder(firstName, lastName, email, password string) User {
	user := User{}

	hash, err := hashPassword(password)

	if err != nil {
		// handle error
	}

	user.SetFirstName(firstName)
	user.SetLastName(lastName)
	user.SetDisplayName(firstName, lastName)
	user.SetEmail(email)
	user.SetPassword(hash)

	return user
}
