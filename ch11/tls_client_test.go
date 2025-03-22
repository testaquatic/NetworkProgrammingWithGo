package ch11

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestClientTLS(t *testing.T) {
	ts := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.TLS == nil {
					u := "https://" + r.Host + r.RequestURI
					http.Redirect(w, r, u, http.StatusMovedPermanently)
					return
				}
				w.WriteHeader(http.StatusOK)
			},
		),
	)
	defer ts.Close()
	// [이 문서](https://pkg.go.dev/net/http/httptest#example-Server-HTTP2)의 설명에 따라서 코드를 변경했다.
	ts.EnableHTTP2 = true
	ts.StartTLS()

	resp, err := ts.Client().Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d; actual status %d", http.StatusOK, resp.StatusCode)
	}

	tp := &http.Transport{
		TLSClientConfig: &tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			MinVersion:       tls.VersionTLS12,
		},
		// [이 문서](https://pkg.go.dev/net/http#Transport)에 따라서 설정했다.
		ForceAttemptHTTP2: true,
		Protocols:         new(http.Protocols),
	}
	tp.Protocols.SetHTTP2(true)
	client2 := &http.Client{
		Transport: tp,
	}

	_, err = client2.Get(ts.URL)
	if err == nil || !strings.Contains(err.Error(), "certificate signed by unknown authority") {
		t.Fatalf("expected unknown authority error; actual: %q", err)
	}

	tp.TLSClientConfig.InsecureSkipVerify = true
	client2 = &http.Client{
		Transport: tp,
	}

	resp, err = client2.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	// 책의 내용과 다르기 때문에 일단 점검하고 넘어가자.
	if resp.ProtoMajor != 2 {
		t.Errorf("expected major version 2; actual: %d", resp.ProtoMajor)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d; actual status %d", http.StatusOK, resp.StatusCode)
	}
}

func TestClientTLSGoogle(t *testing.T) {
	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: 30 * time.Second},
		"tcp",
		"www.google.com:443",
		&tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			MinVersion:       tls.VersionTLS12,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	state := conn.ConnectionState()
	t.Log(tls.VersionName(state.Version))
	t.Log(tls.CipherSuiteName(state.CipherSuite))
	t.Log(state.VerifiedChains[0][0].Issuer.Organization[0])

	_ = conn.Close()
}
