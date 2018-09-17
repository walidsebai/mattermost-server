// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package imageproxy

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/http"

	"github.com/mattermost/mattermost-server/utils"
)

type AtmosCamoBackend struct {
	configService utils.ConfigService
}

func makeAtmosCamoBackend(configService utils.ConfigService) *AtmosCamoBackend {
	return &AtmosCamoBackend{
		configService: configService,
	}
}

func (backend *AtmosCamoBackend) GetImage(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, backend.getImageURL(url), http.StatusFound)
}

func (backend *AtmosCamoBackend) getImageURL(url string) string {
	settings := backend.configService.Config().ImageProxySettings

	mac := hmac.New(sha1.New, []byte(*settings.RemoteImageProxyOptions))
	mac.Write([]byte(url))
	digest := hex.EncodeToString(mac.Sum(nil))

	return *settings.RemoteImageProxyURL + digest + "/" + hex.EncodeToString([]byte(url))
}
