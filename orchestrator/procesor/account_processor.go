package processor

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
)

type (
	AccountCreatedEvent        = accountdomain.AccountCreatedEvent
	AccountFundsWithdrawnEvent = accountdomain.FundsWithdrawnEvent
	AccountFundsDepositedEvent = accountdomain.FundsDepositedEvent
	AccountBlockedEvent        = accountdomain.AccountBlockedEvent
	AccountUnblockedEvent      = accountdomain.AccountUnblockedEvent
)

type accountEventType interface {
	AccountCreatedEvent |
		AccountFundsWithdrawnEvent |
		AccountFundsDepositedEvent |
		AccountBlockedEvent |
		AccountUnblockedEvent
}

// AccountProcessor handles the processing of account-related events
type AccountProcessor struct {
	// orchestrator repository
	orcRepo OrchestratorRepository
	// account repository
	accountRepo AccountRepository
}

// NewAccountProcessor creates a new account event processor
func NewAccountProcessor(orcRepo OrchestratorRepository, accountRepo AccountRepository) *AccountProcessor {
	return &AccountProcessor{
		orcRepo:     orcRepo,
		accountRepo: accountRepo,
	}
}

// Process handles the account event processing
func (p *AccountProcessor) Process(ctx context.Context, event BaseEvent) error {
	switch event.GetType() {
	case accountdomain.AccountCreatedEventType.String():
		event, err := UnmarshalEvent[AccountCreatedEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal account created event: %w", err)
		}

		return p.handleAccountCreatedEvent(ctx, event.Data)

	case accountdomain.AccountFundsWithdrawnEventType.String():
		event, err := UnmarshalEvent[AccountFundsWithdrawnEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal account funds withdrawn event: %w", err)
		}

		return p.handleAccountFundsWithdrawnEvent(ctx, event.Data)

	case accountdomain.AccountFundsDepositedEventType.String():
		event, err := UnmarshalEvent[AccountFundsDepositedEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal account funds deposited event: %w", err)
		}

		return p.handleAccountFundsDepositedEvent(ctx, event.Data)

	case accountdomain.AccountBlockedEventType.String():
		event, err := UnmarshalEvent[AccountBlockedEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal account blocked event: %w", err)
		}

		return p.handleAccountBlockedEvent(ctx, event.Data)

	case accountdomain.AccountUnblockedEventType.String():
		event, err := UnmarshalEvent[AccountUnblockedEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal account unblocked event: %w", err)
		}

		return p.handleAccountUnblockedEvent(ctx, event.Data)

	default:
		if err := p.handleUnknownEvent(ctx, event.GetID()); err != nil {
			return fmt.Errorf("handling unknown account event: %w", err)
		}

		return nil
	}
}

// handleAccountCreated processes account creation events
func (p *AccountProcessor) handleAccountCreatedEvent(_ context.Context, _ AccountCreatedEvent) error {
	// Implement account creation logic
	return nil
}

func (p *AccountProcessor) handleAccountFundsWithdrawnEvent(_ context.Context, _ AccountFundsWithdrawnEvent) error {
	// Implement account funds withdrawn logic
	return nil
}

func (p *AccountProcessor) handleAccountFundsDepositedEvent(_ context.Context, _ AccountFundsDepositedEvent) error {
	// Implement account funds deposited logic
	return nil
}

// handleAccountBlocked processes account blocked events
func (p *AccountProcessor) handleAccountBlockedEvent(_ context.Context, _ AccountBlockedEvent) error {
	// Implement account deletion logic
	return nil
}

func (p *AccountProcessor) handleAccountUnblockedEvent(_ context.Context, _ AccountUnblockedEvent) error {
	// Implement account unblocked logic
	return nil
}

// handleUnknownEvent processes unknown event types
func (p *AccountProcessor) handleUnknownEvent(ctx context.Context, id uuid.UUID) error {
	return p.orcRepo.UpdateEventState(ctx, id, "unprocessable")
}
