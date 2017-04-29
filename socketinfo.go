// Copyright 2017 Julien Schmidt. All rights reserved.
// Use of this source code is governed by MIT license,
// a copy can be found in the LICENSE file.

package socketinfo

import (
	"syscall"
)

type SocketType uint16

const (
	Stream    SocketType = syscall.SOCK_STREAM    // stream (connection) socket
	Datagram  SocketType = syscall.SOCK_DGRAM     // datagram (connection-less) socket
	Raw       SocketType = syscall.SOCK_RAW       // raw socket
	SeqPacket SocketType = syscall.SOCK_SEQPACKET // sequential packet socket
	/* syscall.SOCK_RDM // reliably-delivered message */
)

type ProtocolFamily uint16

const (
	Unspecified ProtocolFamily = syscall.AF_UNSPEC
	Unix        ProtocolFamily = syscall.AF_UNIX  // Unix domain sockets
	IPv4        ProtocolFamily = syscall.AF_INET  // Internet Protocol v4
	IPv6        ProtocolFamily = syscall.AF_INET6 // Internet Protocol v6
)
