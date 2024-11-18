package snmp

import "context"

// Client is the interface which exposes the methods to interact with an SNMP agent.
// Each instance of a Client is associated with a specific agent and authentication credentials.
type Client interface {
	// Get performs a GET request on the agent.
	// The provided OID must be a leaf node.
	Get(ctx context.Context, oid OID) (ScalarValue, error)
	// Set performs a SET request on the agent.
	// The OID of the provided ScalarValue must be a leaf node.
	Set(ctx context.Context, value ScalarValue) error
	// Walk performs a walk of the subtree rooted at the provided OID and calls the provided
	// callback function for each leaf node.
	// The implementation may use GETBULK requests to optimize the walk, but must fall back to
	// GETNEXT requests if the agent does not support GETBULK. The client may cache the GETBULK
	// capability of the agent for future walks.
	//
	// The walk is aborted if:
	// 	- the client receives an error response from the agent or the agent does not respond
	//	- the callback returns an error
	// 	- the context is cancelled
	// In the first two cases, the error is returned by Walk. In the last case, the error is context.Canceled.
	Walk(ctx context.Context, oid OID, callback func(ScalarValue) error) error
}
