package Controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func ReqMain() {
	url := "https:www.baidu.com"
	//method := "GET"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("ayshds")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	defer res.Body.Close()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}

	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	defer res.Body.Close()

	time.Sleep(time.Millisecond * 1)

}
