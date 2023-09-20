package main

import (
	"log"

	device_manager "kubevirt.io/kubevirt/pkg/virt-handler/device-manager"
)

func main() {
	deviceName := "kvm"
	devicePath := "/dev/kvm"
	maxDevices := 100
	device := device_manager.NewGenericDevicePlugin(deviceName, devicePath, maxDevices, "", (deviceName != "kvm"))
	stop := make(chan struct{})
	if err := device.Start(stop); err != nil {
		log.Fatal(err)
	}
}
