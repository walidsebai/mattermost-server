// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package imageproxy

import (
	"net/http"
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/utils"
)

type ImageProxy struct {
	configService    utils.ConfigService
	configListenerId string

	httpService utils.HTTPService

	lock    sync.RWMutex
	backend ImageProxyBackend
}

type ImageProxyBackend interface {
	GetImage(w http.ResponseWriter, r *http.Request, url string)
}

func MakeImageProxy(configService utils.ConfigService, httpService utils.HTTPService) *ImageProxy {
	proxy := &ImageProxy{
		configService: configService,
		httpService:   httpService,
	}

	proxy.lock.Lock()
	defer proxy.lock.Unlock()

	proxy.configListenerId = proxy.configService.AddConfigListener(proxy.OnConfigChange)
	proxy.backend = proxy.makeBackend(*proxy.configService.Config().ImageProxySettings.ImageProxyType)

	return proxy
}

func (proxy *ImageProxy) makeBackend(proxyType string) ImageProxyBackend {
	switch proxyType {
	case model.IMAGE_PROXY_TYPE_LOCAL:
		return makeLocalBackend(proxy.httpService)
	case model.IMAGE_PROXY_TYPE_ATMOS_CAMO:
		return makeAtmosCamoBackend(proxy.configService)
	default:
		return nil
	}
}

func (proxy *ImageProxy) Close() {
	proxy.lock.Lock()
	defer proxy.lock.Unlock()

	proxy.configService.RemoveConfigListener(proxy.configListenerId)
}

func (proxy *ImageProxy) OnConfigChange(oldConfig, newConfig *model.Config) {
	if *oldConfig.ImageProxySettings.ImageProxyType != *newConfig.ImageProxySettings.ImageProxyType {
		proxy.lock.Lock()
		defer proxy.lock.Unlock()

		proxy.backend = proxy.makeBackend(*newConfig.ImageProxySettings.ImageProxyType)
	}
}

func (proxy *ImageProxy) GetImage(w http.ResponseWriter, r *http.Request, url string) {
	proxy.lock.RLock()
	defer proxy.lock.RUnlock()

	proxy.backend.GetImage(w, r, url)
}
