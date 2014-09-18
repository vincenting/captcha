package process

import (
	"testing"
)

func TestCaptchaInsert(t *testing.T) {
	CaptchaContainer.Append("1", "2")
	if len(CaptchaContainer.captchaList) != 2 || CaptchaContainer.pointerIndex != 0 {
		t.Fatalf("Insert with unexcepted")
	}
}

func TestCaptchaNext(t *testing.T) {
	item, err := CaptchaContainer.Next()
	if err != nil {
		t.Fatalf("Next with err: %s", err.Error())
	}
	if item != "1" {
		t.Fatalf("Next with unexcepted")
	}
	item, _ = CaptchaContainer.Next()
	if item != "2" {
		t.Fatalf("Next with unexcepted")
	}
}

func TestCaptchaUpdate(t *testing.T) {
	a := CaptchaContainer.Update("3")
	if a[0] != "1" || len(a) != 1 {
		t.Fatalf("Update return with unexcepted")
	}
	item, _ := CaptchaContainer.Next()
	if item != "3" {
		t.Fatalf("Next with unexcepted")
	}
	item, _ = CaptchaContainer.Next()
	if item != "2" {
		t.Fatalf("Next with unexcepted")
	}
}

func TestCaptchaLock(t *testing.T) {
	CaptchaContainer.Lock()
	go func() {
		CaptchaContainer.Update("4")
		CaptchaContainer.Unlock()
	}()
	item, _ := CaptchaContainer.Next()
	if item != "3" {
		t.Fatalf("Next with unexcepted")
	}
}
