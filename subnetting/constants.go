package subnetting

const (
	// ExitSuccess - standard exit code
	ExitSuccess = 0

	// ExitMissingArgs  - standard exit code
	ExitMissingArgs = 1

	// ExitSubnettingError  - standard exit code
	ExitSubnettingError = 2

	// ErrGeneral - General error formatting string
	ErrGeneral = "Error:%s"

	//ErrMissingArguments - standard error
	ErrMissingArguments = "Error: missing arguments parentCIDR subnetSize"

	// ErrInvalidParentCIDR - standard error
	ErrInvalidParentCIDR = "invalid parent CIDR:%s"

	// ErrInvalidSubnetSize - standard error
	ErrInvalidSubnetSize = "invalid subnet size:%d"

	//MsgIpv4CIDR - standard CIDR format string
	MsgIpv4CIDR = "%d.%d.%d.%d/%d"
)
