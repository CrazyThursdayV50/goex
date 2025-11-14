package klines

type Params struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime int64  `json:"startTime,string,omitempty"`
	EndTime   int64  `json:"endTime,string,omitempty"`
	Limit     int    `json:"limit,string,omitempty"`
}

/*
 [
   [
     1499040000000,      // 开盘时间
     "0.01634790",       // 开盘价
     "0.80000000",       // 最高价
     "0.01575800",       // 最低价
     "0.01577100",       // 收盘价(当前K线未结束的即为最新价)
     "148976.11427815",  // 成交量
     1499644799999,      // 收盘时间
     "2434.19055334",    // 成交额
     308,                // 成交笔数
     "1756.87402397",    // 主动买入成交量
     "28.46694368",      // 主动买入成交额
     "17928899.62484339" // 请忽略该参数
   ]
 ]
*/

type RawKline []any

type Result []RawKline

func (r RawKline) OpenTime() int64 {
	t, ok := r[0].(int64)
	if ok {
		return t
	}
	return int64(r[0].(float64))
}

func (r RawKline) Open() string {
	return r[1].(string)
}

func (r RawKline) High() string {
	return r[2].(string)
}

func (r RawKline) Low() string {
	return r[3].(string)
}

func (r RawKline) Close() string {
	return r[4].(string)
}

func (r RawKline) Volume() string {
	return r[5].(string)
}

func (r RawKline) CloseTime() int64 {
	t, ok := r[6].(int64)
	if ok {
		return t
	}
	return int64(r[6].(float64))
}

func (r RawKline) Amount() string {
	return r[7].(string)
}

func (r RawKline) TradeCount() int64 {
	t, ok := r[8].(int64)
	if ok {
		return t
	}
	return int64(r[8].(float64))
}

func (r RawKline) TakerBuyVolume() string {
	return r[9].(string)
}

func (r RawKline) TakerBuyAmount() string {
	return r[10].(string)
}

func (r RawKline) Ignore() string {
	return r[11].(string)
}
