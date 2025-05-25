package utils

import (
	"context"
)

// 定义非导出 key 类型，防止外部冲突
type ctxKey string

// 定义所有 key（注意：是 ctxKey 类型）
const (
	adminIDKey ctxKey = "admin_id"
	userIDKey  ctxKey = "user_id"
	roleKey    ctxKey = "role"
)

// SetAdminID 将 adminID 写入 context
func SetAdminID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, adminIDKey, id)
}

// GetAdminID 从 context 中获取 adminID
func GetAdminID(ctx context.Context) (int, bool) {
	val := ctx.Value(adminIDKey)
	id, ok := val.(int)
	return id, ok
}

// SetUserID 写入用户ID
func SetUserID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

// GetUserID 读取用户ID
func GetUserID(ctx context.Context) (int, bool) {
	val := ctx.Value(userIDKey)
	id, ok := val.(int)
	return id, ok
}
