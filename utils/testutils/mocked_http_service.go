// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package testutils

import (
	"net/http"
	"net/http/httptest"
	"time"
)

type MockedHTTPService struct {
	Server *httptest.Server

	OverrideTimeout bool
	Timeout         time.Duration
}

func MakeMockedHTTPService(handler http.Handler) *MockedHTTPService {
	return &MockedHTTPService{
		Server: httptest.NewServer(handler),
	}
}

func (h *MockedHTTPService) MakeClient(trustURLs bool) *http.Client {
	client := h.Server.Client()

	if h.OverrideTimeout {
		client.Timeout = h.Timeout
	}

	return client
}

func (h *MockedHTTPService) Close() {
	h.Server.CloseClientConnections()
	h.Server.Close()
}
