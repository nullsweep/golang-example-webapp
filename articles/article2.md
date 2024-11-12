
---

#### **Basic Scans**

**Simple Scan:**
```bash
nmap <target>
```

**Multiple Targets:**
```bash
nmap <target1> <target2> <target3>
```

**Scan a Range:**
```bash
nmap <start-IP>-<end-IP>
```

**Scan a Subnet:**
```bash
nmap <IP/CIDR>
```

**Scan Hosts from a File:**
```bash
nmap -iL <file_with_hosts>
```

**Exclude Hosts from a File:**
```bash
nmap --excludefile <file_with_hosts>
```

**Exclude Specific Hosts:**
```bash
nmap --exclude <host1>,<host2>,<host3>
```

**Scan Random Targets:**
```bash
nmap -iR <number_of_random_hosts>
```

**Scan an Entire Network:**
```bash
nmap <network_address>/<CIDR>
```

**Scan Specific IP Protocols:**
```bash
nmap -sO <target>
```

---

#### **Host Discovery**

**Ping Scan (Discovery Scan):**
```bash
nmap -sn <target>
```

**Disable Ping (Treat all hosts as online):**
```bash
nmap -Pn <target>
```

**ICMP Echo Discovery:**
```bash
nmap -PE <target>
```

**ICMP Timestamp Discovery:**
```bash
nmap -PP <target>
```

**ICMP Address Mask Discovery:**
```bash
nmap -PM <target>
```

**TCP SYN Discovery (on port 443):**
```bash
nmap -PS443 <target>
```

**TCP ACK Discovery (on port 80):**
```bash
nmap -PA80 <target>
```

**UDP Discovery (on port 123):**
```bash
nmap -PU123 <target>
```

**ARP Discovery (local network only):**
```bash
nmap -PR <target>
```

**IPv6 Host Discovery:**
```bash
nmap -6 -sn <target>
```

**Traceroute:**
```bash
nmap --traceroute <target>
```

**Host Discovery with Detailed Output:**
```bash
nmap -sn -v <target>
```

**List Scan (only lists targets to be scanned, no packets sent):**
```bash
nmap -sL <target>
```

**Aggressive Host Discovery (combining multiple techniques):**
```bash
nmap -PE -PP -PS80,443 -PA80,443 -PU123,161 -PR <target>
```

**Scan with Custom Packet Timing:**
```bash
nmap -T<0-5> <target>
```

**Maximum Parallelism:**
```bash
nmap --min-parallelism <number> <target>
nmap --max-parallelism <number> <target>
```

**Max Hostgroup Size:**
```bash
nmap --min-hostgroup <number> <target>
nmap --max-hostgroup <number> <target>
```

---

#### **Port Scanning Techniques**

**TCP SYN Scan (default):**
```bash
nmap -sS <target>
```

**TCP Connect Scan:**
```bash
nmap -sT <target>
```

**UDP Scan:**
```bash
nmap -sU <target>
```

**TCP ACK Scan:**
```bash
nmap -sA <target>
```

**Window Scan:**
```bash
nmap -sW <target>
```

**FIN Scan:**
```bash
nmap -sF <target>
```

**Xmas Scan:**
```bash
nmap -sX <target>
```

**Null Scan:**
```bash
nmap -sN <target>
```

**Idle Scan:**
```bash
nmap -sI <zombie-IP> <target>
```

---

#### **Port Specification and Scan Order**

**Specify Ports:**
```bash
nmap -p <port1>,<port2>,<port3> <target>
```

**Scan Port Range:**
```bash
nmap -p <start_port>-<end_port> <target>
```

**Exclude Ports:**
```bash
nmap --exclude-ports <port1>,<port2> <target>
```

**Fast Scan (100 ports):**
```bash
nmap -F <target>
```

**Scan Top Ports:**
```bash
nmap --top-ports <number> <target>
```

**Randomize Port Order:**
```bash
nmap --randomize-hosts <target>
```

---

#### **Service and Version Detection**

**Service Version Detection:**
```bash
nmap -sV <target>
```

**Aggressive Scan (includes OS detection, version detection, script scanning, and traceroute):**
```bash
nmap -A <target>
```

**Operating System Detection:**
```bash
nmap -O <target>
```

**Enable OS Detection, Version Detection, Script Scanning, and Traceroute:**
```bash
nmap -A <target>
```

---

#### **NSE (Nmap Scripting Engine)**

**Run Default Scripts:**
```bash
nmap -sC <target>
```

**Specify Scripts:**
```bash
nmap --script <script_name> <target>
```

**Run Multiple Scripts:**
```bash
nmap --script <script1>,<script2> <target>
```

**Script Categories:**
```bash
nmap --script <category> <target>
```

**Examples:**
```bash
nmap --script vuln <target>  # Run vulnerability scripts
nmap --script exploit <target>  # Run exploit scripts
nmap --script discovery <target>  # Run discovery scripts
```

---

#### **Output Options**

**Normal Output:**
```bash
nmap -oN <filename> <target>
```

**XML Output:**
```bash
nmap -oX <filename> <target>
```

**Grepable Output:**
```bash
nmap -oG <filename> <target>
```

**All Formats:**
```bash
nmap -oA <basename> <target>
```

**Append Output:**
```bash
nmap --append-output <target>
```

---

#### **Timing and Performance**

**Tuning Scans (1-5, 3 is default):**
```bash
nmap -T<0-5> <target>
```

**Set Max Retries:**
```bash
nmap --max-retries <number> <target>
```

**Set Host Timeout:**
```bash
nmap --host-timeout <time> <target>
```

**Set Scan Delay:**
```bash
nmap --scan-delay <time> <target>
```

**Max Parallelism:**
```bash
nmap --min-parallelism <number> <target>
nmap --max-parallelism <number> <target>
```

---

#### **Advanced Features**

**Detect Firewall:**
```bash
nmap -sA <target>
```

**Evade Firewalls and IDS:**
```bash
nmap -f <target>  # Fragment packets
nmap --data-length <number> <target>  # Append random data
nmap --badsum <target>  # Send packets with bad checksums
```

**Spoof IP Address:**
```bash
nmap -S <spoof-IP> <target>
```

**Decoy Scan:**
```bash
nmap -D <decoy1>,<decoy2>,<decoy3> <target>
```

**Scan Through Proxy:**
```bash
nmap -Pn --proxies <proxy1>,<proxy2> <target>
```

---

#### **Miscellaneous Options**

**Save Scan Results:**
```bash
nmap -oN/-oX/-oG/-oA <file> <target>
```

**Resume Scan:**
```bash
nmap --resume <filename>
```

**IPv6 Scanning:**
```bash
nmap -6 <target>
```

**Scan from a File:**
```bash
nmap -iL <inputfile>
```

**Exclude Hosts:**
```bash
nmap --exclude <host1>,<host2>
```

**Banner Grabbing:**
```bash
nmap -sV --script=banner <target>
```

---