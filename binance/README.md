# Binance Go SDK

重构后的币安 Go SDK，采用模块化设计，业务逻辑更加内聚。

## 项目结构

```
binance/
├── binance.go                    # 主入口文件，暴露各SDK的New方法
├── variables/                    # 全局配置和常量
│   ├── const.go
│   └── utils.go
├── websocket-api/               # WebSocket API SDK
│   ├── api.go                   # API核心实现
│   ├── client.go                # WebSocket客户端
│   ├── handlers.go              # 消息处理器
│   ├── simpleapi.go             # 简化API
│   ├── models/                  # API相关数据模型 (按功能分组)
│   │   ├── common.go            # 通用基础结构
│   │   ├── auth.go              # 认证API模型
│   │   ├── order.go             # 订单API模型
│   │   ├── account.go           # 账户API模型
│   │   ├── exchange_info.go     # 交易所信息API模型 + 过滤器
│   │   └── ping.go              # Ping API模型
│   └── signer/                  # 签名实现
│       ├── signer.go            # 签名接口
│       └── sign.go              # Ed25519签名实现
└── websocket-streams/           # WebSocket 流 SDK
    ├── stream.go                # 流管理器
    ├── client.go                # WebSocket客户端
    ├── handler.go               # 处理器接口
    ├── bookTicker.go            # 最佳买卖价流
    └── models/                  # 流相关数据模型 (按流类型分组)
        ├── depth.go             # 深度数据流模型
        └── book_ticker.go       # 最佳买卖价流模型
```

## 重构后的Models组织原则

### WebSocket API Models (按API功能分组)

- **common.go**: 所有API共用的基础结构
  - `WsAPIParams[T]` - 通用请求参数结构
  - `WsAPIResult` - 通用响应结果结构  
  - `WsAPIRateLimit` - 速率限制信息
  - `Sign` - 签名基础结构

- **auth.go**: 认证API相关模型
  - `WsAPIAuthParamsData` - 认证请求参数
  - `WsAPIAuthResultData` - 认证响应数据

- **order.go**: 订单API相关模型
  - `WsOrderParamsData` - 下单请求参数
  - `WsOrderResultData` - 下单响应数据
  - `OrderSide`, `OrderType`, `TimeInForce` - 订单相关枚举

- **account.go**: 账户API相关模型
  - `WsAccountStatusParamsData` - 账户状态请求参数
  - `WsAccountStatusResultData` - 账户状态响应数据
  - `Balance`, `CommissionRates` - 账户相关结构

- **exchange_info.go**: 交易所信息API相关模型
  - `WsExchangeInfoParamsData` - 交易所信息请求参数
  - `WsExchangeInfoResultData` - 交易所信息响应数据
  - 各种过滤器类型 (`PriceFilter`, `LotSizeFilter`等)

- **ping.go**: Ping API相关模型
  - `WsPingParams` - Ping请求参数

### WebSocket Streams Models (按流类型分组)

- **depth.go**: 深度数据流模型
  - `PartialDepthEvent` - 深度数据事件
  - `PartialDepthData` - 深度数据结构
  - `PartialDepthCombinedEvent` - 组合深度数据事件
  - `orderbook` - 订单簿条目

- **book_ticker.go**: 最佳买卖价流模型
  - `IndividualSymbolBookTicker` - 个股最佳买卖价数据
  - `IndividualSymbolBookTickerEvent` - 个股最佳买卖价事件

## 使用方法

### WebSocket API

```go
import (
    "context"
    "github.com/CrazyThursdayV50/goex/binance"
    "github.com/CrazyThursdayV50/pkgo/log"
)

// 创建 WebSocket API 客户端
ctx := context.Background()
logger := log.New()
apiKey := "your_api_key"
secretKey := "your_secret_key"

wsAPI := binance.NewWebSocketAPI(ctx, logger, apiKey, secretKey)
defer wsAPI.Stop()

// 使用API进行交易
// ...
```

### WebSocket Streams

```go
import (
    "context"
    "github.com/CrazyThursdayV50/goex/binance"
    "github.com/CrazyThursdayV50/goex/binance/websocket-streams/models"
    "github.com/CrazyThursdayV50/pkgo/log"
)

// 创建 WebSocket 流客户端
ctx := context.Background()
logger := log.New()
wsStreams := binance.NewWebSocketStreams()

// 创建深度数据流
depthClient := wsStreams.PartialBookDepth5Stream(ctx, logger, "BTCUSDT", 
    func(data *models.PartialDepthData) {
        // 处理深度数据
    })
defer depthClient.Stop()

// 创建最佳买卖价流
tickerClient := wsStreams.IndividualSymbolBookTickerStream(ctx, logger, 
    []string{"BTCUSDT", "ETHUSDT"}, 
    func(data *models.IndividualSymbolBookTicker) {
        // 处理最佳买卖价数据
    })
defer tickerClient.Stop()
```

## 重构优势

1. **业务内聚**: 每个SDK包含自己的models、client和相关逻辑
2. **模块化**: websocket-api 和 websocket-streams 完全独立
3. **统一入口**: 通过 binance 包暴露各SDK的New方法
4. **类型安全**: 每个SDK使用自己的数据类型，避免类型混淆
5. **易于维护**: 各模块职责清晰，便于独立开发和测试
6. **功能分组**: Models按API功能分组，代码组织更清晰

## 主要功能

### WebSocket API
- 身份认证 (Ed25519签名)
- 订单管理 (下单、撤单、查询)
- 账户信息查询
- 交易所信息查询

### WebSocket Streams  
- 实时市场数据流
- 深度数据流 (5档、10档、20档)
- 最佳买卖价流
- 24小时价格统计流 