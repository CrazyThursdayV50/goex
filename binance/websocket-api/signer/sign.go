package signer

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"slices"
	"strings"
)

// Sign 通用签名函数，使用 Ed25519 私钥对参数进行签名
func SignEd25519(params map[string]string, secretkey ed25519.PrivateKey) string {
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
		if v != "" { // 只包含非空值
			parts = append(parts, fmt.Sprintf("%s=%s", k, v))
		}
	}

	payload := strings.Join(parts, "&")

	signature := ed25519.Sign(secretkey, []byte(payload))
	// 5. 将签名转换为 base64 字符串
	return base64.StdEncoding.EncodeToString(signature)
}
