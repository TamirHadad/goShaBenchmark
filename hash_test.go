package goShaBenchmark

import (
	"testing"
	"github.com/TamirHadad/goShaBenchmark/utils"
	"path/filepath"
	"fmt"
	"os"
)

func TestCorrectness(t *testing.T) {
	path, err := filepath.Abs("resources/1M")
	if err != nil  {
		fmt.Println(err)
		os.Exit(1)
	}
	pSha1, pSha256, pMd5 := utils.CalcHashes(path)
	hashInfo, _ := utils.CreateHashInfoByPath(path)
	if hashInfo.Md5() != pMd5 {
		t.Error("Incorect Md5:", "Xray:", hashInfo.Md5(), "paralle:", pMd5)
	}
	if hashInfo.Sha1() != pSha1 {
		t.Error("Incorect Sha1:", "Xray:", hashInfo.Sha1(), "paralle:", pSha1)
	}
	if hashInfo.Sha256() != pSha256 {
		t.Error("Incorect Sha1:", "Xray:", hashInfo.Sha256(), "paralle:", pSha256)
	}
}

func BenchmarkParallel1M(b *testing.B) {
	path, err := filepath.Abs("resources/1M")
	if err != nil  {
		fmt.Println(err)
		os.Exit(1)
	}

	for n := 0; n < b.N; n++ {
		utils.CalcHashes(path)
	}
}

func BenchmarkXrayCalcParallel1M(b *testing.B) {
	path, err := filepath.Abs("resources/1M")
	if err != nil  {
		fmt.Println(err)
		os.Exit(1)
	}

	for n := 0; n < b.N; n++ {
		utils.CreateHashInfoByPath(path)
	}
}

func BenchmarkParallel1G(b *testing.B) {
	path, err := filepath.Abs("resources/1G")
	if err != nil  {
		fmt.Println(err)
	}

	for n := 0; n < b.N; n++ {
		utils.CalcHashes(path)	}
}

func BenchmarkXrayCalcParallel1G(b *testing.B) {
	path, err := filepath.Abs("resources/1G")
	if err != nil  {
		fmt.Println(err)
	}

	for n := 0; n < b.N; n++ {
		utils.CreateHashInfoByPath(path)	}
}

func BenchmarkParallel6G(b *testing.B) {
	path, err := filepath.Abs("resources/6G")
	if err != nil  {
		fmt.Println(err)
	}

	for n := 0; n < b.N; n++ {
		utils.CalcHashes(path)	}
}

func BenchmarkXrayCalcParallel6G(b *testing.B) {
	path, err := filepath.Abs("resources/6G")
	if err != nil  {
		fmt.Println(err)
	}

	for n := 0; n < b.N; n++ {
		utils.CreateHashInfoByPath(path)	}
}