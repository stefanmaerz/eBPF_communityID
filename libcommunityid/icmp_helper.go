package gommunityid

/*
import (
	"fmt"
	"github.com/google/gopacket/layers"
)
*/

var icmpv4PortEquivalents = map[uint8]uint8{
	8: 0,
	0: 8,
	13: 14,
	14: 13,
	15: 16,
	16: 15,
	10: 9,
	9: 10,
	17: 18,
	18: 17,
}

var icmpv6PortEquivalents = map[uint8]uint8{
	128: 129,
	129: 128,
	133: 134,
	134: 133,
	135: 136,
	136: 135,
	130: 131,
	131: 130,
	144: 145,
	145: 144,
}

// GetICMPv4PortEquivalents returns ICMPv4 codes mapped back to pseudo port
// numbers, as well as a bool indicating whether a communication is one-way.
func GetICMPv4PortEquivalents(p1, p2 uint8) (uint16, uint16, bool) {

	if val, ok := icmpv4PortEquivalents[p1]; ok {
		return uint16(p1), uint16(val), false
	}
	return uint16(p1), uint16(p2), true
}

// GetICMPv6PortEquivalents returns ICMPv6 codes mapped back to pseudo port
// numbers, as well as a bool indicating whether a communication is one-way.
func GetICMPv6PortEquivalents(p1, p2 uint8) (uint16, uint16, bool) {
	if val, ok := icmpv6PortEquivalents[p1]; ok {
		return uint16(p1), uint16(val), false
	}
	return uint16(p1), uint16(p2), true
}
