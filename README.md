# PaymentApi SDK для Go

## Установка

Установите последнюю версию SDK с помощью команды:

```go get github.com/paymentsolution/golang-sdk```

## Использование

Создание экземпляра PaymentSdk
Для создания экземпляра PaymentSdk используйте билдера PaymentApiBuilder:

```go
sdkBuilder := paymentsdk.NewPaymentApiBuilder().
ApiURL("https://google.com").
Secret("your_api_secret")

sdk, err := sdkBuilder.Build()
if err != nil {
    log.Fatalf("Error building SDK: %v", err)
}

```

## P2P Транзакции

### Создание P2P транзакции

```go
//создание тела запроса
p2pTransactionRequest := paymentsdk.NewP2PTransactionRequest(
"a53fb67d-d807-4055-b7b3-56aafd88ff16",
"test1234",
"someId1",
"test11",
"some_type",
"https://localhost/response",
"",
paymentsdk.RUB,
paymentsdk.Card,
2000,
)

//опционально 
p2pTransactionRequest = p2pTransactionRequest.
WithEmail("email").
WithCustomerName("customer_name").
WithPayeerCardNumber("payerCardNumber")

p2prsp, err := sdk.P2P.CreateP2PTransaction(context.Background(), *p2pTransactionRequest)
if err != nil {
//обработать ошибку
}
//Бизнес логика
```

### Получение информации о P2P транзакции

```go
transactionID := "mock_transaction_id"
p2pResponse, err := sdk.P2P.GetP2PTransaction(transactionID)
if err != nil {
//обработать ошибку
}
//обработка логики
```

## Создание диспута
### Пример с использованием фреймворка gin
```go
type DisputRequest struct {
File          *multipart.FileHeader `form:"file"`
SecondFile    *multipart.FileHeader `form:"secondFile"`
Amount        int                   `form:"amount"`
TransactionId string                `form:"transactionId"`
}

func sendDisput(c *gin.Context) {
var req DisputRequest
if err := c.ShouldBind(&req); err != nil {
//обработать ошибку
}
file, err := req.File.Open()
if err != nil {
//обработать ошибку
}
//Второй файл опционален
secondFile, err := req.File.Open()
if err != nil {
//обработать ошибку
}

mockP2PDisputeRequest := bovasdk.NewP2PDisputeRequest(req.TransactionId, req.Amount, req.File.Filename, file)
//Второй файл опционален
mockP2PDisputeRequest = mockP2PDisputeRequest.WithProofImage2(req.SecondFile.Filename, secondFile)

rsp, err := sdk.P2P.CreateP2PDispute(context.Background(), mockP2PDisputeRequest)
if err != nil {
//обработать ошибку
}
//Бизнес логика
}
```

## Массовые Транзакции

### Создание массовой транзакции

```go
massTransactionRequest := bovasdk.NewMassTransactionRequest(
"a53fb67d-d807-4055-b7b3-56aafd88ff16",
"test1234",
"someId1",
"test11",
1000,
bovasdk.RUB,
bovasdk.Card,
)

//опционально
massTransactionRequest = massTransactionRequest.
WithBankName("bank").
WithRecipientFirstName("ivan").
WithRecipientLastName("ivanov")

p2prsp, err := sdk.MassTransaction.CreateMassTransaction(context.Background(), *massTransactionRequest)
if err != nil {
//обработать ошибку
}
//Бизнес логика
```

### Получение информации о массовой транзакции

```go
massTransactionID := "mock_transaction_id"
massTransactionResponse, err := sdk.MassTransaction.GetMassTransaction(massTransactionID)
if err != nil {
//обработать ошибку
}
//обработка логики
```

## callback

проверка подписи заголовка Signature при ответе от сервиса платежа 

```go
requestBody, err := io.ReadAll(c.Request.Body)
if err != nil {
//обработать ошибку
}
signature := c.Request.Header.Get("Signature")
verified := sdk.Encoder.VerifySignature(requestBody, signature)
if !verified {
//обработать невалидную сигнатуру
}
//обработать callback
```

## Опциональные настройки
### Логгирование

По умолчанию библиотека создает свой логгер с логами в формате json, логгер логгирует все входящие и исходящие запросы в
stdout,
Вы можете создать и настроить свой логгер реализовав интерфейс:

```go
type Logger interface {
	Enabled() bool

	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}
```

где параметр Enabled - отвечает за то, надо ли логгировать запросы и ответы или нет.
Вашу реализацию необзодимо подложить в NewPaymentApiBuilder при сборке:

```go
sdkBuilder := paymentsdk.NewPaymentApiBuilder().
ApiURL("https://google.com").
Secret("your_api_secret").
Logger(myCustomLogger)


sdk, err := sdkBuilder.Build()
if err != nil {
    log.Fatalf("Error building SDK: %v", err)
}
```

### Клиент для запросов

Кроме того, вы можете полностю передать свою структуру http.Client со своими настройками(включая логгирование, таймауты и т.д.) для
выполнения запросов

```go
sdkBuilder := paymentsdk.NewPaymentApiBuilder().
ApiURL("https://google.com").
Secret("your_api_secret").
Client(myCustomCLient)


sdk, err := sdkBuilder.Build()
if err != nil {
    log.Fatalf("Error building SDK: %v", err)
}
```