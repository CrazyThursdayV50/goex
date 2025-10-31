package apis

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/account"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/auth"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/exchangeinfo"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/klines"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/order"
	"github.com/CrazyThursdayV50/goex/infra/utils"
	"github.com/CrazyThursdayV50/pkgo/log"
)

var (
	apikey string
	secret string

	logger    log.Logger
	apiClient *API

	symbol string
)

type testLogger struct {
	t *testing.T
}

func (l *testLogger) Debug(args ...any) {
	l.t.Log(args...)
}

func (l *testLogger) Info(args ...any) {
	l.t.Log(args...)
}

func (l *testLogger) Warn(args ...any) {
	l.t.Log(args...)
}

func (l *testLogger) Error(args ...any) {
	l.t.Error(args...)
}

func (l *testLogger) Debugf(fmt string, args ...any) {
	l.t.Logf(fmt, args...)
}

func (l *testLogger) Infof(fmt string, args ...any) {
	l.t.Logf(fmt, args...)
}

func (l *testLogger) Warnf(fmt string, args ...any) {
	l.t.Logf(fmt, args...)
}

func (l *testLogger) Errorf(fmt string, args ...any) {
	l.t.Errorf(fmt, args...)
}

func setupLogger(t *testing.T) {
	logger = &testLogger{t}
}

func setupClient(t *testing.T) {
	symbol = "BTCUSDT"

	var err error
	// 创建客户端，应该自动进行身份验证
	apiClient, err = New(logger, apikey, secret)
	if err != nil {
		t.Fatalf("New apiclient failed: %v", err)
	}

	apiClient.Run(t.Context())

	ping, err := apiClient.Ping(t.Context())
	if err != nil {
		t.Fatalf("Ping failed: %v", err)
	}
	t.Logf("Ping Success: %+v", ping)

	info, err := apiClient.ExchangeInfo(t.Context(), &exchangeinfo.ParamsData{Symbol: symbol})
	if err != nil {
		t.Fatalf("ExchangeInfo failed: %v", err)
	}
	t.Logf("exchange info: %v", info)

	res, err := apiClient.Auth(t.Context(), &auth.ParamsData{RecvWindow: 5000})
	if err != nil {
		t.Fatalf("Auth failed: %v", err)
	}
	t.Logf("Auth Success: %+v", res)

}

func getTestEnv(t *testing.T) {
	apikey = os.Getenv("BN_APIKEY")
	secret = os.Getenv("BN_SECRET")
	if apikey == "" || secret == "" {
		t.Skip("请设置 BINANCE_API_KEY 和 BINANCE_SECRET_KEY 环境变量")
	}
}

func setup(t *testing.T) {
	setupLogger(t)
	getTestEnv(t)
	setupClient(t)
}

func TestGenEd25519Keys(t *testing.T) {
	publickey, privatekey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}

	privateBytes, err := x509.MarshalPKCS8PrivateKey(privatekey)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}

	publicBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}

	privatepem := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privateBytes})
	publicpem := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicBytes})

	// privateStr := base64.StdEncoding.EncodeToString(privateBytes)
	// publicStr := base64.StdEncoding.EncodeToString(publicBytes)

	// t.Logf("private: %s", privateStr)
	// t.Logf("public: %s", publicStr)
	t.Logf("private: %s", privatepem)
	t.Logf("public: %s", publicpem)
}

func TestExchangeInfo(t *testing.T) {
	info, err := apiClient.ExchangeInfo(t.Context(), &exchangeinfo.ParamsData{Symbol: symbol})
	if err != nil {
		t.Errorf("ExchangeInfo failed: %v", err)
		return
	}

	logger.Infof("ExchangeInfo: %+#v", info)
}

func TestAccountStatus(t *testing.T) {
	account, err := apiClient.AccountStatus(t.Context(), &account.ParamsData{OmitZeroBalances: true})
	if err != nil {
		t.Errorf("查询账户信息失败: %v", err)
		return
	}

	logger.Infof("账户信息: %+#v", account)
}

