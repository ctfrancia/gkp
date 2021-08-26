package cli

import (
	"flag"
)

// Ports defines the values stored for port killing
type Ports struct {
	Ports     string
	PortRange string
}

// PortsToKill returns a pointer to the Ports struct
func PortsToKill() *Ports {
	p := new(Ports)

	flag.StringVar(&p.Ports, "p", "", "port(s) wished to kill eg: `3000 3005` or `3000`")
	flag.StringVar(&p.PortRange, "r", "", "range of ports wished to terminate")
	flag.Parse()

	return p
}
