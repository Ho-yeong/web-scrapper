// package nothing

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

type requestResult struct {
	url    string
	status string
}

func nothing() {
	// 맵 타입을 선언할때 make를 사용하지 않으면 nil 으로 선언되기 때문에
	// 맵 안에 변수를 넣을 수 가 없다.
	results := make(map[string]string)
	c := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.facebook.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.instagram.com",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} else {
	}
	c <- requestResult{url: url, status: status}
}

// 함수 앞에 go 를 붙여줌으로써 멀티 스레드 방식으로 동작함
// 하지만 메인 함수가 종료되면 내부 함수의 종료여부는 상관없이 모두 종료됨
// 그래서 채널을 사용하여 내부 함수가 끝날때 까지 기다릴 수 있음
// blocking operation 이라고 부름

// func main() {
// 	go function1()
// 	function2()
// }

// 자바스크립트, 파이썬 타입의 싱글스레드 타입의 처리방식
// go lang은 멀티스레드 방식으로 동작 가능하다

// var errRequestFailed = errors.New("Request Failed")

// func main() {

// 	var results = make(map[string]string)

// 	urls := []string{
// 		"https://www.airbnb.com",
// 		"https://www.google.com",
// 		"https://www.amazon.com",
// 		"https://www.facebook.com",
// 		"https://www.reddit.com",
// 		"https://soundcloud.com",
// 		"https://www.instagram.com",
// 	}

// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}

// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}
// }

// func hitURL(url string) error {
// 	fmt.Println("checking : ", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		return errRequestFailed
// 	}
// 	return nil
// }
