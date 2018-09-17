// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package imageproxy

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/utils"
	"github.com/mattermost/mattermost-server/utils/testutils"
	"github.com/stretchr/testify/assert"
)

func makeAtmosCamoConfigService() utils.ConfigService {
	return &testutils.StaticConfigService{
		Cfg: &model.Config{
			ImageProxySettings: model.ImageProxySettings{
				ImageProxyType:          model.NewString(model.IMAGE_PROXY_TYPE_ATMOS_CAMO),
				RemoteImageProxyURL:     model.NewString("http://images.example.com/"),
				RemoteImageProxyOptions: model.NewString("7e5f3fab20b94782b43cdb022a66985ef28ba355df2c5d5da3c9a05e4b697bac"),
			},
		},
	}
}

func TestAtmosCamoBackend_GetImage(t *testing.T) {
	inputURL := "http://www.mattermost.org/wp-content/uploads/2016/03/logoHorizontalWhite.png"
	expectedURL := "http://images.example.com/62183a1cf0a4927c3b56d249366c2745e34ffe63/687474703a2f2f7777772e6d61747465726d6f73742e6f72672f77702d636f6e74656e742f75706c6f6164732f323031362f30332f6c6f676f486f72697a6f6e74616c57686974652e706e67"

	backend := makeAtmosCamoBackend(makeAtmosCamoConfigService())

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "", nil)
	backend.GetImage(recorder, request, inputURL)
	resp := recorder.Result()

	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, expectedURL, resp.Header.Get("Location"))
}

func TestAtmosCamoBackend_getImageURL(t *testing.T) {
	inputURL := "http://www.mattermost.org/wp-content/uploads/2016/03/logoHorizontal.png"
	expectedURL := "http://images.example.com/5b6f6661516bc837b89b54566eb619d14a5c3eca/687474703a2f2f7777772e6d61747465726d6f73742e6f72672f77702d636f6e74656e742f75706c6f6164732f323031362f30332f6c6f676f486f72697a6f6e74616c2e706e67"

	backend := makeAtmosCamoBackend(makeAtmosCamoConfigService())

	assert.Equal(t, expectedURL, backend.getImageURL(inputURL))
}
