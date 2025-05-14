//go:build unit

package processor

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
	eventdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/event"
	"github.com/stefanowiczd/ddd-case-01/orchestrator/application/processor/mock"
)

func TestAccountProcessor_Process_AccountCreatedEvent(t *testing.T) {

	type testCaseParams struct {
		accountCreatedEvent func() *AccountCreatedEvent

		mockOrchestratorRepository func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository
		mockAccountRepository      func(ctrl *gomock.Controller) *mock.MockAccountRepository
	}

	type testCaseExpected struct {
		wantError bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	testCases := []testCase{
		{
			name: "shouldn't process account created event - invalid data resulting in unmarshal error",
			params: testCaseParams{
				accountCreatedEvent: func() *AccountCreatedEvent {
					acc := &AccountCreatedEvent{
						BaseEvent: eventdomain.BaseEvent{
							Origin: "account",
							Type:   "account.created",
						},
					}

					data, err := json.Marshal([]byte(`{ ... invalid data ... }	`))
					if err != nil {
						t.Fatalf("failed to marshal account created event: %v", err)
					}

					acc.Data = data

					return acc
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					return mock.NewMockOrchestratorRepository(ctrl)
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					return mock.NewMockAccountRepository(ctrl)
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process account created event",
			params: testCaseParams{
				accountCreatedEvent: func() *AccountCreatedEvent {
					acc := &AccountCreatedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							ContextID:   uuid.New(),
							Origin:      "account",
							Type:        "account.created",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						CustomerID:     uuid.New(),
						InitialBalance: 1000,
					}

					data, err := json.Marshal(acc)
					if err != nil {
						t.Fatalf("failed to marshal account created event: %v", err)
					}

					acc.Data = data

					return acc
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(nil)

					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(nil)

					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			processor := NewAccountProcessor(
				testCase.params.mockOrchestratorRepository(ctrl),
				testCase.params.mockAccountRepository(ctrl),
			)

			err := processor.Process(context.Background(), testCase.params.accountCreatedEvent())
			if testCase.expected.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountProcessor_Process_AccountFundsWithdrawnEvent(t *testing.T) {
	type testCaseParams struct {
		accountFundsWithdrawnEvent func() *AccountFundsWithdrawnEvent

		mockOrchestratorRepository func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository
		mockAccountRepository      func(ctrl *gomock.Controller) *mock.MockAccountRepository
	}

	type testCaseExpected struct {
		wantError bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	testCases := []testCase{
		{
			name: "shouldn't process account funds withdrawn event - invalid data resulting in unmarshal error",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() *AccountFundsWithdrawnEvent {
					acc := &AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							Origin: "account",
							Type:   "account.funds.withdrawn",
						},
					}

					data, err := json.Marshal([]byte(`{ ... invalid data ... }	`))
					if err != nil {
						t.Fatalf("failed to marshal account funds withdrawn event: %v", err)
					}

					acc.Data = data

					return acc
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					return mock.NewMockOrchestratorRepository(ctrl)
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					return mock.NewMockAccountRepository(ctrl)
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process account funds withdrawn event",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() *AccountFundsWithdrawnEvent {
					acc := &AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							ContextID:   uuid.New(),
							Origin:      "account",
							Type:        "account.funds.withdrawn",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}

					data, err := json.Marshal(acc)
					if err != nil {
						t.Fatalf("failed to marshal account funds withdrawn event: %v", err)
					}

					acc.Data = data

					return acc
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().WithdrawFunds(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			processor := NewAccountProcessor(
				testCase.params.mockOrchestratorRepository(ctrl),
				testCase.params.mockAccountRepository(ctrl),
			)

			err := processor.Process(context.Background(), testCase.params.accountFundsWithdrawnEvent())
			if testCase.expected.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountProcessor_Process_AccountFundsDepositedEvent(t *testing.T) {
	type testCaseParams struct {
		accountFundsDepositedEvent func() *AccountFundsDepositedEvent
	}

	type testCaseExpected struct {
		wantError bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	testCases := []testCase{
		{
			name: "shouldn't process account funds deposited event - invalid data resulting in unmarshal error",
			params: testCaseParams{
				accountFundsDepositedEvent: func() *AccountFundsDepositedEvent {
					acc := &AccountFundsDepositedEvent{
						BaseEvent: eventdomain.BaseEvent{
							Origin: "account",
							Type:   "account.funds.deposited",
						},
					}

					data, err := json.Marshal([]byte(`{ ... invalid data ... }	`))
					if err != nil {
						t.Fatalf("failed to marshal account funds deposited event: %v", err)
					}

					acc.Data = data

					return acc
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process account funds deposited event",
			params: testCaseParams{
				accountFundsDepositedEvent: func() *AccountFundsDepositedEvent {
					acc := &AccountFundsDepositedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							ContextID:   uuid.New(),
							Origin:      "account",
							Type:        "account.funds.deposited",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}

					data, err := json.Marshal(acc)
					if err != nil {
						t.Fatalf("failed to marshal account funds deposited event: %v", err)
					}

					acc.Data = data

					return acc
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			processor := NewAccountProcessor(
				mock.NewMockOrchestratorRepository(ctrl),
				mock.NewMockAccountRepository(ctrl),
			)

			err := processor.Process(context.Background(), testCase.params.accountFundsDepositedEvent())
			if testCase.expected.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountProcessor_Process_AccountBlockedEvent(t *testing.T) {
	type testCaseParams struct {
		accountBlockedEvent func() *AccountBlockedEvent

		mockOrchestratorRepository func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository
		mockAccountRepository      func(ctrl *gomock.Controller) *mock.MockAccountRepository
	}

	type testCaseExpected struct {
		wantError bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	testCases := []testCase{
		{
			name: "shouldn't process account blocked event - invalid data resulting in unmarshal error",
			params: testCaseParams{
				accountBlockedEvent: func() *AccountBlockedEvent {
					acc := &AccountBlockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							Origin: "account",
							Type:   "account.blocked",
						},
					}

					data, err := json.Marshal([]byte(`{ ... invalid data ... }	`))
					if err != nil {
						t.Fatalf("failed to marshal account blocked event: %v", err)
					}

					acc.Data = data

					return acc
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					return mock.NewMockOrchestratorRepository(ctrl)
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					return mock.NewMockAccountRepository(ctrl)
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process account blocked event",
			params: testCaseParams{
				accountBlockedEvent: func() *AccountBlockedEvent {
					acc := &AccountBlockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							ContextID:   uuid.New(),
							Origin:      "account",
							Type:        "account.blocked",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
					}

					data, err := json.Marshal(acc)
					if err != nil {
						t.Fatalf("failed to marshal account blocked event: %v", err)
					}

					acc.Data = data

					return acc
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().BlockAccount(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			processor := NewAccountProcessor(
				testCase.params.mockOrchestratorRepository(ctrl),
				testCase.params.mockAccountRepository(ctrl),
			)

			err := processor.Process(context.Background(), testCase.params.accountBlockedEvent())
			if testCase.expected.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountProcessor_Process_AccountUnblockedEvent(t *testing.T) {
	type testCaseParams struct {
		accountUnblockedEvent func() *AccountUnblockedEvent
	}

	type testCaseExpected struct {
		wantError bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	testCases := []testCase{
		{
			name: "shouldn't process account unblocked event - invalid data resulting in unmarshal error",
			params: testCaseParams{
				accountUnblockedEvent: func() *AccountUnblockedEvent {
					acc := &AccountUnblockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							Origin: "account",
							Type:   "account.unblocked",
						},
					}

					data, err := json.Marshal([]byte(`{ ... invalid data ... }	`))
					if err != nil {
						t.Fatalf("failed to marshal account unblocked event: %v", err)
					}

					acc.Data = data

					return acc
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process account unblocked event",
			params: testCaseParams{
				accountUnblockedEvent: func() *AccountUnblockedEvent {
					acc := &AccountUnblockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							ContextID:   uuid.New(),
							Origin:      "account",
							Type:        "account.unblocked",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
					}

					data, err := json.Marshal(acc)
					if err != nil {
						t.Fatalf("failed to marshal account unblocked event: %v", err)
					}

					acc.Data = data

					return acc
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			processor := NewAccountProcessor(
				mock.NewMockOrchestratorRepository(ctrl),
				mock.NewMockAccountRepository(ctrl),
			)

			err := processor.Process(context.Background(), testCase.params.accountUnblockedEvent())
			if testCase.expected.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountProcessor_Process_UnknownEvent(t *testing.T) {
	type testCaseParams struct {
		unknownEvent func() *eventdomain.BaseEvent

		orcRepo func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository
	}

	type testCaseExpected struct {
		wantError bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	testCases := []testCase{
		{
			name: "shouldn't process unknown event - internal error",
			params: testCaseParams{
				unknownEvent: func() *eventdomain.BaseEvent {
					return &eventdomain.BaseEvent{
						ID: uuid.New(),
					}
				},
				orcRepo: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventState(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("internal error"))

					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process unknown event",
			params: testCaseParams{
				unknownEvent: func() *eventdomain.BaseEvent {
					return &eventdomain.BaseEvent{
						ID: uuid.New(),
					}
				},
				orcRepo: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventState(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

					return m
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			processor := NewAccountProcessor(
				testCase.params.orcRepo(ctrl),
				mock.NewMockAccountRepository(ctrl),
			)

			err := processor.Process(context.Background(), testCase.params.unknownEvent())
			if testCase.expected.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountProcessor_handleAccountCreatedEvent(t *testing.T) {
	type testCaseParams struct {
		accountCreatedEvent func() AccountCreatedEvent

		mockOrchestratorRepository func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository
		mockAccountRepository      func(ctrl *gomock.Controller) *mock.MockAccountRepository
	}

	type testCaseExpected struct {
		wantError bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	testCases := []testCase{
		{
			name: "shouldn't process account created event - CreateAccount returns internal error, UpdateEventRetry returns nil",
			params: testCaseParams{
				accountCreatedEvent: func() AccountCreatedEvent {
					return AccountCreatedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.created",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						CustomerID:     uuid.New(),
						InitialBalance: 1000,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventRetry(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))

					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
		{
			name: "shouldn't process account created event - CreateAccount returns internal error, UpdateEventRetry returns error",
			params: testCaseParams{
				accountCreatedEvent: func() AccountCreatedEvent {
					return AccountCreatedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.created",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						CustomerID:     uuid.New(),
						InitialBalance: 1000,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventRetry(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("internal error"))

					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))

					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "shouldn't process account created event - CreateAccount returns internal nil, UpdateEventCompletion returns error",
			params: testCaseParams{
				accountCreatedEvent: func() AccountCreatedEvent {
					return AccountCreatedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.created",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						CustomerID:     uuid.New(),
						InitialBalance: 1000,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(nil)

					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process account created event - CreateAccount returns internal nil, UpdateEventCompletion returns nil",
			params: testCaseParams{
				accountCreatedEvent: func() AccountCreatedEvent {
					return AccountCreatedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.created",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						CustomerID:     uuid.New(),
						InitialBalance: 1000,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(nil)

					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
		{
			name: "should process account created event - CreateAccount returns internal ErrAccountAlreadyExists, UpdateEventCompletion returns internal error",
			params: testCaseParams{
				accountCreatedEvent: func() AccountCreatedEvent {
					return AccountCreatedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.created",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						CustomerID:     uuid.New(),
						InitialBalance: 1000,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(accountdomain.ErrAccountAlreadyExists)

					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process account created event - CreateAccount returns internal ErrAccountAlreadyExists, UpdateEventCompletion returns nil",
			params: testCaseParams{
				accountCreatedEvent: func() AccountCreatedEvent {
					return AccountCreatedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.created",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						CustomerID:     uuid.New(),
						InitialBalance: 1000,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(accountdomain.ErrAccountAlreadyExists)

					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			processor := NewAccountProcessor(
				testCase.params.mockOrchestratorRepository(ctrl),
				testCase.params.mockAccountRepository(ctrl),
			)

			err := processor.handleAccountCreatedEvent(context.Background(), testCase.params.accountCreatedEvent())
			if testCase.expected.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountProcessor_handleAccountFundsWithdrawnEvent(t *testing.T) {
	type testCaseParams struct {
		accountFundsWithdrawnEvent func() AccountFundsWithdrawnEvent

		mockOrchestratorRepository func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository
		mockAccountRepository      func(ctrl *gomock.Controller) *mock.MockAccountRepository
	}

	type testCaseExpected struct {
		wantError bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	testCases := []testCase{
		{
			name: "shouldn't process account funds withdrawn event - WithdrawFunds returns ErrAccountInsufficientFunds error, UpdateEventState returns internal error",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() AccountFundsWithdrawnEvent {
					return AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.funds.withdrawn",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventState(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().WithdrawFunds(gomock.Any(), gomock.Any()).Return(accountdomain.ErrAccountInsufficientFunds)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "shouldn't process account funds withdrawn event - WithdrawFunds returns ErrAccountInsufficientFunds error, UpdateEventState returns nil",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() AccountFundsWithdrawnEvent {
					return AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.funds.withdrawn",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventState(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().WithdrawFunds(gomock.Any(), gomock.Any()).Return(accountdomain.ErrAccountInsufficientFunds)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
		{
			name: "shouldn't process account funds withdrawn event - WithdrawFunds returns ErrAccountNotFound error, UpdateEventState returns internal error",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() AccountFundsWithdrawnEvent {
					return AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.funds.withdrawn",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventState(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().WithdrawFunds(gomock.Any(), gomock.Any()).Return(accountdomain.ErrAccountNotFound)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "shouldn't process account funds withdrawn event - WithdrawFunds returns ErrAccountNotFound error, UpdateEventState returns nil",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() AccountFundsWithdrawnEvent {
					return AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.funds.withdrawn",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventState(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().WithdrawFunds(gomock.Any(), gomock.Any()).Return(accountdomain.ErrAccountNotFound)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
		{
			name: "shouldn't process account funds withdrawn event - WithdrawFunds returns internal error, UpdateEventRetry returns internal error",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() AccountFundsWithdrawnEvent {
					return AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.funds.withdrawn",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventRetry(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().WithdrawFunds(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "shouldn't process account funds withdrawn event - WithdrawFunds returns internal error, UpdateEventRetry returns nil",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() AccountFundsWithdrawnEvent {
					return AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.funds.withdrawn",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventRetry(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().WithdrawFunds(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
		{
			name: "shouldn't process account funds withdrawn event - WithdrawFunds returns nil, UpdateEventCompletion returns internal error",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() AccountFundsWithdrawnEvent {
					return AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.funds.withdrawn",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().WithdrawFunds(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process account funds withdrawn event - WithdrawFunds returns nil, UpdateEventCompletion returns nil",
			params: testCaseParams{
				accountFundsWithdrawnEvent: func() AccountFundsWithdrawnEvent {
					return AccountFundsWithdrawnEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.funds.withdrawn",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
						Amount: 100,
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().WithdrawFunds(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			processor := NewAccountProcessor(
				testCase.params.mockOrchestratorRepository(ctrl),
				testCase.params.mockAccountRepository(ctrl),
			)

			err := processor.handleAccountFundsWithdrawnEvent(
				context.Background(),
				testCase.params.accountFundsWithdrawnEvent(),
			)

			if testCase.expected.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAccountProcessor_handleAccountBlockedEvent(t *testing.T) {
	type testCaseParams struct {
		accountBlockedEvent func() AccountBlockedEvent

		mockOrchestratorRepository func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository
		mockAccountRepository      func(ctrl *gomock.Controller) *mock.MockAccountRepository
	}

	type testCaseExpected struct {
		wantError bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	testCases := []testCase{
		{
			name: "shouldn't process account blocked event - BlockAccount returns ErrAccountNotFound error, UpdateEventState returns internal error",
			params: testCaseParams{
				accountBlockedEvent: func() AccountBlockedEvent {
					return AccountBlockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.blocked",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventState(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().BlockAccount(gomock.Any(), gomock.Any()).Return(accountdomain.ErrAccountNotFound)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "shouldn't process account blocked event - BlockAccount returns ErrAccountNotFound error, UpdateEventState returns nil",
			params: testCaseParams{
				accountBlockedEvent: func() AccountBlockedEvent {
					return AccountBlockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.blocked",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventState(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().BlockAccount(gomock.Any(), gomock.Any()).Return(accountdomain.ErrAccountNotFound)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
		{
			name: "shouldn't process account blocked event - BlockAccount returns internal error, UpdateEventRetry returns internal error",
			params: testCaseParams{
				accountBlockedEvent: func() AccountBlockedEvent {
					return AccountBlockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.blocked",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventRetry(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().BlockAccount(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "shouldn't process account blocked event - BlockAccount returns internal error, UpdateEventRetry returns nil",
			params: testCaseParams{
				accountBlockedEvent: func() AccountBlockedEvent {
					return AccountBlockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.blocked",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventRetry(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().BlockAccount(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
		{
			name: "shouldn't process account blocked event - BlockAccount returns nil, UpdateEventCompletion returns internal error",
			params: testCaseParams{
				accountBlockedEvent: func() AccountBlockedEvent {
					return AccountBlockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID:          uuid.New(),
							Origin:      "account",
							Type:        "account.blocked",
							TypeVersion: "1.0.0",
							State:       "created",
							CreatedAt:   time.Now().UTC(),
							ScheduledAt: time.Time{},
							StartedAt:   time.Time{},
							CompletedAt: time.Time{},
							Retry:       0,
							MaxRetry:    3,
							Data:        nil,
						},
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().BlockAccount(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: true,
			},
		},
		{
			name: "should process account blocked event - BlockAccount returns nil, UpdateEventCompletion returns nil",
			params: testCaseParams{
				accountBlockedEvent: func() AccountBlockedEvent {
					return AccountBlockedEvent{
						BaseEvent: eventdomain.BaseEvent{
							ID: uuid.New(),
						},
					}
				},
				mockOrchestratorRepository: func(ctrl *gomock.Controller) *mock.MockOrchestratorRepository {
					m := mock.NewMockOrchestratorRepository(ctrl)
					m.EXPECT().UpdateEventCompletion(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
				mockAccountRepository: func(ctrl *gomock.Controller) *mock.MockAccountRepository {
					m := mock.NewMockAccountRepository(ctrl)
					m.EXPECT().BlockAccount(gomock.Any(), gomock.Any()).Return(nil)
					return m
				},
			},
			expected: testCaseExpected{
				wantError: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			processor := NewAccountProcessor(
				testCase.params.mockOrchestratorRepository(ctrl),
				testCase.params.mockAccountRepository(ctrl),
			)

			err := processor.handleAccountBlockedEvent(
				context.Background(),
				testCase.params.accountBlockedEvent(),
			)

			if testCase.expected.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
