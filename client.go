package snmp

import (
	"context"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"net/netip"
	"time"
)

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

// ClientOptions contains the options for a Client.
type ClientOptions struct {
	// Target is the address and port of the agent.
	Target netip.AddrPort
	// Auth contains the authentication options for the client.
	// The specific type of ClientAuthOptions determines the version of SNMP to use.
	Auth ClientAuthOptions

	// Timeout is the maximum time the client will wait for a response from the agent.
	Timeout time.Duration
	// NumRetries is the number of times the client will retry a request if it does not receive a response
	// from the agent in the Timeout duration.
	NumRetries int
	// BackoffFunc is the function used to calculate the time to wait before retrying a request.
	// The function is called with the number of retries that have been attempted so far.
	BackoffFunc func(retries int) time.Duration

	// TracerProvider to use for creating a tracer for the client.
	// If nil, the global tracer provider will be used.
	TracerProvider trace.TracerProvider
	// MeterProvider with which to create a client-specicic instance of ClientInstruments.
	// If ClientInstruments is non-nil, this field is ignored.
	MeterProvider metric.MeterProvider
	// Instruments of the client to use for metrics. If the metrics should be shared
	// between multiple clients, the same ClientInstruments instance should be passed to each client.
	// If nil, a new instance of ClientInstruments will be created using the MeterProvider.
	// If both MeterProvider and Instruments are nil, no metrics will be recorded.
	Instruments *ClientInstruments
}

// ClientOption is a function that modifies the provided ClientOptions.
type ClientOption func(*ClientOptions)

// ClientAuthOptions contains the authentication options for a Client.
// Exactly one of V2c and V3 must be non-nil.
type ClientAuthOptions struct {
	// V2c contains the authentication options for creating an SNMPv2c client.
	// Implies the use of SNMPv2c.
	V2c *ClientV2cAuthOptions
	// V3 contains the authentication options for creating an SNMPv3 client.
	// Implies the use of SNMPv3.
	V3 *ClientV3AuthOptions
}

// ClientInstruments contains the metrics for one or more Clients.
type ClientInstruments struct {
}

// NewClient creates a new Client with the provided target address using a default set of options.
// The default options may be overridden by providing additional ClientOption functions which modify the options.
func NewClient(target netip.AddrPort, opts ...ClientOption) (Client, error) {
	defaultOpts := ClientOptions{
		Target: target,
		Auth: ClientAuthOptions{
			V2c: &ClientV2cAuthOptions{
				Community: "public",
			},
		},
		Timeout:    3 * time.Second,
		NumRetries: 3,
	}

	WithExponentialBackoff(2)(&defaultOpts)

	for _, opt := range opts {
		opt(&defaultOpts)
	}

	return NewClientWithOptions(defaultOpts)
}

// NewClientWithOptions creates a new Client with the provided options.
// The target address and port must be specified in the options.
func NewClientWithOptions(opts ClientOptions) (Client, error) {
	return nil, nil
}

// WithTimeout is a ClientOption that sets the timeout for the client.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(opts *ClientOptions) {
		opts.Timeout = timeout
	}
}

// WithNumRetries is a ClientOption that sets the number of retries for the client.
func WithNumRetries(numRetries int) ClientOption {
	return func(opts *ClientOptions) {
		opts.NumRetries = numRetries
	}
}

// WithBackoffFunc is a ClientOption that sets the backoff function for the client.
func WithBackoffFunc(backoffFunc func(retries int) time.Duration) ClientOption {
	return func(opts *ClientOptions) {
		opts.BackoffFunc = backoffFunc
	}
}

// WithExponentialBackoff is a ClientOption that sets the backoff function
// to an exponential backoff with the provided base.
func WithExponentialBackoff(base int) ClientOption {
	return WithBackoffFunc(func(retries int) time.Duration {
		return time.Duration(base^retries) * time.Second
	})
}

// WithTracerProvider is a ClientOption that sets the tracer provider for the client.
func WithTracerProvider(provider trace.TracerProvider) ClientOption {
	return func(opts *ClientOptions) {
		opts.TracerProvider = provider
	}
}

// WithMeterProvider is a ClientOption that sets the meter provider for the client.
func WithMeterProvider(provider metric.MeterProvider) ClientOption {
	return func(opts *ClientOptions) {
		opts.MeterProvider = provider
	}
}

// WithCommunity is a ClientOption that sets the community string for the client.
// Implies the use of SNMPv2c.
func WithCommunity(community string) ClientOption {
	return func(opts *ClientOptions) {
		opts.Auth.V2c = &ClientV2cAuthOptions{
			Community: community,
		}
	}
}
