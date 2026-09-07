package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Заранее подготавливаем некоторые ошибки
var (
	ErrUnsupportedPaymentMethod = errors.New("unsupported payment method")
)

// Ошибка нехватки средств
type InsufficientFundsError struct {
	RequestedAmount float64
	MaxAmount       float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("requested amount: %.2f, maximum allowed amount: %.2f", e.RequestedAmount, e.MaxAmount)
}

// PaymentProcessor обрабатывает платежи
type PaymentProcessor struct{}

// NewPaymentProcessor создает новый процессор платежей
func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{}
}

// ProcessPayment обрабатывает платеж
func (pp *PaymentProcessor) ProcessPayment(method string, amount float64) error {
	if method != "карта" && method != "СБП" {
		return ErrUnsupportedPaymentMethod
	}

	// Имитация проверки средств
	maxAmount := 1000.0
	if amount > maxAmount {
		return &InsufficientFundsError{
			RequestedAmount: amount,
			MaxAmount:       maxAmount,
		}
	}

	return nil
}

func HandlePayments() error {
	var method string
	var amountStr string

	// Спросить метод перевода в консоли
	fmt.Print("Введите метод перевода (карта/СБП): ")
	fmt.Scanln(&method)

	// Спросить сумму перевода в консоли
	fmt.Print("Введите сумму перевода: ")
	fmt.Scanln(&amountStr)

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return fmt.Errorf("invalid amount: %w", err)
	}

	pp := NewPaymentProcessor()
	if err := pp.ProcessPayment(method, amount); err != nil {
		return fmt.Errorf("process payment: %w", err)
	}

	return nil
}

func main() {
	if err := HandlePayments(); err != nil {
		if errors.Is(err, ErrUnsupportedPaymentMethod) {
			fmt.Printf("Ошибка: неподдерживаемый метод платежа.\n")
		} else if InsufficientFundsError, ok := errors.AsType[*InsufficientFundsError](err); ok {
			fmt.Printf("Ошибка: недостаточно средств. %s.\n", InsufficientFundsError)
		} else {
			fmt.Printf("Неизвестная ошибка: %s.\n", err)
		}
		os.Exit(1)
	}

	fmt.Println("Платеж успешно обработан!")
}
