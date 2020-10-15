package plugins

// Dynamic plugin function names and types for reflection.
const (
	NameFn    string = "GetName"
	PlugSym   string = "Plugin"
	DriverSym string = "Driver"
)

// SFPluginCache defines an interface for a plugin cache.
type SFPluginCache interface {
	AddDriver(name string, factory interface{})
	AddProcessor(name string, factory interface{})
	AddHandler(name string, factory interface{})
	AddChannel(name string, factory interface{})
}

// SFPluginFactory defines an abstract factory for plugins.
type SFPluginFactory interface {
	Register(pc SFPluginCache)
}
