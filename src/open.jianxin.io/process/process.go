package process

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"open.jianxin.io/captcha"
)

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func randomName() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	return Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(rndNum, 10)))
}

func captchaGenerate(size int) []string {
	s := make([]string, 0)
	for i := 0; i < size; i++ {
		Topic := captcha.RandTopic()
		fileName := randomName() + ".gif"
		captcha.Draw(Topic.Subject, "../tmp/"+fileName)
		s = append(s, fileName+"|"+Topic.Result)
	}
	return s
}

func Start() {
	captchas := captchaGenerate(100)
	CaptchaContainer.Append(captchas...)
	log.Print("Init success.")
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for _ = range ticker.C {
			go func() {
				workder()
			}()
		}
	}()
}

func workder() {
	if !CaptchaContainer.UpdateNeed() {
		return
	}
	captchas := captchaGenerate(100)
	CaptchaContainer.Lock()
	oldCaptchas := CaptchaContainer.Update(captchas...)
	CaptchaContainer.Unlock()
	for _, captcha := range oldCaptchas {
		fileName := strings.Split(captcha, "|")[0]
		os.Remove("../tmp/" + fileName)
	}
	log.Print("Update suceess.")
}
