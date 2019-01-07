package pwn

import (
	"reflect"
	"testing"
)

// TODO: implement mock of []net.Interface for testing
/*
	interfaces := []net.Interface{
		{
			Index: 1,
			MTU: 313,
			Name: "test iface",
			HardwareAddr: net.HardwareAddr,
			Flags: net.Flags,
		}
	}
	ifaces, err := getInterfaceAddrs(interfaces)
	if err != nil {
		t.Fatal(err)
	}
*/

func TestGetAllIfaceAddrs(t *testing.T) {
	// currently only test that the function can be called
	ifaces, err := GetInterfaceAddrs()
	if err != nil {
		t.Fatal(err)
	}
	if len(ifaces) < 1 {
		t.Skip("len(ifaces) < 1")
	}

	iface, err := GetIFaceByName(ifaces[0].Name)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iface, ifaces[0]) {
		t.Fatal("iface != faces[0]")
	}
}
