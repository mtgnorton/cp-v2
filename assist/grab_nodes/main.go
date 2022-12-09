package main

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/shared"

	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"

	"github.com/gogf/gf/v2/os/gfile"
)

func main() {
	ctx := context.Background()
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-local.toml")

	c := gfile.GetContents("/Users/mtgnorton/Coding/go/src/github.com/mtgnorton/gf-admin/assist/grabNodes/nodes.txt")

	r, err := gregex.ReplaceStringFunc("<table(.+?)</table>", c, func(match string) string {

		// 抓取分类
		category, err := gregex.MatchString(`<span class="fade">(.*?)</span>`, match)
		if len(category) == 2 {
		}
		if err != nil {
			g.Log().Fatal(ctx, err)
		}

		var categoryId int64
		// 判断是否存在
		categoryIdVar, err := dao.NodeCategories.Ctx(ctx).Where("name", category[1]).Value("id")
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		categoryId = categoryIdVar.Int64()

		if categoryId == 0 {

			categoryId, err = dao.NodeCategories.Ctx(ctx).InsertAndGetId(g.Map{
				"name":                category[1],
				"is_index_navigation": 1,
			})
			if err != nil {
				g.Log().Fatal(ctx, err)
			}
			g.Log().Infof(ctx, "分类%s抓取完成", category[1])
		}

		nodes, err := gregex.MatchAllString(`<a href="(.*?)" style="font-size: 14px;">(.*?)</a>`, match)

		parentId, err := dao.Nodes.Ctx(ctx).InsertAndGetId(g.Map{
			"category_id": categoryId,
			"name":        category[1],
			"keyword":     category[1],
			"img":         "",
			"description": "",
			"parent_id":   0,
		})

		if err != nil {
			g.Log().Error(ctx, err)
		}
		for _, node := range nodes {

			nodeKeyword := gfile.Basename(node[1])
			nodeUrl := `https://www.v2ex.com` + node[1]
			nodeName := node[2]
			g.Log().Infof(ctx, "节点%s开始抓取", nodeName)

			nodeCountVar, err := dao.Nodes.Ctx(ctx).Where("keyword", nodeKeyword).Value("id")

			if err != nil {
				g.Log().Error(ctx, err)
			}
			if nodeCountVar.Int() > 0 {
				g.Log().Infof(ctx, "节点%s已存在", nodeName)

				continue
			}

			//<div class="cell page-content-header">
			//<img src="https://cdn.v2ex.com/navatar/c399/862d/534_large.png?m=1403278091" border="0" align="default" width="64" alt="soccer" />
			//<div>
			//<div class="title">
			//<div class="node-breadcrumb"><a href="/">V2EX</a> <span class="chevron">&nbsp;›&nbsp;</span> 绿茵场</div>
			//<span class="topic-count">主题总数 <strong>268</strong></span>
			//</div>
			//<div class="intro">Brazil 2014</div>

			html := g.Client().GetContent(ctx, nodeUrl, nil)
			desc, err := gregex.MatchString(`<div class="intro">(.*?)</div>`, html)
			if err != nil {
				g.Log().Fatal(ctx, err)
			}

			description := ""
			if len(desc) == 2 {
				description = desc[1]
			}
			nodeImage := ""

			nodeImgUrl, err := gregex.MatchString(`page-content-header">
<img src="(.*?)"`, html)

			if len(nodeImgUrl) > 0 {
				out, err := shared.Download.Image(ctx, &model.DownloadImageInput{
					Url: nodeImgUrl[1],
					Dir: "node",
				})
				if err != nil {
					g.Log().Error(ctx, err)
				}
				if out != nil {
					nodeImage = out.RelativePath
				}
			}

			_, err = dao.Nodes.Ctx(ctx).InsertAndGetId(g.Map{
				"category_id": categoryId,
				"name":        nodeName,
				"keyword":     nodeKeyword,
				"img":         nodeImage,
				"description": description,
				"parent_id":   parentId,
			})

			if err != nil {
				g.Log().Fatal(ctx, err)
			}

			g.Log().Infof(ctx, "节点%s抓取完成", nodeName)

		}

		return ""
	})
	g.Dump(r, err)
}
