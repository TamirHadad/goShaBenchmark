package utils

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
)

type HashInfo struct {
	md5    string
	sha1   string
	sha256 string
}

func CreateHashInfoByPath(filePath string) (*HashInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return CalculateBasicHashes(bufio.NewReader(file), "", false)
}

func CreateHashInfo(reader io.Reader) (*HashInfo, error) {
	return CalculateBasicHashes(reader, "", false)
}

func WriteStreamToHashAndData(reader io.Reader, path string, includeFile bool) (*HashInfo, error) {
	return CalculateBasicHashes(reader, path, includeFile)
}

func (self HashInfo) Sha1() string {
	return self.sha1
}

func (self HashInfo) Md5() string {
	return self.md5
}

func (self HashInfo) Sha256() string {
	return self.sha256
}

func (self HashInfo) getSha256IfNotExist(sha256 string) string {
	if len(sha256) > 0 {
		return sha256
	} else {
		return self.sha256
	}
}

func CalculateBasicHashes(rd io.Reader, path string, includeFile bool) (*HashInfo, error) {
	md5 := md5.New()
	sha1 := sha1.New()
	sha256 := sha256.New()
	var multiWriter io.Writer
	pagesize := os.Getpagesize()
	reader := bufio.NewReaderSize(rd, pagesize)
	if includeFile {
		file, err := os.Create(path)
		defer file.Close()
		if err != nil {
			fmt.Print(err.Error())
		}
		multiWriter = AsyncMultiWriter(file, sha256, sha1, md5)
	} else {
		multiWriter = AsyncMultiWriter(sha256, sha1, md5)
	}
	_, err := io.Copy(multiWriter, reader)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return &HashInfo{sha1: fmt.Sprintf("%x", sha1.Sum(nil)), sha256: fmt.Sprintf("%x", sha256.Sum(nil)),
		md5: fmt.Sprintf("%x", md5.Sum(nil))}, nil
}

func CalculateSha256(val string) (string, error) {
	return CalculateHash(sha256.New(), "sha256", val)
}

func CalculateSha1(val string) (string, error) {
	return CalculateHash(sha256.New(), "sha1", val)
}

func CalculateMD5(val string) (string, error) {
	return CalculateHash(md5.New(), "md5", val)
}

func CalculateHash(h hash.Hash, hashName, val string) (string, error) {
	n, err := h.Write([]byte(val))
	if err != nil {
		return "", err
	}
	if n != len(val) {
		return "", errors.New("short write for " + hashName)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}