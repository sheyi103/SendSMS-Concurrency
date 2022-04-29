package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	wg.Add(5000)
	for i := 1; i <= 5000; i++ {
		go sendSMS()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}
func sendSMS() {
	counter := 0
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://41.76.195.35:8080/fcgi-bin/jar_http_sai.fcgi?username=accesstrans&password=accesstr&to=2348160584802&text=Debit%250aAmt%253aNGN10.75%250aAcc%253a008******703%250aDesc%253a099MJKL21305Mjpn%252fAmt+includes+COMM+%2526+VAT%252fAccount+Balance%250aTime%253a01%252f11%252f2021%250aAvail+Bal%253aNGN15%252c939.28%250aTotal%253aNGN15%252c989.28&dlr-mask=1&from=AccessBank&smsc=7334725312&dlr-url=http://thevartech.com/newdlr.php%253FsmsID%253D19%2526dlr%253D%2525d", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	//request.Header.Add("Accept", "application/json")
	//request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err.Error())
	}
	if bodyBytes != nil {
		counter++
		fmt.Printf("counter is: %d \n", counter)
	}
	fmt.Println(string(bodyBytes))
	//var responseObject Response
	//err = json.Unmarshal(bodyBytes, &responseObject)
	//if err != nil {
	//	return
	//}
	//fmt.Println("\nresponse", response)
	//fmt.Println("\nresponse.Body", response.Body)
	//fmt.Println("\nbodyBytes", bodyBytes)
	//fmt.Println("\nresponseObject", responseObject)

	//fmt.Println("\n", responseObject.Joke)
	wg.Done()
}
