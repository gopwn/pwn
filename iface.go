package pwn

import (
	"net"
)

type IFace struct {
	Index          int
	MTU            int
	Name           string
	HardwareAddr   net.HardwareAddr
	Flags          net.Flags
	Addrs          []net.Addr
	MulticastAddrs []net.Addr
}

// Returns all ifaces with Addrs and MulticastAddrs
func GetAllIfaceAddrs() ([]IFace, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return []IFace{}, err
	}

	IFaces := []IFace{}

	for i := 0; i < len(ifaces); i++ {
		// Create a new IFaces instance inline and append it to the IFaces slice
		IFaces = append(IFaces, IFace{
			Index:        ifaces[i].Index,
			MTU:          ifaces[i].MTU,
			Name:         ifaces[i].Name,
			HardwareAddr: ifaces[i].HardwareAddr,
			Flags:        ifaces[i].Flags,
		})

		// now you can add the other fields, because append grows the slice

		// if there is an error, addrs will be returned as its zero value
		// aka nil, so the "error checking" was pointless.
		addrs, _ := ifaces[i].Addrs()
		IFaces[i].Addrs = addrs

		multiAddrs, _ := ifaces[i].MulticastAddrs()
		IFaces[i].MulticastAddrs = multiAddrs
	}

	return IFaces, nil
}

// returns given iface
func getIfaceAddrs() error {
	// TODO
	return nil
}
