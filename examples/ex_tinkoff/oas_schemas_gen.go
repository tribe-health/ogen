// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"
)

// Ref: #/components/schemas/BrokerAccountType
type BrokerAccountType string

const (
	BrokerAccountTypeTinkoff    BrokerAccountType = "Tinkoff"
	BrokerAccountTypeTinkoffIis BrokerAccountType = "TinkoffIis"
)

// Ref: #/components/schemas/Candle
type Candle struct {
	Figi     string           "json:\"figi\""
	Interval CandleResolution "json:\"interval\""
	O        float64          "json:\"o\""
	C        float64          "json:\"c\""
	H        float64          "json:\"h\""
	L        float64          "json:\"l\""
	V        int32            "json:\"v\""
	// ISO8601.
	Time time.Time "json:\"time\""
}

// Интервал свечи и допустимый промежуток запроса:
// - 1min [1 minute, 1 day]
// - 2min [2 minutes, 1 day]
// - 3min [3 minutes, 1 day]
// - 5min [5 minutes, 1 day]
// - 10min [10 minutes, 1 day]
// - 15min [15 minutes, 1 day]
// - 30min [30 minutes, 1 day]
// - hour [1 hour, 7 days]
// - day [1 day, 1 year]
// - week [7 days, 2 years]
// - month [1 month, 10 years].
// Ref: #/components/schemas/CandleResolution
type CandleResolution string

const (
	CandleResolution1min  CandleResolution = "1min"
	CandleResolution2min  CandleResolution = "2min"
	CandleResolution3min  CandleResolution = "3min"
	CandleResolution5min  CandleResolution = "5min"
	CandleResolution10min CandleResolution = "10min"
	CandleResolution15min CandleResolution = "15min"
	CandleResolution30min CandleResolution = "30min"
	CandleResolutionHour  CandleResolution = "hour"
	CandleResolutionDay   CandleResolution = "day"
	CandleResolutionWeek  CandleResolution = "week"
	CandleResolutionMonth CandleResolution = "month"
)

// Ref: #/components/schemas/Candles
type Candles struct {
	Figi     string           "json:\"figi\""
	Interval CandleResolution "json:\"interval\""
	Candles  []Candle         "json:\"candles\""
}

// Ref: #/components/schemas/CandlesResponse
type CandlesResponse struct {
	TrackingId string  "json:\"trackingId\""
	Status     string  "json:\"status\""
	Payload    Candles "json:\"payload\""
}

func (*CandlesResponse) marketCandlesGetRes() {}

// Ref: #/components/schemas/Currencies
type Currencies struct {
	Currencies []CurrencyPosition "json:\"currencies\""
}

// Ref: #/components/schemas/Currency
type Currency string

const (
	CurrencyRUB Currency = "RUB"
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
	CurrencyHKD Currency = "HKD"
	CurrencyCHF Currency = "CHF"
	CurrencyJPY Currency = "JPY"
	CurrencyCNY Currency = "CNY"
	CurrencyTRY Currency = "TRY"
)

// Ref: #/components/schemas/CurrencyPosition
type CurrencyPosition struct {
	Currency Currency   "json:\"currency\""
	Balance  float64    "json:\"balance\""
	Blocked  OptFloat64 "json:\"blocked\""
}

// Ref: #/components/schemas/Empty
type Empty struct {
	TrackingId string       "json:\"trackingId\""
	Payload    EmptyPayload "json:\"payload\""
	Status     string       "json:\"status\""
}

func (*Empty) ordersCancelPostRes()             {}
func (*Empty) sandboxClearPostRes()             {}
func (*Empty) sandboxCurrenciesBalancePostRes() {}
func (*Empty) sandboxPositionsBalancePostRes()  {}
func (*Empty) sandboxRemovePostRes()            {}

type EmptyPayload struct{}

// Ref: #/components/schemas/Error
type Error struct {
	TrackingId string       "json:\"trackingId\""
	Status     string       "json:\"status\""
	Payload    ErrorPayload "json:\"payload\""
}

