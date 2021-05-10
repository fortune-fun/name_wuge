package name_wuge

import (
	"strings"

	"github.com/godcong/yi"
)

// WuGe ...
type WuGe struct {
	SanCaiNum *yi.SanCaiNum
	WaiGe     int
	ZongGe    int
}

func (wg *WuGe) GetSanCai() *WuGeSanCai {
	wuGeSanCai := WuGeSanCai{
		sanCaiNum: wg.SanCaiNum,
		tianRenDi: strings.Join([]string{yi.NumberWuXing(wg.SanCaiNum.Tian),
			yi.NumberWuXing(wg.SanCaiNum.Ren), yi.NumberWuXing(wg.SanCaiNum.Di)}, ""),
	}

	return &wuGeSanCai
}
