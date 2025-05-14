package models

import "time"

// Report - отчет по трейсу
type Report struct {
	TraceID       string          `json:"trace_id"`
	Summary       string          `json:"summary"`        // Краткое текстовое описание
	Status        string          `json:"status"`         // Общий статус (success, partial_failure, error)
	Timeline      []TimelineEvent `json:"timeline"`       // Ключевые события по времени
	Errors        []ErrorInfo     `json:"errors"`         // Список ошибок (если есть)
	ExternalCalls []ExternalCall  `json:"external_calls"` // Вызовы внешних сервисов
	LogsLink      string          `json:"logs_link"`      // Ссылка на логи в Opensearch
}

// TimelineEvent - ключевое событие в трейсе
type TimelineEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Event     string    `json:"event"`   // Например: "call_user_service", "cache_miss"
	Details   string    `json:"details"` // Доп. информация
}

// ErrorInfo - информация об ошибке
type ErrorInfo struct {
	Timestamp  time.Time `json:"timestamp"`
	Message    string    `json:"message"`
	StackTrace string    `json:"stack_trace"` // Опционально
	Service    string    `json:"service"`     // Где произошла ошибка
}

// ExternalCall - вызов внешнего сервиса/API
type ExternalCall struct {
	Timestamp  time.Time `json:"timestamp"`
	Service    string    `json:"service"`     // Например: "payment_service"
	Method     string    `json:"method"`      // HTTP-метод (GET, POST)
	Endpoint   string    `json:"endpoint"`    // URL или название метода
	Duration   int64     `json:"duration"`    // Время выполнения (мс)
	StatusCode int       `json:"status_code"` // HTTP-код или аналогичный
}
