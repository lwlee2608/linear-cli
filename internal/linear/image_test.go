package linear

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type trackingReadCloser struct {
	io.Reader
	closed bool
}

func (r *trackingReadCloser) Close() error {
	r.closed = true
	return nil
}

func TestMarkdownImageURLs(t *testing.T) {
	description := `before
![first](https://uploads.linear.app/a)
[normal](https://uploads.linear.app/not-an-image)
![with title](<https://uploads.linear.app/b> "title")
![external](https://example.com/image.png)
` + "`![inline code](https://uploads.linear.app/inline)`" + `
<!-- ![comment](https://uploads.linear.app/comment) -->

` + "```markdown" + `
![fenced](https://uploads.linear.app/fenced)
` + "```" + `

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

func TestWriteNewFileClosesBodyWhenFileExists(t *testing.T) {
	path := filepath.Join(t.TempDir(), "image.png")
	if err := os.WriteFile(path, []byte("existing"), 0o644); err != nil {
		t.Fatalf("create existing file: %v", err)
	}

	body := &trackingReadCloser{Reader: strings.NewReader("replacement")}
	if err := writeNewFile(path, body); err == nil {
		t.Fatal("writeNewFile() overwrote an existing file")
	}
	if !body.closed {
		t.Fatal("writeNewFile() did not close body")
	}
}
