package services

/*import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var params = Argon2Parameters{
	memory:      6 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

func testHashPassword(t *testing.T, password string) (encodedHash string) {
	encodedHash, err := HashPassword(password, params)
	assert.NoError(t, err, "hashing error should be nil")
	assert.NotEqual(t, password, encodedHash, "passwords should not equal")
	return encodedHash
}

func testRandomBytes(t *testing.T, length uint32) (randomBytes string) {
	bytes, err := RandomBytes(length)
	assert.NoError(t, err, "random bytes should be nil")
	return string(bytes)
}

func testDecodeHash(t *testing.T, encodedHash string) (hash []byte) {
	_, _, hash, err := DecodeHash(encodedHash)
	assert.NoError(t, err, "decode hash error should be nil")
	assert.Equal(t)
	return hash
}

func TestHashPassword(t *testing.T) {
	testHashPassword(t, testRandomBytes(t, 16))
}

func TestDecodeHash(t *testing.T) {
	encodedHash := testHashPassword(t, testRandomBytes(t, 16))
	testDecodeHash(t)
}
*/
