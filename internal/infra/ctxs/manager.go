package ctxs

import "context"

const (
	cidKey = "key-cid"
)

func WithCid(ctx context.Context, cid string) context.Context {
	return context.WithValue(ctx, cidKey, cid)
}

func GetCid(ctx context.Context) string {
	key := ctx.Value(cidKey)
	if key == nil {
		return ""
	}

	return key.(string)
}
