// Code generated by ent, DO NOT EDIT.

package placio_api

import (
	"encoding/json"
	"fmt"
	"placio_api/accountsettings"
	"placio_api/business"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// AccountSettings is the model entity for the AccountSettings schema.
type AccountSettings struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// TwoFactorAuthentication holds the value of the "TwoFactorAuthentication" field.
	TwoFactorAuthentication bool `json:"TwoFactorAuthentication,omitempty"`
	// BlockedUsers holds the value of the "BlockedUsers" field.
	BlockedUsers []string `json:"BlockedUsers,omitempty"`
	// MutedUsers holds the value of the "MutedUsers" field.
	MutedUsers []string `json:"MutedUsers,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountSettingsQuery when eager-loading is set.
	Edges                              AccountSettingsEdges `json:"edges"`
	business_business_account_settings *string
	selectValues                       sql.SelectValues
}

// AccountSettingsEdges holds the relations/edges for other nodes in the graph.
type AccountSettingsEdges struct {
	// BusinessAccount holds the value of the business_account edge.
	BusinessAccount *Business `json:"business_account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BusinessAccountOrErr returns the BusinessAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountSettingsEdges) BusinessAccountOrErr() (*Business, error) {
	if e.loadedTypes[0] {
		if e.BusinessAccount == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: business.Label}
		}
		return e.BusinessAccount, nil
	}
	return nil, &NotLoadedError{edge: "business_account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccountSettings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accountsettings.FieldBlockedUsers, accountsettings.FieldMutedUsers:
			values[i] = new([]byte)
		case accountsettings.FieldTwoFactorAuthentication:
			values[i] = new(sql.NullBool)
		case accountsettings.FieldID:
			values[i] = new(sql.NullString)
		case accountsettings.ForeignKeys[0]: // business_business_account_settings
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccountSettings fields.
func (as *AccountSettings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accountsettings.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				as.ID = value.String
			}
		case accountsettings.FieldTwoFactorAuthentication:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field TwoFactorAuthentication", values[i])
			} else if value.Valid {
				as.TwoFactorAuthentication = value.Bool
			}
		case accountsettings.FieldBlockedUsers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field BlockedUsers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &as.BlockedUsers); err != nil {
					return fmt.Errorf("unmarshal field BlockedUsers: %w", err)
				}
			}
		case accountsettings.FieldMutedUsers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field MutedUsers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &as.MutedUsers); err != nil {
					return fmt.Errorf("unmarshal field MutedUsers: %w", err)
				}
			}
		case accountsettings.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_business_account_settings", values[i])
			} else if value.Valid {
				as.business_business_account_settings = new(string)
				*as.business_business_account_settings = value.String
			}
		default:
			as.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AccountSettings.
// This includes values selected through modifiers, order, etc.
func (as *AccountSettings) Value(name string) (ent.Value, error) {
	return as.selectValues.Get(name)
}

// QueryBusinessAccount queries the "business_account" edge of the AccountSettings entity.
func (as *AccountSettings) QueryBusinessAccount() *BusinessQuery {
	return NewAccountSettingsClient(as.config).QueryBusinessAccount(as)
}

// Update returns a builder for updating this AccountSettings.
// Note that you need to call AccountSettings.Unwrap() before calling this method if this AccountSettings
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *AccountSettings) Update() *AccountSettingsUpdateOne {
	return NewAccountSettingsClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the AccountSettings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *AccountSettings) Unwrap() *AccountSettings {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("placio_api: AccountSettings is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *AccountSettings) String() string {
	var builder strings.Builder
	builder.WriteString("AccountSettings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("TwoFactorAuthentication=")
	builder.WriteString(fmt.Sprintf("%v", as.TwoFactorAuthentication))
	builder.WriteString(", ")
	builder.WriteString("BlockedUsers=")
	builder.WriteString(fmt.Sprintf("%v", as.BlockedUsers))
	builder.WriteString(", ")
	builder.WriteString("MutedUsers=")
	builder.WriteString(fmt.Sprintf("%v", as.MutedUsers))
	builder.WriteByte(')')
	return builder.String()
}

// AccountSettingsSlice is a parsable slice of AccountSettings.
type AccountSettingsSlice []*AccountSettings
