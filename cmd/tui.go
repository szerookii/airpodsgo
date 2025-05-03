package main

import (
	"fmt"
	"os"
	"time"
	"tinygo.org/x/bluetooth"

	bubbletea "github.com/charmbracelet/bubbletea"
	"github.com/szerookii/airpodsgo/airpods"
)

var (
	devicesMap = make(map[string]airpods.PairedModeData)
	adapter    = bluetooth.DefaultAdapter
)

type tickMsg struct{}

type model struct {
	devices []airpods.PairedModeData
	errs    []error
}

func initialModel() model {
	return model{devices: []airpods.PairedModeData{}, errs: nil}
}

func tick() bubbletea.Cmd {
	return bubbletea.Tick(time.Second, func(time.Time) bubbletea.Msg {
		return tickMsg{}
	})
}

func (m model) Init() bubbletea.Cmd {
	return tick()
}

func (m model) Update(msg bubbletea.Msg) (bubbletea.Model, bubbletea.Cmd) {
	switch msg := msg.(type) {
	case bubbletea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, bubbletea.Quit
		}
	case tickMsg:
		var list []airpods.PairedModeData
		for _, d := range devicesMap {
			list = append(list, d)
		}
		m.devices = list
		m.errs = make([]error, len(list))
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
	s := "Detected AirPods:\n\n"

	for i, d := range m.devices {
		if m.errs != nil && m.errs[i] != nil {
			s += fmt.Sprintf("  #%d: %s\n    Error: %s\n", i+1, d.DeviceModel.String(), m.errs[i])
			continue
		}

		s += fmt.Sprintf(
			"  #%d: %s\n    Case: %s (%s) | Left: %s (%s) | Right: %s (%s)\n",
			i+1,
			d.DeviceModel.String(),
			d.BatteryIndication.CaseBatteryLevel(),
			chargingStatus(d.BatteryIndication.CaseCharging()),
			d.BatteryIndication.LeftLevel(),
			chargingStatus(d.BatteryIndication.LeftCharging()),
			d.BatteryIndication.RightLevel(),
			chargingStatus(d.BatteryIndication.RightCharging()),
		)
	}

	s += "\n[press q to quit]\n"

	return s
}

func chargingStatus(charging bool) string {
	if charging {
		return "charging"
	}

	return "not charging"
}

func main() {
	if err := adapter.Enable(); err != nil {
		fmt.Printf("Error enabling adapter: %v\n", err)
		os.Exit(1)
	}

	go adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		if len(device.ManufacturerData()) <= 0 {
			return
		}

		manufacturerData := device.ManufacturerData()[0]

		if data, err := airpods.DecodePairedModeData(manufacturerData.Data); err == nil {
			addr := device.Address.String()
			devicesMap[addr] = data
		}
	})

	p := bubbletea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		os.Exit(1)
	}
}
