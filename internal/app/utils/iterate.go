package utils

import (
	"fmt"
	"github.com/RajabovIlyas/ping-ip/internal/app/constants"
)

func ClassChangeToIndex(subnetClass string) int8 {
	switch subnetClass {
	case constants.ClassA:
		return 1
	case constants.ClassB:
		return 2
	case constants.ClassC:
		return 3
	default:
		return 4
	}
}

func GenerateIP(ipBlocks []uint8) string {
	return fmt.Sprintf("%d.%d.%d.%d", ipBlocks[0], ipBlocks[1], ipBlocks[2], ipBlocks[3])
}

func IterateIPs(callback func(ip string) (string, error), ipBlocks []uint8, index int8) ([]string, error) {
	openIPs := []string{}
	if index > 3 || index < 0 {
		result, err := callback(GenerateIP(ipBlocks))
		if err != nil {
			return openIPs, err
		}
		openIPs = append(openIPs, result)
		return openIPs, nil
	}
	for i := 0; i < 255; i++ {
		ipBlocks[index] = uint8(i)
		result, err := IterateIPs(callback, ipBlocks, index+1)
		if err != nil {
			continue
		}
		openIPs = append(openIPs, result...)
	}

	return openIPs, nil
}
