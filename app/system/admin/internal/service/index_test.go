package service

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestIndex_Statistics(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		r, err := Index.Statistics(ctx)
		g.Dump(r, err)
	})
}
