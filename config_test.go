package opencc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func Test_config(t *testing.T) {
	fileName := `s2t.json`
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	var conf *Config
	err = json.Unmarshal(body, &conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)

}

func Test_opencc_s2t(t *testing.T) {
	cc, err := NewOpenCC("s2t")
	if err != nil {
		fmt.Println(err)
		return
	}
	nText, err := cc.ConvertText(`保税工厂声明：本书为无限小说网(txt53.com)以下作品内容之版权与本站无任何关系`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nText)
}

func Test_opencc_t2s(t *testing.T) {
	cc, err := NewOpenCC("t2s")
	if err != nil {
		fmt.Println(err)
		return
	}
	nText, err := cc.ConvertText(`保稅工廠聲明：本書爲無限小說網(txt53.com)以下作品內容之版權與本站無任何關係`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nText)
}

func Test_openccFile(t *testing.T) {
	localdir := filepath.Dir(os.Args[0])

	inFile, err := os.Open(localdir + "/神剑渡魔.txt")
	if err != nil {
		fmt.Println("in:", err)
		return
	}
	outFile, err := os.OpenFile(localdir+"/神剑渡魔_3.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("out:", err)
		return
	}
	cc, err := NewOpenCC("s2t")
	if err != nil {
		fmt.Println("cc:", err)
		return
	}
	startTime := time.Now()
	log.Println("start...")
	err = cc.ConvertFile(inFile, outFile)
	if err != nil {
		fmt.Println("ccf", err)
		return
	}
	log.Println("end...", time.Now().Unix()-startTime.Unix())
}
