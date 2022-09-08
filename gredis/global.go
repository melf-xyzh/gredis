/**
 * @Time    :2022/7/10 16:42
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:global.go
 * @Project :go-redis
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package gredis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	RDB *redis.Client // Redis连接池
	ctx = context.Background()
)
