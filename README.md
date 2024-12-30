# RFC

Read internet RFC documents from your terminal.
This command line tool fetches RFC documents from _https://datatracker.ietf.org/_ and caches the locally, for faster second time reads.

## Usage

Read an RFC by passing its number to the `rfc` tool and using `less` to paginate:

```bash
$ rfc 1035 | less
```

This loads RFC 1035 (Domain names, implementation and specification) in your terminal.

## Installation

Clone the repo and cd into it:

```bash
$ git clone https://github.com/angelsolaorbaiceta/rfc
$ cd rfc
```

Build and install it:

```bash
$ make
```

Make sure the installation was successful by calling `rfc` without arguments to get its help message:

```bash
$ rfc

RFC Reader

USAGE
	rfc <number>

	Where:
		- <number>: the RFC number you want to read

EXAMPLES
  1 GO_FILES = rfc.go

		Read RFC 1035 (DNS)
		$ rfc 1035 | less

  1 GO_FILES = rfc.go
IMPORTANT RFCS

	1. Core Protocols

	RFC 791  - Internet Protocol (IP)
	RFC 2460 - Internet Protocol, Version 6 (IPv6)
	RFC 8200 - Internet Protocol, Version 6 (IPv6, Updated)
	RFC 768 - User Datagram Protocol (UDP)
	RFC 792 - Internet Control Message Protocol (ICMP)
	RFC 4443 - Internet Control Message Protocol (ICMPv6)
	RFC 793 - Transmission Control Protocol (TCP)
	RFC 9293 - Transmission Control Protocol (TCP) (Updated)

	2. Routing Protocols

	RFC 4271 - A Border Gateway Protocol 4 (BGP-4)
	RFC 4760 - Multiprotocol Extensions for BGP-4

	3. Address Resolution

	RFC 826 - An Ethernet Address Resolution Protocol
	RFC 4861 - Neighbor Discovery for IP version 6 (IPv6)

	4. Domain Name System

	RFC 1034 - Domain Names - Concepts and Facilities
	RFC 1035 - Domain Names. Implementation and Specification

	5. Dynamic Host Configuration

	RFC 2131 - Dynamic Host Configuration Protocol (DHCP)
	RFC 8415 - Dynamic Host Configuration Protocol for IPv6 (DHCPv6)

	6. Application Layer Protocols

	RFC 9051 - Internet Message Access Protocol (IMAP) - Version 4rev2
	RFC 1939 - Post Office Protocol - Version 3 (POP3)
	RFC 3501 - Internet Message Access Protocol - Version 4rev1 (IMAP4rev1)
	RFC 5321 - Simple Mail Transfer Protocol (SMTP)
	RFC 5322 - Internet Message Format
	RFC 2616 - Hypertext Transfer Protocol -- HTTP/1.1
	RFC 9113 - HTTP/2 and HPACK
	RFC 9114 - HTTP/3
	RFC 9204 - QPACK: Field Compression for HTTP/3
	RFC 9218 - Extensible Prioritization Scheme for HTTP
```
