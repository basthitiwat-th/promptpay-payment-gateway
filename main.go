package main

import (
	"time"

	"promptpay-payment-gateway/api"
	"promptpay-payment-gateway/client"
	"promptpay-payment-gateway/configs"
	"promptpay-payment-gateway/database"
	"promptpay-payment-gateway/repository"
	"promptpay-payment-gateway/service"
)

func main() {
	conf, err := configs.LoadConfig("./configs")
	if err != nil {
		panic(err)
	}

	mySQL, err := database.NewMySqlDB(conf.MySQL, conf.Secrets)
	if err != nil {
		panic(err)
	}
	defer mySQL.Close()

	repository := repository.NewTransactionsStore(mySQL.Client)
	clientPromptpay := client.NewPromptPayClient(conf.PromptPayClient)
	clientMerchant := client.NewMerchantClient(conf.MerchantClient)

	svc := service.NewPaymentService(configs.Config{}, repository, *clientPromptpay, *clientMerchant)
	handler := api.NewHandler(svc)

	server := api.NewServer(10 * time.Second)
	api.Route(server.Echo, handler)
	server.Start("8080")
}
