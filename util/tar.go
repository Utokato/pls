package util

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func UnarchivedTarGz(source []byte, destination string) error {
	reader := bytes.NewReader(source)
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer gzipReader.Close()
	// 使用 tar Reader 解压
	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			// 读取到文件末尾，解压完成
			break
		}
		if err != nil {
			return fmt.Errorf("error reading tar entry: %w", err)
		}

		// 获取目标路径
		targetPath := filepath.Join(destination, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			// 创建目录
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}

		case tar.TypeReg:
			// 创建文件
			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			defer outFile.Close()
			// 将内容写入文件
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("failed to copy file content: %w", err)
			}
		default:
			fmt.Printf("Unsupported file type: %c in file: %s\n", header.Typeflag, header.Name)
		}
	}

	return nil
}
