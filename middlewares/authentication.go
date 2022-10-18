package middlewares

import (
	"esaku-project/configs"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/securecookie"
	"os"
)

type AuthImpl struct {
	Config configs.Config
}

type Claims struct {
	KodeLokasi string `json:"kode_lokasi"`
	Nik        string `json:"nik"`
	jwt.RegisteredClaims
}

var (
	//accessCookieToken  = os.Getenv("COOKIE_ACCESS_TOKEN")
	//refreshCookieToken = os.Getenv("COOKIE_ACCESS_REFRESH_TOKEN")
	//jwtSecretKey       = os.Getenv("JWT_KEY_TOKEN")
	//jwtRefreshKey      = os.Getenv("JWT_REFRESH_KEY_TOKEN")
	hashCookie  = os.Getenv("COOKIE_HASH_KEY")
	blockCookie = os.Getenv("COOKIE_BLOCK_KEY")
)

func (auth *AuthImpl) GetHashCookie() string {
	return auth.Config.Get("COOKIE_HASH_KEY")
}

func (auth *AuthImpl) setSecureCookie() *securecookie.SecureCookie {
	fmt.Println(hashCookie)
	fmt.Println(auth.GetHashCookie())
	fmt.Println(blockCookie)
	secure := securecookie.New([]byte(hashCookie), []byte(blockCookie))

	return secure
}

