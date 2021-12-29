package helper

import (
	"fmt"
	"strconv"
	"strings"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Status:  status,
		Code:    code,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func GenerateFibonacci(nickname string) string {
	split := strings.Split(nickname, "-")
	length := len(split) - 1
	number := split[length]

	i, _ := strconv.Atoi(number)
	fibonacci := FibonacciRecursion(i)
	fmt.Println(fibonacci)
	return nickname
}

func FibonacciRecursion(number int) int {
	if number <= 1 {
		return number
	}
	return FibonacciRecursion(number-1) + FibonacciRecursion(number-2)
}
