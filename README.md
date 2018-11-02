# free-ip-calculator
The tool can be used to filter the unused ips from the used ip address list
Create a file called ip.csv. Add all the used ipaddress in ip.csv as  comma separated values

# Build Procedure
go build -o ipcalc main.go  

# Example
./ipcalc 192.167.121.34 29     
- 192.167.121.34  -  IP pool first address
- 29 -  Network size

Let the ip.csv contains the following ips, 

```
192.167.121.38,
192.167.121.39,
```

The output will be 

```
Total Allowed Hosts are 6 
The Gateway Ip is 192.167.121.35 
The Broadcast Ip is 192.167.121.42 
The Usable free Ips are
192.167.121.36
192.167.121.37
192.167.121.40
192.167.121.41

```
