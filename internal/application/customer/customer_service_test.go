//go:build unit

package customer

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/stefanowiczd/ddd-case-01/internal/application/customer/mock"
	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
)

func TestCustomerService_CreateCustomer(t *testing.T) {
	type testCaseParams struct {
		dto CreateCustomerDTO

		mockCustomerQueryRepo func(*gomock.Controller) *mock.MockCustomerQueryRepository
		mockCustomerEventRepo func(*gomock.Controller) *mock.MockCustomerEventRepository
	}

	type testCaseExpected struct {
		wantError      bool
		errWantCompare bool
		err            error
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	tests := []testCase{
		{
			name: "shouldn't create customer - customer already exists",
			params: testCaseParams{
				dto: CreateCustomerDTO{
					FirstName:   "John",
					LastName:    "Doe",
					Email:       "john.doe@example.com",
					Phone:       "1234567890",
					DateOfBirth: "1900-01-01",
					Address:     Address{Street: "Street 1", City: "Warsaw", State: "Masovian", PostalCode: "00-000", Country: "Poland"},
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrCustomerAlreadyExists)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					return mock.NewMockCustomerEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: true,
				err:            ErrCustomerAlreadyExists,
			},
		},
		{
			name: "shouldn't create customer - customer query repository error",
			params: testCaseParams{
				dto: CreateCustomerDTO{
					FirstName:   "John",
					LastName:    "Doe",
					Email:       "john.doe@example.com",
					Phone:       "1234567890",
					DateOfBirth: "1900-01-01",
					Address:     Address{Street: "Street 1", City: "Warsaw", State: "Masovian", PostalCode: "00-000", Country: "Poland"},
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					return mock.NewMockCustomerEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: false,
			},
		},
		{
			name: "shouldn't create customer - customer create event repository error",
			params: testCaseParams{
				dto: CreateCustomerDTO{
					FirstName:   "John",
					LastName:    "Doe",
					Email:       "john.doe@example.com",
					Phone:       "1234567890",
					DateOfBirth: "1900-01-01",
					Address:     Address{Street: "Street 1", City: "Warsaw", State: "Masovian", PostalCode: "00-000", Country: "Poland"},
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(errors.New("some error"))

					return mock
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: false,
			},
		},
		{
			name: "should create customer",
			params: testCaseParams{
				dto: CreateCustomerDTO{
					FirstName:   "John",
					LastName:    "Doe",
					Email:       "john.doe@example.com",
					Phone:       "1234567890",
					DateOfBirth: "1900-01-01",
					Address:     Address{Street: "Street 1", City: "Warsaw", State: "Masovian", PostalCode: "00-000", Country: "Poland"},
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)

					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewCustomerService(
				tt.params.mockCustomerQueryRepo(ctrl),
				tt.params.mockCustomerEventRepo(ctrl),
			)

			customer, err := service.CreateCustomer(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)

				if tt.expected.errWantCompare {
					require.ErrorIs(t, err, tt.expected.err)
				}
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, customer)
			}
		})
	}
}

func Test_CustomerService_GetCustomer(t *testing.T) {

	type testCaseParams struct {
		dto                   GetCustomerDTO
		mockCustomerQueryRepo func(*gomock.Controller) *mock.MockCustomerQueryRepository
	}

	type testCaseExpected struct {
		wantError      bool
		errWantCompare bool
		err            error
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	tests := []testCase{
		{
			name: "shouldn't get customer - customer not found",
			params: testCaseParams{
				dto: GetCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrCustomerNotFound)

					return mock
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: true,
				err:            ErrCustomerNotFound,
			},
		},
		{
			name: "should get customer",
			params: testCaseParams{
				dto: GetCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewCustomerService(
				tt.params.mockCustomerQueryRepo(ctrl),
				mock.NewMockCustomerEventRepository(ctrl),
			)

			customer, err := service.GetCustomer(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)

				if tt.expected.errWantCompare {
					require.ErrorIs(t, err, tt.expected.err)
				}
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, customer)
			}
		})
	}
}

