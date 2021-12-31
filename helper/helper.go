package helper

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

func FibonacciRecursion(number int) int {
	if number <= 1 {
		return number
	}
	return FibonacciRecursion(number-1) + FibonacciRecursion(number-2)
}

func CheckPrime(number int) bool {
	isPrime := true
	if number == 0 || number == 1 {
		isPrime = false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
	}
	return isPrime
}
