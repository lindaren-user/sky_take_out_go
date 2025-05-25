package utils

import (
	"net/http"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// JWTAdminMiddleware 校验JWT令牌的中间件
func JWTAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 从请求头中获取令牌
		tokenHeaderKey := viper.GetString("jwt.admin.name")
		tokenString := r.Header.Get(tokenHeaderKey)
		if tokenString == "" {
			http.Error(w, "Authorization header 是必须的", http.StatusUnauthorized)
			return
		}

		// 2. 校验令牌
		claims, err := ParseJWT(tokenString)
		if err != nil {
			Logger.Error("JWT 校验失败", zap.Error(err))
			http.Error(w, "无效 token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// 3. 获取管理员ID并存入context
		adminID := claims.UserID
		ctx := SetAdminID(r.Context(), adminID)

		// 4. 放行
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
