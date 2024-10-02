package helpers

import (
	"errors"
	"fmt"
	"github.com/RajabovIlyas/ping-ip/internal/app/constants"
	"github.com/RajabovIlyas/ping-ip/internal/app/ping"
	"github.com/RajabovIlyas/ping-ip/internal/app/utils"
	"log"
	"net"
	"os"
	"sync"
)

type Ping struct {
	ip          []uint8
	ports       []int
	subnetClass string
}

func NewPing() ping.Helper {
	return &Ping{}
}

func (p Ping) CheckIP(ip string) (string, error) {
	_, err := net.LookupIP(ip)
	if err != nil {
		return "", err
	}
	return ip, nil
}

func (p Ping) CheckPort(ip string, port int) error {
	ipWithPort := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", ipWithPort)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func (p Ping) CheckPorts(ip string, ports []int) error {
	for _, port := range ports {
		err := p.CheckPort(ip, port)
		if err != nil {
			continue
		}
		return nil
	}

	return errors.New("No valid ports found")
}

func (p Ping) Run() error {
	ipArg := os.Args[1]

	if len(os.Args) <= 2 {
		p.subnetClass = constants.OneIP
	} else {
		p.subnetClass = os.Args[2]
	}

	ipBlocks, err := utils.SplitIP(ipArg)
	if err != nil {
		log.Printf("Error while splitting IP: %v", err)
		return err
	}
	p.ip = ipBlocks

	openIPs, err := utils.IterateIPs(p.CheckIP, p.ip, utils.ClassChangeToIndex(p.subnetClass))
	if err != nil {
		log.Printf("Error while iterating IPs: %v", err)
		return err
	}

	for _, openIP := range openIPs {
		log.Printf("Open IP: %s\n", openIP)
	}

	ports := os.Args[2:]

	checkPorts, err := utils.ConvertArrayStrToInt(ports)
	if err != nil {
		log.Printf("Error while converting ports: %v", err)
		return err
	}

	if len(checkPorts) < 1 {
		p.ports = constants.DefaultPorts
	} else {
		p.ports = checkPorts
	}

	openPorts := make([]string, 0, len(openIPs))

	var wg sync.WaitGroup

	for _, ip := range openIPs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := p.CheckPorts(ip, p.ports)
			if err != nil {
				return
			}

			openPorts = append(openPorts, ip)
		}()
	}
	wg.Wait()

	for _, openIP := range openPorts {
		log.Printf("Open IP by ports: %s\n", openIP)
	}

	return nil
}
