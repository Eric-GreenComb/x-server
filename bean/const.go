package bean

import (
	"time"
)

const (
	// JWTSigningKey JWTSigningKey
	JWTSigningKey string = "fifu.io"
	// ExpireTime ExpireTime
	ExpireTime time.Duration = time.Minute * 60
	// Realm Realm
	Realm string = "fifu blockchain"
)
