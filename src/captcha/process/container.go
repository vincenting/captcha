package process

import (
	"errors"
	"sync"

	"captcha/config"
)

type Container struct {
	captchaList  []string
	locker       *sync.Mutex
	pointerIndex int
	consumption  int
}

var CaptchaContainer *Container

func init() {
	CaptchaContainer = &Container{
		[]string{},
		new(sync.Mutex),
		0,
		0,
	}
}

// Append items to the end of the list and remove old items from the front.
// At the same time move pointer
func (c *Container) Update(items ...string) []string {
	itemsSize, listSize := len(items), len(c.captchaList)
	c.Append(items...)
	captchaList := make([]string, len(c.captchaList))
	copy(captchaList, c.captchaList)
	c.captchaList = captchaList[itemsSize:]
	c.pointerIndex = (c.pointerIndex - itemsSize) % listSize
	if c.pointerIndex < 0 {
		c.pointerIndex += listSize
	}
	return captchaList[:itemsSize]
}

func (c *Container) UpdateNeed() bool {
	consumption := c.consumption
	c.consumption = 0
	if consumption < config.GetConfig().Threshold {
		return false
	}
	return true
}

// Append items to the end of the list
func (c *Container) Append(items ...string) {
	c.captchaList = append(c.captchaList, items...)
}

func (c *Container) Lock() {
	c.locker.Lock()
}

func (c *Container) Unlock() {
	c.locker.Unlock()
}

// Get next item by index
func (c *Container) Next() (string, error) {
	c.Lock()
	defer c.Unlock()
	c.consumption += 1
	if len(c.captchaList) == 0 {
		return "", errors.New("No item found")
	}
	index := c.pointerIndex % len(c.captchaList)
	c.pointerIndex = index + 1
	return c.captchaList[index], nil
}
