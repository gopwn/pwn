package pwn

import (
	"reflect"
	"testing"
)

// TODO: implement mock of []net.Interface for testing
// not that they work.
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

// Test IFace NOTE: there is a small chance of this failing if a new interface
// appears while this test is running
func TestIFace(t *testing.T) {
	t.Parallel()
	// currently only test that the function can be called
	ifaces, err := GetInterfaceAddrs()
	if err != nil {
		t.Fatal(err)
	}
	if len(ifaces) < 1 {
		t.Skip("len(ifaces) < 1")
	}

	t.Run("GetIFaceByName", func(t *testing.T) {
		t.Parallel()
		iface, err := GetIFaceByName(ifaces[0].Name)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(iface, ifaces[0]) {
			t.Fatal("iface != ifaces[0]")
		}
	})

	t.Run("GetIFaceByIndex", func(t *testing.T) {
		t.Parallel()
		iface, err := GetIFaceByIndex(1)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(iface, ifaces[0]) {
			t.Fatal("iface != ifaces[0]")
		}
	})
}
