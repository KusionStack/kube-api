# kube-api 仓库改动计划 - Rollout 支持集群维度精细化分批

## 概述

本计划目标是 RolloutSpec 支持内联策略配置。

## 改动范围

根据方案文档，kube-api 仓库仅需修改 Rollout 类型定义，新增内联策略字段。

---

## 步骤 1: 修改 RolloutSpec 类型定义

**目标**: 在 RolloutSpec 中添加 CanaryStrategy 和 BatchStrategy 内联字段

**文件**: `apis/rollout/v1alpha1/types.go`

**改动内容**:

在 RolloutSpec 结构体中新增两个字段：

```go
// RolloutSpec defines the desired state of Rollout
type RolloutSpec struct {
    // Disabled means that rollout will not respond for new events.
    // Default value is false.
    Disabled bool `json:"disabled,omitempty"`

    // HistoryLimit defines the maximum number of completed rolloutRun
    // history records to keep.
    // The HistoryLimit can start from 0 (no retained RolloutRun history).
    // When not set or set to math.MaxInt32, Rollout will keep all RolloutRun history records.
    //
    // +kubebuilder:default=10
    HistoryLimit *int32 `json:"historyLimit,omitempty"`

    // TriggerPolicy defines when rollout will be triggered
    //
    // +kubebuilder:default=Auto
    TriggerPolicy RolloutTriggerPolicy `json:"triggerPolicy,omitempty"`

    // StrategyRef is a reference to a rollout strategy.
    // Mutually exclusive with CanaryStrategy and BatchStrategy.
    // If specified, CanaryStrategy and BatchStrategy must be empty.
    //
    // +kubebuilder:validation:Optional
    StrategyRef string `json:"strategyRef,omitempty"`

    // CanaryStrategy defines the inline canary strategy.
    // This allows specifying canary deployment details directly in Rollout
    // without requiring a separate RolloutStrategy resource.
    // Mutually exclusive with StrategyRef and BatchStrategy.
    //
    // +kubebuilder:validation:Optional
    CanaryStrategy *RolloutRunCanaryStrategy `json:"canaryStrategy,omitempty"`

    // BatchStrategy defines the inline batch strategy.
    // This allows specifying batch deployment details directly in Rollout
    // without requiring a separate RolloutStrategy resource.
    // Mutually exclusive with StrategyRef and CanaryStrategy.
    //
    // +kubebuilder:validation:Optional
    BatchStrategy *RolloutRunBatchStrategy `json:"batchStrategy,omitempty"`

    // WorkloadRef is a reference to a kind of workloads
    WorkloadRef WorkloadRef `json:"workloadRef,omitempty"`

    // TrafficTopologyRefs defines the networking traffic relationships between
    // workloads, backend services, and routes.
    TrafficTopologyRefs []string `json:"trafficTopologyRefs,omitempty"`
}
```

**依赖**: 无

**验收标准**:
- RolloutSpec 包含 `CanaryStrategy` 字段，类型为 `*RolloutRunCanaryStrategy`
- RolloutSpec 包含 `BatchStrategy` 字段，类型为 `*RolloutRunBatchStrategy`
- 字段使用正确的 json tag 和 kubebuilder validation
- 与 StrategyRef 字段互斥（validation 逻辑由 rollout 项目实现）

---

## 步骤 2: 重新生成 CRD 文件

**目标**: 生成包含新字段的 CRD

**命令**:
```bash
make generate
make manifests
```

**依赖**: 步骤 1

**验收标准**:
- Rollout CRD spec 中包含 `canaryStrategy` 字段定义
- Rollout CRD spec 中包含 `batchStrategy` 字段定义

---

## 其他仓库改动说明

以下改动在 **rollout 项目**中实现，不在 kube-api 仓库：

| 步骤 | 工作内容 | 项目 |
|------|----------|------|
| 1 | 在 RolloutSpec 中添加 CanaryStrategy 和 BatchStrategy 字段 | kube-api (本仓库) |
| 2 | 在 OneTimeStrategy 中新增 InlineBatch 字段 | rollout |
| 3 | 修改 validation 逻辑添加互斥校验 | rollout |
| 4 | 新增 inline_strategy.go 文件 | rollout |
| 5 | 修改 constructRolloutRun 支持内联策略 | rollout |
| 6 | 修改 rollout_controller 处理内联策略 | rollout |
| 7 | 修改 applyOneTimeStrategy 适配内联策略 | rollout |
| 8 | 重新生成 CRD 文件 | 多项目 |
| 9 | 编写测试 | rollout |

---

## 关键类型说明

| 类型 | 来源 | 用途 |
|------|------|------|
| RolloutRunCanaryStrategy | 现有，kube-api | 内联 canary 策略 |
| RolloutRunBatchStrategy | 现有，kube-api | 内联 batch 策略 |
| RolloutRunStepTarget | 现有，kube-api | 批量中的目标配置（含 cluster、name、replicas） |

这些类型已存在于 kube-api 仓库中，直接复用即可。

---

*计划创建时间: 2026-03-06*
*参考文档: Rollout支持集群维度精细化分批 - https://yuque.antfin.com/antcloud-paas/dp4wap/hi78idzzn4g8p1qb*