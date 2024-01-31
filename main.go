package main

import (
	"fmt"
	"net"

	"github.com/da0x/golang/olog"
)

type Networks struct {
	Interface string `json:"Interfaces"`
	Ip        string `json:"ip"`
}

func GetAllInterfaces() []Networks {
	ifaces, err := net.Interfaces()
	var nets []Networks
	if err != nil {
		fmt.Print(fmt.Errorf("localAddresses: %+v", err.Error()))
		return nets
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v", err.Error()))
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				// fmt.Printf("%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())

			case *net.IPNet:

				if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {

					if ipnet.IP.To4() != nil {

						n := Networks{
							Ip:        v.IP.String(),
							Interface: i.Name+"  ",
						}
						nets = append(nets, n)

					}

				}

			}
		}

	}

	return nets
}

func main() {
	re := GetAllInterfaces()

	olog.Print(re)

}
