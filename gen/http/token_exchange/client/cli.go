// Code generated by goa v3.16.0, DO NOT EDIT.
//
// tokenExchange HTTP client CLI support package
//
// Command:
// $ goa gen kubernetes-federated-credential-controller/design

package client

import (
	"encoding/json"
	"fmt"
	tokenexchange "kubernetes-federated-credential-controller/gen/token_exchange"
)

// BuildExchangeTokenPayload builds the payload for the tokenExchange
// exchangeToken endpoint from CLI flags.
func BuildExchangeTokenPayload(tokenExchangeExchangeTokenBody string) (*tokenexchange.ExchangeTokenPayload, error) {
	var err error
	var body ExchangeTokenRequestBody
	{
		err = json.Unmarshal([]byte(tokenExchangeExchangeTokenBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"jwt\": \"Vel et ipsam.\",\n      \"namespace\": \"Ex iure voluptas in et similique.\",\n      \"serviceAccount\": \"Architecto tenetur ipsam non est non aut.\"\n   }'")
		}
	}
	v := &tokenexchange.ExchangeTokenPayload{
		JWT:            body.JWT,
		Namespace:      body.Namespace,
		ServiceAccount: body.ServiceAccount,
	}

	return v, nil
}
