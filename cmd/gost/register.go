package main

import (
	// Register connectors
	_ "github.com/liukeqqs/x2-master/connector/direct"
	_ "github.com/liukeqqs/x2-master/connector/forward"
	_ "github.com/liukeqqs/x2-master/connector/http"
	_ "github.com/liukeqqs/x2-master/connector/http2"
	_ "github.com/liukeqqs/x2-master/connector/relay"
	_ "github.com/liukeqqs/x2-master/connector/serial"
	_ "github.com/liukeqqs/x2-master/connector/sni"
	_ "github.com/liukeqqs/x2-master/connector/socks/v4"
	_ "github.com/liukeqqs/x2-master/connector/socks/v5"
	_ "github.com/liukeqqs/x2-master/connector/ss"
	_ "github.com/liukeqqs/x2-master/connector/ss/udp"
	_ "github.com/liukeqqs/x2-master/connector/sshd"
	_ "github.com/liukeqqs/x2-master/connector/tcp"
	_ "github.com/liukeqqs/x2-master/connector/tunnel"
	_ "github.com/liukeqqs/x2-master/connector/unix"

	// Register dialers
	_ "github.com/liukeqqs/x2-master/dialer/direct"
	_ "github.com/liukeqqs/x2-master/dialer/dtls"
	_ "github.com/liukeqqs/x2-master/dialer/ftcp"
	_ "github.com/liukeqqs/x2-master/dialer/grpc"
	_ "github.com/liukeqqs/x2-master/dialer/http2"
	_ "github.com/liukeqqs/x2-master/dialer/http2/h2"
	_ "github.com/liukeqqs/x2-master/dialer/http3"
	_ "github.com/liukeqqs/x2-master/dialer/http3/wt"
	_ "github.com/liukeqqs/x2-master/dialer/icmp"
	_ "github.com/liukeqqs/x2-master/dialer/kcp"
	_ "github.com/liukeqqs/x2-master/dialer/mtcp"
	_ "github.com/liukeqqs/x2-master/dialer/mtls"
	_ "github.com/liukeqqs/x2-master/dialer/mws"
	_ "github.com/liukeqqs/x2-master/dialer/obfs/http"
	_ "github.com/liukeqqs/x2-master/dialer/obfs/tls"
	_ "github.com/liukeqqs/x2-master/dialer/pht"
	_ "github.com/liukeqqs/x2-master/dialer/quic"
	_ "github.com/liukeqqs/x2-master/dialer/serial"
	_ "github.com/liukeqqs/x2-master/dialer/ssh"
	_ "github.com/liukeqqs/x2-master/dialer/sshd"
	_ "github.com/liukeqqs/x2-master/dialer/tcp"
	_ "github.com/liukeqqs/x2-master/dialer/tls"
	_ "github.com/liukeqqs/x2-master/dialer/udp"
	_ "github.com/liukeqqs/x2-master/dialer/unix"
	_ "github.com/liukeqqs/x2-master/dialer/ws"

	// Register handlers
	_ "github.com/liukeqqs/x2-master/handler/auto"
	_ "github.com/liukeqqs/x2-master/handler/dns"
	_ "github.com/liukeqqs/x2-master/handler/file"
	_ "github.com/liukeqqs/x2-master/handler/forward/local"
	_ "github.com/liukeqqs/x2-master/handler/forward/remote"
	_ "github.com/liukeqqs/x2-master/handler/http"
	_ "github.com/liukeqqs/x2-master/handler/http2"
	_ "github.com/liukeqqs/x2-master/handler/http3"
	_ "github.com/liukeqqs/x2-master/handler/metrics"
	_ "github.com/liukeqqs/x2-master/handler/redirect/tcp"
	_ "github.com/liukeqqs/x2-master/handler/redirect/udp"
	_ "github.com/liukeqqs/x2-master/handler/relay"
	_ "github.com/liukeqqs/x2-master/handler/serial"
	_ "github.com/liukeqqs/x2-master/handler/sni"
	_ "github.com/liukeqqs/x2-master/handler/socks/v4"
	_ "github.com/liukeqqs/x2-master/handler/socks/v5"
	_ "github.com/liukeqqs/x2-master/handler/ss"
	_ "github.com/liukeqqs/x2-master/handler/ss/udp"
	_ "github.com/liukeqqs/x2-master/handler/sshd"
	_ "github.com/liukeqqs/x2-master/handler/tap"
	_ "github.com/liukeqqs/x2-master/handler/tun"
	_ "github.com/liukeqqs/x2-master/handler/tunnel"
	_ "github.com/liukeqqs/x2-master/handler/unix"

	// Register listeners
	_ "github.com/liukeqqs/x2-master/listener/dns"
	_ "github.com/liukeqqs/x2-master/listener/dtls"
	_ "github.com/liukeqqs/x2-master/listener/ftcp"
	_ "github.com/liukeqqs/x2-master/listener/grpc"
	_ "github.com/liukeqqs/x2-master/listener/http2"
	_ "github.com/liukeqqs/x2-master/listener/http2/h2"
	_ "github.com/liukeqqs/x2-master/listener/http3"
	_ "github.com/liukeqqs/x2-master/listener/http3/h3"
	_ "github.com/liukeqqs/x2-master/listener/http3/wt"
	_ "github.com/liukeqqs/x2-master/listener/icmp"
	_ "github.com/liukeqqs/x2-master/listener/kcp"
	_ "github.com/liukeqqs/x2-master/listener/mtcp"
	_ "github.com/liukeqqs/x2-master/listener/mtls"
	_ "github.com/liukeqqs/x2-master/listener/mws"
	_ "github.com/liukeqqs/x2-master/listener/obfs/http"
	_ "github.com/liukeqqs/x2-master/listener/obfs/tls"
	_ "github.com/liukeqqs/x2-master/listener/pht"
	_ "github.com/liukeqqs/x2-master/listener/quic"
	_ "github.com/liukeqqs/x2-master/listener/redirect/tcp"
	_ "github.com/liukeqqs/x2-master/listener/redirect/udp"
	_ "github.com/liukeqqs/x2-master/listener/rtcp"
	_ "github.com/liukeqqs/x2-master/listener/rudp"
	_ "github.com/liukeqqs/x2-master/listener/serial"
	_ "github.com/liukeqqs/x2-master/listener/ssh"
	_ "github.com/liukeqqs/x2-master/listener/sshd"
	_ "github.com/liukeqqs/x2-master/listener/tap"
	_ "github.com/liukeqqs/x2-master/listener/tcp"
	_ "github.com/liukeqqs/x2-master/listener/tls"
	_ "github.com/liukeqqs/x2-master/listener/tun"
	_ "github.com/liukeqqs/x2-master/listener/udp"
	_ "github.com/liukeqqs/x2-master/listener/unix"
	_ "github.com/liukeqqs/x2-master/listener/ws"
)
