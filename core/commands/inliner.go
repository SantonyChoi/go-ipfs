package commands

import (
	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
	cid "gx/ipfs/QmYVNvtQkeZ6AKSwDrjQTs432QtL6umrrK41EBq3cu7iSP/go-cid"
)

type Inliner struct {
	base  cid.Builder
	limit int
}

func (p Inliner) GetCodec() uint64 {
	return p.base.GetCodec()
}

func (p Inliner) WithCodec(c uint64) cid.Builder {
	return Inliner{p.base.WithCodec(c), p.limit}
}

func (p Inliner) Sum(data []byte) (*cid.Cid, error) {
	if len(data) > p.limit {
		return p.base.Sum(data)
	}
	return cid.V1Builder{Codec: p.base.GetCodec(), MhType: mh.ID}.Sum(data)
}
