package app

import (
	"GoProgrammingJourney/blog_service/global"
	"GoProgrammingJourney/blog_service/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	// payload字段
	jwt.StandardClaims
}

// 返回秘钥
// 不能返回string类型, 在GenerateToken()中使用时, 会报 key is of invalid type 错误
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// 用来生成JWT Token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	// 过期时间
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			// 签发者
			Issuer: global.JWTSetting.Issuer,
		},
	}

	// 创建Token实例
	// 1: 加密类型
	// 2: 加密数据
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 加密后的字符串
	token, err := tokenClaims.SignedString(GetJWTSecret())
	global.Logger.Infof("GenerateToken, secret: %v", GetJWTSecret())
	global.Logger.Infof("GenerateToken, token: %v, err: %v", token, err)

	return token, err
}

// 解析, 校验 JWT Token
func ParseToken(token string) (*Claims, error) {
	// 解析鉴权, 内部是解码和校验的过程
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		// 验证基于时间的证明
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
