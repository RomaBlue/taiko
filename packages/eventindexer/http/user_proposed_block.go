package http

import (
	"net/http"

	"github.com/cyberhorsey/webutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer"
)

func (srv *Server) UserProposedBlock(c echo.Context) error {
	delegate, err := srv.addressDelegater.ProposerToDelegate(nil, common.HexToAddress(c.QueryParam("address")))
	if err != nil {
		return webutils.LogAndRenderErrors(c, http.StatusUnprocessableEntity, err)
	}

	if delegate.Delegate.Hex() == "" {
		return c.JSON(http.StatusOK, &galaxeAPIResponse{
			Data: galaxeData{
				IsOK: false,
			},
		})
	}

	event, err := srv.eventRepo.FirstByAddressAndEventName(
		c.Request().Context(),
		delegate.Delegate.Hex(),
		eventindexer.EventNameBlockProposed,
	)
	if err != nil {
		return webutils.LogAndRenderErrors(c, http.StatusUnprocessableEntity, err)
	}

	var found bool = false

	if event != nil {
		found = true
	}

	return c.JSON(http.StatusOK, &galaxeAPIResponse{
		Data: galaxeData{
			IsOK: found,
		},
	})
}
