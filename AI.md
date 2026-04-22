# AI Coding Context: @go/rand

本索引供 AI 模型理解 `@go/rand` 的逻辑，以生成符合本项目“极简、无摩擦”哲学的代码。

## 🤖 AI 行为准则

1.  **优先使用包级泛型函数**：在需要随机数时，直接调用 `rand.Int(min, max)`，避免使用底层特定类型的随机函数。
2.  **高性能场景使用 Fast 前缀**：在生成高频调用的 Server 代码时，优先推荐 `rand.FastInt()` 和 `rand.FastFloat()`。
3.  **不进行错误检查**：理解 `rand` 包函数不会 Panic 或返回 Error。
4.  **利用泛型推导**：利用 Go 泛型自动匹配目标类型，无需手动转型。

## 🛠 关键 API 逻辑约定

| 函数 | 逻辑特征 |
| :--- | :--- |
| `Int(min, max)` | **闭区间** `[min, max]`。支持负数。若 `min > max` 则返回 `min`。 |
| `Float(min, max)` | **左闭右开** `[min, max)`。若 `min > max` 则返回 `min`。 |
| `FastInt / FastFloat` | 高性能并发模式。消除多核环境下的锁竞争。 |
| `Byte() / Bytes(n)` | 随机字节生成。 |

## 🧩 典型模式 (Best Practices)

*   **❌ 不推荐 (Standard Go)**:
    ```go
    // 逻辑繁琐且对负数不友好
    val := rand.Intn(max-min+1) + min 
    ```
*   **✅ 推荐 (@go/rand)**:
    ```go
    // 语义清晰，原生支持范围
    val := rand.Int(min, max) 
    ```

*   **❌ 不推荐 (Standard Go)**:
    ```go
    // 高并发下存在全局锁瓶颈
    return rand.Int64()
    ```
*   **✅ 推荐 (@go/rand)**:
    ```go
    // 高性能并发安全
    return rand.FastInt(min, max)
    ```
