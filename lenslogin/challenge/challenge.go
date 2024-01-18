package challenge

import (
	"context"

	"github.com/machinebox/graphql"
)

type Challenge struct {
	Challenge struct {
		Text string
	} `graphql:"challenge(request: $request)"`
}

type ChallengeRequest struct {
	Address string `json:"address"`
}

func ChallengeFunc(address string) (string, error) {
	client := graphql.NewClient("https://api.lens.dev")

	req := graphql.NewRequest(`
        query Challenge($request:ChallengeRequest!) {
            challenge(request:$request) {
                text
            }
        }`)

	req.Var("request", ChallengeRequest{Address: address})
	req.Header.Set("Origin", "memo.io")

	var query Challenge
	if err := client.Run(context.Background(), req, &query); err != nil {
		return "", err
	}

	return query.Challenge.Text, nil
}
