package scanner

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"time"
)

// SCANNER INTERFACE FOR NullInt16
type NullInt16 sql.NullInt16

func (ni *NullInt16) Scan(value interface{}) error {
	var i sql.NullInt16
	if err := i.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*ni = NullInt16{i.Int16, false}
	} else {
		*ni = NullInt16{i.Int16, true}
	}
	return nil
}

// SCANNER INTERFACE FOR NullString
type NullString sql.NullString

func (ns *NullString) Scan(value interface{}) error {
	var i sql.NullString
	if err := i.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*ns = NullString{i.String, false}
	} else {
		*ns = NullString{i.String, true}
	}
	return nil
}

// SCANNER INTERFACE FOR NullTime
type NullTime sql.NullTime

func (nt *NullTime) Scan(value interface{}) error {
	var i sql.NullTime
	if err := i.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*nt = NullTime{i.Time, false}
	} else {
		*nt = NullTime{i.Time, true}
	}
	return nil
}

// JSON MARSHALLING Int16
func (ni *NullInt16) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int16)
}

// JSON UNMARSHALLING Int16
func (ni *NullInt16) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int16)
	ni.Valid = (err == nil)
	return err
}

// JSON MARSHALLING String
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// JSON UNMARSHALLING String
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

// JSON MARSHALLING TIME
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	date := nt.Time.Format("2006-01-02")
	return json.Marshal(date)
}

// JSON UNMARSHALLING TIME
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	nt.Time = t
	nt.Valid = (err == nil)
	return err
}
