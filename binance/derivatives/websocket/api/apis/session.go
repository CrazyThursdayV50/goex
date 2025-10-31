package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/session"
)

type Session struct{ api *API }

func (api *API) Session() *Session {
	return &Session{api}
}

func (s *Session) Logon(ctx context.Context, data *session.LogonData) (result *session.LogonResultData, err error) {
	req := session.Logon()
	req.Params, err = s.api.Sign(data)
	if err != nil {
		return
	}

	err = request(ctx, s.api.api, req, &result)
	return
}

func (s *Session) Status(ctx context.Context) (result *session.StatusResultData, err error) {
	req := session.Status()
	err = request(ctx, s.api.api, req, &result)
	return
}

func (s *Session) Logout(ctx context.Context) (result *session.LogoutResultData, err error) {
	req := session.Logout()
	err = request(ctx, s.api.api, req, &result)
	return
}
