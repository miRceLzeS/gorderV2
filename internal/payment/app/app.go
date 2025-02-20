package app

import "github.com/miRceLzeS/gorder-v2/payment/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	CreatePayment command.CreatePaymentHandler
}
