package processor

import (
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"sync"
)

type SFChannel struct {
	In chan *sfgo.SysFlow
}

type SFProcessor interface {
	//Process(record <-chan *sfgo.SysFlow, wg *sync.WaitGroup)
	Process(record interface{}, wg *sync.WaitGroup)
	SetOutChan(ch interface{})
	Init(conf map[string]string) error
	Cleanup()
}
