package client

import (
	"encoding/json"
	"hospital-api/internal/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHospitalClient_GetPatientByID_Success(t *testing.T) {
	// mock HIS server
	mockPatient := dto.HospitalClientPatientResponse{
		FirstNameTH: "สมชาย",
		LastNameTH:  "ใจดี",
		FirstNameEN: "Somchai",
		LastNameEN:  "Jaidee",
		PatientHN:   "HN001",
		NationalID:  "1234567890123",
		PassportID:  "A1234567",
		Gender:      "M",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockPatient)
	}))
	defer server.Close()

	client := NewHospitalClient(server.URL)

	resp, err := client.GetPatientByID("any-id")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "HN001", resp.PatientHN)
	assert.Equal(t, "Somchai", resp.FirstNameEN)
}

func TestHospitalClient_GetPatientByID_HIS_Error(t *testing.T) {
	// mock server returns 500
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewHospitalClient(server.URL)

	resp, err := client.GetPatientByID("any-id")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "HIS api error")
}

func TestHospitalClient_GetPatientByID_Invalid_JSON(t *testing.T) {
	// mock server returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{invalid-json}"))
	}))
	defer server.Close()

	client := NewHospitalClient(server.URL)

	resp, err := client.GetPatientByID("any-id")
	assert.Error(t, err)
	assert.Nil(t, resp)
}
