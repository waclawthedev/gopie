package gopie_test

import (
	"testing"

	"github.com/waclawthedev/gopie"
)

func TestMD5(t *testing.T) {
	shouldEqual(t, gopie.MD5("hello world"), "5eb63bbbe01eeed093cb22bb8f5acdc3")

	// Two different strings have the same md5, don't use md5 for cryptography :)
	shouldEqual(t,
		gopie.MD5("TEXTCOLLBYfGiJUETHQ4hAcKSMd5zYpgqf1YRDhkmxHkhPWptrkoyz28wnI9V0aHeAuaKnak"),
		gopie.MD5("TEXTCOLLBYfGiJUETHQ4hEcKSMd5zYpgqf1YRDhkmxHkhPWptrkoyz28wnI9V0aHeAuaKnak"))
}
