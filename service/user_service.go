package service

import (
	"api-gateway/config"
	"io/ioutil"
	"net/http"
	"strings"

	"encoding/json"

	"github.com/sirupsen/logrus"
)

// MyService struct represents a service
type UserService struct {
	Config *config.Config
}

// NewMyService creates a new MyService instance
func NewUserService(cfg *config.Config) *UserService {
	return &UserService{Config: cfg}
}

func (s *UserService) CreateUser(data string) (string, error) {

	// url
	partnerURL := s.Config.TransactionBaseURL
	endpoint := "/create-user"
	url := partnerURL + endpoint
	payload := strings.NewReader("")

	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		logrus.Errorf("Error creating HTTP request: %v", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error making HTTP POST request: %v", err)
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("Error reading response body: %v", err)
		return "", err
	}

	// Unmarshal the response body to a map for better JSON formatting
	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		logrus.Errorf("Error unmarshalling response body: %v", err)
		return "", err
	}

	formattedJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		logrus.Errorf("Error formatting response body as JSON: %v", err)
		return "", err
	}

	logrus.Infof("HTTP POST Request to %s with data: %s", url, data)
	logrus.Infof("HTTP Response Status: %s", response.Status)
	logrus.Infof("HTTP Response Body (JSON):\n%s", formattedJSON)

	return string(formattedJSON), nil
}

func (s *UserService) DeleteUser(data string) (string, error) {

	// url
	partnerURL := s.Config.TransactionBaseURL
	endpoint := "/delete-user"
	url := partnerURL + endpoint
	payload := strings.NewReader("")

	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		logrus.Errorf("Error creating HTTP request: %v", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error making HTTP POST request: %v", err)
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("Error reading response body: %v", err)
		return "", err
	}

	// Unmarshal the response body to a map for better JSON formatting
	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		logrus.Errorf("Error unmarshalling response body: %v", err)
		return "", err
	}

	formattedJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		logrus.Errorf("Error formatting response body as JSON: %v", err)
		return "", err
	}

	logrus.Infof("HTTP POST Request to %s with data: %s", url, data)
	logrus.Infof("HTTP Response Status: %s", response.Status)
	logrus.Infof("HTTP Response Body (JSON):\n%s", formattedJSON)

	return string(formattedJSON), nil
}

func (s *UserService) Login(data string) (string, error) {

	// url
	partnerURL := s.Config.TransactionBaseURL
	endpoint := "/login"
	url := partnerURL + endpoint
	payload := strings.NewReader("")

	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		logrus.Errorf("Error creating HTTP request: %v", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error making HTTP POST request: %v", err)
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("Error reading response body: %v", err)
		return "", err
	}

	// Unmarshal the response body to a map for better JSON formatting
	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		logrus.Errorf("Error unmarshalling response body: %v", err)
		return "", err
	}

	formattedJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		logrus.Errorf("Error formatting response body as JSON: %v", err)
		return "", err
	}

	logrus.Infof("HTTP POST Request to %s with data: %s", url, data)
	logrus.Infof("HTTP Response Status: %s", response.Status)
	logrus.Infof("HTTP Response Body (JSON):\n%s", formattedJSON)

	return string(formattedJSON), nil
}

func (s *UserService) DetailUser(data string) (string, error) {

	// url
	partnerURL := s.Config.TransactionBaseURL
	endpoint := "/detail-user"
	url := partnerURL + endpoint
	payload := strings.NewReader("")

	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		logrus.Errorf("Error creating HTTP request: %v", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error making HTTP POST request: %v", err)
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("Error reading response body: %v", err)
		return "", err
	}

	// Unmarshal the response body to a map for better JSON formatting
	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		logrus.Errorf("Error unmarshalling response body: %v", err)
		return "", err
	}

	formattedJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		logrus.Errorf("Error formatting response body as JSON: %v", err)
		return "", err
	}

	logrus.Infof("HTTP POST Request to %s with data: %s", url, data)
	logrus.Infof("HTTP Response Status: %s", response.Status)
	logrus.Infof("HTTP Response Body (JSON):\n%s", formattedJSON)

	return string(formattedJSON), nil
}

func (s *UserService) UpdateUser(data string) (string, error) {

	// url
	partnerURL := s.Config.TransactionBaseURL
	endpoint := "/update-user"
	url := partnerURL + endpoint
	payload := strings.NewReader("")

	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		logrus.Errorf("Error creating HTTP request: %v", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error making HTTP POST request: %v", err)
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("Error reading response body: %v", err)
		return "", err
	}

	// Unmarshal the response body to a map for better JSON formatting
	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		logrus.Errorf("Error unmarshalling response body: %v", err)
		return "", err
	}

	formattedJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		logrus.Errorf("Error formatting response body as JSON: %v", err)
		return "", err
	}

	logrus.Infof("HTTP POST Request to %s with data: %s", url, data)
	logrus.Infof("HTTP Response Status: %s", response.Status)
	logrus.Infof("HTTP Response Body (JSON):\n%s", formattedJSON)

	return string(formattedJSON), nil
}

func (s *UserService) ListUser(data string) (string, error) {

	// url
	partnerURL := s.Config.TransactionBaseURL
	endpoint := "/list-user"
	url := partnerURL + endpoint
	payload := strings.NewReader("")

	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		logrus.Errorf("Error creating HTTP request: %v", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error making HTTP POST request: %v", err)
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("Error reading response body: %v", err)
		return "", err
	}

	// Unmarshal the response body to a map for better JSON formatting
	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		logrus.Errorf("Error unmarshalling response body: %v", err)
		return "", err
	}

	formattedJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		logrus.Errorf("Error formatting response body as JSON: %v", err)
		return "", err
	}

	logrus.Infof("HTTP POST Request to %s with data: %s", url, data)
	logrus.Infof("HTTP Response Status: %s", response.Status)
	logrus.Infof("HTTP Response Body (JSON):\n%s", formattedJSON)

	return string(formattedJSON), nil
}
