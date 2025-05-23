// Code generated by ogen, DO NOT EDIT.

package opensearch

import (
	"time"
)

type IndexSearchPostReq struct {
	Query OptIndexSearchPostReqQuery `json:"query"`
}

// GetQuery returns the value of Query.
func (s *IndexSearchPostReq) GetQuery() OptIndexSearchPostReqQuery {
	return s.Query
}

// SetQuery sets the value of Query.
func (s *IndexSearchPostReq) SetQuery(val OptIndexSearchPostReqQuery) {
	s.Query = val
}

type IndexSearchPostReqQuery struct {
	Term OptIndexSearchPostReqQueryTerm `json:"term"`
}

// GetTerm returns the value of Term.
func (s *IndexSearchPostReqQuery) GetTerm() OptIndexSearchPostReqQueryTerm {
	return s.Term
}

// SetTerm sets the value of Term.
func (s *IndexSearchPostReqQuery) SetTerm(val OptIndexSearchPostReqQueryTerm) {
	s.Term = val
}

type IndexSearchPostReqQueryTerm struct {
	TraceIDDotKeyword OptString `json:"trace_id.keyword"`
}

// GetTraceIDDotKeyword returns the value of TraceIDDotKeyword.
func (s *IndexSearchPostReqQueryTerm) GetTraceIDDotKeyword() OptString {
	return s.TraceIDDotKeyword
}

// SetTraceIDDotKeyword sets the value of TraceIDDotKeyword.
func (s *IndexSearchPostReqQueryTerm) SetTraceIDDotKeyword(val OptString) {
	s.TraceIDDotKeyword = val
}

// Ref: #/components/schemas/LogEntry
type LogEntry struct {
	Source OptLogEntrySource `json:"_source"`
}

// GetSource returns the value of Source.
func (s *LogEntry) GetSource() OptLogEntrySource {
	return s.Source
}

// SetSource sets the value of Source.
func (s *LogEntry) SetSource(val OptLogEntrySource) {
	s.Source = val
}

type LogEntrySource struct {
	Timestamp OptDateTime `json:"timestamp"`
	Message   OptString   `json:"message"`
	TraceID   OptString   `json:"trace_id"`
}

// GetTimestamp returns the value of Timestamp.
func (s *LogEntrySource) GetTimestamp() OptDateTime {
	return s.Timestamp
}

// GetMessage returns the value of Message.
func (s *LogEntrySource) GetMessage() OptString {
	return s.Message
}

// GetTraceID returns the value of TraceID.
func (s *LogEntrySource) GetTraceID() OptString {
	return s.TraceID
}

// SetTimestamp sets the value of Timestamp.
func (s *LogEntrySource) SetTimestamp(val OptDateTime) {
	s.Timestamp = val
}

// SetMessage sets the value of Message.
func (s *LogEntrySource) SetMessage(val OptString) {
	s.Message = val
}

