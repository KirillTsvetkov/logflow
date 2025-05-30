// Code generated by ogen, DO NOT EDIT.

package opensearch

import (
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
)

// Encode implements json.Marshaler.
func (s *IndexSearchPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *IndexSearchPostReq) encodeFields(e *jx.Encoder) {
	{
		if s.Query.Set {
			e.FieldStart("query")
			s.Query.Encode(e)
		}
	}
}

var jsonFieldsNameOfIndexSearchPostReq = [1]string{
	0: "query",
}

// Decode decodes IndexSearchPostReq from json.
func (s *IndexSearchPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode IndexSearchPostReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "query":
			if err := func() error {
				s.Query.Reset()
				if err := s.Query.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"query\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode IndexSearchPostReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *IndexSearchPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *IndexSearchPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *IndexSearchPostReqQuery) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *IndexSearchPostReqQuery) encodeFields(e *jx.Encoder) {
	{
		if s.Term.Set {
			e.FieldStart("term")
			s.Term.Encode(e)
		}
	}
}

var jsonFieldsNameOfIndexSearchPostReqQuery = [1]string{
	0: "term",
}

// Decode decodes IndexSearchPostReqQuery from json.
func (s *IndexSearchPostReqQuery) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode IndexSearchPostReqQuery to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "term":
			if err := func() error {
				s.Term.Reset()
				if err := s.Term.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"term\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode IndexSearchPostReqQuery")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *IndexSearchPostReqQuery) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *IndexSearchPostReqQuery) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *IndexSearchPostReqQueryTerm) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *IndexSearchPostReqQueryTerm) encodeFields(e *jx.Encoder) {
	{
		if s.TraceIDDotKeyword.Set {
			e.FieldStart("trace_id.keyword")
			s.TraceIDDotKeyword.Encode(e)
		}
	}
}

var jsonFieldsNameOfIndexSearchPostReqQueryTerm = [1]string{
	0: "trace_id.keyword",
}

// Decode decodes IndexSearchPostReqQueryTerm from json.
func (s *IndexSearchPostReqQueryTerm) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode IndexSearchPostReqQueryTerm to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "trace_id.keyword":
			if err := func() error {
				s.TraceIDDotKeyword.Reset()
				if err := s.TraceIDDotKeyword.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"trace_id.keyword\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode IndexSearchPostReqQueryTerm")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *IndexSearchPostReqQueryTerm) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *IndexSearchPostReqQueryTerm) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LogEntry) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LogEntry) encodeFields(e *jx.Encoder) {
	{
		if s.Source.Set {
			e.FieldStart("_source")
			s.Source.Encode(e)
		}
	}
}

var jsonFieldsNameOfLogEntry = [1]string{
	0: "_source",
}

// Decode decodes LogEntry from json.
func (s *LogEntry) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LogEntry to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "_source":
			if err := func() error {
				s.Source.Reset()
				if err := s.Source.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"_source\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LogEntry")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LogEntry) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LogEntry) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LogEntrySource) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LogEntrySource) encodeFields(e *jx.Encoder) {
	{
		if s.Timestamp.Set {
			e.FieldStart("timestamp")
			s.Timestamp.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
	{
		if s.TraceID.Set {
			e.FieldStart("trace_id")
			s.TraceID.Encode(e)
		}
	}
}

var jsonFieldsNameOfLogEntrySource = [3]string{
	0: "timestamp",
	1: "message",
	2: "trace_id",
}

// Decode decodes LogEntrySource from json.
func (s *LogEntrySource) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LogEntrySource to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "timestamp":
			if err := func() error {
				s.Timestamp.Reset()
				if err := s.Timestamp.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"timestamp\"")
			}
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "trace_id":
			if err := func() error {
				s.TraceID.Reset()
				if err := s.TraceID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"trace_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LogEntrySource")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LogEntrySource) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LogEntrySource) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDateTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDateTime)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDateTime)
}

// Encode encodes IndexSearchPostReq as json.
func (o OptIndexSearchPostReq) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes IndexSearchPostReq from json.
func (o *OptIndexSearchPostReq) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptIndexSearchPostReq to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptIndexSearchPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptIndexSearchPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes IndexSearchPostReqQuery as json.
func (o OptIndexSearchPostReqQuery) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes IndexSearchPostReqQuery from json.
func (o *OptIndexSearchPostReqQuery) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptIndexSearchPostReqQuery to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptIndexSearchPostReqQuery) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptIndexSearchPostReqQuery) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes IndexSearchPostReqQueryTerm as json.
func (o OptIndexSearchPostReqQueryTerm) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes IndexSearchPostReqQueryTerm from json.
func (o *OptIndexSearchPostReqQueryTerm) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptIndexSearchPostReqQueryTerm to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptIndexSearchPostReqQueryTerm) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptIndexSearchPostReqQueryTerm) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes LogEntrySource as json.
func (o OptLogEntrySource) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes LogEntrySource from json.
func (o *OptLogEntrySource) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptLogEntrySource to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptLogEntrySource) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptLogEntrySource) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SearchResponseHits as json.
func (o OptSearchResponseHits) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes SearchResponseHits from json.
func (o *OptSearchResponseHits) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSearchResponseHits to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSearchResponseHits) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSearchResponseHits) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SearchResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SearchResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Hits.Set {
			e.FieldStart("hits")
			s.Hits.Encode(e)
		}
	}
}

var jsonFieldsNameOfSearchResponse = [1]string{
	0: "hits",
}

// Decode decodes SearchResponse from json.
func (s *SearchResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "hits":
			if err := func() error {
				s.Hits.Reset()
				if err := s.Hits.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hits\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SearchResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SearchResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SearchResponseHits) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SearchResponseHits) encodeFields(e *jx.Encoder) {
	{
		if s.Hits != nil {
			e.FieldStart("hits")
			e.ArrStart()
			for _, elem := range s.Hits {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfSearchResponseHits = [1]string{
	0: "hits",
}

// Decode decodes SearchResponseHits from json.
func (s *SearchResponseHits) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchResponseHits to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "hits":
			if err := func() error {
				s.Hits = make([]LogEntry, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem LogEntry
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Hits = append(s.Hits, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hits\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SearchResponseHits")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SearchResponseHits) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchResponseHits) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
