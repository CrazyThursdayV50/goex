package sign

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/CrazyThursdayV50/goex/infra/iface"
	"github.com/CrazyThursdayV50/goex/infra/utils"
)

// Sign 通用签名函数，使用 Ed25519 私钥对参数进行签名
func Ed25519(params map[string]any, secretkey ed25519.PrivateKey) {
	// 1. 按照键名对参数排序
	var keys []string
	for k := range params {
		if k != "signature" { // 排除 signature 字段
			keys = append(keys, k)
		}
	}

	slices.Sort(keys)

	// 2. 构建签名字符串
	var parts []string
	for _, k := range keys {
		v := params[k]
		// if v != "" { // 只包含非空值
		parts = append(parts, fmt.Sprintf("%s=%v", k, v))
		// }
	}

	payload := strings.Join(parts, "&")

	signature := ed25519.Sign(secretkey, []byte(payload))
	// 5. 将签名转换为 base64 字符串
	sign64 := base64.StdEncoding.EncodeToString(signature)
	params["signature"] = sign64
}

func ParseSecretEd25519(apikey, secret string) (ed25519.PrivateKey, error) {
	prv, err := base64.RawStdEncoding.DecodeString(secret)
	if err != nil {
		return nil, err
	}

	privatekey, err := x509.ParsePKCS8PrivateKey(prv)
	if err != nil {
		return nil, err
	}

	return privatekey.(ed25519.PrivateKey), nil
}

func NewSignerFuncEd25519(apikey string, secretKey ed25519.PrivateKey) iface.SignerFunc {
	return func(data any) (map[string]any, error) {
		params, err := utils.MapAny(data)
		if err != nil {
			return nil, err
		}

		params["timestamp"] = time.Now().UnixMilli()
		params["apiKey"] = apikey
		Ed25519(params, secretKey)
		return params, nil
	}
}
