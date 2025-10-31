package variables

import "time"

// 是否测试环境
var isTest = false

func IsTest() bool { return isTest }
func SetIsTest()   { isTest = true }

// 代理
var proxy = ""

func GetProxy() string    { return proxy }
func SetProxy(url string) { proxy = url }

// PING PONG 超时时间
var writeControlTimeout = time.Second

func WriteControlTimeout() time.Duration           { return writeControlTimeout }
func SetWriteControlTimeout(timeout time.Duration) { writeControlTimeout = timeout }