func (*Error) marketBondsGetRes()               {}
func (*Error) marketCandlesGetRes()             {}
func (*Error) marketCurrenciesGetRes()          {}
func (*Error) marketEtfsGetRes()                {}
func (*Error) marketOrderbookGetRes()           {}
func (*Error) marketSearchByFigiGetRes()        {}
func (*Error) marketSearchByTickerGetRes()      {}
func (*Error) marketStocksGetRes()              {}
func (*Error) operationsGetRes()                {}
func (*Error) ordersCancelPostRes()             {}
func (*Error) ordersGetRes()                    {}
func (*Error) ordersLimitOrderPostRes()         {}
func (*Error) ordersMarketOrderPostRes()        {}
func (*Error) portfolioCurrenciesGetRes()       {}
func (*Error) portfolioGetRes()                 {}
func (*Error) sandboxClearPostRes()             {}
func (*Error) sandboxCurrenciesBalancePostRes() {}
func (*Error) sandboxPositionsBalancePostRes()  {}
func (*Error) sandboxRegisterPostRes()          {}
func (*Error) sandboxRemovePostRes()            {}
func (*Error) userAccountsGetRes()              {}

type ErrorPayload struct {
	Message OptString "json:\"message\""
	Code    OptString "json:\"code\""
}

// Ref: #/components/schemas/InstrumentType
type InstrumentType string

const (
	InstrumentTypeStock    InstrumentType = "Stock"
	InstrumentTypeCurrency InstrumentType = "Currency"
	InstrumentTypeBond     InstrumentType = "Bond"
	InstrumentTypeEtf      InstrumentType = "Etf"
)

// Ref: #/components/schemas/LimitOrderRequest
type LimitOrderRequest struct {
	Lots      int32         "json:\"lots\""
	Operation OperationType "json:\"operation\""
	Price     float64       "json:\"price\""
}

// Ref: #/components/schemas/LimitOrderResponse
type LimitOrderResponse struct {
	TrackingId string           "json:\"trackingId\""
	Status     string           "json:\"status\""
	Payload    PlacedLimitOrder "json:\"payload\""
}

func (*LimitOrderResponse) ordersLimitOrderPostRes() {}

// Ref: #/components/schemas/MarketInstrument
type MarketInstrument struct {
	Figi   string    "json:\"figi\""
	Ticker string    "json:\"ticker\""
	Isin   OptString "json:\"isin\""
	// Шаг цены.
	MinPriceIncrement OptFloat64 "json:\"minPriceIncrement\""
	Lot               int32      "json:\"lot\""
	// Минимальное число инструментов для покупки должно
	// быть не меньше, чем размер лота х количество лотов.
	MinQuantity OptInt32       "json:\"minQuantity\""
	Currency    OptCurrency    "json:\"currency\""
	Name        string         "json:\"name\""
	Type        InstrumentType "json:\"type\""
}

// Ref: #/components/schemas/MarketInstrumentList
type MarketInstrumentList struct {
	Total       int32              "json:\"total\""
	Instruments []MarketInstrument "json:\"instruments\""
}

// Ref: #/components/schemas/MarketInstrumentListResponse
type MarketInstrumentListResponse struct {
	TrackingId string               "json:\"trackingId\""
	Status     string               "json:\"status\""
	Payload    MarketInstrumentList "json:\"payload\""
}

func (*MarketInstrumentListResponse) marketBondsGetRes()          {}
func (*MarketInstrumentListResponse) marketCurrenciesGetRes()     {}
func (*MarketInstrumentListResponse) marketEtfsGetRes()           {}
func (*MarketInstrumentListResponse) marketSearchByTickerGetRes() {}
func (*MarketInstrumentListResponse) marketStocksGetRes()         {}

// Ref: #/components/schemas/MarketOrderRequest
type MarketOrderRequest struct {
	Lots      int32         "json:\"lots\""
	Operation OperationType "json:\"operation\""
}

// Ref: #/components/schemas/MarketOrderResponse
type MarketOrderResponse struct {
	TrackingId string            "json:\"trackingId\""
	Status     string            "json:\"status\""
	Payload    PlacedMarketOrder "json:\"payload\""
}

func (*MarketOrderResponse) ordersMarketOrderPostRes() {}

// Ref: #/components/schemas/MoneyAmount
type MoneyAmount struct {
	Currency Currency "json:\"currency\""
	Value    float64  "json:\"value\""
}

// Ref: #/components/schemas/Operation
type Operation struct {
	ID         string           "json:\"id\""
	Status     OperationStatus  "json:\"status\""
	Trades     []OperationTrade "json:\"trades\""
	Commission OptMoneyAmount   "json:\"commission\""
	Currency   Currency         "json:\"currency\""
	Payment    float64          "json:\"payment\""
	Price      OptFloat64       "json:\"price\""
	// Число инструментов в выставленной заявке.
	Quantity OptInt32 "json:\"quantity\""
	// Число инструментов, исполненных в заявке.
	QuantityExecuted OptInt32          "json:\"quantityExecuted\""
	Figi             OptString         "json:\"figi\""
	InstrumentType   OptInstrumentType "json:\"instrumentType\""
	IsMarginCall     bool              "json:\"isMarginCall\""
	// ISO8601.
	Date          time.Time                      "json:\"date\""
	OperationType OptOperationTypeWithCommission "json:\"operationType\""
}

