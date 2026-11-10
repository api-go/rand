# Rand 模块测试报告

## v1.5.4 验证
- 独立模块解析、全量测试与基准测试均通过。

## 概述
本模块提供基于 `math/rand/v2` 的高性能随机数生成，包含以下核心功能：
- **Int/FastInt**: 支持泛型的范围随机整数生成。
- **Float/FastFloat**: 支持泛型的范围随机浮点数生成。
- **Byte/Bytes**: 随机字节及切片生成。
- **Perm/Shuffle**: 集合随机化操作。

## 功能测试
- `TestGenericInt`: 验证整数生成覆盖正数、负数区间及边界逻辑（`max < min`）。
- `TestFastInt`: 验证高并发下的随机整数生成。
- `TestBytes`: 验证边界长度（<=0）处理及生成切片正确性。

## 性能 (Benchmark)
通过 `bench_test.go` 对比了普通生成与 `sync.Pool` 优化后的性能：

| 测试用例 | 耗时 (ns/op) | 优化说明 |
| :--- | :--- | :--- |
| `BenchmarkInt` | 7.816 ns/op | 标准生成 |
| `BenchmarkFastIntParallel` | 2.301 ns/op | 基于 sync.Pool 并发优化 |

性能分析显示，使用 `sync.Pool` 管理 `rand.Rand` 实例在高并发场景下表现显著优于全局锁或频繁创建实例。

## GitHub migration validation (pending release)

Module paths and internal dependency versions are migrated. Development verification uses GOWORK=off and explicit sibling replacements; remove these replacements before release. See migration validation logs for actual results. Existing unrelated working changes are preserved and require separate review.

API inventory verification: `GOWORK=off go doc` succeeded; output retained in the migration validation directory. Source exports were preserved by the path migration.

## GitHub release verification

GOWORK=off go test -v ./... and go test -bench=. ./... passed using published GitHub dependencies, without local replacements. go mod verify passed.
