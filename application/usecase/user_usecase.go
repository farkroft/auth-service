package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/farkroft/auth-service/application/request"
	"gitlab.com/farkroft/auth-service/application/response"
	"gitlab.com/farkroft/auth-service/external/constants"
	"gitlab.com/farkroft/auth-service/external/log"
	"gitlab.com/farkroft/auth-service/external/util"
	"golang.org/x/crypto/bcrypt"
)

// UserRegister usecase for register user
func (u *UseCase) UserRegister(req request.UserRequest) (int, string, error) {
	if req.Password == "" {
		err := fmt.Errorf("password is empty")
		return http.StatusBadRequest, "bad request", err
	}

	if req.Username == "" {
		err := fmt.Errorf("username is empty")
		return http.StatusBadRequest, "bad request", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("hash", err)
		return http.StatusInternalServerError, "hash", err
	}

	req.Password = string(hashedPassword)

	_, err = u.UserRepo.RegisterUser(req)
	if err != nil {
		return http.StatusInternalServerError, "database", err
	}

	return http.StatusCreated, "OK", nil
}

// UserLogin usecase for user login
func (u *UseCase) UserLogin(req request.UserRequest) (int, string, interface{}, error) {
	if req.Password == "" {
		err := fmt.Errorf("password is empty")
		return http.StatusBadRequest, "bad request", nil, err
	}

	if req.Username == "" {
		err := fmt.Errorf("username is empty")
		return http.StatusBadRequest, "bad request", nil, err
	}

	user, err := u.UserRepo.GetUser(req)
	if err != nil {
		if util.IsErrorRecordNotFound(err) {
			return http.StatusNotFound, "user not found", nil, err
		}

		log.Errorf("get user", err)
		return http.StatusInternalServerError, "get user", nil, err
	}

	fmt.Println(u.Cfg.GetInt(constants.EnvJWTPeriod))

	now := util.WIBTimezone(util.Now())
	expiredAt := now.Add(time.Minute * time.Duration(u.Cfg.GetInt(constants.EnvJWTPeriod))).Unix()
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Errorf("hash pass", err)
		return http.StatusBadRequest, "Bad Credentials", nil, err
	}

	userClaims := jwt.MapClaims{
		"UserID":   user.ID,
		"Username": user.Username,
		"StandardClaims": &jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenStr, err := token.SignedString([]byte(u.Cfg.GetString(constants.EnvJWTSecret)))
	if err != nil {
		log.Errorf("token decode", err)
		return http.StatusInternalServerError, "token decode", nil, err
	}

	err = u.Redis.Set(context.TODO(), user.ID.String(), tokenStr, u.Cfg.GetInt(constants.EnvJWTPeriod), "min")
	if err != nil {
		log.Errorf("save token to redis", err)
		return http.StatusInternalServerError, "save token to redis", nil, err
	}

	resp := response.LoginResponse{
		Token: tokenStr,
	}

	return http.StatusOK, "OK", resp, nil
}

// UserAuthVerify to verify token valid or not
func (u *UseCase) UserAuthVerify(str string) (*jwt.Token, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.Cfg.GetString(constants.EnvJWTSecret)), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	exp := extractClaims(claims)
	isExpired := isTokenExpired(exp)
	if isExpired {
		err := fmt.Errorf("token expired")
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	tokenRedis, err := u.Redis.Get(context.TODO(), extractUserID(claims))
	if err != nil {
		return nil, err
	}

	tokenStr, err := token.SignedString([]byte(u.Cfg.GetString(constants.EnvJWTSecret)))
	if err != nil {
		return nil, err
	}

	if tokenRedis != tokenStr {
		return nil, errors.New("token is invalid")
	}

	return token, nil
}

func extractClaims(claims jwt.MapClaims) int64 {
	stdClaims := claims["StandardClaims"].(map[string]interface{})
	exp := stdClaims["exp"].(float64)
	return int64(exp)
}

func isTokenExpired(expiredAt int64) bool {
	nowUnix := util.WIBTimezone(util.Now()).Unix()
	return expiredAt < nowUnix
}

func extractUserID(claims jwt.MapClaims) string {
	userID := claims["UserID"]

	switch v := userID.(type) {
	case string:
		return v
	case nil:
		return ""
	}

	return ""
}
