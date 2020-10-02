package file

import (
	"fmt"
	"github.com/wgj6112345/go_crawl/logger"
	"github.com/wgj6112345/go_crawl/model/baidu"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// 标题，简介，url，时间
var (
	f      *excelize.File
	lineNo int = 2
)

func InitExcel() {
	var (
		titles []string
		index  int
	)

	titles = []string{
		"标题", "简介", "url", "时间",
	}

	f = excelize.NewFile()
	index = f.NewSheet("Sheet1")
	// 设置标题
	for i, v := range titles {
		position := div(i+1) + "1"
		f.SetCellValue("Sheet1", position, v)
	}

	f.SetActiveSheet(index)
	logger.Logger.Infof("init excel file...\n")
}

func SaveToExcel(baiduItem baidu.BaiduItem) {
	// 批量插入数据
	f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", div(1), lineNo), baiduItem.Title)
	f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", div(2), lineNo), baiduItem.Intro)
	f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", div(3), lineNo), baiduItem.Url)
	f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", div(4), lineNo), baiduItem.UpdateTime)

	fmt.Println("Sheet1 ", fmt.Sprintf("%s%d", div(1), lineNo), baiduItem.Title)
	lineNo++
}

func SaveExcel(name string) {
	var (
		err error
	)

	if err = f.SaveAs(fmt.Sprintf("./%s.xlsx", name)); err != nil {
		fmt.Println("save xlsx file failed, err : ", err)
		return
	}

	fmt.Println("save xlsx file success")
}

// Div 数字转字母
func div(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}
