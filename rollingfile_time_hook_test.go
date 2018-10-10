package rollingfile

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewRollingFileTimeHook(t *testing.T) {
	hook, err := NewRollingFileTimeHook("./test.log",
		"2006-01-02", 5)
	if err != nil {
		t.Errorf("NewRollingFileTimeHook: %s", err)
	}
	defer hook.Close()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.AddHook(hook)

	for i:=0; i < 1000; i++ {
		logrus.Debugf("TestNewRollingFileTimeHook: %d",i)
		time.Sleep(time.Millisecond * 10)
	}
}