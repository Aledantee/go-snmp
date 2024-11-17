package snmp

// LeafNode represents a node in a MIB tree that contains an Object Identifier (OID) and its associated SNMP value.
// May be used to return agent values and/or set a value.
type LeafNode struct {
	// OId is the Object Identifier of the leaf node.
	OId OID
	// Value is the current or value to set at the leaf node.
	Value Value
}

func (l LeafNode) String() string {
	return l.OId.String() + " = " + l.Value.String()
}
