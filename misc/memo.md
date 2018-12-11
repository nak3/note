memo
---

## DNS record

DNS record is a mapping record of IP address against domain name.

| record      | explain                                         |
|:------------|--------------------------------------------------|
| A record    | a record points a domain to an IPv4 IP address.  |
| AAAA record | a record points a domain to an IPv6 IP address. (IPV6 is 128 bit = four times IPv4's 32 bit, so 4 As) |
| CNAME       | a record points a domain to on alias. |
| NS record   | a record points a domain to name servers that the domain delegates to. |
| MX record   | a record points a domain to a mail server |
| SOA record  | a record points a domain to an information of its authoritative name server. It stands for Start Of a zone Of Authority. |



## IP address

network, hosts, subnet mask, and so on.

- When CIDR is `/24`, subnet mask is 255.255.255.0. Total IPs are `2^8 = 256`. `254 (=256-2)` are available IPs. -2 is from Broadcast IP and network IP.
- When CIDR is `/16`, subnet mask is 255.255.0.0. Total IPs are `2^16 = 256*256`. `65534 (=65536-2)` are available IPs.



## TCP basic

### connection establishment

- 3 way handshake, used SYN, SYN+ACK, ACK packets

- A server becomes Listen status and waits for open request from a client. The server must bind to and listen at port to open it up for connections (passive open). The client initiates the connection.

This picture is from [wikipedia](https://en.wikipedia.org/wiki/Transmission_Control_Protocol#Connection_establishment)

![tcp establishment](/images/Tcp_state_diagram_fixed_new.svg)


### Basic feature of TCP

- sequence number
- retransmission
- flow control
- Window scaling
