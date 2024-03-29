package main

import (
	"root/Controllers"
)

func main() {
	//Controllers.TryAnonymous()

	//Controllers.GoMain()

	//Controllers.ChannelMain()

	//Controllers.RegexMain()

	//fmt.Println(Controllers.ReverseStringWithSameWordLength("bat and ball"))

	//Controllers.CreateConnection()

	//Controllers.CreateClient()

	//Controllers.ProfitMain()

	//Controllers.SelfMultiplierMain()
	// nums := []int{1, 1, 1, 2, 2, 3}
	// output, _ := Controllers.FrequentInteger(nums, 2)
	// fmt.Println(output)

	//Controllers.GoRoMain()

	//Controllers.WeatherStation("others")

	//Controllers.ShieldMain()

	//Controllers.GetDataFromElastic("yash")

	//Controllers.RedisUse()
	//fmt.Println(Controllers.GetDataFromElastic("topic"))

	// Controllers.Intro()

	// doneChan := make(chan bool)

	// go readInput(doneChan)

	// <-doneChan

	// close(doneChan)

	//fmt.Println(Controllers.IsBalanced("{{[]}}()"))

	Controllers.GraphMain()
}

// func readInput(doneChan chan bool) {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		res, done := CheckNo(scanner)

// 		if done {
// 			doneChan <- true
// 			return
// 		}

// 		fmt.Println(res)
// 		Controllers.Prompt()
// 	}
// }

// func CheckNo(scanner *bufio.Scanner) (string, bool)
// 	scanner.Scan()

// 	if strings.EqualFold(scanner.Text(), "q") {
// 		return "", true
// 	}

// 	num, err := strconv.Atoi(scanner.Text())
// 	if err != nil {
// 		return "please enter a whole num", false
// 	}

// 	_, msg := Controllers.IsPrime(num)

// 	return msg, false
// }
