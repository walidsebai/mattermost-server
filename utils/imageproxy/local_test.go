// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package imageproxy

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mattermost/mattermost-server/utils/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocalBackend_GetImage(t *testing.T) {
	t.Run("image", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "max-age=2592000, private")
			w.Header().Set("Content-Type", "image/png")
			w.Header().Set("Content-Length", "10")

			w.Write([]byte("1111111111"))

			w.WriteHeader(http.StatusOK)
		})

		httpService := testutils.MakeMockedHTTPService(handler)
		defer httpService.Close()

		backend := makeLocalBackend(httpService)

		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "", nil)
		backend.GetImage(recorder, request, httpService.Server.URL+"/image.png")
		resp := recorder.Result()

		require.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "max-age=2592000, private", resp.Header.Get("Cache-Control"))
		assert.Equal(t, "10", resp.Header.Get("Content-Length"))

		respBody, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, []byte("1111111111"), respBody)
	})

	t.Run("not an image", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotAcceptable)
		})

		httpService := testutils.MakeMockedHTTPService(handler)
		defer httpService.Close()

		backend := makeLocalBackend(httpService)

		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "", nil)
		backend.GetImage(recorder, request, httpService.Server.URL+"/file.pdf")
		resp := recorder.Result()

		require.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("not an image, but remote server ignores accept header", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "max-age=2592000, private")
			w.Header().Set("Content-Type", "application/pdf")
			w.Header().Set("Content-Length", "10")

			w.Write([]byte("1111111111"))

			w.WriteHeader(http.StatusOK)
		})

		httpService := testutils.MakeMockedHTTPService(handler)
		defer httpService.Close()

		backend := makeLocalBackend(httpService)

		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "", nil)
		backend.GetImage(recorder, request, httpService.Server.URL+"/file.pdf")
		resp := recorder.Result()

		require.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("not found", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		httpService := testutils.MakeMockedHTTPService(handler)
		defer httpService.Close()

		backend := makeLocalBackend(httpService)

		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "", nil)
		backend.GetImage(recorder, request, httpService.Server.URL+"/image.png")
		resp := recorder.Result()

		require.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("other server error", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})

		httpService := testutils.MakeMockedHTTPService(handler)
		defer httpService.Close()

		backend := makeLocalBackend(httpService)

		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "", nil)
		backend.GetImage(recorder, request, httpService.Server.URL+"/image.png")
		resp := recorder.Result()

		require.Equal(t, http.StatusBadGateway, resp.StatusCode)
	})

	t.Run("image with etag", func(t *testing.T) {
		// TODO
		// handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 	w.Header().Set("Cache-Control", "max-age=2592000, private")
		// 	w.Header().Set("Content-Type", "image/png")
		// 	w.Header().Set("Content-Length", "10")

		// 	w.Write([]byte("1111111111"))

		// 	w.WriteHeader(http.StatusOK)
		// })

		// httpService := testutils.MakeMockedHTTPService(handler)
		// defer httpService.Close()

		// backend := makeLocalBackend(httpService)

		// recorder := httptest.NewRecorder()
		// request, _ := http.NewRequest(http.MethodGet, "", nil)
		// backend.GetImage(recorder, request, httpService.Server.URL+"/image.png")
		// resp := recorder.Result()

		// require.Equal(t, http.StatusOK, resp.StatusCode)
		// assert.Equal(t, "max-age=2592000, private", resp.Header.Get("Cache-Control"))
		// assert.Equal(t, "10", resp.Header.Get("Content-Length"))

		// respBody, _ := ioutil.ReadAll(resp.Body)
		// assert.Equal(t, []byte("1111111111"), respBody)
	})

	t.Run("timeout", func(t *testing.T) {
		wait := make(chan bool, 1)

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			<-wait
			w.WriteHeader(http.StatusNotFound)
		})

		httpService := testutils.MakeMockedHTTPService(handler)
		defer httpService.Close()
		httpService.OverrideTimeout = true
		httpService.Timeout = time.Second

		backend := makeLocalBackend(httpService)

		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "", nil)
		backend.GetImage(recorder, request, httpService.Server.URL+"/image.png")
		resp := recorder.Result()

		require.Equal(t, http.StatusGatewayTimeout, resp.StatusCode)

		wait <- true
	})
}
