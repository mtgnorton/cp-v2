package utility

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/test/gtest"
)

//
//func TestPager(t *testing.T) {
//	html := Pager(1, 20, 100, "/page/%d", 5)
//	g.Dump(html)
//
//	html = Pager(30, 20, 10000, "/page/%d", 10)
//	g.Dump(html)
//
//	html = Pager(1, 20, 100, "/page/%d", 4)
//	g.Dump(html)
//
//	html = Pager(7, 20, 100, "/page/%d", 5)
//	g.Dump(html)
//}

func TestReplaceWarp(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		str, err := ReplaceWarp("123\r\n456")
		t.Assert(err, nil)
		g.Dump(str)
	})
}
