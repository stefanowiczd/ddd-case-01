package orchestrator

// Orchestrator is the main struct for the orchestrator
type Orchestrator struct {
	orcRepo      OrchestratorRepository
	accountRepo  AccountRepository
	customerRepo CustomerRepository
}

// NewOrchestrator creates a new Orchestrator
func NewOrchestrator(
	orcRepo OrchestratorRepository,
	accountRepo AccountRepository,
	customerRepo CustomerRepository,
) *Orchestrator {
	return &Orchestrator{orcRepo, accountRepo, customerRepo}
}
