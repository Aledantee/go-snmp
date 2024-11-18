package snmp

import (
	"context"
	"fmt"
)

// A ScalarNode value may be used to perform type-safe GET/SET requests on a client.
type ScalarNode[V Value] struct {
	// The Name of the scalar node.
	Name string
	// The Description of the scalar node.
	Description string
	// The Object Identifier of the scalar node.
	OID OID
}

// A ScalarValue is a Value associated with an OID.
// Used by a client as a return value for GET requests and parameter for SET requests.
type ScalarValue struct {
	// The OID of the scalar value.
	OID OID
	// The Value of the scalar value.
	Value Value
}

// A ScalarNodeValue is a Value associated with a ScalarNode.
// In contrast to a ScalarValue, a ScalarNodeValue is used by a client to perform type-safe GET/SET requests.
type ScalarNodeValue[V Value] struct {
	// The Node this value is associated with.
	Node *ScalarNode[V]
	// The Value.
	Value V
}

// GetScalar performs a GET request on the client for the provided ScalarNode.
// Used to implement type-safe GET requests.
func GetScalar[V Value](ctx context.Context, c Client, node *ScalarNode[V]) (ScalarNodeValue[V], error) {
	scalarValue, err := c.Get(ctx, node.OID)
	if err != nil {
		return ScalarNodeValue[V]{}, err
	}

	// The agent might return a value of a different type than expected.
	// In this case, we return an error. No attempt is made to convert the value.
	if v, ok := scalarValue.Value.(V); !ok {
		var v V
		//goland:noinspection GoDfaNilDereference
		expectedType := v.Type()

		return ScalarNodeValue[V]{}, fmt.Errorf("unexpected value type in response: expected %q, got %q", expectedType, scalarValue.Value)
	} else {
		return ScalarNodeValue[V]{Node: node, Value: v}, nil
	}
}

// SetScalar performs a SET request on the client for the provided ScalarNodeValue.
// Used to implement type-safe SET requests.
func SetScalar[V Value](ctx context.Context, c Client, node ScalarNodeValue[V]) error {
	return c.Set(ctx, ScalarValue{OID: node.Node.OID, Value: node.Value})
}
