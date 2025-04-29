//go:build unit

package account

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/stefanowiczd/ddd-case-01/internal/application/account/mock"
	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
)

func TestAccountService_CreateAccount(t *testing.T) {
	type testCaseParams struct {
		dto                  CreateAccountDTO
		mockAccountEventRepo func(*gomock.Controller) *mock.MockAccountEventRepository
	}

	type testCaseExpected struct {
		wantError bool
		err       error
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "should create account successfully",
			params: testCaseParams{
				dto: CreateAccountDTO{
					CustomerID:     "00000000-0000-0000-0000-000000000000",
					InitialBalance: 0,
					Currency:       "USD",
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					mock := mock.NewMockAccountEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
				err:       nil,
			},
		},
		{
			name: "should return error for negative balance",
			params: testCaseParams{
				dto: CreateAccountDTO{
					CustomerID:     "00000000-0000-0000-0000-000000000000",
					InitialBalance: -100,
					Currency:       "USD",
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					return mock.NewMockAccountEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError: true,
				err:       ErrInvalidAmount,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewService(
				mock.NewMockAccountQueryRepository(ctrl),
				mock.NewMockCustomerQueryRepository(ctrl),
				tt.params.mockAccountEventRepo(ctrl),
			)
			account, err := service.CreateAccount(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)
				require.Equal(t, tt.expected.err, err)
				require.Empty(t, account)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, account)
			}
		})
	}
}

func TestAccountService_GetAccount(t *testing.T) {
	type testCaseParams struct {
		dto                  GetAccountDTO
		mockAccountQueryRepo func(*gomock.Controller) *mock.MockAccountQueryRepository
	}

	type testCaseExpected struct {
		wantError bool
		err       error
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "shouldn't get account",
			params: testCaseParams{
				dto: GetAccountDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrAccountNotFound)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError: true,
				err:       ErrAccountNotFound,
			},
		},
		{
			name: "should get account successfully",
			params: testCaseParams{
				dto: GetAccountDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&Account{}, nil)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
				err:       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewService(
				tt.params.mockAccountQueryRepo(ctrl),
				mock.NewMockCustomerQueryRepository(ctrl),
				mock.NewMockAccountEventRepository(ctrl),
			)
			account, err := service.GetAccount(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.expected.err)
				require.Empty(t, account)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, account)
			}
		})
	}
}

