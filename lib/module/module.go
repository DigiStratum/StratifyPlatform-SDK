package module

import(
	lib "github.com/DigiStratum/GoLib"
)

// Required: Module public interface
type ModuleIfc interface {
	GetName() string
	GetConfig() lib.ConfigIfc
}

// Optional: Configurable Module public interface
type ConfigurableModuleIfc interface {
	Configure(moduleConfig lib.ConfigIfc) error
}

