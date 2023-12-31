package mock

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var SuccessHeader = [32]byte{0x1}

type HeaderSyncer struct {
	Fail bool
}

func (h *HeaderSyncer) GetCrossChainBlockHash(opts *bind.CallOpts, blockId uint64) ([32]byte, error) {
	if h.Fail {
		return [32]byte{}, errors.New("fail")
	}

	return SuccessHeader, nil
}
