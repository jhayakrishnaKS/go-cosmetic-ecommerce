package token

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetAccessAndRefreshToken(l int) (string, string) {
	accessToken := randStringBytes(l)
	refreshToken := randStringBytes(l)

	return accessToken, refreshToken
}

func GetAuthorizationCode(l int)string{
	return randStringBytes(l)
}