func TestKlines(t *testing.T) {
	klinesRequest := apiClient.Klines().
		Symbol("BTCUSDT").
		Interval(klines.Min1.String()).
		StartTime(1753601220000).
		EndTime(1753602719999).
		Limit(1000)
	klines, err := klinesRequest.Do(t.Context())
	if err != nil {
		t.Errorf("获取Klines失败: %v", err)
		return
	}

	logger.Infof("Klines: %+#v", klines)
}

func TestOrderTest(t *testing.T) {
	// 测试下单
	price := "100000"
	quantity := "0.0001"
	order, err := apiClient.Order().Test(t.Context(), &order.ParamsData{
		Symbol:      "BTCUSDT",
		Side:        order.BUY,
		Type:        order.LIMIT,
		TimeInForce: utils.Ptr(order.GTC),
		Price:       &price,
		Quantity:    &quantity,
	})
	if err != nil {
		t.Errorf("测试下单失败: %v", err)
		return
	}
	logger.Infof("测试下单成功: %+v", order)
}

func TestAll(t *testing.T) {
	setup(t)
	t.Run("ExchangeInfo", TestExchangeInfo)
	t.Run("AccountStatus", TestAccountStatus)
	t.Run("Klines", TestKlines)
	t.Run("PlaceOrderTest", TestOrderTest)
}

// func TestWsAPI(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("跳过需要环境变量和网络连接的测试")
// 	}

// 	ctx := context.TODO()
// 	logger := sugar.New(sugar.DefaultConfig())

// 	// 获取 API Key 和 Secret Key
// 	apiKey, secretKey := getTestEnv(t)

// 	// 创建客户端，使用身份验证
// 	client := New(logger, apiKey, secretKey)
// 	client.Run(ctx)
// 	defer client.Stop()

// 	// 测试下单
// 	// price := "99000"
// 	// quantity := "0.0001"
// 	// newOrder, err := client.Order(ctx, &models.WsOrderParamsData{
// 	// 	Symbol:      "BTCUSDT",
// 	// 	Side:        models.BUY,
// 	// 	Type:        models.LIMIT,
// 	// 	TimeInForce: models.GTC,
// 	// 	Price:       &price,
// 	// 	Quantity:    &quantity,
// 	// })
// 	newOrder, err := client.Order().Place(ctx, &orderModels.ParamsData{
// 		Symbol: "BTCUSDT",
// 		Side:   orderModels.BUY,
// 		Type:   orderModels.MARKET,
// 		// TimeInForce:   variables.Ptr(models.FOK),
// 		QuoteOrderQty: utils.Ptr("10"),
// 	})
// 	if err != nil {
// 		t.Errorf("下单失败: %v", err)
// 		return
// 	}
// 	logger.Infof("下单成功: %+#v", newOrder)
// }

// func TestAuth(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("跳过需要环境变量和网络连接的测试")
// 	}

// 	apiKey, secretKey := getTestEnv(t)

// 	// 创建日志记录器
// 	logger := sugar.New(sugar.DefaultConfig())

// 	// 测试创建客户端
// 	tests := []struct {
// 		name      string
// 		apiKey    string
// 		secretKey string
// 		wantErr   bool
// 	}{
// 		{
// 			name:      "正常创建",
// 			apiKey:    apiKey,
// 			secretKey: secretKey,
// 			wantErr:   false,
// 		},
// 		// {
// 		// 	name:      "正常创建",
// 		// 	apiKey:    "",
// 		// 	secretKey: "",
// 		// 	wantErr:   false,
// 		// },
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 			defer cancel()

// 			api := New(logger, tt.apiKey, tt.secretKey)
// 			if api == nil && !tt.wantErr {
// 				t.Errorf("New() 返回了nil，但期望不返回错误")
// 				return
// 			}
// 			if api != nil && tt.wantErr {
// 				t.Errorf("New() 返回了非nil，但期望返回错误")
// 				return
// 			}

// 			// 如果创建成功，测试基本功能
// 			if api != nil {
// 				api.Run(ctx)

// 				// 测试 Ping
// 				_, err := api.Ping(ctx)
// 				if err != nil {
// 					t.Errorf("Ping() 错误 = %v", err)
// 				}

// 				// 清理资源
// 				api.Stop()
// 			}
// 		})
// 	}
// }
