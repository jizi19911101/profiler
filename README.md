# Profiler

[![Go Report Card](https://goreportcard.com/badge/github.com/xyctruth/profiler?x=xyctruth)](https://goreportcard.com/report/github.com/xyctruth/profiler)
[![codecov](https://codecov.io/gh/xyctruth/profiler/branch/master/graph/badge.svg?token=YWNYJK9KQW)](https://codecov.io/gh/xyctruth/profiler)
[![Build status](https://img.shields.io/github/workflow/status/xyctruth/profiler/Server-Build/master)](https://github.com/xyctruth/profiler/actions/workflows/server-build.yml)

> [Demo](https://profiling.jia-huang.com)

## `profiler` 基于 `pprof` 与 `go trace` 持续性能剖析工具

- 根据配置文件收集目标服务的样本, 收集程序会监听配置文件的变化,即时应用变化后的配置文件
- 支持的样本 `trace` `fgprof` `profile` `mutex` `heap` `goroutine` `allocs` `block` `threadcreate`
- 提供 `chart` 观测服务性能指标的趋势，找出性能问题的时间点
- 点击 `chart` 中的气泡跳转到 `pprof` 与 `trace` 的详细页面进行进一步详细的分析

<table>
  <tr>
      <td width="50%" align="center"><b>pprof chart</b></td>
      <td width="50%" align="center"><b>点击气泡跳转pprof详情</b></td>
  </tr>
  <tr>
     <td><img src="https://xtruth.oss-cn-shenzhen.aliyuncs.com/profiler.png"/></td>
     <td><img src="https://xtruth.oss-cn-shenzhen.aliyuncs.com/profiler-pprof.png"/></td>
  </tr>
  <tr>
      <td width="50%" align="center"><b>trace chart</b></td>
      <td width="50%" align="center"><b>点击气泡跳转trace详情</b></td>
  </tr>
  <tr>
     <td><img src="https://xtruth.oss-cn-shenzhen.aliyuncs.com/profiler.png"/></td>
     <td><img src="https://xtruth.oss-cn-shenzhen.aliyuncs.com/profiler-trace.png"/></td>
  </tr>
</table>

## 快速入门

### 本地启动
```bash
# run server :8080
go run server/main.go 

# run ui :80
cd ui && npm install --registry=https://registry.npm.taobao.org && npm run dev --base_api_url=http://localhost:8080 
```

### In Docker

```bash
# 无持久化
docker run -d -p 80:80 --name profiler xyctruth/profiler:latest

# 持久化
mkdir -vp ~/profiler/config/
cp ./collector.yaml ~/profiler/config/
docker run -d -p 80:80 -v ~/profiler/data/:/profiler/data/ -v ~/profiler/config/:/profiler/config/ --name profiler xyctruth/profiler:latest
```

## 抓取配置

需要被收集分析的 `golang` 程序,需要提供 `net/http/pprof` 端点，并配置在 `./collector.yaml` 配置文件中。

配置文件可以在线更新，收集程序会监听配置文件的变化，即时应用变化后的配置文件。

### `collector.yaml`

```yaml
collector:
  targetConfigs:

    profiler-server:        # 目标名称
      interval: 15s         # 抓取间隔
      expiration: 0         # 无过期时间
      host: localhost:9000  # 目标服务host
      profileConfigs:       # 使用默认配置 
        
    server2:
      interval: 10s
      expiration: 168h      # 过期时间7天
      host: localhost:9000
      profileConfigs:       # 覆盖部分默认配置字段
        trace:
          enable: false
        fgprof:
          enable: false
        profile:
          path: /debug/pprof/profile?seconds=10
          enable: false
        heap:
          path: /debug/pprof/heap

```

### `profileConfigs` 默认配置

```yaml
profileConfigs:
  profile:
    Path:   "/debug/pprof/profile?seconds=10"
    Enable: true
  fgprof:
    Path:   "/debug/fgprof?seconds=10"
    Enable: true
  trace:
    Path:   "/debug/pprof/trace?seconds=10"
    Enable: true
  mutex:
    Path:   "/debug/pprof/mutex"
    Enable: true
  heap:
    Path:   "/debug/pprof/heap"
    Enable: true
  goroutine:
    Path:   "/debug/pprof/goroutine"
    Enable: true
  allocs:
    Path:   "/debug/pprof/allocs"
    Enable: true
  block:
    Path:   "/debug/pprof/block"
    Enable: true
  threadcreate:
    Path:   "/debug/pprof/threadcreate"
    Enable: true
```