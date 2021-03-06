// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/daymenu/gostudy/examples/ent/ent/admin"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// Admin is the model entity for the Admin schema.
type Admin struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Rank holds the value of the "rank" field.
	Rank float64 `json:"rank,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// URL holds the value of the "url" field.
	URL *url.URL `json:"url,omitempty"`
	// Strings holds the value of the "strings" field.
	Strings []string `json:"strings,omitempty"`
	// State holds the value of the "state" field.
	State admin.State `json:"state,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Admin) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullInt64{},   // age
		&sql.NullFloat64{}, // rank
		&sql.NullBool{},    // active
		&sql.NullString{},  // name
		&sql.NullTime{},    // created_at
		&[]byte{},          // url
		&[]byte{},          // strings
		&sql.NullString{},  // state
		&uuid.UUID{},       // uuid
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Admin fields.
func (a *Admin) assignValues(values ...interface{}) error {
	if m, n := len(values), len(admin.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[0])
	} else if value.Valid {
		a.Age = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field rank", values[1])
	} else if value.Valid {
		a.Rank = value.Float64
	}
	if value, ok := values[2].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field active", values[2])
	} else if value.Valid {
		a.Active = value.Bool
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[3])
	} else if value.Valid {
		a.Name = value.String
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[4])
	} else if value.Valid {
		a.CreatedAt = value.Time
	}

	if value, ok := values[5].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field url", values[5])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &a.URL); err != nil {
			return fmt.Errorf("unmarshal field url: %v", err)
		}
	}

	if value, ok := values[6].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field strings", values[6])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &a.Strings); err != nil {
			return fmt.Errorf("unmarshal field strings: %v", err)
		}
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[7])
	} else if value.Valid {
		a.State = admin.State(value.String)
	}
	if value, ok := values[8].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field uuid", values[8])
	} else if value != nil {
		a.UUID = *value
	}
	return nil
}

// Update returns a builder for updating this Admin.
// Note that, you need to call Admin.Unwrap() before calling this method, if this Admin
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Admin) Update() *AdminUpdateOne {
	return (&AdminClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Admin) Unwrap() *Admin {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Admin is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Admin) String() string {
	var builder strings.Builder
	builder.WriteString("Admin(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", a.Age))
	builder.WriteString(", rank=")
	builder.WriteString(fmt.Sprintf("%v", a.Rank))
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", a.Active))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", url=")
	builder.WriteString(fmt.Sprintf("%v", a.URL))
	builder.WriteString(", strings=")
	builder.WriteString(fmt.Sprintf("%v", a.Strings))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", a.State))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", a.UUID))
	builder.WriteByte(')')
	return builder.String()
}

// Admins is a parsable slice of Admin.
type Admins []*Admin

func (a Admins) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
