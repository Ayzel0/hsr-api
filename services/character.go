package services

import (
	"encoding/json"
	"fmt"
	"hsr-api/client"
	"hsr-api/models"
)

type CharacterService struct {
	client *client.Client
}

// constructor
func NewCharacterService(c *client.Client) *CharacterService {
	return &CharacterService{
		client: c,
	}
}

func (service *CharacterService) GetCharacter(id int) (*models.Character, error) {
	// make the request
	endpointSuffix := fmt.Sprintf("/avatar/%d", id)
	response, err := service.client.MakeRequest(endpointSuffix)

	// check for error
	if err != nil {
		return nil, fmt.Errorf("GetCharacter(): error making request: %w", err)
	}

	// parse the json response
	var apiResponse struct {
		ResponseCode int                  `json:"response"`
		Data         models.CharacterJSON `json:"data"`
	}

	if err := json.Unmarshal([]byte(response), &apiResponse); err != nil {
		return nil, fmt.Errorf("GetCharacter(): error parsing json: %w", err)
	}

	// now that the json is parsed, create a character model and return it
	character := apiResponse.Data.ToCharacter()
	return &character, nil
}
