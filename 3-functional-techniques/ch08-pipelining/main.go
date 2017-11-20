package main

import (
	"log"
	"fmt"
	gc "github.com/go-goodies/go_currency"
	gu "github.com/go-goodies/go_utils"
	"strings"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"bytes"
)

func main() {
	pipeline := BuildPipeline(Decrypt{}, Charge{}, Authenticate{})

	go func(){
		orders := GetOrders()
		for _, order := range orders {
			fmt.Printf("order: %v\n", order)
			pipeline.Send(*order)
		}
		log.Println("Close Pipeline")
		pipeline.Close()
	}()

	pipeline.Receive(func(o Order){
		log.Printf("Received: %v", o)
	})
}


type LineItem struct {
	Description string
	Count       int
	PriceUSD    gc.USD
}

type Order struct {
	OrderNumber int
	IsAuthenticated bool
	IsDecrypted bool
	Credentials string
	CCardNumber string
	CCardExpDate string
	LineItems []LineItem
}

func GetOrders() []*Order {

	order1 := &Order{
		10001,
		false,
		false,
		"alice,secret",
		"7b/HWvtIB9a16AYk+Yv6WWwer3GFbxpjoR+GO9iHIYY=",
		"0922",
		[]LineItem{
			{"Apples", 1, gc.USD{4, 50}},
			{"Oranges", 4, gc.USD{12, 00}},
		},
	}

	order2 := &Order{
		10002,
		false,
		false,
		"bob,secret",
		"EOc3kF/OmxY+dRCaYRrey8h24QoGzVU0/T2QKVCHb1Q=",
		"0123",
		[]LineItem{
			{"Milk", 2, gc.USD{8, 00}},
			{"Sugar", 1, gc.USD{2, 25}},
			{"Salt", 3, gc.USD{3, 75}},
		},
	}
	orders := []*Order{order1, order2}
	return orders
}

var AESEncryptionKey = "a very very very very secret key"

func encrypt(rawString string) (string, error) {
	rawBytes := []byte(rawString)
	block, err := aes.NewCipher([]byte(AESEncryptionKey))
	if err != nil {
		return "", err
	}
	if len(rawBytes)%aes.BlockSize != 0 {
		padding := aes.BlockSize - len(rawBytes)%aes.BlockSize
		padText := bytes.Repeat([]byte{byte(0)}, padding)
		rawBytes = append(rawBytes, padText...)
	}
	ciphertext := make([]byte, aes.BlockSize+len(rawBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], rawBytes)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(encodedValue string) (string, error) {
	block, err := aes.NewCipher([]byte(AESEncryptionKey))
	if err != nil {
		return "", err
	}
	b, err := base64.StdEncoding.DecodeString(encodedValue)
	if err != nil {
		return "", err
	}
	if len(b) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := b[:aes.BlockSize]
	b = b[aes.BlockSize:]
	if len(b)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(b, b)
	b = bytes.TrimRight(b, "\x00")
	return string(b), nil
}

type Filterer interface {
	Filter(input chan Order) chan Order
}

type Authenticate struct {}
func (a Authenticate) Filter(input chan Order) chan Order {
	output := make(chan Order)
	go func(){
		for order := range input {
			usernamePwd := strings.Split(order.Credentials, ",")
			if usernamePwd[1] == "secret" {
				order.IsAuthenticated = true
				output <- order
			} else {
				order.IsAuthenticated = false
				errMsg := fmt.Sprintf("Error: Invalid password for order Id: %d", order.OrderNumber)
				log.Println("Error:", errors.New(errMsg))
				output <- order
			}
		}
		close(output)
	}()
	return output
}

type Decrypt struct {}
func (d Decrypt) Filter(input chan Order) chan Order {
	output := make(chan Order)
	go func(){
		for order := range input {
			creditCardNo, err := decrypt(order.CCardNumber)
			if err != nil {
				order.IsDecrypted = false
				log.Println("Error:", err.Error())
			} else {
				order.IsDecrypted = true
				order.CCardNumber = creditCardNo
				output <- order
			}
		}
		close(output)
	}()
	return output
}

func ChargeCard(ccardNo string, amount gc.USD) {
	fmt.Printf("Credit card %v%v charged %v\n", gu.Dashes(len(ccardNo)-4, "X"), ccardNo[len(ccardNo)-4:], amount)
}

type Charge struct {}
func (c Charge) Filter(input chan Order) chan Order {
	output := make(chan Order)
	go func(){
		for order := range input {
			if order.IsAuthenticated && order.IsDecrypted {
				total := gc.USD{0, 0}
				for _, li := range order.LineItems {
					total, _ = total.Add(li.PriceUSD)
				}
				ChargeCard(order.CCardNumber, total)
				output <- order
			} else {
				errMsg := fmt.Sprintf("Error: Unable to charge order Id: %d", order.OrderNumber)
				log.Println("Error:", errors.New(errMsg))
			}
		}
		close(output)
	}()
	return output
}

func BuildPipeline(filters ...Filterer) Filter {
	source := make(chan Order)
	var nextFilter chan Order
	for _, filter := range filters {
		if nextFilter == nil {
			nextFilter = filter.Filter(source)
		} else {
			nextFilter = filter.Filter(nextFilter)
		}
	}
	return Filter{ input: source, output: nextFilter }
}

type Filter struct {
	input  chan Order
	output chan Order
}

func (f *Filter) Send(order Order) {
	f.input <- order
}

func (f *Filter) Receive(callback func(Order)){
	for o := range f.output {
		callback(o)
	}
}

func (f *Filter) Close() {
	close(f.input)
}