// Статус заявки.
// Ref: #/components/schemas/OperationStatus
type OperationStatus string

const (
	OperationStatusDone     OperationStatus = "Done"
	OperationStatusDecline  OperationStatus = "Decline"
	OperationStatusProgress OperationStatus = "Progress"
)

// Ref: #/components/schemas/OperationTrade
type OperationTrade struct {
	TradeId string "json:\"tradeId\""
	// ISO8601.
	Date     time.Time "json:\"date\""
	Price    float64   "json:\"price\""
	Quantity int32     "json:\"quantity\""
}

// Ref: #/components/schemas/OperationType
type OperationType string

const (
	OperationTypeBuy  OperationType = "Buy"
	OperationTypeSell OperationType = "Sell"
)

// Ref: #/components/schemas/OperationTypeWithCommission
type OperationTypeWithCommission string

const (
	OperationTypeWithCommissionBuy                OperationTypeWithCommission = "Buy"
	OperationTypeWithCommissionBuyCard            OperationTypeWithCommission = "BuyCard"
	OperationTypeWithCommissionSell               OperationTypeWithCommission = "Sell"
	OperationTypeWithCommissionBrokerCommission   OperationTypeWithCommission = "BrokerCommission"
	OperationTypeWithCommissionExchangeCommission OperationTypeWithCommission = "ExchangeCommission"
	OperationTypeWithCommissionServiceCommission  OperationTypeWithCommission = "ServiceCommission"
	OperationTypeWithCommissionMarginCommission   OperationTypeWithCommission = "MarginCommission"
	OperationTypeWithCommissionOtherCommission    OperationTypeWithCommission = "OtherCommission"
	OperationTypeWithCommissionPayIn              OperationTypeWithCommission = "PayIn"
	OperationTypeWithCommissionPayOut             OperationTypeWithCommission = "PayOut"
	OperationTypeWithCommissionTax                OperationTypeWithCommission = "Tax"
	OperationTypeWithCommissionTaxLucre           OperationTypeWithCommission = "TaxLucre"
	OperationTypeWithCommissionTaxDividend        OperationTypeWithCommission = "TaxDividend"
	OperationTypeWithCommissionTaxCoupon          OperationTypeWithCommission = "TaxCoupon"
	OperationTypeWithCommissionTaxBack            OperationTypeWithCommission = "TaxBack"
	OperationTypeWithCommissionRepayment          OperationTypeWithCommission = "Repayment"
	OperationTypeWithCommissionPartRepayment      OperationTypeWithCommission = "PartRepayment"
	OperationTypeWithCommissionCoupon             OperationTypeWithCommission = "Coupon"
	OperationTypeWithCommissionDividend           OperationTypeWithCommission = "Dividend"
	OperationTypeWithCommissionSecurityIn         OperationTypeWithCommission = "SecurityIn"
	OperationTypeWithCommissionSecurityOut        OperationTypeWithCommission = "SecurityOut"
)

// Ref: #/components/schemas/Operations
type Operations struct {
	Operations []Operation "json:\"operations\""
}

// Ref: #/components/schemas/OperationsResponse
type OperationsResponse struct {
	TrackingId string     "json:\"trackingId\""
	Status     string     "json:\"status\""
	Payload    Operations "json:\"payload\""
}

func (*OperationsResponse) operationsGetRes() {}

// NewOptBrokerAccountType returns new OptBrokerAccountType with value set to v.
func NewOptBrokerAccountType(v BrokerAccountType) OptBrokerAccountType {
	return OptBrokerAccountType{
		Value: v,
		Set:   true,
	}
}

// OptBrokerAccountType is optional BrokerAccountType.
type OptBrokerAccountType struct {
	Value BrokerAccountType
	Set   bool
}

// IsSet returns true if OptBrokerAccountType was set.
func (o OptBrokerAccountType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBrokerAccountType) Reset() {
	var v BrokerAccountType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBrokerAccountType) SetTo(v BrokerAccountType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBrokerAccountType) Get() (v BrokerAccountType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBrokerAccountType) Or(d BrokerAccountType) BrokerAccountType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCurrency returns new OptCurrency with value set to v.
