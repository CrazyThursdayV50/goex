package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models"
	"github.com/CrazyThursdayV50/goex/infra/iface"
	"github.com/CrazyThursdayV50/pkgo/builtin"
	"github.com/CrazyThursdayV50/pkgo/json"
)

type UnWrapper[T any] = builtin.UnWrapper[T]

func request[Data any](
	ctx context.Context,
	api iface.WebsocketAPI[*models.WsAPIRequest, *models.WsAPIResult],
	params *models.WsAPIRequest,
	result *Data,
) error {
	res, err := api.Send(ctx, params)
	if err != nil {
		return err
	}

	err = json.JSON().Unmarshal(res.Unwrap().Data(), result)
	if err != nil {
		return err
	}

	return nil
}

// func (api *API) GetResult(ctx context.Context, id string) *models.WsAPIResult {
// 	ctx, cancel := context.WithTimeout(ctx, api.resultTimeout)
// 	defer cancel()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return nil

// 		default:
// 			if result := api.resultMap.Get(id).Unwrap(); result != nil {
// 				return result
// 			}

// 			time.Sleep(time.Millisecond)
// 		}
// 	}
// }

// func (api *API) reqID() string {
// 	return uuid.New().String()
// }

// func (api *API) Sign(signerData iface.SignerData) {
// 	signerData.SetAPIKEY(api.api.APIKEY())
// 	signerData.SetTimestamp()
// 	paramsmap := signerData.Map()
// 	signature := sign.Ed25519(paramsmap, api.secretKey)
// 	signerData.SetSignature(signature)
// }
