# Lesson01 · LV3 作业索引

## 一、题目要求

课件仓库 `lesson01/README.md` 第 772 行（不在本仓库内）：

> 给你这节课写的函数补一个 `_test.go` 测试文件，至少包含：
>
> - 一个功能测试；
> - 一个表驱动测试（覆盖正常值和边界值，比如除数为 0）。

教材是同一文件 569–684 行「五、go test」。

## 二、答案文件在哪里

**本目录只有这份索引，没有 `.go` 文件。** 测试代码紧贴被测函数，放在各自的同级目录：

| 被测函数 | 实现文件 | 测试文件（相对本目录） |
| --- | --- | --- |
| `Work`（LV1 频次统计） | [`../LV1/main.go`](../LV1/main.go) | [`../LV1/LV1_test.go`](../LV1/LV1_test.go) |
| `CalcFactory`（LV2 闭包工厂） | [`../LV2/main.go`](../LV2/main.go) | [`../LV2/LV2_test.go`](../LV2/LV2_test.go) |

> **为什么测试不放进 LV3**：Go 要求测试文件与被测函数**同包、同目录**。`Work` 和 `CalcFactory` 分别定义在 LV1、LV2，只把 `_test.go` 搬进 LV3 会编译报 `undefined: Work`；连实现一起搬，又会因为两个 `package main` 各带一个 `main()` 而报 `main redeclared`。按 Go 的规则原地存放才是对的，LV3 只做索引。

## 三、各文件包含什么

### [`../LV1/LV1_test.go`](../LV1/LV1_test.go)

- `TestWork`：**表驱动功能测试**，`t.Run` 跑 3 组 —— 常规 `1,1,2,3,3,3` / 空切片（边界）/ 含零和负数（边界），用 `reflect.DeepEqual` 比对整个 `map`，会同时抓出"算错"和"返回空 map"。
- `ExampleWork`：带 `// Output:` 的示例测试（注释写了才会真的比对标准输出）。
- `BenchmarkWork`：用 `b.N` 的基准测试（对应 LVX 选做）。

### [`../LV2/LV2_test.go`](../LV2/LV2_test.go)

- `TestCalcFactory`：**表驱动功能测试**，5 组 —— 加 / 减 / 乘 / 除 / **除数为 0**（边界值，本实现约定返回 `0`，不 panic）；每组先判 `CalcFactory` 是否返回 `nil`（用 `t.Fatalf` 立刻中止），避免 nil 调用导致 panic。
- `TestCalcFactoryUnknownOp`：**独立功能测试**，覆盖非法操作符返回 `nil` 的边界。
- `BenchmarkCalcFactory`：用 `b.N` 的基准测试。

对照题目「一个 `_test.go` 里同时包含功能测试和表驱动测试（含边界值）」：`../LV2/LV2_test.go` 两类齐全且边界用例完整；`../LV1/LV1_test.go` 的 `TestWork` 表驱动每组都在验证功能，并覆盖了空切片、零、负数。

## 四、怎么跑（演示 / 验收请直接复制）

**在仓库根目录执行**（包路径写全，不要写成 `./...`）：

```bash
go test -v ./Homework_lesson01/LV1 ./Homework_lesson01/LV2
```

预期：`TestWork` 3 个子用例、`ExampleWork`、`TestCalcFactory` 5 个子用例、`TestCalcFactoryUnknownOp` 全部 `PASS`，两个包都是 `ok`。

其余按需：

```bash
# 覆盖率（main() 里的 fmt.Scan 交互测不到，30%~45% 属正常）
go test -cover ./Homework_lesson01/LV1 ./Homework_lesson01/LV2

# 只跑某几个用例（-run 后是正则，不是通配符）
go test -v -run 'TestCalcFactory' ./Homework_lesson01/LV2

# 基准测试（普通 go test 不会跑 Benchmark）
go test -bench=. -run '^$' ./Homework_lesson01/LV1 ./Homework_lesson01/LV2

# 改了代码却"结果没变"时禁用测试缓存
go test -count=1 -v ./Homework_lesson01/LV1 ./Homework_lesson01/LV2
```

也可以进到单个目录直接敲 `go test -v`，但上面这种写全包路径的方式最不容易出错。

> ### 为什么不要用 `go test ./...`
>
> `./...` 不是「跑所有测试」，而是**遍历当前 module 下每一个包并逐个编译**；只要有一个包编译不过，整条命令就 FAIL，哪怕 LV1 / LV2 的测试全过。目前会被两处与本作业无关的历史问题拖红：
>
> 1. `Homework_lesson00/Test/` 下 `Test1_0.go` ~ `Test1_7.go` 8 个文件都是 `package main`、各带一个 `func main()` → 同包函数名重复，`main redeclared in this block`；
> 2. `Homework_lesson02/LV0/Source02/main.go` 是空文件 → `expected 'package', found 'EOF'`。
>
> 另外 `Homework_lesson01/LV0_14/` 自带 `go.mod`，属于**嵌套 module**，`./...` 按规则跳过它，这不是错误。所以演示时用上面两条明确包路径的命令即可。
