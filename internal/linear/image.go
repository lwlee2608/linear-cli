package linear

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	pkglinear "github.com/lwlee2608/linear-cli/pkg/linear"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func (s *Service) DownloadIssueImages(ctx context.Context, issue *pkglinear.Issue, dir string) ([]string, error) {
	urls := markdownImageURLs(issue.Description)
	if len(urls) == 0 {
		return nil, nil
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("create image directory: %w", err)
	}

	paths := make([]string, 0, len(urls))
	for i, imageURL := range urls {
		body, contentType, err := s.client.DownloadImage(ctx, imageURL)
		if err != nil {
			return paths, fmt.Errorf("download image %d: %w", i+1, err)
		}

		ext := imageExtension(contentType)
		path := filepath.Join(dir, fmt.Sprintf("image-%02d%s", i+1, ext))
		if err := writeNewFile(path, body); err != nil {
			return paths, err
		}
		paths = append(paths, path)
	}
	return paths, nil
}

func markdownImageURLs(description string) []string {
	source := []byte(description)
	document := goldmark.DefaultParser().Parse(text.NewReader(source))
	urls := make([]string, 0)
	seen := make(map[string]struct{})
	_ = ast.Walk(document, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		image, ok := node.(*ast.Image)
		if !entering || !ok {
			return ast.WalkContinue, nil
		}
		imageURL := string(image.Destination)
		if !isLinearUploadURL(imageURL) {
			return ast.WalkContinue, nil
		}
		if _, ok := seen[imageURL]; ok {
			return ast.WalkContinue, nil
		}
		seen[imageURL] = struct{}{}
		urls = append(urls, imageURL)
		return ast.WalkContinue, nil
	})
	return urls
}

func isLinearUploadURL(imageURL string) bool {
	parsedURL, err := url.Parse(imageURL)
	return err == nil && parsedURL.Scheme == "https" && strings.EqualFold(parsedURL.Hostname(), "uploads.linear.app")
}

func imageExtension(contentType string) string {
	switch strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0])) {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/svg+xml":
		return ".svg"
	case "image/avif":
		return ".avif"
	default:
		return ""
	}
}

func writeNewFile(path string, body interface {
	Read([]byte) (int, error)
	Close() error
}) (err error) {
	defer body.Close()

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return fmt.Errorf("create image file %s: %w", path, err)
	}
	defer func() {
		if closeErr := file.Close(); err == nil && closeErr != nil {
			err = fmt.Errorf("close image file %s: %w", path, closeErr)
		}
		if err != nil {
			os.Remove(path)
		}
	}()

	if _, err = file.ReadFrom(body); err != nil {
		return fmt.Errorf("write image file %s: %w", path, err)
	}
	return nil
}
