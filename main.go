package main

import (
	"log"

	device_manager "kubevirt.io/kubevirt/pkg/virt-handler/device-manager"
)

func main() {
	const (
		deviceName = "kvm"
		devicePath = "/dev/kvm"

		// From KubeVirt:

		// TODO: the Device Plugin API does not allow for infinitely available (shared) devices
		// so the current approach is to register an arbitrary number.
		// This should be deprecated if the API allows for shared resources in the future
		maxDevices = 1000
	)
	device := device_manager.NewGenericDevicePlugin(deviceName, devicePath, maxDevices, "", (deviceName != "kvm"))
	stop := make(chan struct{})
	if err := device.Start(stop); err != nil {
		log.Fatal(err)
	}
}
