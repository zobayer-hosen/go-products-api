package auth
import auth

import(
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zobayer-hosen/go-products-api/internal/model"

)

// jwtCustomerClaims embeds the standard RegisteredClaims and adds our customer fields
type jwtCustomClaims struct {
	model.Claims
	jwt.RegisteredClaims
}

//Manager handles all JWT operations using a typed struct, not a global variable
type Manager struct{
	secretKey []byte
	expiryHours int
}

//newManager creates a new JWT Manager
func NewManager(secret string, expirHours int) *Manager{
	return &Manager{
		secretKey: []byte(secret),
		expiryHours: expiryHours,
	}
}

//Generate creates a signed jwt string for a given user
func (m *Manager) Generate(userID int64, email,role string)(string,error){
	claims := jwtCustomClaims{
		Claims:model.Claims{
			UserID: userID,
			Email: email,
			Role:role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericData(time.Now().Add(time.Duration(m.expiryHours)*time.Hour)),
			IssuedAt: jwt.NewNumericData(time.Now()),
			Issuer:"go-products-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", fmt.Errorf("auth:faild to sign token: %w",err)
	}
	return signedToken, nil
}
//Vlaidate parses and validates a JWT string, returning the embedded claims
func (m *Manager) Validate(tokenStr string)(*model.Claims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &jwtCustomClaims{}, func(token *jwt.Token)(any,error){
		if _,ok := token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil, fmt.Errorf("auth: unexpected signing method: %v")
		}
		return m.secretKey,nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired){
			return nil, fmt.Errorf("auth: token has expired")
		}
		return nil, fmt.Errorf("auth:Invalid token:%w",err)
	}
	claims, ok := token.Claims.(*jwtCustomClaims)
	if !of || !token.Valid {
		return nil, fmt.Errorf("auth: token claims are invalid")
	}

	return &claims.Claims, nil

}