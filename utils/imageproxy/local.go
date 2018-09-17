// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package imageproxy

import (
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"

	"github.com/mattermost/mattermost-server/mlog"
	"github.com/mattermost/mattermost-server/utils"
)

type LocalBackend struct {
	httpService utils.HTTPService
}

func makeLocalBackend(httpService utils.HTTPService) *LocalBackend {
	return &LocalBackend{
		httpService: httpService,
	}
}

func (backend *LocalBackend) GetImage(w http.ResponseWriter, r *http.Request, imageURL string) {
	req, err := http.NewRequest(http.MethodGet, imageURL, nil)
	if err != nil {
		// http.NewRequest should only return an error on an invalid URL
		mlog.Debug("Failed to create request for proxied image", mlog.String("url", imageURL), mlog.Any("err", err))

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	req.Header.Set("Accept", "image/*")

	resp, err := backend.httpService.MakeClient(false).Get(imageURL)
	if err != nil {
		mlog.Debug("Failed to get proxied image", mlog.String("url", imageURL), mlog.Any("err", err))

		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusBadGateway)
		}
		w.Write([]byte{})

		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte{})
		return
	} else if resp.StatusCode >= 400 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}

	// Double check that the returned content is an image
	contentType, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if !strings.HasPrefix(contentType, "image/") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	copyProxiedResponse(w, resp)
}

func copyProxiedResponse(w http.ResponseWriter, remoteResp *http.Response) {
	for _, key := range []string{
		"Content-Type",
		"Content-Length",
		"Cache-Control",
		"Last-Modified",
		"Expires",
		"Etag",
		"Link",
	} {
		w.Header().Add(key, remoteResp.Header.Get(key))
	}

	w.WriteHeader(remoteResp.StatusCode)

	io.Copy(w, remoteResp.Body)
}
