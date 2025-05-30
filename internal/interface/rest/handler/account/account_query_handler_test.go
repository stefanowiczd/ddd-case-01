//go:build unit

package account

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/stefanowiczd/ddd-case-01/internal/application/account"
	"github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/account/mock"
)

func TestAccountHandler_GetAccount(t *testing.T) {

	type testCaseParams struct {
		req                     GetAccountRequest // TODO: marshal request
		mockAccountQueryService func(*gomock.Controller) *mock.MockAccountQueryService
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
			name: "invalid account id format in request path",
			params: testCaseParams{
				req: GetAccountRequest{
					AccountID: "0000",
				},
				mockAccountQueryService: func(m *gomock.Controller) *mock.MockAccountQueryService {
					return mock.NewMockAccountQueryService(m)
				},
			},
			expected: testCaseExpected{
				wantError:  true,
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "unsuccessful account retrieval",
			params: testCaseParams{
				req: GetAccountRequest{
					AccountID: "00000000-0000-0000-0000-000000000000",
				},
				mockAccountQueryService: func(m *gomock.Controller) *mock.MockAccountQueryService {
					mock := mock.NewMockAccountQueryService(m)
					mock.EXPECT().
						GetAccount(gomock.Any(), gomock.Any()).
						Return(account.AccountResponseDTO{}, errors.New("error"))
					return mock
				},
			},
			expected: testCaseExpected{
				wantError:  true,
				statusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "success account retrieval",
			params: testCaseParams{
				req: GetAccountRequest{
					AccountID: "00000000-0000-0000-0000-000000000000",
				},
				mockAccountQueryService: func(m *gomock.Controller) *mock.MockAccountQueryService {
					mock := mock.NewMockAccountQueryService(m)
					mock.EXPECT().
						GetAccount(gomock.Any(), gomock.Any()).
						Return(account.AccountResponseDTO{}, nil)

					return mock
				},
			},
			expected: testCaseExpected{
				wantError:  true,
				statusCode: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := NewAccountQueryHandler(tt.params.mockAccountQueryService(ctrl))

			req := httptest.NewRequest(http.MethodGet, "/account/{id}", nil)
			req.SetPathValue("id", tt.params.req.AccountID)

			w := httptest.NewRecorder()

			handler.GetAccount(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)

				// TODO: add further validation of response body
			}
		})
	}
}

func TestAccountHandler_GetCustomerAccounts(t *testing.T) {

	type testCaseParams struct {
		req                     GetCustomerAccountsRequest // TODO: marshal request
		mockAccountQueryService func(*gomock.Controller) *mock.MockAccountQueryService
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
			name: "invalid customer id format in request path",
			params: testCaseParams{
				req: GetCustomerAccountsRequest{
					CustomerID: "0000",
				},
				mockAccountQueryService: func(m *gomock.Controller) *mock.MockAccountQueryService {
					return mock.NewMockAccountQueryService(m)
				},
			},
			expected: testCaseExpected{
				wantError:  true,
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "unsuccessful customer accounts retrieval",
			params: testCaseParams{
				req: GetCustomerAccountsRequest{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockAccountQueryService: func(m *gomock.Controller) *mock.MockAccountQueryService {
					mock := mock.NewMockAccountQueryService(m)
					mock.EXPECT().
						GetCustomerAccounts(gomock.Any(), gomock.Any()).
						Return(account.GetCustomerAccountsResponseDTO{}, errors.New("error"))
					return mock
				},
			},
			expected: testCaseExpected{
				wantError:  true,
				statusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "success customer accounts retrieval",
			params: testCaseParams{
				req: GetCustomerAccountsRequest{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockAccountQueryService: func(m *gomock.Controller) *mock.MockAccountQueryService {
					mock := mock.NewMockAccountQueryService(m)
					mock.EXPECT().
						GetCustomerAccounts(gomock.Any(), gomock.Any()).
						Return(account.GetCustomerAccountsResponseDTO{}, nil)

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

			handler := NewAccountQueryHandler(tt.params.mockAccountQueryService(ctrl))

			req := httptest.NewRequest(http.MethodGet, "/customers/{customerId}/accounts", nil)
			req.SetPathValue("customerId", tt.params.req.CustomerID)

			w := httptest.NewRecorder()

			handler.GetCustomerAccounts(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)

				// TODO: add further validation of response body
			}
		})
	}

}
