package handlers

import (
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

type FlatChannel struct {
	In chan *sfgo.FlatRecord
}
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
	SetOutChan(ch interface{})
	Init(conf map[string]string) error
	Cleanup()
}
