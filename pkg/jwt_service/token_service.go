package jwt_service

import (
	"dnd-api/db/models"
	"dnd-api/pkg/file_service"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const AccessTokenDuration = time.Minute * 10

type TokenType string

var TokenTypeUser TokenType = "user"
var TokenTypeFiles TokenType = "files"

type JwtCustomClaims struct {
	ID         string    `json:"id"`
	Type       TokenType `json:"type"`
	Authorised bool      `json:"authorised"`
	OwnerUid   string    `json:"owner_uid"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID   string    `json:"id"`
	Type TokenType `json:"type"`
	jwt.StandardClaims
}

type JwtFilesCustomClaims struct {
	ID    string                              `json:"id"`
	Type  TokenType                           `json:"type"`
	Paths map[string]file_service.FilesAccess `json:"paths"`
	jwt.StandardClaims
}

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) CreateUserAccessToken(user *models.User, authorised bool) (accessToken string, exp int64, err error) {
	exp = time.Now().Add(AccessTokenDuration).Unix()
	claims := &JwtCustomClaims{
		user.Uid,
		TokenTypeUser,
		authorised,
		"",
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", 0, err
	}

	return t, exp, err
}

func (s *TokenService) CreateUserRefreshToken(user *models.User, durationMinutes uint) (string, *time.Time, error) {
	refreshExpires := time.Now().Add(time.Minute * time.Duration(durationMinutes))
	claimsRefresh := &JwtCustomRefreshClaims{
		ID:   user.Uid,
		Type: TokenTypeUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpires.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return "", nil, err
	}
	return rt, &refreshExpires, err
}

func (s *TokenService) CreateUserFilesToken(userUid string, fileAccess map[string]file_service.FilesAccess, durationMinutes uint) (string, *time.Time, error) {
	expires := time.Now().Add(time.Minute * time.Duration(durationMinutes))
	claims := &JwtFilesCustomClaims{
		ID:    userUid,
		Type:  TokenTypeFiles,
		Paths: fileAccess,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
		},
	}
	filesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := filesToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return "", nil, err
	}

	return rt, &expires, err
}

func (s *TokenService) GetFilesCookie(filesToken string, filesExpires time.Time) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "files"
	cookie.Value = filesToken
	cookie.Expires = filesExpires
	cookie.HttpOnly = true
	cookie.Domain = "." + os.Getenv("PORTAL_HOST")
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Path = "/files"
	if os.Getenv("ENVIRONMENT") == "production" {
		cookie.Secure = true
	} else {
		cookie.Secure = false
	}
	return cookie
}

func (s *TokenService) AttachFilesCookie(c echo.Context, filesToken string, filesExpires time.Time) {
	cookie := s.GetFilesCookie(filesToken, filesExpires)
	c.SetCookie(cookie)
}

func (s *TokenService) AttachRefreshCookie(c echo.Context, refreshToken string, refreshExpires time.Time) {
	cookie := s.GetRefreshCookie(refreshToken, refreshExpires)
	c.SetCookie(cookie)
}

func (s *TokenService) GetRefreshCookie(refreshToken string, refreshExpires time.Time) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "refresh"
	cookie.Value = refreshToken
	cookie.Expires = refreshExpires
	cookie.HttpOnly = true
	cookie.Path = "/auth/refresh"
	cookie.SameSite = http.SameSiteStrictMode
	if os.Getenv("ENVIRONMENT") == "production" {
		cookie.Secure = true
	} else {
		cookie.Secure = false
	}
	return cookie
}
