// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/kindlyfire/longa/ent/link"
)

// Link is the model entity for the Link schema.
type Link struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Link holds the value of the "link" field.
	Link string `json:"link,omitempty"`
	// Target holds the value of the "target" field.
	Target string `json:"target,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Link) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case link.FieldID:
			values[i] = new(sql.NullInt64)
		case link.FieldLink, link.FieldTarget:
			values[i] = new(sql.NullString)
		case link.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Link", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Link fields.
func (l *Link) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case link.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case link.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				l.Link = value.String
			}
		case link.FieldTarget:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field target", values[i])
			} else if value.Valid {
				l.Target = value.String
			}
		case link.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Link.
// Note that you need to call Link.Unwrap() before calling this method if this Link
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Link) Update() *LinkUpdateOne {
	return (&LinkClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Link entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Link) Unwrap() *Link {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Link is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Link) String() string {
	var builder strings.Builder
	builder.WriteString("Link(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", link=")
	builder.WriteString(l.Link)
	builder.WriteString(", target=")
	builder.WriteString(l.Target)
	builder.WriteString(", createdAt=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Links is a parsable slice of Link.
type Links []*Link

func (l Links) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
