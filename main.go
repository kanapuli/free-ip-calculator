package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("Please specify the starting ipadress and subnet range \n For Example: ./ipcalc 192.168.23.34 29")
	}
	usedIPList := readUsedIPs()

	networkSize, _ := strconv.Atoi(args[2])
	allowedNetworkSize := int(math.Pow(float64(2), float64(32-networkSize)))
	fmt.Printf("Total Allowed Hosts are %d \n", allowedNetworkSize-2)

	ipParts := strings.Split(args[1], ".")
	if len(ipParts) != 4 {
		fmt.Println("The given ipaddress is not in proper format")
		os.Exit(1)
	}
	lastIPPart, _ := strconv.Atoi(ipParts[3])
	var ipPool []string
	for i := 1; i <= allowedNetworkSize; i++ {
		if (lastIPPart + i) > 255 {
			break
		}
		poolAddress := fmt.Sprintf("%s.%s.%s.%v", ipParts[0], ipParts[1], ipParts[2], lastIPPart+i)
		ipPool = append(ipPool, poolAddress)
	}
	fmt.Printf("The Gateway Ip is %s \nThe Broadcast Ip is %s \n", ipPool[0], ipPool[len(ipPool)-1])
	fmt.Println("The Usable free Ips are")
	//remove the first and last element
	ipPool = ipPool[1:]
	ipPool = ipPool[:len(ipPool)-1]

	for _, allowedIP := range ipPool {
		found := false
		for _, usedIP := range usedIPList {
			if usedIP == allowedIP {
				found = true
				break
			}
		}
		if found == false {
			fmt.Println(allowedIP)
		}
	}
}
func readUsedIPs() []string {
	var usedIPs []string
	ipFile, err := os.Open("ip.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := csv.NewReader(bufio.NewReader(ipFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		usedIPs = append(usedIPs, line[0])
	}
	return usedIPs
}
