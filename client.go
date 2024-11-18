package snmp

// Client is the interface which exposes the methods to interact with an SNMP agent.
// Each instance of a Client is associated with a specific agent and authentication credentials.
type Client interface {
	// Get performs a GET request on the agent.
	// The provided OID must be a leaf node.
	Get(oid OID) (ScalarValue, error)
	// Set performs a SET request on the agent.
	// The OID of the provided ScalarValue must be a leaf node.
	Set(value ScalarValue) error
}
