package module

import(
	cfg "github.com/DigiStratum/GoLib/Config"
)

// Required: Module public interface
type ModuleIfc interface {
	GetName() string
	GetConfig() lib.ConfigIfc
}

// Optional: Configurable Module public interface
type ConfigurableModuleIfc interface {
	Configure(moduleConfig cfg.ConfigIfc) error
}

