//go:build unit

package customer

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.uber.org/mock/gomock"

	customerapplication "github.com/stefanowiczd/ddd-case-01/internal/application/customer"
	"github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/customer/mock"
	"github.com/stretchr/testify/require"
)

func TestCustomerHandler_CreateCustomer(t *testing.T) {
	type testCaseParams struct {
		req                 CreateCustomerRequest
		reqBody             func(r CreateCustomerRequest) io.Reader
		mockCustomerService func(*gomock.Controller) *mock.MockCustomerService
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

	// TODO: add tests for payload validation
	tests := []testCase{
		{
			name: "should return 500 - customer creation failed",
			params: testCaseParams{
				req: CreateCustomerRequest{
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john.doe@example.com",
					Phone:     "1234567890",
					Address: Address{
						Street:     "Street 1",
						City:       "Warsaw",
						State:      "Masovian",
						PostalCode: "00-000",
						Country:    "Poland",
					},
				},
				reqBody: func(r CreateCustomerRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().CreateCustomer(gomock.Any(), gomock.Any()).
						Return(
							customerapplication.CreateCustomerResponseDTO{},
							errors.New("customer creation failed"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "should return 201 - customer created successfully",
			params: testCaseParams{
				req: CreateCustomerRequest{
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john.doe@example.com",
					Phone:     "1234567890",
					Address: Address{
						Street:     "Street 1",
						City:       "Warsaw",
						State:      "Masovian",
						PostalCode: "00-000",
						Country:    "Poland",
					},
				},
				reqBody: func(r CreateCustomerRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().CreateCustomer(gomock.Any(), gomock.Any()).
						Return(
							customerapplication.CreateCustomerResponseDTO{
								Customer: customerapplication.CustomerResponseDTO{
									ID:        "123",
									FirstName: "John",
									LastName:  "Doe",
									Email:     "john.doe@example.com",
									Phone:     "1234567890",
									Address: customerapplication.Address{
										Street:     "Street 1",
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
				statusCode: http.StatusCreated,
				wantError:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := NewCustomerHandler(tt.params.mockCustomerService(ctrl))

			req := httptest.NewRequest(http.MethodPost, "/customer", tt.params.reqBody(tt.params.req))
			w := httptest.NewRecorder()

			handler.CreateCustomer(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)

				// TODO: add further validation of response body
			}
		})
	}
}

func TestCustomerHandler_UpdateCustomer(t *testing.T) {
	type testCaseParams struct {
		req                 UpdateCustomerRequest
		reqBody             func(r UpdateCustomerRequest) io.Reader
		mockCustomerService func(*gomock.Controller) *mock.MockCustomerService
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

	// TODO: add tests for payload validation
	tests := []testCase{
		{
			name: "should return 500 - customer update failed",
			params: testCaseParams{
				req: UpdateCustomerRequest{
					CustomerID: "123",
					FirstName:  "John",
					LastName:   "Doe",
					Email:      "john.doe@example.com",
					Phone:      "1234567890",
					Address: Address{
						Street:     "Street 1",
						City:       "Warsaw",
						State:      "Masovian",
						PostalCode: "00-000",
						Country:    "Poland",
					},
				},
				reqBody: func(r UpdateCustomerRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().UpdateCustomer(gomock.Any(), gomock.Any()).
						Return(errors.New("customer update failed"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "should return 204 - customer updated successfully",
			params: testCaseParams{
				req: UpdateCustomerRequest{
					CustomerID: "123",
					FirstName:  "John",
					LastName:   "Doe",
					Email:      "john.doe@example.com",
					Phone:      "1234567890",
					Address: Address{
						Street:     "Street 1",
						City:       "Warsaw",
						State:      "Masovian",
						PostalCode: "00-000",
						Country:    "Poland",
					},
				},
				reqBody: func(r UpdateCustomerRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().UpdateCustomer(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusNoContent,
				wantError:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := NewCustomerHandler(tt.params.mockCustomerService(ctrl))

			req := httptest.NewRequest(http.MethodPut, "/customer/{customerId}", tt.params.reqBody(tt.params.req))
			req.SetPathValue("customerId", tt.params.req.CustomerID)

			w := httptest.NewRecorder()

			handler.UpdateCustomer(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)
			}
		})
	}
}

func TestCustomerHandler_BlockCustomer(t *testing.T) {
	type testCaseParams struct {
		req                 BlockCustomerRequest
		reqBody             func(r BlockCustomerRequest) io.Reader
		mockCustomerService func(*gomock.Controller) *mock.MockCustomerService
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

	// TODO: add tests for payload validation
	tests := []testCase{
		{
			name: "should return 500 - customer block failed",
			params: testCaseParams{
				req: BlockCustomerRequest{
					CustomerID: "00000000-0000-0000-0000-000000000000",
					Reason:     "Blocked by user",
				},
				reqBody: func(r BlockCustomerRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().BlockCustomer(gomock.Any(), gomock.Any()).
						Return(errors.New("customer blocking failed"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "should return 204 - customer block successfully",
			params: testCaseParams{
				req: BlockCustomerRequest{
					CustomerID: "00000000-0000-0000-0000-000000000000",
					Reason:     "Blocked by user",
				},
				reqBody: func(r BlockCustomerRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().BlockCustomer(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusNoContent,
				wantError:  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := NewCustomerHandler(tt.params.mockCustomerService(ctrl))

			req := httptest.NewRequest(http.MethodPost, "/customer/{customerId}/block", tt.params.reqBody(tt.params.req))
			req.SetPathValue("customerId", tt.params.req.CustomerID)

			w := httptest.NewRecorder()

			handler.BlockCustomer(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			}
		})
	}
}

func TestCustomerHandler_UnblockCustomer(t *testing.T) {
	type testCaseParams struct {
		customerID          string
		mockCustomerService func(*gomock.Controller) *mock.MockCustomerService
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

	// TODO: add tests for payload validation
	tests := []testCase{
		{
			name: "should return 500 - customer unblock failed",
			params: testCaseParams{
				customerID: "00000000-0000-0000-0000-000000000000",
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().UnblockCustomer(gomock.Any(), gomock.Any()).
						Return(errors.New("customer unblocking failed"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "should return 204 - customer unblocked successfully",
			params: testCaseParams{
				customerID: "00000000-0000-0000-0000-000000000000",
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().UnblockCustomer(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusNoContent,
				wantError:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := NewCustomerHandler(tt.params.mockCustomerService(ctrl))

			req := httptest.NewRequest(http.MethodPost, "/customer/{customerId}/unblock", nil)
			req.SetPathValue("customerId", tt.params.customerID)

			w := httptest.NewRecorder()

			handler.UnblockCustomer(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)
			}
		})
	}

}

func TestCustomerHandler_DeleteCustomer(t *testing.T) {
	type testCaseParams struct {
		customerID          string
		mockCustomerService func(*gomock.Controller) *mock.MockCustomerService
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

	// TODO: add tests for payload validation
	tests := []testCase{
		{
			name: "should return 500 - customer deletion failed",
			params: testCaseParams{
				customerID: "00000000-0000-0000-0000-000000000000",
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().DeleteCustomer(gomock.Any(), gomock.Any()).
						Return(errors.New("customer deletion failed"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "should return 204 - customer deleted successfully",
			params: testCaseParams{
				customerID: "00000000-0000-0000-0000-000000000000",
				mockCustomerService: func(c *gomock.Controller) *mock.MockCustomerService {
					mock := mock.NewMockCustomerService(c)
					mock.EXPECT().DeleteCustomer(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusNoContent,
				wantError:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := NewCustomerHandler(tt.params.mockCustomerService(ctrl))

			req := httptest.NewRequest(http.MethodDelete, "/customer/{customerId}", nil)
			req.SetPathValue("customerId", tt.params.customerID)

			w := httptest.NewRecorder()

			handler.DeleteCustomer(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)
			}
		})
	}
}
