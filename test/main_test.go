package test

import (
	"bytes"
	"encoding/json"
	"govtech-onecv/internal/controller"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	baseURL = "http://localhost:8080/api"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestHealth(t *testing.T) {
	resp, err := http.Get(baseURL + "/health")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// register
func TestRegister(t *testing.T) {

	request1 := controller.Request{
		Teacher:  "teacherken@gmail.com",
		Students: []string{"studentmary@gmail.com", "studentjon@gmail.com"},
	}

	request2 := controller.Request{
		Teacher:  "teacherjoe@gmail.com",
		Students: []string{"studentjon@gmail.com", "studentbob@gmail.com", "studentmiche@gmail.com"},
	}

	// Marshal the request bodies
	requestBody1, err := json.Marshal(request1)
	assert.Nil(t, err)

	requestBody2, err := json.Marshal(request2)
	assert.Nil(t, err)

	// Send POST requests
	resp1, err := http.Post(baseURL+"/register", "application/json", bytes.NewBuffer(requestBody1))
	assert.Nil(t, err)

	resp2, err := http.Post(baseURL+"/register", "application/json", bytes.NewBuffer(requestBody2))
	assert.Nil(t, err)

	// Check response status codes
	assert.Equal(t, http.StatusNoContent, resp1.StatusCode)
	assert.Equal(t, http.StatusNoContent, resp2.StatusCode)
}

// commonstudents
func TestCommonStudents(t *testing.T) {
	request3 := controller.Request{
		Student: "teacherjoe@gmail.com",
	}

	// Marshal the request bodies
	requestBody3, err := json.Marshal(request3)
	assert.Nil(t, err)

	// Send POST requests
	resp3, err := http.Post(baseURL+"/commonstudents?teacher=teacherken@gmail.com&teacher=teacherjoe@gmail.com", "application/json", bytes.NewBuffer(requestBody3))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp3.StatusCode)

	body, err := ioutil.ReadAll(resp3.Body)
	assert.Nil(t, err)

	var response controller.Response
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)
	assert.Equal(t, controller.Response{Students: []string{"studentjon@gmail.com"}}, response)
}

// retrievefornotifications
func TestRetrieveForNotifications1(t *testing.T) {
	request4 := controller.Request{
		Teacher:      "teacherken@gmail.com",
		Notification: "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com",
	}

	// Marshal the request bodies
	requestBody4, err := json.Marshal(request4)
	assert.Nil(t, err)

	// Send POST requests
	resp4, err := http.Post(baseURL+"/retrievefornotifications", "application/json", bytes.NewBuffer(requestBody4))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp4.StatusCode)

	body, err := ioutil.ReadAll(resp4.Body)
	assert.Nil(t, err)

	var response controller.Response
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)
	assert.Equal(t, controller.Response{Recipients: []string{"studentmary@gmail.com", "studentjon@gmail.com", "studentmiche@gmail.com"}}, response)
}

func TestRetrieveForNotifications2(t *testing.T) {
	request4 := controller.Request{
		Teacher:      "teacherken@gmail.com",
		Notification: "Hello students!",
	}

	// Marshal the request bodies
	requestBody4, err := json.Marshal(request4)
	assert.Nil(t, err)

	// Send POST requests
	resp4, err := http.Post(baseURL+"/retrievefornotifications", "application/json", bytes.NewBuffer(requestBody4))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp4.StatusCode)

	body, err := ioutil.ReadAll(resp4.Body)
	assert.Nil(t, err)

	var response controller.Response
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)
	assert.Equal(t, controller.Response{Recipients: []string{"studentmary@gmail.com", "studentjon@gmail.com"}}, response)
}

// suspend
func TestSuspension(t *testing.T) {
	request5 := controller.Request{
		Student: "studentmary@gmail.com",
	}

	// Marshal the request bodies
	requestBody5, err := json.Marshal(request5)
	assert.Nil(t, err)

	// Send POST requests
	resp5, err := http.Post(baseURL+"/suspend", "application/json", bytes.NewBuffer(requestBody5))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, resp5.StatusCode)
}

// retrievefornotifications
func TestRetrieveForNotifications3(t *testing.T) {
	request4 := controller.Request{
		Teacher:      "teacherken@gmail.com",
		Notification: "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com",
	}

	// Marshal the request bodies
	requestBody4, err := json.Marshal(request4)
	assert.Nil(t, err)

	// Send POST requests
	resp4, err := http.Post(baseURL+"/retrievefornotifications", "application/json", bytes.NewBuffer(requestBody4))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp4.StatusCode)

	body, err := ioutil.ReadAll(resp4.Body)
	assert.Nil(t, err)

	var response controller.Response
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)
	assert.Equal(t, controller.Response{Recipients: []string{"studentjon@gmail.com", "studentmiche@gmail.com"}}, response)
}

func TestRetrieveForNotifications4(t *testing.T) {
	request4 := controller.Request{
		Teacher:      "teacherken@gmail.com",
		Notification: "Hello students!",
	}

	// Marshal the request bodies
	requestBody4, err := json.Marshal(request4)
	assert.Nil(t, err)

	// Send POST requests
	resp4, err := http.Post(baseURL+"/retrievefornotifications", "application/json", bytes.NewBuffer(requestBody4))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp4.StatusCode)

	body, err := ioutil.ReadAll(resp4.Body)
	assert.Nil(t, err)

	var response controller.Response
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)
	assert.Equal(t, controller.Response{Recipients: []string{"studentjon@gmail.com"}}, response)
}
