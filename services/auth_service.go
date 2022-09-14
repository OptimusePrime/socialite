package services

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"math/rand"
	"strings"
	"time"
)

type Argon2Parameters struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func RandomBytes(n uint32) (bytes []byte, err error) {
	bytes = make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	_, err = rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func HashPassword(password string, p Argon2Parameters) (encodedHash string, err error) {
	salt, err := RandomBytes(16)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		p.iterations,
		p.memory,
		p.parallelism,
		p.keyLength)

	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, base64Salt, base64Hash)

	return encodedHash, nil
}

func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	params, salt, hash, err := DecodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, params.iterations, params.memory, params.parallelism, params.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func DecodeHash(encodedHash string) (params *Argon2Parameters, salt, hash []byte, err error) {
	values := strings.Split(encodedHash, "$")

	if len(values) != 6 {
		return nil, nil, nil, ErrInvalidPasswordHash
	}

	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}

	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleArgon2Version
	}

	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &params.memory, params.iterations, params.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}

	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}

	params.keyLength = uint32(len(hash))

	return params, salt, hash, nil
}

/*func LoginUser(ctx echo.Context, db *ent.Client, c context.Context) error {

}*/
