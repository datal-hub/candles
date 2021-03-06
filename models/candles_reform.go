// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type candleTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *candleTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("candles").
func (v *candleTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *candleTableType) Columns() []string {
	return []string{"id", "high", "low", "open", "close", "ticker_id", "begin_dttm", "end_dttm", "interval"}
}

// NewStruct makes a new struct for that view or table.
func (v *candleTableType) NewStruct() reform.Struct {
	return new(Candle)
}

// NewRecord makes a new record for that table.
func (v *candleTableType) NewRecord() reform.Record {
	return new(Candle)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *candleTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// CandleTable represents candles view or table in SQL database.
var CandleTable = &candleTableType{
	s: parse.StructInfo{Type: "Candle", SQLSchema: "", SQLName: "candles", Fields: []parse.FieldInfo{{Name: "ID", Type: "int64", Column: "id"}, {Name: "High", Type: "float64", Column: "high"}, {Name: "Low", Type: "float64", Column: "low"}, {Name: "Open", Type: "float64", Column: "open"}, {Name: "Close", Type: "float64", Column: "close"}, {Name: "TickerId", Type: "int64", Column: "ticker_id"}, {Name: "BeginDttm", Type: "int64", Column: "begin_dttm"}, {Name: "EndDttm", Type: "int64", Column: "end_dttm"}, {Name: "Interval", Type: "int64", Column: "interval"}}, PKFieldIndex: 0},
	z: new(Candle).Values(),
}

// String returns a string representation of this struct or record.
func (s Candle) String() string {
	res := make([]string, 9)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "High: " + reform.Inspect(s.High, true)
	res[2] = "Low: " + reform.Inspect(s.Low, true)
	res[3] = "Open: " + reform.Inspect(s.Open, true)
	res[4] = "Close: " + reform.Inspect(s.Close, true)
	res[5] = "TickerId: " + reform.Inspect(s.TickerId, true)
	res[6] = "BeginDttm: " + reform.Inspect(s.BeginDttm, true)
	res[7] = "EndDttm: " + reform.Inspect(s.EndDttm, true)
	res[8] = "Interval: " + reform.Inspect(s.Interval, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Candle) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.High,
		s.Low,
		s.Open,
		s.Close,
		s.TickerId,
		s.BeginDttm,
		s.EndDttm,
		s.Interval,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Candle) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.High,
		&s.Low,
		&s.Open,
		&s.Close,
		&s.TickerId,
		&s.BeginDttm,
		&s.EndDttm,
		&s.Interval,
	}
}

// View returns View object for that struct.
func (s *Candle) View() reform.View {
	return CandleTable
}

// Table returns Table object for that record.
func (s *Candle) Table() reform.Table {
	return CandleTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Candle) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Candle) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Candle) HasPK() bool {
	return s.ID != CandleTable.z[CandleTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Candle) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int64(i64)
	} else {
		s.ID = pk.(int64)
	}
}

// check interfaces
var (
	_ reform.View   = CandleTable
	_ reform.Struct = (*Candle)(nil)
	_ reform.Table  = CandleTable
	_ reform.Record = (*Candle)(nil)
	_ fmt.Stringer  = (*Candle)(nil)
)

func init() {
	parse.AssertUpToDate(&CandleTable.s, new(Candle))
}
