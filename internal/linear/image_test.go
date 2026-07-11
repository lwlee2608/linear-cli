package linear

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestMarkdownImageURLs(t *testing.T) {
	description := `before
![first](https://uploads.linear.app/a)
[normal](https://uploads.linear.app/not-an-image)
![with title](<https://uploads.linear.app/b> "title")
![duplicate](https://uploads.linear.app/a)`

	want := []string{
		"https://uploads.linear.app/a",
		"https://uploads.linear.app/b",
	}
	if got := markdownImageURLs(description); !reflect.DeepEqual(got, want) {
		t.Fatalf("markdownImageURLs() = %v, want %v", got, want)
	}
}

func TestImageExtension(t *testing.T) {
	tests := map[string]string{
		"image/png":                  ".png",
		"image/jpeg; charset=binary": ".jpg",
		"image/svg+xml":              ".svg",
		"image/x-custom":             "",
	}
	for contentType, want := range tests {
		t.Run(contentType, func(t *testing.T) {
			if got := imageExtension(contentType); got != want {
				t.Fatalf("imageExtension(%q) = %q, want %q", contentType, got, want)
			}
		})
	}
}

func TestWriteNewFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "image.png")
	body := io.NopCloser(strings.NewReader("image data"))
	if err := writeNewFile(path, body); err != nil {
		t.Fatalf("writeNewFile() error = %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read image: %v", err)
	}
	if string(data) != "image data" {
		t.Fatalf("image contents = %q", data)
	}

	if err := writeNewFile(path, io.NopCloser(strings.NewReader("replacement"))); err == nil {
		t.Fatal("writeNewFile() overwrote an existing file")
	}
}
