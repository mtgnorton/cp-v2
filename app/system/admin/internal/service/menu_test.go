package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestMenu_Tree(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tree, err := Menu.Tree(ctx, 0)
		t.AssertNil(err)
		g.Dump(tree)
	})

}
func TestMenu_Tree2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {

		tree, err := Menu.Tree(ctx, 0, 1)
		t.AssertNil(err)
		g.Dump(tree)

	})
}
func TestMenu_MiniTree(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		miniTree, err := Menu.MiniTree(ctx)
		t.AssertNil(err)
		g.Dump(miniTree)
	})
}

func TestMenu_FrontRoutes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		routes, err := Menu.FrontRoutes(ctx, 1)
		t.AssertNil(err)
		g.Dump("TestMenu_FrontRoutes:", routes)
	})
}
