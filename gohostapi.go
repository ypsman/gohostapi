// testapi
package main

import (
	"encoding/json"
	"net/http"
	"os"
	"os/exec"

	//"time"
	//"fmt"
	//"reflect"
)

type MyIface struct {
	IfaceName string
	IfaceHAW  string
	IfaceIPv4 string
}

type ServerInfos struct {
	Name      string
	Completed bool
	Uptime    string
	MyIface
}

type Server []ServerInfos

func main() {
	http.HandleFunc("/serverinfos", InfoIndex)
	http.ListenAndServe(":8080", nil)
}

func InfoIndex(w http.ResponseWriter, r *http.Request) {
	servers := Server{
		ServerInfos{
			Name:      getMyHostname(),
			Completed: true,
			Uptime:    getUptime(),
			MyIface: MyIface{
				IfaceName: "foobarName",
				IfaceHAW:  "foobarHAW",
				IfaceIPv4: "foobarIPv4",
			}},
	}

	json.NewEncoder(w).Encode(servers)
}

func getUptime() string {
	out, err := exec.Command("uptime", "-ps").Output()
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Online since:", out)
	//fmt.Println(reflect.TypeOf(out))
	xx := string(out)
	return xx
}

func getMyHostname() string {
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}
	return name
}

// ### Tests Needed
/*
func getMyIfaces() {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		addrs, err := inter.Addrs()
		if err != nil {
			panic(err)
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					c := MyIface{inter.Name, ipnet.IP.String(), inter.HardwareAddr.String()}
					fmt.Println(c.IfaceName, "::", c.IfaceHAW, "::", c.IfaceIPv4)
				}
			}
		}
	}
}
*/
