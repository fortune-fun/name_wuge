package name_wuge

import (
	"strings"

	"github.com/godcong/yi"
)

type WuGeSanCai struct {
	sanCaiNum *yi.SanCaiNum
	tianRenDi string //三才
}

func NewWuGeSanCai(sancai_num *yi.SanCaiNum) *WuGeSanCai {
	wuGeSanCai := WuGeSanCai{
		sanCaiNum: sancai_num,
		tianRenDi: strings.Join([]string{yi.NumberWuXing(sancai_num.Tian),
			yi.NumberWuXing(sancai_num.Ren), yi.NumberWuXing(sancai_num.Di)}, ""),
	}

	return &wuGeSanCai
}

//五格三才幸运信息
type WuGeSanCaiFortune struct {
	fortune string //吉凶
	comment string //说明
}

//key为tianRenDi
var sanCaiList map[string]*WuGeSanCaiFortune = map[string]*WuGeSanCaiFortune{}

func init() {
	file_3wuxing, err := DataFiles.Open("data/3wuxing.csv")
	if err != nil {
		panic(err)
	}

	records, err := readCSV(file_3wuxing)

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		sancai := WuGeSanCaiFortune{
			fortune: record[1],
			comment: record[2],
		}

		sanCaiList[record[0]] = &sancai
	}
}

// 由三才数获得描述
func (sc *WuGeSanCai) GetFortune() *WuGeSanCaiFortune {
	if sanCaiList[sc.tianRenDi] == nil {
		panic(sc.tianRenDi)
	}

	return sanCaiList[sc.tianRenDi]
}

//Check 检查三才吉凶
func (sc *WuGeSanCaiFortune) IsLucky() bool {
	return sc.fortune == "吉"
}
