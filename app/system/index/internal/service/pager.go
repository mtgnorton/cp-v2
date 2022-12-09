package service

import (
	"fmt"
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

var Pager = pager{}

type pager struct {
}

//  <div class="pager">
//                <a class="pager-current" href="">1</a>
//                <a class="pager-item" href="">2</a>
//                <a class="pager-item" href="">3</a>
//                <a class="pager-item" href="">4</a>
//                <a class="pager-item" href="">5</a>
//                <span>...</span>
//                <a class="pager-item" href="">1080</a>
//                <input type="text" value="1">
//            </div>

// Pager 生成分页html代码
func (p *pager) Pager(req *model.PagerReq) (res *model.PagerRes) {
	res = &model.PagerRes{}

	res.CurrentPage = req.CurrentPage
	res.Size = req.Size
	res.TotalRow = req.TotalRow
	res.ShowPageAmount = req.ShowPageAmount

	if req.TotalRow == 0 {
		return
	}

	res.BeginIndex = (req.CurrentPage-1)*req.Size + 1
	res.EndIndex = req.CurrentPage * req.Size
	if res.EndIndex > req.TotalRow {
		res.EndIndex = req.TotalRow
	}

	if req.TotalRow < req.Size {
		return res
	}

	totalPage := req.TotalRow / req.Size

	if req.TotalRow%req.Size > 0 {
		totalPage++
	}
	if req.CurrentPage > totalPage {
		req.CurrentPage = totalPage
	}
	var (
		leftPage, rightPage int
	)

	if totalPage <= req.ShowPageAmount {
		leftPage = 1
		rightPage = totalPage
	} else {

		showPageAmountHalf := req.ShowPageAmount / 2
		// 避免uint越界
		if req.CurrentPage < showPageAmountHalf {
			leftPage = 1
		} else {
			leftPage = req.CurrentPage - showPageAmountHalf
		}

		rightPage = req.CurrentPage + showPageAmountHalf

		if rightPage > totalPage {
			rightPage = totalPage
		}
	}

	g.Dump(fmt.Sprintf("当前页为：%d,总页数为：%d,左边界为：%d,右边界为：%d", req.CurrentPage, totalPage, leftPage, rightPage))
	for i := leftPage; i <= rightPage; i++ {
		g.Dump(fmt.Sprintf("循环页面为:%d", i))
		if i == req.CurrentPage {
			res.Html += fmt.Sprintf(` <a class="pager-current" href="%s">%d</a>`, fmt.Sprintf(req.Url, i), i)
		} else {
			res.Html += fmt.Sprintf(` <a class="pager-item" href="%s">%d</a>`, fmt.Sprintf(req.Url, i), i)
		}

	}
	if leftPage != 1 {
		res.Html = fmt.Sprintf(` <a class="pager-item" href="%s">1</a><span>...</span>%s`, fmt.Sprintf(req.Url, 1), res.Html)
	}
	if rightPage != totalPage {
		res.Html += fmt.Sprintf(` <span>...</span><a class="pager-item" href="%s">%d</a>`, fmt.Sprintf(req.Url, totalPage), totalPage)
	}

	res.Html = fmt.Sprintf(`<div class="pager bottom-line" >%s</div>`, res.Html)

	return
}
