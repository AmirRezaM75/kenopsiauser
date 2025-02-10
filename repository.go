package kenopsiauser

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type UserRepository struct {
	baseUrl string
	token   string
}

func NewUserRepository(baseUrl, token string) UserRepository {
	return UserRepository{
		baseUrl: baseUrl,
		token:   token,
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

type UserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Verified bool   `json:"verified"`
	AvatarId uint8  `json:"avatarId"`
}

func (userRepository UserRepository) GetByIds(userIds []string) ([]UserResponse, error) {
	var users = make([]UserResponse, 0, len(userIds))

	if len(userIds) == 0 {
		return users, nil
	}

	query := "ids[]=" + strings.Join(userIds, "&ids[]=")

	url := fmt.Sprintf("%s/users?%s", userRepository.baseUrl, query)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return users, fmt.Errorf("could not create new reques: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Token", userRepository.token)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return users, fmt.Errorf("could not send request: %w", err)
	}

	if response.StatusCode != 200 {
		return users, fmt.Errorf("response status %d is not ok: %w", response.StatusCode, err)
	}

	err = json.NewDecoder(response.Body).Decode(&users)

	_ = response.Body.Close()

	if err != nil {
		return users, fmt.Errorf("could not decode response body: %w", err)
	}

	return users, nil
}
