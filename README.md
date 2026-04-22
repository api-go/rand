# 关于本项目

本项目完全由 AI 维护。代码源自 github.com/ssgo/u 的重构。

# @go/rand

`@go/rand` 是一个专注于极简调用与极致性能的 Go 随机数库。它通过泛型消除了传统库中繁琐的类型分支，并针对高并发场景进行了底层优化。

## 🎯 设计哲学

*   **极简 API**：不再需要区分类型名。一个 `Int(min, max)` 搞定所有。
*   **需求导向**：所有函数默认支持范围随机，无需手动计算偏移。
*   **零摩擦防御**：遵循“No-Panic”准则，自动处理反向区间等异常输入。
*   **极致性能**：通过 `Fast` 系列函数彻底解决多核环境下的全局锁竞争瓶颈。

## 🛠 API Reference

### 基础随机函数 (基于标准库全局生成器)
- `func Int[T IntegerType](min, max T) T`：获取 `[min, max]` 闭区间随机整数。
- `func Float[T FloatType](min, max T) T`：获取 `[min, max)` 左闭右开区间随机浮点数。

### 高性能随机函数 (基于 PCG 无锁池)
在高并发 Server 场景下，推荐使用以下函数以获得极致吞吐量。
- `func FastInt[T IntegerType](min, max T) T`
- `func FastFloat[T FloatType](min, max T) T`

### 字节与切片
- `func Byte() byte`
- `func Bytes(n int) []byte`
- `func Perm(n int) []int`
- `func Shuffle(n int, swap func(i, j int))`

## 📦 安装

```bash
go get apigo.cc/go/rand
```

## 💡 快速开始

```go
import "apigo.cc/go/rand"

// 基础随机
num := rand.Int(1, 100)        // 返回 [1, 100] 的 int

// 高并发场景 (推荐)
userID := rand.FastInt(1000, 9999)

// 字节处理
key := rand.Bytes(16)
```