func TestAccountService_Deposit(t *testing.T) {
	type testCaseParams struct {
		dto                  DepositDTO
		mockAccountQueryRepo func(*gomock.Controller) *mock.MockAccountQueryRepository
		mockAccountEventRepo func(*gomock.Controller) *mock.MockAccountEventRepository
	}

	type testCaseExpected struct {
		wantError bool
		err       error
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "shouldn't deposit - invalid amount",
			params: testCaseParams{
				dto: DepositDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					Amount:    -100,
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					return mock.NewMockAccountQueryRepository(m)
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					return mock.NewMockAccountEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError: true,
				err:       ErrInvalidAmount,
			},
		},
		{
			name: "shouldn't deposit - account not found",
			params: testCaseParams{
				dto: DepositDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					Amount:    100,
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrAccountNotFound)
					return mock
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					return mock.NewMockAccountEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError: true,
				err:       ErrAccountNotFound,
			},
		},
		{
			name: "should deposit successfully",
			params: testCaseParams{
				dto: DepositDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					Amount:    100,
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&Account{}, nil)
					return mock
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					mock := mock.NewMockAccountEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
				err:       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewService(
				tt.params.mockAccountQueryRepo(ctrl),
				mock.NewMockCustomerQueryRepository(ctrl),
				tt.params.mockAccountEventRepo(ctrl),
			)
			err := service.Deposit(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.expected.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountService_Withdraw(t *testing.T) {
	type testCaseParams struct {
		dto                  WithdrawDTO
		mockAccountQueryRepo func(*gomock.Controller) *mock.MockAccountQueryRepository
		mockAccountEventRepo func(*gomock.Controller) *mock.MockAccountEventRepository
	}

	type testCaseExpected struct {
		wantError bool
		err       error
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "shouldn't withdraw - invalid amount",
			params: testCaseParams{
				dto: WithdrawDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					Amount:    -100,
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					return mock.NewMockAccountQueryRepository(m)
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					return mock.NewMockAccountEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError: true,
				err:       ErrInvalidAmount,
			},
		},
		{
			name: "shouldn't withdraw - account not found",
			params: testCaseParams{
				dto: WithdrawDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					Amount:    100,
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrAccountNotFound)
					return mock
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					return mock.NewMockAccountEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError: true,
				err:       ErrAccountNotFound,
			},
		},
		{
			name: "should withdraw successfully",
			params: testCaseParams{
				dto: WithdrawDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					Amount:    100,
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&Account{}, nil)
					return mock
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					mock := mock.NewMockAccountEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
				err:       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewService(
				tt.params.mockAccountQueryRepo(ctrl),
				mock.NewMockCustomerQueryRepository(ctrl),
				tt.params.mockAccountEventRepo(ctrl),
			)
			err := service.Withdraw(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.expected.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountService_BlockAccount(t *testing.T) {
	type testCaseParams struct {
		dto                  BlockAccountDTO
		mockAccountQueryRepo func(*gomock.Controller) *mock.MockAccountQueryRepository
		mockAccountEventRepo func(*gomock.Controller) *mock.MockAccountEventRepository
	}

	type testCaseExpected struct {
		wantError bool
		err       error
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "shouldn't block account - account not found",
			params: testCaseParams{
				dto: BlockAccountDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrAccountNotFound)
					return mock
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					return mock.NewMockAccountEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError: true,
				err:       ErrAccountNotFound,
			},
		},
		{
			name: "should block account successfully",
			params: testCaseParams{
				dto: BlockAccountDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&Account{}, nil)
					return mock
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					mock := mock.NewMockAccountEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
				err:       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewService(
				tt.params.mockAccountQueryRepo(ctrl),
				mock.NewMockCustomerQueryRepository(ctrl),
				tt.params.mockAccountEventRepo(ctrl),
			)
			err := service.BlockAccount(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.expected.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountService_UnblockAccount(t *testing.T) {
	type testCaseParams struct {
		dto                  UnblockAccountDTO
		mockAccountQueryRepo func(*gomock.Controller) *mock.MockAccountQueryRepository
		mockAccountEventRepo func(*gomock.Controller) *mock.MockAccountEventRepository
	}

	type testCaseExpected struct {
		wantError bool
		err       error
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "shouldn't unblock account - account not found",
			params: testCaseParams{
				dto: UnblockAccountDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrAccountNotFound)
					return mock
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					return mock.NewMockAccountEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError: true,
				err:       ErrAccountNotFound,
			},
		},
		{
			name: "should unblock account successfully",
			params: testCaseParams{
				dto: UnblockAccountDTO{
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&Account{}, nil)
					return mock
				},
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					mock := mock.NewMockAccountEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
				err:       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewService(
				tt.params.mockAccountQueryRepo(ctrl),
				mock.NewMockCustomerQueryRepository(ctrl),
				tt.params.mockAccountEventRepo(ctrl),
			)
			err := service.UnblockAccount(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.expected.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountService_GetCustomerAccounts(t *testing.T) {
	type testCaseParams struct {
		dto                   GetCustomerAccountsDTO
		mockAccountQueryRepo  func(*gomock.Controller) *mock.MockAccountQueryRepository
		mockCustomerQueryRepo func(*gomock.Controller) *mock.MockCustomerQueryRepository
	}

	type testCaseExpected struct {
		wantError bool
		err       error
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "shouldn't get customer accounts - customer not found",
			params: testCaseParams{
				dto: GetCustomerAccountsDTO{
					CustomerID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					return mock.NewMockAccountQueryRepository(m)
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrCustomerNotFound)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError: true,
				err:       ErrCustomerNotFound,
			},
		},
		{
			name: "should get customer accounts successfully",
			params: testCaseParams{
				dto: GetCustomerAccountsDTO{
					CustomerID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				},
				mockAccountQueryRepo: func(m *gomock.Controller) *mock.MockAccountQueryRepository {
					mock := mock.NewMockAccountQueryRepository(m)
					mock.EXPECT().FindByCustomerID(gomock.Any(), gomock.Any()).Return([]*Account{}, nil)
					return mock
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)
					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
				err:       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewService(
				tt.params.mockAccountQueryRepo(ctrl),
				tt.params.mockCustomerQueryRepo(ctrl),
				mock.NewMockAccountEventRepository(ctrl),
			)
			accounts, err := service.GetCustomerAccounts(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.expected.err)
				require.Empty(t, accounts)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, accounts)
			}
		})
	}
}
