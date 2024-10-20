package kenopsiauser

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type UserRepository struct {
	baseUrl string
}

func NewUserRepository(baseUrl string) UserRepository {
	return UserRepository{
		baseUrl: baseUrl,
	}
}

func (userRepository UserRepository) AcquireUserId(ticketId string) (userId string, err error) {
	url := fmt.Sprintf("%s/tickets/acquire", userRepository.baseUrl)
	payload := []byte(fmt.Sprintf(`{"ticketId": "%s"}`, ticketId))
	reader := bytes.NewReader(payload)

	response, err := http.Post(url, "application/json", reader)

	defer func(Body io.ReadCloser) {
		if e := Body.Close(); e != nil {
			err = e
		}
	}(response.Body)

	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return "", errors.New(string(body))
	}

	var output struct {
		UserId string `json:"userId"`
	}

	err = json.NewDecoder(response.Body).Decode(&output)

	if err != nil {
		return "", err
	}

	return output.UserId, nil
}
