package plugins

// Dynamic plugin function names and types for reflection.
const (
	NameFn  string = "GetName"
	PlugSym string = "Plugin"
)

// SFPluginCache defines an interface for a plugin cache.
type SFPluginCache interface {
	AddProcessor(name string, factory interface{})
	AddHandler(name string, factory interface{})
	AddChannel(name string, factory interface{})
}

// SFPluginFactory defines an abstract factory for plugins.
type SFPluginFactory interface {
	Register(pc SFPluginCache)
}
