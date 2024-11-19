package snmp

// ClientV2cAuthOptions contains the authentication options for creating an SNMPv2c client.
type ClientV2cAuthOptions struct {
	// Community is the SNMPv2c community string to use for authentication.
	Community string
}
