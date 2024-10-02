package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SplitIP(ip string) ([]uint8, error) {
	ipBlocksStr := strings.Split(ip, ".")
	if len(ipBlocksStr) != 4 {

		return []uint8{}, fmt.Errorf("ip format error: %s", ip)
	}

	ipBlocks := make([]uint8, len(ipBlocksStr))

	for i, str := range ipBlocksStr {
		uintVal, _ := strconv.ParseUint(str, 10, 8)
		ipBlocks[i] = uint8(uintVal)
	}

	return ipBlocks, nil
}
