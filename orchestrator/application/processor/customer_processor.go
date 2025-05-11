package processor

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
)

type (
	CustomerCreatedEvent     = customerdomain.CustomerCreatedEvent
	CustomerActivatedEvent   = customerdomain.CustomerActivatedEvent
	CustomerDeactivatedEvent = customerdomain.CustomerDeactivatedEvent
	CustomerBlockedEvent     = customerdomain.CustomerBlockedEvent
	CustomerUnblockedEvent   = customerdomain.CustomerUnblockedEvent
)

type customerEventType interface {
	CustomerCreatedEvent |
		CustomerActivatedEvent |
		CustomerDeactivatedEvent |
		CustomerBlockedEvent |
		CustomerUnblockedEvent
}

type CustomerProcessor struct {
	orcRepo      OrchestratorRepository
	customerRepo CustomerRepository
}

func NewCustomerProcessor(orcRepo OrchestratorRepository, customerRepo CustomerRepository) *CustomerProcessor {
	return &CustomerProcessor{
		orcRepo:      orcRepo,
		customerRepo: customerRepo,
	}
}

func (p *CustomerProcessor) Process(ctx context.Context, event BaseEvent) error {
	switch event.GetType() {
	case customerdomain.CustomerCreatedEventType.String():
		event, err := UnmarshalEvent[CustomerCreatedEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal customer created event: %w", err)
		}

		return p.handleCustomerCreatedEvent(ctx, event.Data)

	case customerdomain.CustomerActivatedEventType.String():
		event, err := UnmarshalEvent[CustomerActivatedEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal customer activated event: %w", err)
		}

		return p.handleCustomerActivatedEvent(ctx, event.Data)

	case customerdomain.CustomerDeactivatedEventType.String():
		event, err := UnmarshalEvent[CustomerDeactivatedEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal customer deactivated event: %w", err)
		}

		return p.handleCustomerDeactivatedEvent(ctx, event.Data)

	case customerdomain.CustomerBlockedEventType.String():
		event, err := UnmarshalEvent[CustomerBlockedEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal customer blocked event: %w", err)
		}

		return p.handleCustomerBlockedEvent(ctx, event.Data)

	case customerdomain.CustomerUnblockedEventType.String():
		event, err := UnmarshalEvent[CustomerUnblockedEvent](event.GetEventData())
		if err != nil {
			return fmt.Errorf("unmarshal customer unblocked event: %w", err)
		}

		return p.handleCustomerUnblockedEvent(ctx, event.Data)

	default:
		if err := p.handleUnknownEvent(ctx, event.GetID()); err != nil {
			return fmt.Errorf("handling unknown customer event: %w", err)
		}
	}
	return nil
}

func (p *CustomerProcessor) handleCustomerCreatedEvent(ctx context.Context, event CustomerCreatedEvent) error {
	return nil
}

func (p *CustomerProcessor) handleCustomerActivatedEvent(ctx context.Context, event CustomerActivatedEvent) error {
	return nil
}

func (p *CustomerProcessor) handleCustomerDeactivatedEvent(ctx context.Context, event CustomerDeactivatedEvent) error {
	return nil
}

func (p *CustomerProcessor) handleCustomerBlockedEvent(ctx context.Context, event CustomerBlockedEvent) error {
	return nil
}

func (p *CustomerProcessor) handleCustomerUnblockedEvent(ctx context.Context, event CustomerUnblockedEvent) error {
	return nil

}

func (p *CustomerProcessor) handleUnknownEvent(ctx context.Context, id uuid.UUID) error {
	return p.orcRepo.UpdateEventState(ctx, id, "unprocessable")
}
