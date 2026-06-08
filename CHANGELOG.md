# CHANGELOG

## v1.5.1 (2026-06-08)
- **JS 对齐**: 将所有注册到 `jsmod` 的导出方法名统一为 PascalCase（如 `Int`, `FastInt`, `Byte`），以消除 JS 与 Go 调用体感上的摩擦。

## v1.5.0 (2026-05-10)
- 优化 `Bytes` 函数，提升随机生成健壮性。
- 完善 `TEST.md` 性能基准报告与功能测试覆盖。
