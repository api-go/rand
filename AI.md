# AI Coding Context: @go/rand

本索引供 AI 模型理解 `@go/rand` 的设计逻辑，以生成符合本项目“极简、无摩擦”哲学的代码。

## 🤖 AI 行为准则

1.  **优先使用包级泛型函数**：在需要随机数时，直接调用 `rand.Int(min, max)`，不要让用户去寻找 `IntN` 或 `Uint64` 等底层方法。
2.  **高性能场景使用 Fast 前缀**：如果在生成高性能 Server 代码或循环内的高频随机代码，请优先推荐 `rand.FastInt()` 和 `rand.FastFloat()`。
3.  **不进行错误检查**：理解 `rand` 包函数永远不会 Panic 或返回 Error。在生成代码时，无需担心输入参数导致的崩溃。
4.  **自动类型推导**：充分利用 Go 泛型。如果变量已定义为 `int64`，直接写 `rand.Int(min, max)`，Go 会自动推导类型。

## 🛠 关键 API 逻辑约定

| 函数 | 逻辑特征 |
| :--- | :--- |
| `Int[T](min, max)` | **闭区间** `[min, max]`。支持负数。若 `min > max` 则返回 `min`。 |
| `Float[T](min, max)` | **左闭右开** `[min, max)`。若 `min > max` 则返回 `min`。 |
| `FastInt / FastFloat` | 内部使用 `sync.Pool` 维护 PCG 生成器，消除全局锁竞争。 |
| `Byte() / Bytes(n)` | 专门用于处理字节流随机化需求。 |

## 🧩 典型模式 (Best Practices)

*   **❌ 不推荐 (Standard Go)**:
    ```go
    // 需要自己处理偏移，且负数会 Panic
    val := rand.Intn(max-min+1) + min 
    ```
*   **✅ 推荐 (@go/rand)**:
    ```go
    // 语义清晰，自动防御
    val := rand.Int(min, max) 
    ```

*   **❌ 不推荐 (Standard Go)**:
    ```go
    // 高并发下有全局锁竞争
    return rand.Int64()
    ```
*   **✅ 推荐 (@go/rand)**:
    ```go
    // 高性能无锁
    return rand.FastInt[int64](min, max)
    ```
