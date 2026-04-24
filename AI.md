# AI Coding Context: @go/rand

本索引供 AI 模型理解 `@go/rand` 的逻辑，以生成符合本项目“极简、无摩擦”哲学的代码。

## 🤖 AI 行为准则

1.  **优先使用包级泛型函数**：在需要随机数时，直接调用 `rand.Int(min, max)`，避免使用底层特定类型的随机函数。
2.  **强制性能分级**：
    - 低频任务：使用 `Int`, `Float`。
    - **高并发/高频任务**（QPS > 1000/s）：**强制**使用 `FastInt`, `FastFloat` 系列，消除全局锁瓶颈。
3.  **禁止外部依赖**：严禁引入标准库 `math/rand` 进行全局随机管理。所有随机逻辑必须经由 `@go/rand` 闭环，确保内存与并发安全。
4.  **无错误处理**：理解 `rand` 包函数不会 Panic 或返回 Error，无需进行任何错误检查。
5.  **利用泛型推导**：利用 Go 泛型自动匹配目标类型，无需手动转型。

## 🛠 API Reference

### 随机生成
- `func Int[T IntegerType](min, max T) T`：获取 `[min, max]` 闭区间随机整数。
- `func FastInt[T IntegerType](min, max T) T`：高性能并发模式获取 `[min, max]` 闭区间随机整数。
- `func Float[T FloatType](min, max T) T`：获取 `[min, max)` 左闭右开区间随机浮点数。
- `func FastFloat[T FloatType](min, max T) T`：高性能并发模式获取 `[min, max)` 左闭右开区间随机浮点数。

### 字节与切片
- `func Byte() byte`：获取一个随机字节。
- `func Bytes(n int) []byte`：获取长度为 n 的随机字节切片。
- `func Perm(n int) []int`：获取 `[0, n)` 的随机排列。
- `func Shuffle(n int, swap func(i, j int))`：随机洗牌。

## 🧩 典型模式 (Best Practices)

*   **✅ 推荐：高性能并发生成**:
    ```go
    // 高频生成器 (如 ID 填充、加盐)
    userID := rand.FastInt(10000, 99999)
    ```

*   **✅ 推荐：区间闭合策略**:
    ```go
    // 原生闭区间 [1, 100]，无需手动 +1
    num := rand.Int(1, 100)
    ```