func NewOptCurrency(v Currency) OptCurrency {
	return OptCurrency{
		Value: v,
		Set:   true,
	}
}

// OptCurrency is optional Currency.
type OptCurrency struct {
	Value Currency
	Set   bool
}

// IsSet returns true if OptCurrency was set.
func (o OptCurrency) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCurrency) Reset() {
	var v Currency
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCurrency) SetTo(v Currency) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCurrency) Get() (v Currency, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCurrency) Or(d Currency) Currency {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInstrumentType returns new OptInstrumentType with value set to v.
func NewOptInstrumentType(v InstrumentType) OptInstrumentType {
	return OptInstrumentType{
		Value: v,
		Set:   true,
	}
}

// OptInstrumentType is optional InstrumentType.
type OptInstrumentType struct {
	Value InstrumentType
	Set   bool
}

// IsSet returns true if OptInstrumentType was set.
func (o OptInstrumentType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInstrumentType) Reset() {
	var v InstrumentType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInstrumentType) SetTo(v InstrumentType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInstrumentType) Get() (v InstrumentType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInstrumentType) Or(d InstrumentType) InstrumentType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMoneyAmount returns new OptMoneyAmount with value set to v.
func NewOptMoneyAmount(v MoneyAmount) OptMoneyAmount {
	return OptMoneyAmount{
		Value: v,
		Set:   true,
	}
}

// OptMoneyAmount is optional MoneyAmount.
type OptMoneyAmount struct {
	Value MoneyAmount
	Set   bool
}

// IsSet returns true if OptMoneyAmount was set.
func (o OptMoneyAmount) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMoneyAmount) Reset() {
	var v MoneyAmount
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMoneyAmount) SetTo(v MoneyAmount) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMoneyAmount) Get() (v MoneyAmount, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMoneyAmount) Or(d MoneyAmount) MoneyAmount {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptOperationTypeWithCommission returns new OptOperationTypeWithCommission with value set to v.
func NewOptOperationTypeWithCommission(v OperationTypeWithCommission) OptOperationTypeWithCommission {
	return OptOperationTypeWithCommission{
		Value: v,
		Set:   true,
	}
}

// OptOperationTypeWithCommission is optional OperationTypeWithCommission.
type OptOperationTypeWithCommission struct {
	Value OperationTypeWithCommission
	Set   bool
}

// IsSet returns true if OptOperationTypeWithCommission was set.
func (o OptOperationTypeWithCommission) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptOperationTypeWithCommission) Reset() {
	var v OperationTypeWithCommission
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptOperationTypeWithCommission) SetTo(v OperationTypeWithCommission) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptOperationTypeWithCommission) Get() (v OperationTypeWithCommission, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptOperationTypeWithCommission) Or(d OperationTypeWithCommission) OperationTypeWithCommission {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSandboxRegisterRequest returns new OptSandboxRegisterRequest with value set to v.
func NewOptSandboxRegisterRequest(v SandboxRegisterRequest) OptSandboxRegisterRequest {
	return OptSandboxRegisterRequest{
		Value: v,
		Set:   true,
	}
}

// OptSandboxRegisterRequest is optional SandboxRegisterRequest.
type OptSandboxRegisterRequest struct {
	Value SandboxRegisterRequest
	Set   bool
}

// IsSet returns true if OptSandboxRegisterRequest was set.
func (o OptSandboxRegisterRequest) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSandboxRegisterRequest) Reset() {
	var v SandboxRegisterRequest
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSandboxRegisterRequest) SetTo(v SandboxRegisterRequest) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSandboxRegisterRequest) Get() (v SandboxRegisterRequest, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSandboxRegisterRequest) Or(d SandboxRegisterRequest) SandboxRegisterRequest {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Order
type Order struct {
	OrderId       string        "json:\"orderId\""
	Figi          string        "json:\"figi\""
	Operation     OperationType "json:\"operation\""
	Status        OrderStatus   "json:\"status\""
	RequestedLots int32         "json:\"requestedLots\""
	ExecutedLots  int32         "json:\"executedLots\""
	Type          OrderType     "json:\"type\""
	Price         float64       "json:\"price\""
}

// Ref: #/components/schemas/OrderResponse
type OrderResponse struct {
	Price    float64 "json:\"price\""
	Quantity int32   "json:\"quantity\""
}

// Статус заявки.
// Ref: #/components/schemas/OrderStatus
type OrderStatus string

const (
	OrderStatusNew            OrderStatus = "New"
	OrderStatusPartiallyFill  OrderStatus = "PartiallyFill"
	OrderStatusFill           OrderStatus = "Fill"
	OrderStatusCancelled      OrderStatus = "Cancelled"
	OrderStatusReplaced       OrderStatus = "Replaced"
	OrderStatusPendingCancel  OrderStatus = "PendingCancel"
	OrderStatusRejected       OrderStatus = "Rejected"
	OrderStatusPendingReplace OrderStatus = "PendingReplace"
	OrderStatusPendingNew     OrderStatus = "PendingNew"
)

// Тип заявки.
// Ref: #/components/schemas/OrderType
type OrderType string

const (
	OrderTypeLimit  OrderType = "Limit"
	OrderTypeMarket OrderType = "Market"
)

// Ref: #/components/schemas/Orderbook
type Orderbook struct {
	Figi        string          "json:\"figi\""
	Depth       int32           "json:\"depth\""
	Bids        []OrderResponse "json:\"bids\""
	Asks        []OrderResponse "json:\"asks\""
	TradeStatus TradeStatus     "json:\"tradeStatus\""
	// Шаг цены.
	MinPriceIncrement float64 "json:\"minPriceIncrement\""
	// Номинал для облигаций.
	FaceValue  OptFloat64 "json:\"faceValue\""
	LastPrice  OptFloat64 "json:\"lastPrice\""
	ClosePrice OptFloat64 "json:\"closePrice\""
	// Верхняя граница цены.
	LimitUp OptFloat64 "json:\"limitUp\""
	// Нижняя граница цены.
	LimitDown OptFloat64 "json:\"limitDown\""
}

// Ref: #/components/schemas/OrderbookResponse
type OrderbookResponse struct {
	TrackingId string    "json:\"trackingId\""
	Status     string    "json:\"status\""
	Payload    Orderbook "json:\"payload\""
}

func (*OrderbookResponse) marketOrderbookGetRes() {}

// Ref: #/components/schemas/OrdersResponse
type OrdersResponse struct {
	TrackingId string  "json:\"trackingId\""
	Status     string  "json:\"status\""
	Payload    []Order "json:\"payload\""
}

func (*OrdersResponse) ordersGetRes() {}

// Ref: #/components/schemas/PlacedLimitOrder
type PlacedLimitOrder struct {
	OrderId      string        "json:\"orderId\""
	Operation    OperationType "json:\"operation\""
	Status       OrderStatus   "json:\"status\""
	RejectReason OptString     "json:\"rejectReason\""
	// Сообщение об ошибке.
	Message       OptString      "json:\"message\""
	RequestedLots int            "json:\"requestedLots\""
	ExecutedLots  int            "json:\"executedLots\""
	Commission    OptMoneyAmount "json:\"commission\""
}

// Ref: #/components/schemas/PlacedMarketOrder
type PlacedMarketOrder struct {
	OrderId      string        "json:\"orderId\""
	Operation    OperationType "json:\"operation\""
	Status       OrderStatus   "json:\"status\""
	RejectReason OptString     "json:\"rejectReason\""
	// Сообщение об ошибке.
	Message       OptString      "json:\"message\""
	RequestedLots int            "json:\"requestedLots\""
	ExecutedLots  int            "json:\"executedLots\""
	Commission    OptMoneyAmount "json:\"commission\""
}

// Ref: #/components/schemas/Portfolio
type Portfolio struct {
	Positions []PortfolioPosition "json:\"positions\""
}

// Ref: #/components/schemas/PortfolioCurrenciesResponse
type PortfolioCurrenciesResponse struct {
	TrackingId string     "json:\"trackingId\""
	Status     string     "json:\"status\""
	Payload    Currencies "json:\"payload\""
}

func (*PortfolioCurrenciesResponse) portfolioCurrenciesGetRes() {}

// Ref: #/components/schemas/PortfolioPosition
type PortfolioPosition struct {
	Figi                      string         "json:\"figi\""
	Ticker                    OptString      "json:\"ticker\""
	Isin                      OptString      "json:\"isin\""
	InstrumentType            InstrumentType "json:\"instrumentType\""
	Balance                   float64        "json:\"balance\""
	Blocked                   OptFloat64     "json:\"blocked\""
	ExpectedYield             OptMoneyAmount "json:\"expectedYield\""
	Lots                      int32          "json:\"lots\""
	AveragePositionPrice      OptMoneyAmount "json:\"averagePositionPrice\""
	AveragePositionPriceNoNkd OptMoneyAmount "json:\"averagePositionPriceNoNkd\""
	Name                      string         "json:\"name\""
}

// Ref: #/components/schemas/PortfolioResponse
type PortfolioResponse struct {
	TrackingId string    "json:\"trackingId\""
	Status     string    "json:\"status\""
	Payload    Portfolio "json:\"payload\""
}

func (*PortfolioResponse) portfolioGetRes() {}

type SSOAuth struct {
	Token string
}

// Ref: #/components/schemas/SandboxAccount
type SandboxAccount struct {
	BrokerAccountType BrokerAccountType "json:\"brokerAccountType\""
	BrokerAccountId   string            "json:\"brokerAccountId\""
}

// Ref: #/components/schemas/SandboxCurrency
type SandboxCurrency string

const (
	SandboxCurrencyRUB SandboxCurrency = "RUB"
	SandboxCurrencyUSD SandboxCurrency = "USD"
	SandboxCurrencyEUR SandboxCurrency = "EUR"
	SandboxCurrencyGBP SandboxCurrency = "GBP"
	SandboxCurrencyHKD SandboxCurrency = "HKD"
	SandboxCurrencyCHF SandboxCurrency = "CHF"
	SandboxCurrencyJPY SandboxCurrency = "JPY"
	SandboxCurrencyCNY SandboxCurrency = "CNY"
	SandboxCurrencyTRY SandboxCurrency = "TRY"
)

// Ref: #/components/schemas/SandboxRegisterRequest
type SandboxRegisterRequest struct {
	BrokerAccountType OptBrokerAccountType "json:\"brokerAccountType\""
}

// Ref: #/components/schemas/SandboxRegisterResponse
type SandboxRegisterResponse struct {
	TrackingId string         "json:\"trackingId\""
	Status     string         "json:\"status\""
	Payload    SandboxAccount "json:\"payload\""
}

func (*SandboxRegisterResponse) sandboxRegisterPostRes() {}

// Ref: #/components/schemas/SandboxSetCurrencyBalanceRequest
type SandboxSetCurrencyBalanceRequest struct {
	Currency SandboxCurrency "json:\"currency\""
	Balance  float64         "json:\"balance\""
}

// Ref: #/components/schemas/SandboxSetPositionBalanceRequest
type SandboxSetPositionBalanceRequest struct {
	Figi    OptString "json:\"figi\""
	Balance float64   "json:\"balance\""
}

// Ref: #/components/schemas/SearchMarketInstrument
type SearchMarketInstrument struct {
	Figi   string    "json:\"figi\""
	Ticker string    "json:\"ticker\""
	Isin   OptString "json:\"isin\""
	// Шаг цены.
	MinPriceIncrement OptFloat64     "json:\"minPriceIncrement\""
	Lot               int32          "json:\"lot\""
	Currency          OptCurrency    "json:\"currency\""
	Name              string         "json:\"name\""
	Type              InstrumentType "json:\"type\""
}

// Ref: #/components/schemas/SearchMarketInstrumentResponse
type SearchMarketInstrumentResponse struct {
	TrackingId string                 "json:\"trackingId\""
	Status     string                 "json:\"status\""
	Payload    SearchMarketInstrument "json:\"payload\""
}

func (*SearchMarketInstrumentResponse) marketSearchByFigiGetRes() {}

// Ref: #/components/schemas/TradeStatus
type TradeStatus string

const (
	TradeStatusNormalTrading          TradeStatus = "NormalTrading"
	TradeStatusNotAvailableForTrading TradeStatus = "NotAvailableForTrading"
)

// Ref: #/components/schemas/UserAccount
type UserAccount struct {
	BrokerAccountType BrokerAccountType "json:\"brokerAccountType\""
	BrokerAccountId   string            "json:\"brokerAccountId\""
}

// Ref: #/components/schemas/UserAccounts
type UserAccounts struct {
	Accounts []UserAccount "json:\"accounts\""
}

// Ref: #/components/schemas/UserAccountsResponse
type UserAccountsResponse struct {
	TrackingId string       "json:\"trackingId\""
	Status     string       "json:\"status\""
	Payload    UserAccounts "json:\"payload\""
}

func (*UserAccountsResponse) userAccountsGetRes() {}
