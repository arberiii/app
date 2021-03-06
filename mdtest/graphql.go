package mdtest

import (
	"net/http/httptest"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/short-d/app/fw"
)

func IsGraphQlAPIValid(api fw.GraphQLAPI) bool {
	_, err := graphql.ParseSchema(api.GetSchema(), api.GetResolver())
	return err == nil
}

type GraphQLServerFake struct {
	server *httptest.Server
}

func (g GraphQLServerFake) URL() string {
	return g.server.URL
}

func NewGraphQLServerFake(api fw.GraphQLAPI) GraphQLServerFake {
	schema := graphql.MustParseSchema(api.GetSchema(), api.GetResolver())
	relayHandler := relay.Handler{
		Schema: schema,
	}

	server := httptest.NewServer(&relayHandler)
	return GraphQLServerFake{
		server: server,
	}
}
