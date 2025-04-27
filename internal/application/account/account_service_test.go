//go:build unit

package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/stefanowiczd/ddd-case-01/internal/application/account/mock"
	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
)

func TestAccountService_CreateAccount(t *testing.T) {

	type testCaseParams struct {
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
				mockAccountEventRepo: func(m *gomock.Controller) *mock.MockAccountEventRepository {
					mock := mock.NewMockAccountEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil) // TODO: add more specific expectations for the events.
					return mock
				},
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
			account, err := service.CreateAccount(context.Background(), "00000000-0000-0000-0000-000000000000", 0, "USD")

			if tt.expected.wantError {
				require.Error(t, err)
				require.Equal(t, tt.expected.err, err)
				require.Nil(t, account)
			} else {
				require.NoError(t, err)
				require.NotNil(t, account)
			}
		})
	}
}

func TestAccountService_GetAccount(t *testing.T) {

	type testCaseParams struct {
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
			account, err := service.GetAccount(context.Background(), "00000000-0000-0000-0000-000000000000")

			if tt.expected.wantError {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.expected.err)
				require.Nil(t, account)
			} else {
				require.NoError(t, err)
				require.NotNil(t, account)
			}
		})
	}
}

func TestAccountService_Deposit(t *testing.T) {

	type testCaseParams struct {
		mockAccountQueryRepo func(*gomock.Controller) *mock.MockAccountQueryRepository
		mockAccountEventRepo func(*gomock.Controller) *mock.MockAccountEventRepository
		amount               float64
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
				amount: -100,
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
				amount: 100,
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
				amount: 100,
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
			err := service.Deposit(context.Background(), "00000000-0000-0000-0000-000000000000", tt.params.amount)

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
		mockAccountQueryRepo func(*gomock.Controller) *mock.MockAccountQueryRepository
		mockAccountEventRepo func(*gomock.Controller) *mock.MockAccountEventRepository
		amount               float64
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
				amount: -100,
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
				amount: 100,
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
				amount: 100,
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
			err := service.Withdraw(context.Background(), "00000000-0000-0000-0000-000000000000", tt.params.amount)

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
			err := service.BlockAccount(context.Background(), "00000000-0000-0000-0000-000000000000")

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
			err := service.UnblockAccount(context.Background(), "00000000-0000-0000-0000-000000000000")

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
			accounts, err := service.GetCustomerAccounts(context.Background(), "00000000-0000-0000-0000-000000000000")

			if tt.expected.wantError {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.expected.err)
				require.Nil(t, accounts)
			} else {
				require.NoError(t, err)
				require.NotNil(t, accounts)
			}
		})
	}
}
