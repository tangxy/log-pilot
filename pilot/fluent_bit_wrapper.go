package pilot

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
)

const PILOT_FLUENT_BIT = "fluent-bit"
const FLUENT_BIT_CONF_HOME = "/fluent-bit/etc/conf.d"

var fluentBit *exec.Cmd

type FluentBitPiloter struct {
	name string
}

func NewFluentBitPiloter() (Piloter, error) {
	return &FluentBitPiloter{
		name: PILOT_FLUENT_BIT,
	}, nil
}

func (p *FluentBitPiloter) Start() error {
	if fluentBit != nil {
		return fmt.Errorf(ERR_ALREADY_STARTED)
	}

	log.Info("start fluent-bit")
	fluentBit = exec.Command("/fluent-bit/bin/fluent-bit", "-e","/fluent-bit/bin/out_grafana_loki.so",
	"-c", "/fluent-bit/etc/fluent-bit.conf")
	fluentBit.Stderr = os.Stderr
	fluentBit.Stdout = os.Stdout
	err := fluentBit.Start()
	if err != nil {
		log.Error(err)
	}
	go func() {
		err := fluentBit.Wait()
		if err != nil {
			log.Error(err)
		}
	}()
	return err
}

func (p *FluentBitPiloter) Stop() error {
	return nil
}

func (p *FluentBitPiloter) Reload() error {
	if fluentBit == nil {
		err := fmt.Errorf("fluent-bit have not started")
		log.Error(err)
		return err
	}

	log.Info("reload fluent-bit")
	return nil
}

func (p *FluentBitPiloter) ConfPathOf(container string) string {
	return fmt.Sprintf("%s/%s.conf", FLUENT_BIT_CONF_HOME, container)
}


func (p *FluentBitPiloter) ConfHome() string {
	return FLUENT_BIT_CONF_HOME
}

func (p *FluentBitPiloter) Name() string {
	return p.name
}

func (p *FluentBitPiloter) OnDestroyEvent(container string) error {
	log.Info("refactor in the future!!!")
	return nil
}
