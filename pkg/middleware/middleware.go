package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/types"
)

var SecretKey string

func init() {
	SecretKey = model.JwtSecretKey()
}

func createToken(user types.User) string {
	var claims = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"userid":   user.Userid,
		"IsAdmin":  user.IsAdmin,
		"expiry":   time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		log.Fatalf("Failed to sign token: %v", err)
	}
	return token
}

func SendCookie(w http.ResponseWriter, user types.User) {
	token := createToken(user)
	cookie := http.Cookie{
		Name:     "JWT",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("JWT")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username := claims["username"].(string)
			userid := claims["userid"].(float64)
			isAdmin := claims["IsAdmin"].(bool)
			expiry := claims["expiry"].(float64)
			if int64(expiry) < time.Now().Unix() {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			user := types.User{
				Username: username,
				Userid:   userid,
				IsAdmin:  isAdmin,
			}
			ctx := context.WithValue(r.Context(), "user", user)
			req := strings.Split(r.URL.String(), "/")[1]
			if !(user.IsAdmin) && req == "admin" {
				http.Redirect(w, r, "/user/home", http.StatusSeeOther)
				return
			} else if user.IsAdmin && req == "user" {
				http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	})
}

func GetUser(r *http.Request) types.User {
	user := r.Context().Value("user").(types.User)
	return user
}

func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("JWT")

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username := claims["username"].(string)
			userid := claims["userid"].(float64)
			isAdmin := claims["IsAdmin"].(bool)
			expiry := claims["expiry"].(float64)
			if int64(expiry) < time.Now().Unix() {
				next.ServeHTTP(w, r)
				return
			}
			user := types.User{
				Username: username,
				Userid:   userid,
				IsAdmin:  isAdmin,
			}
			ctx := context.WithValue(r.Context(), "user", user)

			if user.IsAdmin {
				http.Redirect(w, r.WithContext(ctx), "/admin/home", http.StatusSeeOther)
				return
			} else {
				http.Redirect(w, r.WithContext(ctx), "/user/home", http.StatusSeeOther)
				return
			}
		} else {
			next.ServeHTTP(w, r)
			return
		}
	})
}
