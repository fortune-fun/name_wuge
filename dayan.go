package name_wuge

import (
	"strings"

	"github.com/godcong/yi"
)

type WuGeDaYan struct {
	TianDaYan string
	RenDaYan  string
	DiDaYan   string
	WaiDaYan  string
	ZongDaYan string
}

func (dy *WuGeDaYan) IsLucky(filter_hard bool) bool {
	var is_lucky bool = false

	// ignore gd.TianDaYan

	if strings.Contains(dy.RenDaYan, "吉") && !strings.Contains(dy.RenDaYan, "凶") &&
		strings.Contains(dy.DiDaYan, "吉") && !strings.Contains(dy.DiDaYan, "凶") &&
		strings.Contains(dy.ZongDaYan, "吉") && !strings.Contains(dy.ZongDaYan, "凶") {

		if filter_hard {
			if !strings.Contains(dy.ZongDaYan, "半") && !strings.Contains(dy.ZongDaYan, "平") &&
				strings.Contains(dy.WaiDaYan, "吉") && !strings.Contains(dy.WaiDaYan, "凶") {
				is_lucky = true
			}
		} else {
			is_lucky = true
		}

	}

	return is_lucky
}

func (ge *WuGe) GetWuGeDaYan(sex yi.Sex) *WuGeDaYan {
	wuGeDaYan := WuGeDaYan{}

	if daYan := yi.GetDaYan(ge.SanCaiNum.Tian); sex == yi.SexGirl && strings.Contains(daYan.NvMing, "凶") {
		wuGeDaYan.TianDaYan = daYan.NvMing
	} else {
		wuGeDaYan.TianDaYan = daYan.Lucky
	}

	if daYan := yi.GetDaYan(ge.SanCaiNum.Ren); sex == yi.SexGirl && strings.Contains(daYan.NvMing, "凶") {
		wuGeDaYan.RenDaYan = daYan.NvMing
	} else {
		wuGeDaYan.RenDaYan = daYan.Lucky
	}

	if daYan := yi.GetDaYan(ge.SanCaiNum.Di); sex == yi.SexGirl && strings.Contains(daYan.NvMing, "凶") {
		wuGeDaYan.DiDaYan = daYan.NvMing
	} else {
		wuGeDaYan.DiDaYan = daYan.Lucky
	}

	if daYan := yi.GetDaYan(ge.WaiGe); sex == yi.SexGirl && strings.Contains(daYan.NvMing, "凶") {
		wuGeDaYan.WaiDaYan = daYan.NvMing
	} else {
		wuGeDaYan.WaiDaYan = daYan.Lucky
	}

	if daYan := yi.GetDaYan(ge.ZongGe); sex == yi.SexGirl && strings.Contains(daYan.NvMing, "凶") {
		wuGeDaYan.ZongDaYan = daYan.NvMing
	} else {
		wuGeDaYan.ZongDaYan = daYan.Lucky
	}

	return &wuGeDaYan
}
