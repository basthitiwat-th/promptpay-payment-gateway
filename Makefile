start-promptpay:
	go run mock-merchant/main.go


start-merchant:
	go run mock-promptpay/main.go


start-payment-gateway:
	go run main.go


