package Controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ShieldMain() {

	CreateShieldDb()

	fmt.Println("=======------S.H.I.E.L.D =========")
	fmt.Println("Welcome to Captain Fury.")
	fmt.Println("1. Check the missions")
	fmt.Println("2. Assign mission to Avengers")
	fmt.Println("3. Check mission’s details")
	fmt.Println("4. Check Avenger’s details")
	fmt.Println("5. Update Mission’s status")
	fmt.Println("6. List Avengers")
	fmt.Println("7. Assign avenger to mission.")
	for {
		var selectedOption string
		fmt.Print("Enter your choice (or 'q' to quit): ")
		fmt.Scanln(&selectedOption)

		handleInput(selectedOption)
	}
}

func handleInput(option string) {
	switch option {
	case "1":
		fmt.Println("Checking the missions...")
		missionListing()
	case "2":
		missionCreate()
	case "3":
		missionGet()
	case "4":
		fmt.Println("Checking Avenger’s details...")
	case "5":
		fmt.Println("Updating Mission’s status...")
	case "6":
		fmt.Println("Listing Avengers...")
	case "7":
		fmt.Println("Assigning avenger to mission...")
	case "q":
		defer dbShield.Close()
		fmt.Println("Exiting the program.")
		os.Exit(0)
	default:
		fmt.Println("Invalid option. Please try again.")
	}
}

func missionGet() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Mission Name:")
	missionName, _ := reader.ReadString('\n')
	rows, err := dbShield.Query("SELECT md.description, md.status, GROUP_CONCAT(av.name) FROM mission_details md JOIN avenger_mission_mapping amm ON md.id = amm.mission_id JOIN avengers av ON av.id = amm.avenger_id where md.name = ? GROUP BY md.id", strings.TrimSpace(missionName))
	if err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		fmt.Printf("No mission found with name %v\n", missionName)
	}

	for rows.Next() {
		var missionDescription, status, avengers string
		err = rows.Scan(&missionDescription, &status, &avengers)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Mission Details:", missionDescription)
		fmt.Println("Mission Status:", status)
		fmt.Println("Avengers:", avengers)
	}
	defer rows.Close()

}

func missionCreate() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Avengers:")
	avengerName, _ := reader.ReadString('\n')

	fmt.Print("Enter Mission Name: ")
	missionName, _ := reader.ReadString('\n')

	fmt.Print("Enter Mission Details: ")
	missionDescription, _ := reader.ReadString('\n')

	_, err := dbShield.Exec("INSERT INTO mission_details (name, description) VALUES (?, ?)", strings.TrimSpace(missionName), strings.TrimSpace(missionDescription))
	if err != nil {
		log.Fatal(err)
	}

	avengerList := strings.Split(avengerName, ", ")

	for _, avenger := range avengerList {
		var avengerId, missionId int
		fmt.Println(avenger)
		rowA, err := dbShield.Query("SELECT id FROM avengers WHERE name=?", strings.TrimSpace(avenger))
		if err != nil {
			log.Fatal(err)
		}

		for rowA.Next() {
			err = rowA.Scan(&avengerId)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(avengerId)
		defer rowA.Close()
		rowM, err := dbShield.Query("SELECT id FROM mission_details WHERE name=?", missionName)
		if err != nil {
			log.Fatal(err)
		}

		for rowM.Next() {
			err = rowM.Scan(&missionId)
			if err != nil {
				log.Fatal(err)
			}
		}
		defer rowM.Close()

		_, err = dbShield.Exec("INSERT INTO `avenger_mission_mapping`(`avenger_id`, `mission_id`) VALUES (?,?)", avengerId, missionId)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func missionListing() {
	rows, err := dbShield.Query("SELECT md.name, md.status, GROUP_CONCAT(av.name) FROM mission_details md JOIN avenger_mission_mapping amm ON md.id = amm.mission_id JOIN avengers av ON av.id = amm.avenger_id GROUP BY md.id")
	if err != nil {
		log.Fatal(err)
	}
	if !rows.Next() {
		fmt.Println("No Mission has been assigned to an Avenger.")
	}

	defer rows.Close()
	fmt.Println("Mission Name\tStatus\tAvengers")
	for rows.Next() {
		var missionName, status, avengers string
		err = rows.Scan(&missionName, &status, &avengers)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(missionName, status, avengers)
	}

}
