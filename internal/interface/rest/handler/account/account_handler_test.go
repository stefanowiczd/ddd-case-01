//go:build unit

package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/stefanowiczd/ddd-case-01/internal/application/account"
	"github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/account/mock"
)

func TestAccountHandler_CreateAccount(t *testing.T) {
	type testCaseParams struct {
		req                CreateAccountRequest
		reqBody            func(r CreateAccountRequest) io.Reader
		mockAccountService func(*gomock.Controller) *mock.MockAccountService
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
			name: "invalid request body",
			params: testCaseParams{
				req: CreateAccountRequest{
					CustomerID:     "customer123",
					InitialBalance: 100.0,
					Currency:       "USD",
				},
				reqBody: func(_ CreateAccountRequest) io.Reader {
					return bytes.NewBuffer([]byte(`{ ... invalid json ... `))
				},
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					return mock.NewMockAccountService(m)
				},
			},
			expected: testCaseExpected{
				wantError:  true,
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "unsuccessful create account",
			params: testCaseParams{
				req: CreateAccountRequest{
					CustomerID:     "customer123",
					InitialBalance: 100.0,
					Currency:       "USD",
				},
				reqBody: func(r CreateAccountRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						CreateAccount(gomock.Any(), gomock.Any()).
						Return(
							account.CreateAccountResponseDTO{},
							errors.New("error"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "successful account creation",
			params: testCaseParams{
				req: CreateAccountRequest{
					CustomerID:     "customer123",
					InitialBalance: 100.0,
					Currency:       "USD",
				},
				reqBody: func(r CreateAccountRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						CreateAccount(
							gomock.Any(),
							account.CreateAccountDTO{
								CustomerID:     "customer123",
								InitialBalance: 100.0,
								Currency:       "USD",
							}).
						Return(account.CreateAccountResponseDTO{
							AccountResponseDTO: account.AccountResponseDTO{
								ID:            "00000000-0000-0000-0000-000000000000",
								AccountNumber: "1234567890",
								CustomerID:    "customer123",
								Balance:       100.0,
								Currency:      "USD",
								Status:        "active",
							},
						}, nil)

					return mock
				},
			},
			expected: testCaseExpected{
				wantError:  false,
				statusCode: http.StatusCreated,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := NewHandler(tt.params.mockAccountService(ctrl))

			req := httptest.NewRequest(http.MethodPost, "/account", tt.params.reqBody(tt.params.req))
			w := httptest.NewRecorder()

			handler.CreateAccount(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)

				// TODO: add further validation of response body
			}
		})
	}
}

func TestAccountHandler_Deposit(t *testing.T) {
	type testCaseParams struct {
		accountID          string
		req                DepositRequest
		reqBody            func(r DepositRequest) io.Reader
		mockAccountService func(*gomock.Controller) *mock.MockAccountService
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
			name: "invalid request body",
			params: testCaseParams{
				accountID: "acc123",
				req: DepositRequest{
					Amount: 100.0,
				},
				reqBody: func(_ DepositRequest) io.Reader {
					return bytes.NewBuffer([]byte(`{ ... invalid json ... `))
				},
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					return mock.NewMockAccountService(m)
				},
			},
			expected: testCaseExpected{
				wantError:  true,
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "unsuccessful deposit",
			params: testCaseParams{
				accountID: "acc123",
				req: DepositRequest{
					Amount: 100.0,
				},
				reqBody: func(r DepositRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						Deposit(gomock.Any(), gomock.Any()).
						Return(errors.New("error"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "successful deposit",
			params: testCaseParams{
				accountID: "acc123",
				req: DepositRequest{
					Amount: 100.0,
				},
				reqBody: func(r DepositRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						Deposit(
							gomock.Any(),
							account.DepositDTO{
								AccountID: "acc123",
								Amount:    100.0,
							}).
						Return(nil)

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

			handler := NewHandler(tt.params.mockAccountService(ctrl))

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/accounts/%s/deposit", tt.params.accountID), tt.params.reqBody(tt.params.req))
			req = mux.SetURLVars(req, map[string]string{"id": tt.params.accountID})
			w := httptest.NewRecorder()

			handler.Deposit(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)
			}
		})
	}
}

func TestAccountHandler_Withdraw(t *testing.T) {
	type testCaseParams struct {
		accountID          string
		req                WithdrawRequest
		reqBody            func(r WithdrawRequest) io.Reader
		mockAccountService func(*gomock.Controller) *mock.MockAccountService
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
			name: "invalid request body",
			params: testCaseParams{
				accountID: "acc123",
				req: WithdrawRequest{
					Amount: 50.0,
				},
				reqBody: func(_ WithdrawRequest) io.Reader {
					return bytes.NewBuffer([]byte(`{ ... invalid json ... `))
				},
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					return mock.NewMockAccountService(m)
				},
			},
			expected: testCaseExpected{
				wantError:  true,
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "unsuccessful withdrawal",
			params: testCaseParams{
				accountID: "acc123",
				req: WithdrawRequest{
					Amount: 50.0,
				},
				reqBody: func(r WithdrawRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						Withdraw(gomock.Any(), gomock.Any()).
						Return(errors.New("error"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "successful withdrawal",
			params: testCaseParams{
				accountID: "acc123",
				req: WithdrawRequest{
					Amount: 50.0,
				},
				reqBody: func(r WithdrawRequest) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						Withdraw(
							gomock.Any(),
							account.WithdrawDTO{
								AccountID: "acc123",
								Amount:    50.0,
							}).
						Return(nil)

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

			handler := NewHandler(tt.params.mockAccountService(ctrl))

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/accounts/%s/withdraw", tt.params.accountID), tt.params.reqBody(tt.params.req))
			req = mux.SetURLVars(req, map[string]string{"id": tt.params.accountID})
			w := httptest.NewRecorder()

			handler.Withdraw(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)
			}
		})
	}
}

func TestAccountHandler_BlockAccount(t *testing.T) {
	type testCaseParams struct {
		accountID          string
		mockAccountService func(*gomock.Controller) *mock.MockAccountService
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
			name: "unsuccessful account blocking",
			params: testCaseParams{
				accountID: "acc123",
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						BlockAccount(gomock.Any(), gomock.Any()).
						Return(errors.New("error"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "successful account blocking",
			params: testCaseParams{
				accountID: "acc123",
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						BlockAccount(
							gomock.Any(),
							account.BlockAccountDTO{
								AccountID: "acc123",
							}).
						Return(nil)

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

			handler := NewHandler(tt.params.mockAccountService(ctrl))

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/accounts/%s/block", tt.params.accountID), nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.params.accountID})
			w := httptest.NewRecorder()

			handler.BlockAccount(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)
			}
		})
	}
}

func TestAccountHandler_UnblockAccount(t *testing.T) {
	type testCaseParams struct {
		accountID          string
		mockAccountService func(*gomock.Controller) *mock.MockAccountService
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
			name: "unsuccessful account unblocking",
			params: testCaseParams{
				accountID: "acc123",
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						UnblockAccount(gomock.Any(), gomock.Any()).
						Return(errors.New("error"))

					return mock
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusInternalServerError,
				wantError:  true,
			},
		},
		{
			name: "successful account unblocking",
			params: testCaseParams{
				accountID: "acc123",
				mockAccountService: func(m *gomock.Controller) *mock.MockAccountService {
					mock := mock.NewMockAccountService(m)
					mock.EXPECT().
						UnblockAccount(
							gomock.Any(),
							account.UnblockAccountDTO{
								AccountID: "acc123",
							}).
						Return(nil)

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

			handler := NewHandler(tt.params.mockAccountService(ctrl))

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/accounts/%s/unblock", tt.params.accountID), nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.params.accountID})
			w := httptest.NewRecorder()

			handler.UnblockAccount(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)
			}
		})
	}
}
