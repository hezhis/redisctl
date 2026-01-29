# RedisCtl

RedisCtl 是一个简单易用的 Redis 操作工具库，提供了类型安全的 Redis 命令封装和便捷的值转换功能。

## 特性

- **类型安全**: 提供强类型的值转换方法，避免运行时错误
- **超时控制**: 所有操作都支持自定义超时时间
- **错误处理**: 统一的错误处理机制，区分业务错误和 Redis Nil 情况
- **便捷转换**: 支持多种数据类型转换（int32, int64, uint32, uint64, bool, string, bytes）

## 安装
```bash
go get github.com/hezhis/redisctl
```
