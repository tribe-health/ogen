// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"
)

type MarketCandlesGetParams struct {
	// FIGI.
	Figi string
	// Начало временного промежутка.
	From time.Time
	// Конец временного промежутка.
	To time.Time
	// Интервал свечи.
	Interval CandleResolution
}

type MarketOrderbookGetParams struct {
	// FIGI.
	Figi string
	// Глубина стакана [1..20].
	Depth int32
}

type MarketSearchByFigiGetParams struct {
	// FIGI.
	Figi string
}

type MarketSearchByTickerGetParams struct {
	// Тикер инструмента.
	Ticker string
}

type OperationsGetParams struct {
	// Начало временного промежутка.
	From time.Time
	// Конец временного промежутка.
	To time.Time
	// Figi инструмента для фильтрации.
	Figi OptString
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type OrdersCancelPostParams struct {
	// ID заявки.
	OrderId string
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type OrdersGetParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type OrdersLimitOrderPostParams struct {
	// FIGI инструмента.
	Figi string
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type OrdersMarketOrderPostParams struct {
	// FIGI инструмента.
	Figi string
	// Уникальный идентификатор счета (по умолчанию -
	// Тинькофф).
	BrokerAccountId OptString
}

type PortfolioCurrenciesGetParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type PortfolioGetParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type SandboxClearPostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type SandboxCurrenciesBalancePostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type SandboxPositionsBalancePostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type SandboxRemovePostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}
