// Code generated by entc, DO NOT EDIT.

package group

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNillableName holds the string denoting the nillable_name field in the database.
	FieldNillableName = "nillable_name"
	// FieldOptionalName holds the string denoting the optional_name field in the database.
	FieldOptionalName = "optional_name"
	// FieldCreationDate holds the string denoting the creation_date field in the database.
	FieldCreationDate = "creation_date"

	// Table holds the table name of the group in the database.
	Table = "groups"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNillableName,
	FieldOptionalName,
	FieldCreationDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
