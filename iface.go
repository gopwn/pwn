package pwn

import (
	"net"
	"fmt"
)


type IFace struct {
	Index	int
	MTU 	int
	Name 	string
	HardwareAddr net.HardwareAddr
	Flags 	net.Flags
	Addrs 	[]net.Addr
	MulticastAddrs 	[]net.Addr
}

// Returns all ifaces with Addrs and MulticastAddrs
func GetAllIfaceAddrs() ([]IFace, error) {
	tmp_iface, err := net.Interfaces()
	if err != nil {
		return []IFace{}, err
	}

	IFace := []IFace{}

	for i := 0; i < len(tmp_iface); i++ {
		fmt.Println("I: ", i)
		fmt.Printf("len(tmp_iface) %d\n", len(tmp_iface))
		fmt.Printf("cap(tmp_iface) %d\n", cap(tmp_iface))
		fmt.Println(tmp_iface[i].Addrs())
		fmt.Println(tmp_iface[i].MulticastAddrs())

		IFace[i].Index = tmp_iface[i].Index
		IFace[i].MTU = tmp_iface[i].MTU
		IFace[i].Name = tmp_iface[i].Name
		IFace[i].HardwareAddr = tmp_iface[i].HardwareAddr
		IFace[i].Flags = tmp_iface[i].Flags

		addrs, err := tmp_iface[i].Addrs()
		if err != nil {
			IFace[i].Addrs = nil
		}

		IFace[i].Addrs = addrs

		multiAddrs, err := tmp_iface[i].MulticastAddrs()
		if err != nil {
			IFace[i].MulticastAddrs = nil
		}

		IFace[i].MulticastAddrs = multiAddrs


	}

	return IFace, nil
}

//returns given iface
func getIfaceAddrs() (error) {
	// TODO
	return nil
}

func main() {
	GetAllIfaceAddrs()
}
