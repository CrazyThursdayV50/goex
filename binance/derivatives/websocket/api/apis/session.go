package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/session"
)

type Session struct{ api *API }

func (api *API) Session() *Session {
	return &Session{api}
}

func (s *Session) Logon(ctx context.Context) (result *session.LogonResultData, err error) {
	var data session.LogonData
	s.api.Sign(&data)
	params := session.Logon()
	params.Params = &data
	err = request(ctx, s.api, params, &result)
	return
}

func (s *Session) Status(ctx context.Context) (result *session.StatusResultData, err error) {
	params := session.Status()
	err = request(ctx, s.api, params, &result)
	return
}

func (s *Session) Logout(ctx context.Context) (result *session.LogoutResultData, err error) {
	params := session.Logout()
	err = request(ctx, s.api, params, &result)
	return
}
