package peers

import (
	"github.com/gotd/td/constant"
)

type peerIDSet struct {
	m map[constant.TDLibPeerID]struct{}
}

func newPeerIDSet() peerIDSet {
	return peerIDSet{
		m: make(map[constant.TDLibPeerID]struct{}),
	}
}

func (p *peerIDSet) add(ids ...constant.TDLibPeerID) {
	for _, id := range ids {
		p.m[id] = struct{}{}
	}
}

func (p *peerIDSet) delete(ids ...constant.TDLibPeerID) {
	for _, id := range ids {
		delete(p.m, id)
	}
}

func (p *peerIDSet) has(id constant.TDLibPeerID) bool {
	_, ok := p.m[id]
	return ok
}
