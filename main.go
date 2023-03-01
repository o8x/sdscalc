package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/o8x/jk/args"
	"github.com/o8x/jk/signal"
)

var out *os.File

func init() {
	f, err := os.OpenFile("sds.out", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err == nil {
		out = f
	}
}

func write(s string) {
	os.Stdout.WriteString(s)

	if out != nil {
		out.WriteString(s)
	}
}

func writeln(s string) {
	write(s)
	write("\r\n")
}

func main() {
	a := args.Args{
		App: &args.App{
			Name:      "SDS 测试评分程序",
			Usage:     "./sdscalc [-d data.yaml]",
			Copyright: "Alex. stdout.com.cn",
		},
		Flags: []*args.Flag{
			{
				Name:        []string{"-d", "-data", "--datasource"},
				Description: "数据源 yaml 文件",
				Default:     []string{"datasource.yaml"},
				Required:    true,
				SingleValue: true,
			},
		},
	}

	if err := a.Parse(); err != nil {
		a.PrintHelpExit(err)
	}

	InitConfig(a.GetX("data"))

	writeln("抑郁自评量表（Self-rating depression scale，SDS），是含有20个项目，分为4级评分的自评量表，原型是W.K.Zung编制的抑郁量表（1965）。\r\n其特点是使用简便，并能相当直观地反映抑郁患者的主观感受及其在治疗中的变化。\r\n主要适用于具有抑郁症状的成年人，包括门诊及住院患者。")
	writeln("")
	writeln("接下来将开始测试，请认真阅读每一个问题并根据你最近一周的实际情况，选择合适的答案。")
	writeln("本次测试过程以及结果，将会完整记录到文件 sds.out 中。")
	writeln("在测试过程中，你可以随时使用 Ctrl+C 停止测试。")
	writeln("")

	type answer struct {
		Index  int    `json:"index"`
		Answer string `json:"answer"`
	}

	var answers []*answer
	optionMap := []string{"A", "B", "C", "D"}

	subjects := c.Subject

	for i, it := range subjects {

		write(fmt.Sprintf("%s (%d/%d)\r\n", it.Title, i+1, len(subjects)))
		write(fmt.Sprintf("此问题用于评估是否：%s\r\n", it.Symptom))
		writeln("选项：")

		for j, o := range c.Options {
			write(fmt.Sprintf("    %s. %s\r\n", optionMap[j], o))
		}
		a2 := &answer{Index: i, Answer: ""}
		answers = append(answers, a2)
		for {
			write(fmt.Sprintf("对于第 %d 题你的选择是？", i+1))
			if _, err := fmt.Scanf("%s", &a2.Answer); err != nil {
				writeln(fmt.Sprintf("输入错误，请重新输入, %s", err))
				continue
			}

			if out != nil {
				// 将来自 stdin 的输入也写入文件
				out.WriteString(a2.Answer)
				out.WriteString("\r\n")
			}

			a2.Answer = strings.ToUpper(a2.Answer)
			right := false
			for _, s := range optionMap {
				if s == a2.Answer {
					right = true
					break
				}
			}
			if right {
				writeln("")
				break
			}
			writeln("选项仅为 A/a B/b C/c D/d 中的一个，请重新输入你的选择")
		}
	}

	writeln("")
	writeln("你已经回答了本次 SDS 测试询问的所有问题。")
	writeln("")
	writeln("你的答案：")

	positiveScoreMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}
	scoreMap := map[string]int{"A": 4, "B": 3, "C": 2, "D": 1}
	optionIndexMap := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}

	score := 0
	for i, it := range answers {
		write(fmt.Sprintf("%2d. %s: %s\r\n", i+1, c.Options[optionIndexMap[it.Answer]], subjects[it.Index].Title))

		if subjects[it.Index].Positive {
			score += positiveScoreMap[it.Answer]
		} else {
			score += scoreMap[it.Answer]
		}
	}

	writeln("")
	writeln(`得分标准：
标准分分界值为 53 分，总粗分的正常上限为 41 分，分值越低状态越好。
标准分为总粗分乘以 1.25 结果的整数部分。
我国以标准分 ≥50 为有抑郁症状。
  - 53-62 分为轻度抑郁
  - 63-72 分为中度抑郁
  - 73 分以上为重度抑郁`)
	writeln("")
	writeln(`抑郁严重度标准：
各条目累计分 / 80
  - 0.5 以下为无抑郁
  - 0.5—0.59 轻微至轻度抑郁
  - 0.6—0.69 中至重度
  - 0.7 以上为重度抑郁`)

	stdScore := math.Ceil(float64(score) * 1.25)
	severity := float64(score) / 80
	writeln("")
	write(fmt.Sprintf("分数计算结果为粗分：%d, 标准分: %.0f，抑郁严重度: %.2f。\r\n", score, stdScore, severity))
	writeln("")

	signal.Wait()
}
