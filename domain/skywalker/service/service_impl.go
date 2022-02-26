package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/martian/log"
	"rings/domain/skywalker/domain"
	"rings/http"
)

type ServiceImpl struct {
	httpClient http.Client
}

func NewServiceImpl(httpClient http.Client) *ServiceImpl {
	return &ServiceImpl{
		httpClient: httpClient,
	}
}

func (s *ServiceImpl) GetMessage() (*domain.Response, error) {
	searchEndpoints := fmt.Sprintf("https://swapi.dev/api/people?search=skywalker")
	people, err := s.getCharacters(searchEndpoints)
	if err != nil {
		return nil, err
	}

	finalResponse := &domain.Response{}
	for _, peopleC := range people.PeopleCharacter {

		character := domain.NewCharacter(peopleC.Name, []string{})

		for _, film := range peopleC.Films {
			rawFilm, err := s.httpClient.Get(film)
			if err != nil {
				log.Errorf("Error getting the response for film %s, err: %s", film, err.Error())
				return nil, err
			}

			filmStruct := &domain.Film{}
			if err := json.Unmarshal(rawFilm, filmStruct); err != nil {
				log.Errorf("Error unmarshalling the response, err: %s", err.Error())
				return nil, err
			}

			character.Films = append(character.Films, filmStruct.Title)
		}

		finalResponse.Characters = append(finalResponse.Characters, character)

	}
	return finalResponse, nil
}

func (s *ServiceImpl) getCharacters(searchEndpoints string) (*domain.People, error) {
	rawResponse, err := s.httpClient.Get(searchEndpoints)
	if err != nil {
		log.Errorf("Error getting the response for skywalker, err: %s", err.Error())
		return nil, err
	}

	response := &domain.People{}
	if err := json.Unmarshal(rawResponse, response); err != nil {
		log.Errorf("Error unmarshaling the response, err: %s", err.Error())
		return nil, err
	}
	return response, nil
}
