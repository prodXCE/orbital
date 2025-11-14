package downloader

import (
	"archive/tar"   // Standard library for reading .tar files
	"compress/gzip" // Standard library for un-gzipping
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const imageDir = "./.orbital/images"

var knownImages = map[string]string{
	"alpine-arm": "https://dl-cdn.alpinelinux.org/alpine/v3.18/releases/aarch64/alpine-minirootfs-3.18.4-aarch64.tar.gz",
	"alpine-amd": "https://dl-cdn.alpinelinux.org/alpine/v3.18/releases/x86_64/alpine-minirootfs-3.18.4-x86_64.tar.gz",
}

func Pull(imageName string) error {
	url, ok := knownImages[imageName]
	if !ok {
		return fmt.Errorf("unknown image: %s (try 'alpine-arm' or 'alpine-amd')", imageName)
	}
	fmt.Printf("Downloading from %s...\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http get error: %v", err)
	}
	defer resp.Body.Close()

	targetPath := filepath.Join(imageDir, imageName)
	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return fmt.Errorf("mkdir error: %v", err)
	}
	fmt.Printf("Extracting to %s...\n", targetPath)

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("gzip reader error: %v", err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("tar reader error: %v", err)
		}

		path := filepath.Join(targetPath, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return fmt.Errorf("mkdir error: %v", err)
			}
		case tar.TypeReg:
			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, header.FileInfo().Mode())
			if err != nil {
				return fmt.Errorf("create file error: %v", err)
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return fmt.Errorf("file copy error: %v", err)
			}
			outFile.Close()

		case tar.TypeSymlink:
			// TEMP FIX :
			// It's a symlink, so create it.
			// 'header.Linkname' is the target (e.g., "busybox")
			// 'path' is the name of the link (e.g., ".../bin/sh")
			if err := os.Symlink(header.Linkname, path); err != nil {
				return fmt.Errorf("symlink create error: %v", err)
			}

		default:
			// last fix for now
		}
	}

	fmt.Printf("Successfully downloaded and extracted '%s'.\n", imageName)
	return nil
}

func GetImagePath(imageName string) (string, bool) {
	path := filepath.Join(imageDir, imageName)

	if _, err := os.Stat(path); err == nil {
		return path, true
	}

	return "", false
}
