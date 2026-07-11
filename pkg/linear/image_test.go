package linear

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestDownloadImage(t *testing.T) {
	client := NewClient("secret")
	client.httpClient = &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if got := req.Header.Get("Authorization"); got != "secret" {
			t.Fatalf("Authorization = %q, want secret", got)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"image/png"}},
			Body:       io.NopCloser(strings.NewReader("png")),
			Request:    req,
		}, nil
	})}

	body, contentType, err := client.DownloadImage(context.Background(), "https://uploads.linear.app/image")
	if err != nil {
		t.Fatalf("DownloadImage() error = %v", err)
	}
	defer body.Close()
	if contentType != "image/png" {
		t.Fatalf("content type = %q", contentType)
	}
	data, err := io.ReadAll(body)
	if err != nil || string(data) != "png" {
		t.Fatalf("body = %q, error = %v", data, err)
	}
}

func TestDownloadImageRejectsUnsupportedURL(t *testing.T) {
	client := NewClient("secret")
	for _, imageURL := range []string{
		"http://uploads.linear.app/image",
		"https://example.com/image",
	} {
		if _, _, err := client.DownloadImage(context.Background(), imageURL); err == nil {
			t.Fatalf("DownloadImage(%q) accepted unsupported URL", imageURL)
		}
	}
}

func TestDownloadImageRejectsNonImageResponse(t *testing.T) {
	client := NewClient("secret")
	client.httpClient = &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"text/html"}},
			Body:       io.NopCloser(strings.NewReader("not an image")),
			Request:    req,
		}, nil
	})}

	if _, _, err := client.DownloadImage(context.Background(), "https://uploads.linear.app/image"); err == nil {
		t.Fatal("DownloadImage() accepted a non-image response")
	}
}
