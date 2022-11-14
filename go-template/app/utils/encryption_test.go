package utils

import "testing"

func TestMD5Encrypt(t *testing.T) {
	encrypt := MD5Encrypt("admin")
	if encrypt != "21232f297a57a5a743894a0e4a801fc3" {
		t.Error("md5 admin error")
	}
}
