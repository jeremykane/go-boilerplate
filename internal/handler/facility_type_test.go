package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jeremykane/go-boilerplate/internal/entity"
	"github.com/jeremykane/go-boilerplate/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestFacilityTypeHandler_GetAll(t *testing.T) {
	tests := []struct {
		name                string
		mockFacilityTypeSvc func(ctrl *gomock.Controller) service.FacilityTypeService
		expectedHTTPStatus  int
		expectedResult      string
	}{
		{
			name: "success case should return 200 and facility type",
			mockFacilityTypeSvc: func(ctrl *gomock.Controller) service.FacilityTypeService {
				mock := service.NewMockFacilityTypeService(ctrl)
				mock.EXPECT().GetAll(gomock.Any()).Return([]entity.FacilityType{}, nil)
				return mock
			},
			expectedHTTPStatus: http.StatusOK,
			expectedResult: func() string {
				res := map[string]interface{}{
					"data": []entity.FacilityType{},
				}
				jsonData, _ := json.Marshal(res)

				return fmt.Sprintf("%s\n", string(jsonData))
			}(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/facility-types", nil)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := &FacilityTypeHandler{
				facilityTypeSvc: tc.mockFacilityTypeSvc(ctrl),
			}

			err := h.GetAll(c)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedHTTPStatus, rec.Code)
			assert.Equal(t, tc.expectedResult, rec.Body.String())
		})
	}
}