func Test_CustomerService_UpdateCustomer(t *testing.T) {

	type testCaseParams struct {
		dto UpdateCustomerDTO

		mockCustomerQueryRepo func(*gomock.Controller) *mock.MockCustomerQueryRepository
		mockCustomerEventRepo func(*gomock.Controller) *mock.MockCustomerEventRepository
	}

	type testCaseExpected struct {
		wantError      bool
		errWantCompare bool
		err            error
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	tests := []testCase{
		{
			name: "shouldn't update customer - customer not found",
			params: testCaseParams{
				dto: UpdateCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrCustomerNotFound)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					return mock.NewMockCustomerEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: true,
				err:            ErrCustomerNotFound,
			},
		},
		{
			name: "shouldn't update customer - customer event repository error",
			params: testCaseParams{
				dto: UpdateCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(errors.New("some error"))

					return mock
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: false,
			},
		},
		{
			name: "should update customer",
			params: testCaseParams{
				dto: UpdateCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)

					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewCustomerService(
				tt.params.mockCustomerQueryRepo(ctrl),
				tt.params.mockCustomerEventRepo(ctrl),
			)

			err := service.UpdateCustomer(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)

				if tt.expected.errWantCompare {
					require.ErrorIs(t, err, tt.expected.err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_CustomerService_BlockCustomer(t *testing.T) {

	type testCaseParams struct {
		dto BlockCustomerDTO

		mockCustomerQueryRepo func(*gomock.Controller) *mock.MockCustomerQueryRepository
		mockCustomerEventRepo func(*gomock.Controller) *mock.MockCustomerEventRepository
	}

	type testCaseExpected struct {
		wantError      bool
		errWantCompare bool
		err            error
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	tests := []testCase{
		{
			name: "shouldn't block customer - customer not found",
			params: testCaseParams{
				dto: BlockCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
					Reason:     "some reason",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrCustomerNotFound)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					return mock.NewMockCustomerEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: true,
				err:            ErrCustomerNotFound,
			},
		},
		{
			name: "shouldn't block customer - customer event repository error",
			params: testCaseParams{
				dto: BlockCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
					Reason:     "some reason",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(errors.New("some error"))

					return mock
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: false,
			},
		},
		{
			name: "should block customer",
			params: testCaseParams{
				dto: BlockCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
					Reason:     "some reason",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)

					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewCustomerService(
				tt.params.mockCustomerQueryRepo(ctrl),
				tt.params.mockCustomerEventRepo(ctrl),
			)

			err := service.BlockCustomer(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)

				if tt.expected.errWantCompare {
					require.ErrorIs(t, err, tt.expected.err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_CustomerService_UnblockCustomer(t *testing.T) {

	type testCaseParams struct {
		dto UnblockCustomerDTO

		mockCustomerQueryRepo func(*gomock.Controller) *mock.MockCustomerQueryRepository
		mockCustomerEventRepo func(*gomock.Controller) *mock.MockCustomerEventRepository
	}

	type testCaseExpected struct {
		wantError      bool
		errWantCompare bool
		err            error
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	tests := []testCase{
		{
			name: "shouldn't unblock customer - customer not found",
			params: testCaseParams{
				dto: UnblockCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrCustomerNotFound)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					return mock.NewMockCustomerEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: true,
				err:            ErrCustomerNotFound,
			},
		},
		{
			name: "shouldn't unblock customer - customer event repository error",
			params: testCaseParams{
				dto: UnblockCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(errors.New("some error"))

					return mock
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: false,
			},
		},
		{
			name: "should unblock customer",
			params: testCaseParams{
				dto: UnblockCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)

					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewCustomerService(
				tt.params.mockCustomerQueryRepo(ctrl),
				tt.params.mockCustomerEventRepo(ctrl),
			)

			err := service.UnblockCustomer(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)

				if tt.expected.errWantCompare {
					require.ErrorIs(t, err, tt.expected.err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_CustomerService_DeleteCustomer(t *testing.T) {

	type testCaseParams struct {
		dto DeleteCustomerDTO

		mockCustomerQueryRepo func(*gomock.Controller) *mock.MockCustomerQueryRepository
		mockCustomerEventRepo func(*gomock.Controller) *mock.MockCustomerEventRepository
	}

	type testCaseExpected struct {
		wantError      bool
		errWantCompare bool
		err            error
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	tests := []testCase{
		{
			name: "shouldn't delete customer - customer not found",
			params: testCaseParams{
				dto: DeleteCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrCustomerNotFound)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					return mock.NewMockCustomerEventRepository(m)
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: true,
				err:            ErrCustomerNotFound,
			},
		},
		{
			name: "shouldn't delete customer - customer event repository error",
			params: testCaseParams{
				dto: DeleteCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(errors.New("some error"))

					return mock
				},
			},
			expected: testCaseExpected{
				wantError:      true,
				errWantCompare: false,
			},
		},
		{
			name: "should delete customer",
			params: testCaseParams{
				dto: DeleteCustomerDTO{
					CustomerID: "00000000-0000-0000-0000-000000000000",
				},
				mockCustomerQueryRepo: func(m *gomock.Controller) *mock.MockCustomerQueryRepository {
					mock := mock.NewMockCustomerQueryRepository(m)
					mock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&customerdomain.Customer{}, nil)

					return mock
				},
				mockCustomerEventRepo: func(m *gomock.Controller) *mock.MockCustomerEventRepository {
					mock := mock.NewMockCustomerEventRepository(m)
					mock.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(nil)

					return mock
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewCustomerService(
				tt.params.mockCustomerQueryRepo(ctrl),
				tt.params.mockCustomerEventRepo(ctrl),
			)

			err := service.DeleteCustomer(context.Background(), tt.params.dto)

			if tt.expected.wantError {
				require.Error(t, err)

				if tt.expected.errWantCompare {
					require.ErrorIs(t, err, tt.expected.err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}
