// (c) 2019 Christian Bargmann
//
// This project serves for teaching purposes in the CloudWP with Stefan Sarstedt
// at the University of Applied Sciences in Hamburg. The project provides a basic framework for a Restful API,
// which can be used to manage courses of study via simple web calls.
//
package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetStudiengang(t *testing.T) {
	want := Studiengang{
		ID:           "1",
		Name:         "Angewandte Informatik",
		Beschreibung: "Programmieren, programmieren, programmieren...",
		Kontakt: &Professor{
			Vorname:  "Stefan",
			Nachname: "Sarstedt",
		},
	}

	var got Studiengang
	response, err := http.Get("http://localhost:8080/studiengaenge/1")
	if err != nil {
		t.Fatalf("failed to get json, %s", err)
	}

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	err = json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed to parse json, %s", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Test failed, want %+v, got %+v", want, got)
	}
}

func TestGetStudiengaenge(t *testing.T) {
	// not yet implemented
}

func TestCreateStudiengaeng(t *testing.T) {
	// not yet implemented
}

func TestUpdateStudiengaeng(t *testing.T) {
	// not yet implemented
}

func TestDeleteStudiengaeng(t *testing.T) {
	// not yet implemented
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
