package snmp

import "context"

// A TableNode value may be used to interact with a table on an SNMP agent in a type-safe manner.
type TableNode[Idx, Row any] struct {
	// The Name of the table.
	Name string
	// The Description of the table.
	Description string
	// The Object Identifier of the table.
	OID OID
}

// A TableRow is a row in a table with its index.
type TableRow[Idx, Row any] struct {
	// The Index of the row.
	Index Idx
	// The Row.
	Row Row
}

// GetTableRow returns the row of the table with the given index.
func GetTableRow[Idx, Row any](ctx context.Context, c Client, node *TableNode[Idx, Row], index Idx) (TableRow[Idx, Row], error) {
	panic("not implemented")
}

// SetTableRow sets the row of the table with the given index.
func SetTableRow[Idx, Row any](ctx context.Context, c Client, node *TableNode[Idx, Row], row TableRow[Idx, Row]) error {
	panic("not implemented")
}
