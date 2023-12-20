package main

import (
	"encoding/binary"
)

const (
	IKCP_RTO_NDL      = 30      // no delay min rto
	IKCP_RTO_MIN      = 100     // normal min rto
	IKCP_RTO_DEF      = 200
	IKCP_RTO_MAX      = 60000
	IKCP_CMD_PUSH     = 81      // cmd: push data
	IKCP_CMD_ACK      = 82      // cmd: ack
	IKCP_CMD_WASK     = 83      // cmd: window probe (ask)
	IKCP_CMD_WINS     = 84      // cmd: window size (tell)
	IKCP_ASK_SEND     = 1       // need to send IKCP_CMD_WASK
	IKCP_ASK_TELL     = 2       // need to send IKCP_CMD_WINS
	IKCP_WND_SND      = 32
	IKCP_WND_RCV      = 128     // must >= max fragment size
	IKCP_MTU_DEF      = 1400
	IKCP_ACK_FAST     = 3
	IKCP_INTERVAL     = 100
	IKCP_OVERHEAD     = 24
	IKCP_DEADLINK     = 20
	IKCP_THRESH_INIT  = 2
	IKCP_THRESH_MIN   = 2
	IKCP_PROBE_INIT   = 7000    // 7 secs to probe window size
	IKCP_PROBE_LIMIT  = 120000  // up to 120 secs to probe window
	IKCP_FASTACK_LIMIT = 5       // max times to trigger fastack
)

func ikcp_encode8u(p []byte, c byte) []byte {
	p[0] = c
	return p[1:]
}

func ikcp_decode8u(p []byte, c *byte) []byte {
	*c = p[0]
	return p[1:]
}

func ikcp_encode16u(p []byte, w uint16) []byte {
	binary.LittleEndian.PutUint16(p, w)
	return p[2:]
}

func ikcp_decode16u(p []byte, w *uint16) []byte {
	*w = binary.LittleEndian.Uint16(p)
	return p[2:]
}

func ikcp_encode32u(p []byte, l uint32) []byte {
	binary.LittleEndian.PutUint32(p, l)
	return p[4:]
}

func ikcp_decode32u(p []byte, l *uint32) []byte {
	*l = binary.LittleEndian.Uint32(p)
	return p[4:]
}

func _imin_(a, b uint32) uint32 {
	if a <= b {
		return a
	}
	return b
}

func _imax_(a, b uint32) uint32 {
	if a >= b {
		return a
	}
	return b
}

func _ibound_(lower, middle, upper uint32) uint32 {
	return _imin_(_imax_(lower, middle), upper)
}

func _itimediff(later, earlier uint32) int32 {
	return int32(later - earlier)
}