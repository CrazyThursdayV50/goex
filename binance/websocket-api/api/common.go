package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models"
	"github.com/CrazyThursdayV50/goex/binance/websocket-api/signer"
	"github.com/CrazyThursdayV50/pkgo/builtin"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/google/uuid"
)

type UnWrapper[T any] = builtin.UnWrapper[T]

func request[reqData any, resultData any](ctx context.Context, api *API, params *models.WsAPIParams[reqData], dst *resultData) error {
	params.Id = api.reqID()

	data, err := params.BinaryMarshal()
	if err != nil {
		return err
	}

	err = api.client.Send(data)
	if err != nil {
		return err
	}

	result := api.GetResult(ctx, params.Id)
	if result == nil {
		return errors.New("request timeout")
	}

	if result.Status == 200 {
		err = json.JSON().Unmarshal(result.Result, dst)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("request failed with status: %d, error: %s", result.Status, result.Error())
}

func (api *API) GetResult(ctx context.Context, id string) *models.WsAPIResult {
	ctx, cancel := context.WithTimeout(ctx, api.resultTimeout)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return nil

		default:
			if result := api.resultMap.Get(id).Unwrap(); result != nil {
				return result
			}

			time.Sleep(time.Millisecond)
		}
	}
}

func (api *API) reqID() string {
	return uuid.New().String()
}

func (api *API) Sign(signerData signer.SignerData) {
	signerData.SetApiKey(api.apiKey)
	signerData.SetTimestamp()
	paramsmap := signerData.Map()
	signature := signer.SignEd25519(paramsmap, api.secretKey)
	signerData.SetSignature(signature)
}