// SetTraceID sets the value of TraceID.
func (s *LogEntrySource) SetTraceID(val OptString) {
	s.TraceID = val
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptIndexSearchPostReq returns new OptIndexSearchPostReq with value set to v.
func NewOptIndexSearchPostReq(v IndexSearchPostReq) OptIndexSearchPostReq {
	return OptIndexSearchPostReq{
		Value: v,
		Set:   true,
	}
}

// OptIndexSearchPostReq is optional IndexSearchPostReq.
type OptIndexSearchPostReq struct {
	Value IndexSearchPostReq
	Set   bool
}

// IsSet returns true if OptIndexSearchPostReq was set.
func (o OptIndexSearchPostReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptIndexSearchPostReq) Reset() {
	var v IndexSearchPostReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptIndexSearchPostReq) SetTo(v IndexSearchPostReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptIndexSearchPostReq) Get() (v IndexSearchPostReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptIndexSearchPostReq) Or(d IndexSearchPostReq) IndexSearchPostReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptIndexSearchPostReqQuery returns new OptIndexSearchPostReqQuery with value set to v.
func NewOptIndexSearchPostReqQuery(v IndexSearchPostReqQuery) OptIndexSearchPostReqQuery {
	return OptIndexSearchPostReqQuery{
		Value: v,
		Set:   true,
	}
}

// OptIndexSearchPostReqQuery is optional IndexSearchPostReqQuery.
type OptIndexSearchPostReqQuery struct {
	Value IndexSearchPostReqQuery
	Set   bool
}

// IsSet returns true if OptIndexSearchPostReqQuery was set.
func (o OptIndexSearchPostReqQuery) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptIndexSearchPostReqQuery) Reset() {
	var v IndexSearchPostReqQuery
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptIndexSearchPostReqQuery) SetTo(v IndexSearchPostReqQuery) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptIndexSearchPostReqQuery) Get() (v IndexSearchPostReqQuery, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptIndexSearchPostReqQuery) Or(d IndexSearchPostReqQuery) IndexSearchPostReqQuery {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptIndexSearchPostReqQueryTerm returns new OptIndexSearchPostReqQueryTerm with value set to v.
func NewOptIndexSearchPostReqQueryTerm(v IndexSearchPostReqQueryTerm) OptIndexSearchPostReqQueryTerm {
	return OptIndexSearchPostReqQueryTerm{
		Value: v,
		Set:   true,
	}
}

// OptIndexSearchPostReqQueryTerm is optional IndexSearchPostReqQueryTerm.
type OptIndexSearchPostReqQueryTerm struct {
	Value IndexSearchPostReqQueryTerm
	Set   bool
}

// IsSet returns true if OptIndexSearchPostReqQueryTerm was set.
func (o OptIndexSearchPostReqQueryTerm) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptIndexSearchPostReqQueryTerm) Reset() {
	var v IndexSearchPostReqQueryTerm
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptIndexSearchPostReqQueryTerm) SetTo(v IndexSearchPostReqQueryTerm) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptIndexSearchPostReqQueryTerm) Get() (v IndexSearchPostReqQueryTerm, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptIndexSearchPostReqQueryTerm) Or(d IndexSearchPostReqQueryTerm) IndexSearchPostReqQueryTerm {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptLogEntrySource returns new OptLogEntrySource with value set to v.
func NewOptLogEntrySource(v LogEntrySource) OptLogEntrySource {
	return OptLogEntrySource{
		Value: v,
		Set:   true,
	}
}

// OptLogEntrySource is optional LogEntrySource.
type OptLogEntrySource struct {
	Value LogEntrySource
	Set   bool
}

// IsSet returns true if OptLogEntrySource was set.
func (o OptLogEntrySource) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptLogEntrySource) Reset() {
	var v LogEntrySource
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptLogEntrySource) SetTo(v LogEntrySource) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptLogEntrySource) Get() (v LogEntrySource, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptLogEntrySource) Or(d LogEntrySource) LogEntrySource {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSearchResponseHits returns new OptSearchResponseHits with value set to v.
func NewOptSearchResponseHits(v SearchResponseHits) OptSearchResponseHits {
	return OptSearchResponseHits{
		Value: v,
		Set:   true,
	}
}

// OptSearchResponseHits is optional SearchResponseHits.
type OptSearchResponseHits struct {
	Value SearchResponseHits
	Set   bool
}

// IsSet returns true if OptSearchResponseHits was set.
func (o OptSearchResponseHits) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSearchResponseHits) Reset() {
	var v SearchResponseHits
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSearchResponseHits) SetTo(v SearchResponseHits) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSearchResponseHits) Get() (v SearchResponseHits, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSearchResponseHits) Or(d SearchResponseHits) SearchResponseHits {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/SearchResponse
type SearchResponse struct {
	Hits OptSearchResponseHits `json:"hits"`
}

// GetHits returns the value of Hits.
func (s *SearchResponse) GetHits() OptSearchResponseHits {
	return s.Hits
}

// SetHits sets the value of Hits.
func (s *SearchResponse) SetHits(val OptSearchResponseHits) {
	s.Hits = val
}

type SearchResponseHits struct {
	Hits []LogEntry `json:"hits"`
}

// GetHits returns the value of Hits.
func (s *SearchResponseHits) GetHits() []LogEntry {
	return s.Hits
}

// SetHits sets the value of Hits.
func (s *SearchResponseHits) SetHits(val []LogEntry) {
	s.Hits = val
}
