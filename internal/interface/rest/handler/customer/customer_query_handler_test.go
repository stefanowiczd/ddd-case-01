//go:build unit

package customer

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/stefanowiczd/ddd-case-01/internal/application/customer"
	"github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/customer/mock"
)

func TestCustomerQueryHandler_GetCustomer(t *testing.T) {
	type testCaseParams struct {
		customerID               string
		mockCustomerQueryService func(*gomock.Controller) *mock.MockCustomerQueryService
	}

	type testCaseExpected struct {
		statusCode int
		wantError  bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	tests := []testCase{
		{
			name: "unsuccessful account retrieval",
			params: testCaseParams{
				customerID: "00000000-0000-0000-0000-000000000000",
				mockCustomerQueryService: func(m *gomock.Controller) *mock.MockCustomerQueryService {
					mock := mock.NewMockCustomerQueryService(m)
					mock.EXPECT().
						GetCustomer(gomock.Any(), gomock.Any()).
						Return(customer.GetCustomerResponseDTO{}, errors.New("error"))
					return mock
				},
			},
			expected: testCaseExpected{
				wantError:  true,
				statusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "unsuccessful account retrieval",
			params: testCaseParams{
				customerID: "00000000-0000-0000-0000-000000000000",
				mockCustomerQueryService: func(m *gomock.Controller) *mock.MockCustomerQueryService {
					mock := mock.NewMockCustomerQueryService(m)
					mock.EXPECT().
						GetCustomer(gomock.Any(), gomock.Any()).
						Return(customer.GetCustomerResponseDTO{
							Customer: customer.CustomerResponseDTO{
								ID:        "00000000-0000-0000-0000-000000000000",
								FirstName: "John",
								LastName:  "Doe",
								Email:     "john.doe@example.com",
								Phone:     "1234567890",
								Address: customer.Address{
									Street:     "street 1",
									City:       "Warsaw",
									State:      "Masovian",
									PostalCode: "00-000",
									Country:    "Poland",
								},
								Status:    "active",
								CreatedAt: time.Now().UTC(),
								UpdatedAt: time.Now().UTC(),
							},
						}, nil)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError:  false,
				statusCode: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := NewCustomerQueryHandler(tt.params.mockCustomerQueryService(ctrl))

			req := httptest.NewRequest(http.MethodGet, "/customer/{customerId}", nil)
			req.SetPathValue("customerId", tt.params.customerID)

			w := httptest.NewRecorder()

			handler.GetCustomer(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)

				// TODO: add further validation of response body
			}
		})
	}
}
