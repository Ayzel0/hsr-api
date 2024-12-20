package main

import (
	"fmt"
	"hsr-api/client"
	"hsr-api/enums"
	"hsr-api/services"
)

func main() {
	// create a new client with given language
	language := enums.English
	apiClient := client.NewClient(language)
	characterService := services.NewCharacterService(apiClient)

	// make a request for rappa
	char, err := characterService.GetCharacter(1317)
	if err != nil {
		fmt.Printf("main(): Error getting character with id %d: %v", 1317, err)
	}

	// print out the name of the character
	fmt.Printf("Name of character with id %d: %s\n", 1317, char.Name)
}
