package linear

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	pkglinear "github.com/lwlee2608/linear-cli/pkg/linear"
)

var markdownImagePattern = regexp.MustCompile(`!\[[^\]]*\]\(\s*(?:<([^>]+)>|([^\s)]+))(?:\s+(?:"[^"]*"|'[^']*'|\([^)]*\)))?\s*\)`)

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
	matches := markdownImagePattern.FindAllStringSubmatch(description, -1)
	urls := make([]string, 0, len(matches))
	seen := make(map[string]struct{}, len(matches))
	for _, match := range matches {
		imageURL := match[1]
		if imageURL == "" {
			imageURL = match[2]
		}
		if _, ok := seen[imageURL]; ok {
			continue
		}
		seen[imageURL] = struct{}{}
		urls = append(urls, imageURL)
	}
	return urls
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
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return fmt.Errorf("create image file %s: %w", path, err)
	}
	defer func() {
		body.Close()
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
