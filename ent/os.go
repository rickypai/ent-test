// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/rickypai/ent-test/ent/os"
)

// OS is the model entity for the OS schema.
type OS struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OS) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case os.FieldID:
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type OS", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OS fields.
func (o *OS) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case os.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this OS.
// Note that you need to call OS.Unwrap() before calling this method if this OS
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *OS) Update() *OSUpdateOne {
	return (&OSClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the OS entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *OS) Unwrap() *OS {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: OS is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *OS) String() string {
	var builder strings.Builder
	builder.WriteString("OS(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteByte(')')
	return builder.String()
}

// OSs is a parsable slice of OS.
type OSs []*OS

func (o OSs) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}