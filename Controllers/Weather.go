package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type CityInfo struct {
	Name    string `json:"name"`
	Weather string `json:"weather"`
	Status  struct {
		Wind     string `json:"Wind"`
		Humidity string `json:"Humidity"`
	} `json:"status"`
}

func extractNumericalValues(s string) string {
	// Use regular expression to extract numerical values
	re := regexp.MustCompile(`(\d+(\.\d+)?)`)
	match := re.FindString(s)
	if match == "" {
		return ""
	}

	// Convert the matched string to float64
	return match
}

func WeatherStation(keyword string) ([]string, error) {
	url := "https://jsonmock.hackerrank.com/api/weather/search?name="
	req, err := http.NewRequest("GET", url+keyword, nil)
	if err != nil {
		return nil, err
	}

	// Make the HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	// Extract and format the information
	var result []string
	//var cities []CityInfo
	if data, ok := response["data"].([]interface{}); ok {

		for _, cityData := range data {
			if city, ok := cityData.(map[string]interface{}); ok {

				name := city["name"].(string)
				weather := city["weather"].(string)
				status := city["status"].([]interface{})
				wind := extractNumericalValues(status[0].(string))
				humdity := extractNumericalValues(status[1].(string))
				res := fmt.Sprintf("%s,%s,%s,%s", name, strings.TrimSpace(wind), strings.TrimSpace(humdity), weather[:2])
				result = append(result, res)
			}

		}
	}
	fmt.Println(result)

	return result, nil
}
