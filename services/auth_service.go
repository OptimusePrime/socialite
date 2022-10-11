package services

import (
	"context"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/argon2"
	"math/rand"
	"os"
	"socialite/dto"
	"socialite/ent"
	"socialite/ent/user"
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
	params = &Argon2Parameters{}

	if len(values) != 6 {
		return nil, nil, nil, ErrInvalidPasswordHash
	}
	fmt.Println("Value 1:", values[0])
	fmt.Println("Value 2:", values[1])
	fmt.Println("Value 3:", values[2])
	fmt.Println("Value 4:", values[3])

	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}

	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleArgon2Version
	}

	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &params.memory, &params.iterations, &params.parallelism)
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

func GenerateAccessToken(accessClaimsDto dto.UserJWTAccessTokenClaims) (err error, accessTokenString string) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["userId"] = accessClaimsDto.UserID
	accessClaims["exp"] = time.Now().Add(15 * time.Minute).Unix()

	accessTokenString, err = accessToken.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return err, ""
	}

	return nil, accessTokenString
}

func GenerateRefreshToken(refreshClaimsDto dto.UserJWTRefreshTokenClaims) (err error, refreshTokenString string) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["userId"] = refreshClaimsDto.UserID
	refreshClaims["exp"] = time.Now().Add(168 * time.Hour).Unix()

	refreshTokenString, err = refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return err, ""
	}

	return nil, refreshTokenString
}

func GenerateJWTPair(accessClaimsDto dto.UserJWTAccessTokenClaims, refreshClaimsDto dto.UserJWTRefreshTokenClaims) (err error, accessTokenString string, refreshTokenString string) {
	err, accessTokenString = GenerateAccessToken(accessClaimsDto)
	if err != nil {
		return err, "", ""
	}

	err, refreshTokenString = GenerateRefreshToken(refreshClaimsDto)
	if err != nil {
		return err, "", ""
	}

	return nil, accessTokenString, refreshTokenString
}

func validateJWT /*[T jwt.StandardClaims]*/ (tokenString, secretKey string, signingMethod jwt.SigningMethod) (err error, isValid bool, userId *uuid.UUID) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingMethod {
			return nil, ErrUnexpectedJWTSigningMethod
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return err, false, nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return err, false, nil
	}

	id, err := uuid.Parse(claims["userId"].(string))
	if err != nil {
		return err, false, nil
	}

	return nil, true, &id
}

func ValidateJWTAccessToken(tokenString string) (err error, isValid bool, userId *uuid.UUID) {
	return validateJWT /*[dto.UserJWTAccessTokenClaims]*/ (tokenString, os.Getenv("JWT_ACCESS_SECRET"), jwt.SigningMethodHS256)
}

func ValidateJWTRefreshToken(tokenString string) (err error, isValid bool, userId *uuid.UUID) {
	return validateJWT /*[dto.UserJWTRefreshTokenClaims]*/ (tokenString, os.Getenv("JWT_REFRESH_SECRET"), jwt.SigningMethodHS256)
}

func LoginUser(db *ent.Client, ctx context.Context, loginInfo dto.LoginUserDTO) (err error, accessToken, refreshToken string, isMatch bool) {
	foundUser, err := db.User.Query().Where(user.EmailEQ(loginInfo.Email)).First(ctx)
	fmt.Println(foundUser, err, loginInfo)
	if err != nil {
		return err, "", "", false
	}

	isMatch, err = ComparePasswordAndHash(loginInfo.Password, foundUser.Password)
	if err != nil || !isMatch {
		return err, "", "", false
	}

	err, accessToken, refreshToken = GenerateJWTPair(dto.UserJWTAccessTokenClaims{
		UserID: foundUser.ID,
	}, dto.UserJWTRefreshTokenClaims{
		UserID: foundUser.ID,
	})
	if err != nil {
		return err, "", "", false
	}

	return nil, accessToken, refreshToken, true
}

func RefreshUserAccessToken(db *ent.Client, ctx context.Context, refreshToken string) (err error, isValid bool, accessToken string) {
	err, isValid, userId := ValidateJWTRefreshToken(refreshToken)
	if err != nil || !isValid {
		return err, false, ""
	}

	_, err = db.User.Query().Where(user.ID(*userId)).First(ctx)
	if err != nil {
		return err, false, ""
	}

	err, accessToken = GenerateAccessToken(dto.UserJWTAccessTokenClaims{
		UserID: *userId,
	})
	if err != nil {
		return err, false, ""
	}

	return nil, true, accessToken
}

func GetBearerToken(ctx echo.Context) (err error, token string) {
	auth := ctx.Request().Header.Get("Authorization")
	split := strings.Split(auth, " ")
	if split[0] != "Bearer" {
		return ErrInvalidBearerToken, ""
	}
	token = split[1]

	return nil, token
}
