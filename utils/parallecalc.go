package utils

import (
	"os"
	"io"
	"crypto/sha1"
	"encoding/hex"
	"crypto/md5"
	"crypto/sha256"
	"sync"
)

func CalcHashes(filePath string) (string, string ,string) {
	var sha1, sha256, md5 string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		sha1, _ = CalcSha1(filePath)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		sha256, _ = CalcSha256(filePath)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		md5, _ = CalcMd5(filePath)
		wg.Done()
	}()

	wg.Wait()
	return sha1, sha256, md5
}

func CalcSha1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return GetSha1(file)
}

func GetSha1(input io.Reader) (string, error) {
	var resSha1 []byte
	hashSha1 := sha1.New()
	_, err := io.Copy(hashSha1, input)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hashSha1.Sum(resSha1)), nil
}

func CalcMd5(filePath string) (string, error) {
	var err error
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return GetMd5(file)
}

func GetMd5(input io.Reader) (string, error) {
	var resMd5 []byte
	hashMd5 := md5.New()
	_, err := io.Copy(hashMd5, input)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hashMd5.Sum(resMd5)), nil
}

func CalcSha256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return GetSha256(file)
}

func GetSha256(input io.Reader) (string, error) {
	var resSha1 []byte
	hashSha256 := sha256.New()
	_, err := io.Copy(hashSha256, input)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hashSha256.Sum(resSha1)), nil
}