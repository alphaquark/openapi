package upbit

type Response struct {
	Market             string  `json:"market"`                // 종목 구분 코드
	TradeDate          string  `json:"trade_date"`            // 최근 거래 일자 (UTC)
	TradeTime          string  `json:"trade_time"`            // 최근 거래 시각 (UTC)
	TradeDateKST       string  `json:"trade_date_kst"`        // 최근 거래 일자 (KST)
	TradeTimeKST       string  `json:"trade_time_kst"`        // 최근 거래 시각 (KST)
	TradeTimestamp     uint    `json:"trade_timestamp"`       //
	OpeningPrice       float32 `json:"opening_price"`         // 시가
	HighPrice          float32 `json:"high_price"`            // 고가
	LowPrice           float32 `json:"low_price"`             // 저가
	TradePrice         float32 `json:"trade_price"`           // 종가
	PrevClosingPrice   float32 `json:"prev_closing_price"`    // 전일종가
	Change             string  `json:"change"`                // 가격변화
	ChangePrice        float32 `json:"change_price"`          // 변화액의 절대값
	ChangeRate         float32 `json:"change_rate"`           // 번화율의 절대값
	SignedChangePrice  float32 `json:"signed_change_price"`   // 부호가 있는 변화액
	SignedChangeRate   float32 `json:"signed_change_rate"`    // 부호가 있는 변화율
	TradeVolume        float32 `json:"trade_volume"`          // 가장 최근 거래량
	AccTradePrice      float32 `json:"acc_trade_price"`       // 누적 거래대금 (UTC 0시 기준)
	AccTradePrice24H   float32 `json:"acc_trade_price_24h"`   // 24시간 누적 거래대금
	AccTradeVolume     float32 `json:"acc_trade_volume"`      // 누적 거래량 (UTC 0시 기준)
	AccTradeVolume24H  float32 `json:"acc_trade_volume_24h"`  // 24시간 누적 거래대금
	Highest52WeekPrice float32 `json:"highest_52_week_price"` // 52주 신고가
	Highest52WeekDate  string  `json:"highest_52_week_date"`  // 52주 신고가 달성일
	Lowest52WeekPrice  float32 `json:"lowest_52_week_price"`  // 52주 신저가
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`   // 52주 신저가 달성일
	Timestamp          uint    `json:"timestamp"`             // 타임스탬프
}
