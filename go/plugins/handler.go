package plugins

import (
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// SFHandler defines the SysFlow handler interface.
type SFHandler interface {
	IsEntityEnabled() bool
	HandleHeader(hdr *sfgo.SFHeader) error
	HandleContainer(hdr *sfgo.SFHeader, cont *sfgo.Container) error
	HandleProcess(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process) error
	HandleFile(hdr *sfgo.SFHeader, cont *sfgo.Container, file *sfgo.File) error
	HandleNetFlow(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, nf *sfgo.NetworkFlow) error
	HandleFileFlow(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file *sfgo.File, ff *sfgo.FileFlow) error
	HandleFileEvt(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file1 *sfgo.File, file2 *sfgo.File, fe *sfgo.FileEvent) error
	HandleProcEvt(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, pe *sfgo.ProcessEvent) error
	Init(conf map[string]interface{}) error
	RegisterChannel(pc SFPluginCache)
	RegisterHandler(hc SFHandlerCache)
	SetOutChan(ch []interface{})
	Cleanup()
}
