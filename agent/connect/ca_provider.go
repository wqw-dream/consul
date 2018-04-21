package connect

import (
	"crypto/x509"

	"github.com/hashicorp/consul/agent/structs"
)

// CAProvider is the interface for Consul to interact with
// an external CA that provides leaf certificate signing for
// given SpiffeIDServices.
type CAProvider interface {
	ActiveRoot() (*structs.CARoot, error)
	ActiveIntermediate() (*structs.CARoot, error)
	GenerateIntermediate() (*structs.CARoot, error)
	Sign(*SpiffeIDService, *x509.CertificateRequest) (*structs.IssuedCert, error)
	//SignCA(*x509.CertificateRequest) (*structs.IssuedCert, error)
	Teardown() error
}
