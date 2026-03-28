package handler

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// patient service
type mockPatientService struct {
	SearchFunc func(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error)
}

func (m *mockPatientService) Search(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error) {
	if m.SearchFunc != nil {
		return m.SearchFunc(hospitalID, req)
	}
	return nil, nil
}

// helper
func setStaffTestContext(c *gin.Context, staff model.Staff) {
	c.Set("staff", &staff) // same key in middleware.staffContext
}

func TestPatientHandler_PatientSearch_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSrv := &mockPatientService{
		SearchFunc: func(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error) {
			return []model.Patient{
				{
					PatientHN:   "HN1",
					FirstNameTH: "สมชาย",
					LastNameTH:  "ใจดี",
					FirstNameEN: "Somchai",
					LastNameEN:  "Jaidee",
					HospitalID:  hospitalID,
				},
			}, nil
		},
	}

	handler := NewPatientHandler(mockSrv)

	body := `{"first_name":"สมชาย"}`
	req := httptest.NewRequest(http.MethodPost, "/patients/search", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// set staff context
	staff := model.Staff{HospitalID: "HOSP1", Username: "staff1"}
	setStaffTestContext(c, staff)

	handler.PatientSearch(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "HN1")
	assert.Contains(t, w.Body.String(), "Somchai")
}

func TestPatientHandler_PatientSearch_NoStaffContext(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSrv := &mockPatientService{}

	handler := NewPatientHandler(mockSrv)

	body := `{}`
	req := httptest.NewRequest(http.MethodPost, "/patients/search", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// no staff context set
	handler.PatientSearch(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestPatientHandler_PatientSearch_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSrv := &mockPatientService{
		SearchFunc: func(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error) {
			return nil, assert.AnError
		},
	}

	handler := NewPatientHandler(mockSrv)

	body := `{}`
	req := httptest.NewRequest(http.MethodPost, "/patients/search", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// set staff context
	staff := model.Staff{HospitalID: "HOSP1", Username: "staff1"}
	setStaffTestContext(c, staff)

	handler.PatientSearch(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), assert.AnError.Error())
}
