package processor

import (
	"coolscanner/pkg/models"
	"coolscanner/pkg/protobuf"
)

type Scanner interface {
	Process(*models.SystemData) ([]models.Problem, error)
}

type Processor struct {
	scanners map[string]Scanner
}

func New() *Processor {
	return &Processor{
		scanners: make(map[string]Scanner),
	}
}

func (p *Processor) AddScanner(n string, s Scanner) {
	p.scanners[n] = s
}

func (p *Processor) Process(data *protobuf.SystemInfo) ([]models.Problem, error) {
	// TODO: complicate this
	problems := make([]models.Problem, 0)
	for _, scanner := range p.scanners {
		p, err := scanner.Process(nil)
		if err != nil {
			// TODO: multi-error
			return nil, err
		}
		problems = append(problems, p...)
	}
	return problems, nil
}
