package configuration

import (
	"ServiceTools/comms"
	"ServiceTools/firewall"
	"ServiceTools/health"
	"ServiceTools/register"
)

type Config struct {
	Firewall  firewall.Config
	Health    health.Config
	Register  register.Config
	Transport comms.Config
}
