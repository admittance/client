// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package libkb

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestEnvDarwin(t *testing.T) {
	env := newEnv(nil, nil, "darwin", makeLogGetter(t))

	sockFile, err := env.GetSocketBindFile()
	if err != nil {
		t.Fatal(err)
	}

	cacheDir := env.GetSandboxCacheDir()
	expectedSockFile := filepath.Join(cacheDir, "keybased.sock")
	if sockFile != expectedSockFile {
		t.Fatalf("Clients expect sock file to be %s", expectedSockFile)
	}
}

type MockedConfigReader struct {
	NullConfiguration
}

var globalTorMode = TorNone

func (nc MockedConfigReader) GetTorMode() (TorMode, error) {
	return globalTorMode, nil
}

var globalProxyType = ""

func (nc MockedConfigReader) GetProxyType() string {
	return globalProxyType
}

var globalProxyAddress = ""

func (nc MockedConfigReader) GetProxy() string {
	return globalProxyAddress
}

var globalIsCertPinningEnabled = true

func (nc MockedConfigReader) IsCertPinningEnabled() bool {
	return globalIsCertPinningEnabled
}

func resetGlobals() {
	globalTorMode = TorNone
	globalProxyType = ""
	globalProxyAddress = ""
	globalIsCertPinningEnabled = true
}

func TestTorMode(t *testing.T) {
	resetGlobals()
	os.Clearenv()

	mockedEnv := NewEnv(MockedConfigReader{}, MockedConfigReader{}, makeLogGetter(t))

	// Test that when tor mode is enabled, a socks proxy is properly configured
	require.Equal(t, noProxy, mockedEnv.GetProxyType())
	require.Equal(t, "", mockedEnv.GetProxy())

	globalTorMode = TorLeaky
	require.Equal(t, socks, mockedEnv.GetProxyType())
	require.Equal(t, "localhost:9050", mockedEnv.GetProxy())

	globalTorMode = TorStrict
	require.Equal(t, socks, mockedEnv.GetProxyType())
	require.Equal(t, "localhost:9050", mockedEnv.GetProxy())

	// Test that tor mode overrides proxy settings
	globalProxyType = "http"
	globalProxyAddress = "localhost:8080"
	require.Equal(t, socks, mockedEnv.GetProxyType())
	require.Equal(t, "localhost:9050", mockedEnv.GetProxy())
}

func TestGetProxyType(t *testing.T) {
	resetGlobals()
	os.Clearenv()

	defaultEnv := NewEnv(nil, nil, makeLogGetter(t))
	require.Equal(t, noProxy, defaultEnv.GetProxyType())

	mockedEnv := NewEnv(MockedConfigReader{}, MockedConfigReader{}, makeLogGetter(t))
	require.Equal(t, noProxy, mockedEnv.GetProxyType())

	globalProxyType = "socks"
	require.Equal(t, socks, mockedEnv.GetProxyType())
	globalProxyType = "SOCKS"
	require.Equal(t, socks, mockedEnv.GetProxyType())
	globalProxyType = "SoCkS"
	require.Equal(t, socks, mockedEnv.GetProxyType())

	globalProxyType = "http_connect"
	require.Equal(t, httpConnect, mockedEnv.GetProxyType())
	globalProxyType = "HTTP_CONNECT"
	require.Equal(t, httpConnect, mockedEnv.GetProxyType())

	globalProxyType = "BOGUS"
	require.Equal(t, noProxy, mockedEnv.GetProxyType())

	resetGlobals()
	require.Equal(t, noProxy, mockedEnv.GetProxyType())

	os.Setenv("PROXY_TYPE", "socks")
	require.Equal(t, socks, mockedEnv.GetProxyType())
	os.Setenv("PROXY_TYPE", "http_connect")
	require.Equal(t, httpConnect, mockedEnv.GetProxyType())
}

func TestGetProxy(t *testing.T) {
	resetGlobals()
	os.Clearenv()

	defaultEnv := NewEnv(nil, nil, makeLogGetter(t))
	require.Equal(t, "", defaultEnv.GetProxy())

	mockedEnv := NewEnv(MockedConfigReader{}, MockedConfigReader{}, makeLogGetter(t))
	require.Equal(t, "", mockedEnv.GetProxy())

	globalProxyAddress = "localhost:8090"
	require.Equal(t, "localhost:8090", mockedEnv.GetProxy())

	resetGlobals()
	require.Equal(t, "", defaultEnv.GetProxy())

	os.Setenv("PROXY", "localhost:8080")
	require.Equal(t, "localhost:8080", mockedEnv.GetProxy())
	os.Clearenv()

	os.Setenv("HTTP_PROXY", "localhost:8081")
	require.Equal(t, "localhost:8081", mockedEnv.GetProxy())
	os.Clearenv()

	os.Setenv("HTTPS_PROXY", "localhost:8082")
	require.Equal(t, "localhost:8082", mockedEnv.GetProxy())
	os.Clearenv()
}

func TestIsCertPinningEnabled(t *testing.T) {
	resetGlobals()
	os.Clearenv()

	defaultEnv := NewEnv(nil, nil, makeLogGetter(t))
	require.Equal(t, true, defaultEnv.IsCertPinningEnabled())

	mockedEnv := NewEnv(MockedConfigReader{}, MockedConfigReader{}, makeLogGetter(t))
	require.Equal(t, true, mockedEnv.IsCertPinningEnabled())

	globalIsCertPinningEnabled = false
	require.Equal(t, false, mockedEnv.IsCertPinningEnabled())

	globalIsCertPinningEnabled = true
	require.Equal(t, true, mockedEnv.IsCertPinningEnabled())

	os.Setenv("DISABLE_SSL_PINNING", "true")
	require.Equal(t, false, mockedEnv.IsCertPinningEnabled())
	os.Setenv("DISABLE_SSL_PINNING", "no")
	require.Equal(t, true, mockedEnv.IsCertPinningEnabled())
	os.Clearenv()
	require.Equal(t, true, mockedEnv.IsCertPinningEnabled())
}
