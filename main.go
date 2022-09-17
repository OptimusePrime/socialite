package main

import "socialite/cmd"

func main() {
	cmd.Execute()
	/*	id := uuid.New()
		accessClaims := dto.UserJWTAccessTokenClaims{
			UserID: id,
		}
		refreshClaims := dto.UserJWTRefreshTokenClaims{
			UserID: id,
		}
		accessToken, refreshToken, err := services.GenerateJWTPair(accessClaims, refreshClaims)
		if err != nil {
			fmt.Println("Generate error:", err)
		}
		fmt.Println("Access token:", accessToken)
		fmt.Println("Refresh token:", refreshToken)

		isValid, err := services.ValidateJWTAccessToken(accessToken)
		if err != nil {
			log.Fatalf("Access error: %v", err)
		}
		fmt.Println("Is valid access:", isValid)
		isValid, err = services.ValidateJWTRefreshToken(refreshToken)
		if err != nil {
			log.Fatalf("Refresh error: %v", err)
		}
		fmt.Println("Is valid refresh:", isValid)*/

	/*	t, err := services.CreateJWTAccessToken(dto.JWTUserClaimsDTO{
			UserID: id,
		}, jwt.SigningMethodHS256)
		if err != nil {
			log.Fatal("Err 1:", err)
		}
		fmt.Println("Access Token:", t)
		fmt.Println("UserID 1:", id)

		str, err := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.JWTUserClaimsDTO{
			UserID: uuid.New(),
		}).SignedString([]byte("jeff"))
		_, userId, err := services.ValidateJWTAccessToken(str, jwt.SigningMethodHS256)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("UserID 2:", userId)*/
}
