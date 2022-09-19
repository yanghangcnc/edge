package edge

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	CentOSDefaultDNSConfigPath = "/etc/resolv.conf"
)

type Nic struct {
	NicName string            `json:"nicName" describe:"网卡名称"`
	IPAddr  string            `json:"IPAddr" describe:"网卡IP"`
	NetMask string            `json:"netMask" describe:"子网掩码"`
	Gateway string            `json:"gateway" describe:"网关"`
	Dns     map[string]string `json:"dns" describe:"DNS服务器"`
	Status  string            `json:"status" describe:"网卡状态:UP & DOWN"`
}


// GetDNS is used to get the configuration of the dns server.
// The parameter 'path' is optional.
// If you don't specify its value, it defaults to "/etc/resolv.conf".
func GetDNS(path string) (map[string]string, error) {
	var (
		res = make([]string, 0)
		dns = make(map[string]string)
	)

	if path == "" {
		path = CentOSDefaultDNSConfigPath
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		slice := strings.Split(line, " ")
		if len(slice) < 2 {
			continue
		}
		if slice[0] == "nameserver" {
			d := strings.TrimSpace(strings.Join(slice[1:], ""))
			res = append(res, d)
		}
	}

	for i, v := range res {
		dns["DNS"+strconv.Itoa(i+1)] = v
	}
	return dns, nil
}

// GetNicStatus is used to get the status of the network card.
// It will return UP, DOWN or UNKNOWN.
func GetNicStatus(name string) (string, error) {
	var cmd = "ip addr show " + name + " | grep 'state' | head -n1 | awk '{print $9}"
	buf, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	res := strings.Trim(string(buf), "\n")
	return res, nil
}

func EchoHello()  {
	fmt.Println("hello edge")
}

func GetNicList() {

}

func GetIPAddr() {

}

func GetPrefix() {

}

func GetGateway() {

}

func GetNetMask() {

}

func GetBootProto() {

}

