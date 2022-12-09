package utility

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
)

func TestDFAMatcher(t *testing.T) {

	SensitiveInspector.AddSensitiveWords([]string{"妓女", "文件"})

	sensitiveWords, after := SensitiveInspector.MatchAndReplace("谁是 妓 女a妓女")

	g.Dump(sensitiveWords, after)

}
