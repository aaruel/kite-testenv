package repo

import (
	"github.com/aaruel/kite-testenv/utils"
	"github.com/koding/kite"
	"github.com/koding/kite/config"
	"github.com/koding/kite/dnode"
	"github.com/koding/kite/protocol"
)

var (
	kontrolUrl   string                = "http://localhost:6789/kite"
	repoName     string                = "gateway"
	version      string                = "1.0.0"
	serviceNames utils.ServiceToMethod = utils.NewServiceToMethod(
		utils.NewSTM("math", "square"),
	)
)

type Repo struct {
	connector *kite.Kite
	clients   map[string]*kite.Client
}

func New() *Repo {
	nrepo := &Repo{}
	nrepo.setup()
	nrepo.start()
	return nrepo
}

func checkValidService(name string) {
	if !serviceNames.Contains(name) {
		panic("Service `" + name + "` does not exist")
	}
}

func checkValidMethod(service string, method string) {
	if !serviceNames.ContainsMethod(service, method) {
		panic("Method `" + method + "` does not exist in service `" + service + "`")
	}
}

func (r *Repo) GetSTMStructure() utils.ServiceToMethod {
	return serviceNames
}

func (r *Repo) getService(name string) {
	checkValidService(name)

	var c *kite.Kite = r.connector
	kites, err := c.GetKites(&protocol.KontrolQuery{
		Username:    c.Config.Username,
		Environment: c.Config.Environment,
		Name:        name,
	})
	if err != nil {
		panic(err)
	}

	connErr := kites[0].Dial()
	if connErr != nil {
		panic(connErr)
	}

	r.clients[name] = kites[0]
}

func (r *Repo) Query(client string, method string, args ...interface{}) *dnode.Partial {
	checkValidService(client)
	checkValidMethod(client, method)

	response, err := r.clients[client].Tell(method, args...)
	if err != nil {
		panic(err)
	}

	return response
}

func (r *Repo) setup() {
	r.connector = kite.New(repoName, version)
	r.connector.Config = config.MustGet()
	r.connector.Config.KontrolURL = kontrolUrl
	r.clients = make(map[string]*kite.Client)
}

func (r *Repo) start() {
	for service := range serviceNames {
		r.getService(service)
	}
}
