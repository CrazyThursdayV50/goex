package api

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models/account"
	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models/exchangeinfo"
	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models/klines"
	orderModels "github.com/CrazyThursdayV50/goex/binance/websocket-api/models/order"
	"github.com/CrazyThursdayV50/pkgo/log/sugar"
)

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

func TestAccountStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要环境变量和网络连接的测试")
	}

	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())

	// 获取 API Key 和 Secret Key
	apiKey, secretKey := getTestEnv(t)

	// 创建客户端，应该自动进行身份验证
	client := New(logger, apiKey, secretKey)
	client.Run(ctx)
	defer client.Stop()

	account, err := client.AccountStatus(ctx, &account.ParamsData{OmitZeroBalances: true})

	if err != nil {
		t.Errorf("查询账户信息失败: %v", err)
		return
	}

	logger.Infof("账户信息: %+#v", account)
}

func TestWsAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要环境变量和网络连接的测试")
	}

	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())

	// 获取 API Key 和 Secret Key
	apiKey, secretKey := getTestEnv(t)

	// 创建客户端，使用身份验证
	client := New(logger, apiKey, secretKey)
	client.Run(ctx)
	defer client.Stop()

	// 测试下单
	// price := "99000"
	// quantity := "0.0001"
	// newOrder, err := client.Order(ctx, &models.WsOrderParamsData{
	// 	Symbol:      "BTCUSDT",
	// 	Side:        models.BUY,
	// 	Type:        models.LIMIT,
	// 	TimeInForce: models.GTC,
	// 	Price:       &price,
	// 	Quantity:    &quantity,
	// })
	newOrder, err := client.Order().Place(ctx, &orderModels.ParamsData{
		Symbol: "BTCUSDT",
		Side:   orderModels.BUY,
		Type:   orderModels.MARKET,
		// TimeInForce:   variables.Ptr(models.FOK),
		QuoteOrderQty: variables.Ptr("10"),
	})
	if err != nil {
		t.Errorf("下单失败: %v", err)
		return
	}
	logger.Infof("下单成功: %+#v", newOrder)
}

func getTestEnv(t *testing.T) (string, string) {
	apiKey := os.Getenv("BN_APIKEY")
	secretKey := os.Getenv("BN_SECRET")
	if apiKey == "" || secretKey == "" {
		t.Skip("请设置 BINANCE_API_KEY 和 BINANCE_SECRET_KEY 环境变量")
	}

	return apiKey, secretKey
}

func TestAuth(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要环境变量和网络连接的测试")
	}

	apiKey, secretKey := getTestEnv(t)

	// 创建日志记录器
	logger := sugar.New(sugar.DefaultConfig())

	// 测试创建客户端
	tests := []struct {
		name      string
		apiKey    string
		secretKey string
		wantErr   bool
	}{
		{
			name:      "正常创建",
			apiKey:    apiKey,
			secretKey: secretKey,
			wantErr:   false,
		},
		// {
		// 	name:      "正常创建",
		// 	apiKey:    "",
		// 	secretKey: "",
		// 	wantErr:   false,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			api := New(logger, tt.apiKey, tt.secretKey)
			if api == nil && !tt.wantErr {
				t.Errorf("New() 返回了nil，但期望不返回错误")
				return
			}
			if api != nil && tt.wantErr {
				t.Errorf("New() 返回了非nil，但期望返回错误")
				return
			}

			// 如果创建成功，测试基本功能
			if api != nil {
				api.Run(ctx)

				// 测试 Ping
				_, err := api.Ping(ctx)
				if err != nil {
					t.Errorf("Ping() 错误 = %v", err)
				}

				// 清理资源
				api.Stop()
			}
		})
	}
}

func TestExchangeInfo(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要环境变量和网络连接的测试")
	}

	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())

	variables.SetIsTest()
	client := New(logger, "", "")
	client.Run(ctx)
	defer client.Stop()

	info, err := client.ExchangeInfo(ctx, &exchangeinfo.ParamsData{SymbolStatus: "TRADING", Permissions: []string{"SPOT"}})
	if err != nil {
		t.Errorf("获取交易所信息失败: %v", err)
		return
	}

	logger.Infof("交易所信息: %+#v", info)
}

func TestKlines(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要环境变量和网络连接的测试")
	}

	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())

	variables.SetIsTest()
	client := New(logger, "", "")
	client.Run(ctx)
	defer client.Stop()

	klinesRequest := client.Klines().
		Symbol("BTCUSDT").
		Interval(klines.Min1.String()).
		StartTime(1753601220000).
		EndTime(1753602719999).
		Limit(1000)
	klines, err := klinesRequest.Do(ctx)
	if err != nil {
		t.Errorf("获取Klines失败: %v", err)
		return
	}

	logger.Infof("Klines: %+#v", klines)
}

func TestOrderTest(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要环境变量和网络连接的测试")
	}

	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())

	// 获取 API Key 和 Secret Key
	apiKey, secretKey := getTestEnv(t)

	// 创建客户端，应该自动进行身份验证
	client := New(logger, apiKey, secretKey)
	client.Run(ctx)
	defer client.Stop()

	// 测试下单
	price := "100000"
	quantity := "0.0001"
	order, err := client.Order().Test(ctx, &orderModels.ParamsData{
		Symbol:      "BTCUSDT",
		Side:        orderModels.BUY,
		Type:        orderModels.LIMIT,
		TimeInForce: variables.Ptr(orderModels.GTC),
		Price:       &price,
		Quantity:    &quantity,
	})
	if err != nil {
		t.Errorf("测试下单失败: %v", err)
		return
	}
	logger.Infof("测试下单成功: %+v", order)
}
