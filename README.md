# gredis
go-redis的常用封装

### 使用

```bash
go get github.com/melf-xyzh/gredis
```

### 示例

初始化Redis连接池

```go
func InitRedis(ip, port, password string, db int) (rdb *redis.Client, err error)
```

### 公共方法

获取所有Keys

```go
func Keys(pattern string) (keys []string, err error)
```

利用Scan获取Keys（生产环境用于替换Keys）

```go
KeysByScan(cursor uint64, match string, count int64) (keys []string, err error)
```

获取Keys的数量

```go
func Size() (size int64, err error)
```

删除键

```go
func Delete(keys ...string) (err error) 
```

判断键是否存在

```go
func Exists(key string) (exists bool, err error)
```

设置Key的有效期

```go
func Expire(key string, expiration time.Duration) (err error)
```

查看Key剩余的有效期

```go
func TTL(key string) (expiration time.Duration, err error)
```

### String

设置指定 key 的值

```go
func Set(key string, value interface{}, expiration time.Duration) (err error) 
```

批量设置指定 key 的值

```go
func MSet(data map[string]interface{}) (err error) 
```

获取指定 key 的值

```go
func Get(key string) (value string, err error) 
```

批量获取 key 对应的值

```go
func MGet(keys ...string) (values []interface{}, err error)
```

获取指定 key 的值（GetInt64）

```go
func GetInt64(key string) (value int64, err error) 
```

获取指定 key 的值（GetInt）

```go
func GetInt(key string) (value int, err error)
```

获取指定 key 的值（GetUint64）

```go
func GetUint64(key string) (value uint64, err error)
```

获取指定 key 的值（GetBool）

```go
func GetBool(key string) (value bool, err error)
```

获取指定 key 的值（GetFloat32）

```go
func GetFloat32(key string) (value float32, err error) 
```

获取指定 key 的值（GetFloat）

```go
func GetFloat(key string) (value float64, err error)
```

自增1

```go
func Incr(key string) (err error)
```

自增指定步长（整型）

```go
func IncrBy(key string, value int64) (err error)
```

自增指定步长(浮点型)

```go
func IncrByFloat(key string, value float64) (err error)
```

# sorted set（有序集合）

向有序集合添加一个元素

```go
func ZAdd(key string, score float64, member string) (err error)
```

返回集合元素个数

```go
func ZCard(key string) (result int64, err error)
```

删除集合元素

```go
func ZRem(key string, members []string) (err error)
```

返回指的范围的元素

```go
func ZRange(key string, start, end int64) (result []string, err error)
```

计算指定字典区间内成员数量

```go
func ZLexCount(key string, min, max string) (result int64, err error)
```

指定区间分数的成员数

```go
func ZCount(key string, min, max string) (result int64, err error)
```

查询范围内容的设备

```go
func ZRangeByScore(key string, min, max string) (result []string, err error)
```

删除指定分数之间的元素

```go
func ZRemRangeByScore(key string, min, max string)(err error)
```

### 未完待续……



