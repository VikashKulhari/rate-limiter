package ratelimiter

import (
	"errors"
	"net"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	UserID string
	Role   string
	Email  string
	Expiry int64
}

// ExtractClaimsFromJWT fetches the userid from request
func ExtractClaimsFromJWT(r *http.Request) (*JWTClaims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("no token provided")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	userID, _ := claims["sub"].(string)
	role, _ := claims["role"].(string)
	email, _ := claims["email"].(string)
	exp, _ := claims["exp"].(float64)      // JWT stores `exp` as float

	return &JWTClaims{
		UserID: userID,
		Role:   role,
		Email:  email,
		Expiry: int64(exp),
	}, nil
}

// GetIPAddress fetches IP address
func GetIPAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
		if strings.Contains(ip, ":") {
			ip, _, _ = net.SplitHostPort(ip)
		}
	}
	return ip
}
