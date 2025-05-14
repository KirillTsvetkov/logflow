package logging

import (
	opensearch "logParser/internal/adapters/logging/opensearch_gen"
)

// OpenSearchAdapter - адаптер для работы с Opensearch.
type OpenSearchAdapter struct {
	client *opensearch.Client
}
