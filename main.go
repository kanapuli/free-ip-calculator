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

var usedIPList []string

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("Please specify the starting ipadress and subnet range \n For Example: ./ipcalc 192.168.23.34 29")
	}
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
		usedIPList = append(usedIPList, line[0])
	}
	subnetRange, _ := strconv.Atoi(args[2])
	allowedIPRange := int(math.Pow(float64(2), float64(32-subnetRange)))
	fmt.Printf("Total Allowed Hosts are %d \n", allowedIPRange-2)
	splittedIP := strings.Split(args[1], ".")
	if len(splittedIP) != 4 {
		fmt.Println("The given ipaddress is not in proper format")
		os.Exit(1)
	}
	lastIPValue, _ := strconv.Atoi(splittedIP[3])
	var ipPool []string
	for i := 1; i <= allowedIPRange; i++ {
		if (lastIPValue + i) > 255 {
			break
		}
		poolAddress := fmt.Sprintf("%s.%s.%s.%v", splittedIP[0], splittedIP[1], splittedIP[2], lastIPValue+i)
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
			}
		}
		if found == false {
			fmt.Println(allowedIP)
		}
	}
}
