package api

import (
	"crypto/tls"
	"net/http"

	"github.com/esnet/gdg/internal/config/domain"

	"github.com/carlmjohnson/requests"
	"github.com/esnet/gdg/internal/config"
)

// Most of these methods are here due to limitations in existing libraries being used.
type ExtendedApi struct {
	grafanaCfg *domain.GrafanaConfig
	debug      bool
}

func NewExtendedApi() *ExtendedApi {
	cfg := config.Config()
	o := ExtendedApi{
		grafanaCfg: cfg.GetDefaultGrafanaConfig(),
		debug:      cfg.IsApiDebug(),
	}
	return &o
}

func (extended *ExtendedApi) getRequestBuilder() *requests.Builder {
	req := requests.URL(extended.grafanaCfg.GetURL())
	if config.Config().IgnoreSSL() {
		customTransport := http.DefaultTransport.(*http.Transport).Clone()
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402
		req = req.Transport(customTransport)
	}

	if extended.grafanaCfg.GetAPIToken() != "" {
		req.Header("Authorization", "Bearer "+extended.grafanaCfg.GetAPIToken())
	} else {
		req.BasicAuth(extended.grafanaCfg.UserName, extended.grafanaCfg.GetPassword())
	}

	return req
}
