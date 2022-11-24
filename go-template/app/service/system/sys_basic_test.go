package system

import (
	"testing"
)

func TestBasicInfoService_GetMonitorInfo(t *testing.T) {
	b := BasicInfoService{}
	b.GetBasicInfo()
	b.GetMonitorInfo()

}