//func getSecureCookie(name string, request *http.Request) (types.M, error) {
//	secure := setSecureCookie()
//
//	cookie, err := request.Cookie(name)
//
//	helpers.PanicIfError(err)
//
//	data := types.M{}
//
//	if err = secure.Decode(name, cookie.Value, &data); err == nil {
//		return data, nil
//	}
//
//	return nil, err
//}
//
//func setTokenCookie(name string, data types.M, expiration time.Time, writer http.ResponseWriter) {
//	secure := setSecureCookie()
//
//	encoded, err := secure.Encode(name, data)
//	helpers.PanicIfError(err)
//
//	cookie := new(http.Cookie)
//	cookie.Name = name
//	cookie.Value = encoded
//	cookie.Expires = expiration
//	cookie.Path = "/"
//	cookie.HttpOnly = true
//
//	http.SetCookie(writer, cookie)
//}
//
//func generateToken(login web.LoginResponse, expiration time.Time, secret []byte) (string, time.Time, error) {
//	claims := &Claims{
//		KodeLokasi: login.KodeLokasi,
//		Nik:        login.Nik,
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: &jwt.NumericDate{Time: expiration},
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	tokenStr, err := token.SignedString(secret)
//
//	if err != nil {
//		return "", time.Now(), err
//	}
//
//	return tokenStr, expiration, nil
//}
//
//func generateAccessToken(login web.LoginResponse) (string, time.Time, error) {
//	expiration := time.Now().Add(1 * time.Hour)
//
//	return generateToken(login, expiration, []byte(jwtSecretKey))
//}
//
//func generateRefreshToken(login web.LoginResponse) (string, time.Time, error) {
//	expiration := time.Now().Add(24 * time.Hour)
//
//	return generateToken(login, expiration, []byte(jwtRefreshKey))
//}
//
//func GenerateTokenAndCookie(login web.LoginResponse, writer http.ResponseWriter) error {
//	token, expiredAccess, err := generateAccessToken(login)
//	helpers.PanicIfError(err)
//
//	dataAccess := types.M{"value": token}
//	setTokenCookie(accessCookieToken, dataAccess, expiredAccess, writer)
//
//	refreshToken, expiredRefresh, err := generateRefreshToken(login)
//	helpers.PanicIfError(err)
//
//	dataRefresh := types.M{"value": refreshToken}
//	setTokenCookie(refreshCookieToken, dataRefresh, expiredRefresh, writer)
//
//	return nil
//}
//
//func DeleteCookie(writer http.ResponseWriter) error {
//	dataAccess := types.M{}
//	setTokenCookie(accessCookieToken, dataAccess, time.Unix(0, 0), writer)
//
//	dataRefresh := types.M{}
//	setTokenCookie(refreshCookieToken, dataRefresh, time.Unix(0, 0), writer)
//
//	return nil
//}
//
//func MiddlewareRefreshToken(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		tokenCookie, _ := request.Cookie(accessCookieToken)
//
//		if tokenCookie == nil {
//			next.ServeHTTP(writer, request)
//		}
//
//		data, err := getSecureCookie(accessCookieToken, request)
//
//		if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
//			panic(errors.New(err.Error()))
//		}
//
//		tokenAuth := data["value"].(string)
//		verifyKey := []byte(os.Getenv("JWT_KEY_TOKEN"))
//
//		claims := Claims{}
//		token, err := jwt.ParseWithClaims(tokenAuth, claims, func(token *jwt.Token) (interface{}, error) {
//			return verifyKey, nil
//		})
//
//		if err != nil {
//			if err == jwt.ErrSignatureInvalid {
//				panic(exceptions.NewErrorUnauthorized("signature token is invalid"))
//				return
//			}
//
//			panic(exceptions.NewErrorUnauthorized(err.Error()))
//			return
//		}
//
//		if !token.Valid {
//			panic(exceptions.NewErrorUnauthorized("token is invalid"))
//		}
//
//		if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) < 15*time.Minute {
//			refreshCookie, err := request.Cookie(refreshCookieToken)
//
//			if err == nil && refreshCookie != nil {
//				verifyKey := []byte(os.Getenv("JWT_REFRESH_KEY_TOKEN"))
//
//				data, err = getSecureCookie(accessCookieToken, request)
//				tokenRefresh := data["value"].(string)
//
//				token, err = jwt.ParseWithClaims(tokenRefresh, claims, func(token *jwt.Token) (interface{}, error) {
//					return verifyKey, nil
//				})
//
//				login := web.LoginResponse{
//					Nik:        claims.Nik,
//					KodeLokasi: claims.KodeLokasi,
//				}
//
//				if token != nil && token.Valid {
//					_ = GenerateTokenAndCookie(login, writer)
//				}
//
//			}
//		}
//
//		next.ServeHTTP(writer, request)
//	})
//}
//
//func MiddlewareCookie(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		tokenCookie, _ := request.Cookie(accessCookieToken)
//
//		if tokenCookie == nil {
//			next.ServeHTTP(writer, request)
//		}
//
//		data, err := getSecureCookie(accessCookieToken, request)
//
//		if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
//			panic(errors.New(err.Error()))
//		}
//
//		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", data["value"]))
//
//		next.ServeHTTP(writer, request)
//	})
//}
//
//func MiddlewareAuthorization(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		verifyKey := []byte(os.Getenv("JWT_KEY_TOKEN"))
//
//		authHeader := strings.Split(request.Header.Get("Authorization"), " ")
//		if len(authHeader) != 2 {
//			panic(exceptions.NewErrorUnauthorized("malformed jwt token"))
//		}
//		tokenBearer := authHeader[1]
//
//		claims := Claims{}
//		token, err := jwt.ParseWithClaims(tokenBearer, claims, func(token *jwt.Token) (interface{}, error) {
//			return verifyKey, nil
//		})
//
//		if err != nil {
//			if err == jwt.ErrSignatureInvalid {
//				panic(exceptions.NewErrorUnauthorized("signature token is invalid"))
//				return
//			}
//
//			panic(exceptions.NewErrorUnauthorized(err.Error()))
//			return
//		}
//
//		if !token.Valid {
//			panic(exceptions.NewErrorUnauthorized("token is invalid"))
//		}
//
//		data := types.M{"kode_lokasi": claims.KodeLokasi, "nik_input": claims.Nik}
//
//		ctx := context.WithValue(request.Context(), "pic", data)
//
//		next.ServeHTTP(writer, request.WithContext(ctx))
//	})
//}
