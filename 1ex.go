package main

import (
	"fmt"
	"strings"
)

var m map[int]string

func main() {
	var a = 0
	var b string = "鸡你的太美"
	var d string = "鸡你的太美"
LOOP:
	fmt.Println("请选择你的技能类型")
	fmt.Println("1.正常型，2.玄幻小说型，3.装逼型,4.自定义")
	fmt.Scan(&a)
	if a > 4 || a < 0 {
		fmt.Println("超过你的能力啦")
	} else {
		switch a {
		case 0:
			fmt.Println("你啥也没干啊")
			goto LOOP
		case 1:
			fmt.Println("你使用的技能名为")
			fmt.Scan(&b)
			if charge(b) == true {
				goto LOOP
			}
			ReleaseSkill(b, func(skillName string) {
				fmt.Println("尝尝我的厉害吧！", skillName)
			})
		case 2:
			fmt.Println("你使用的技能名为")

			fmt.Scan(&b)
			if charge(b) == true {
				goto LOOP
			}
			ReleaseSkill(b, func(skillName string) {
				fmt.Println("蝼蚁！", skillName)
			})
		case 3:
			fmt.Println("你使用的技能名为")

			fmt.Scan(&b)
			if charge(b) == true {
				goto LOOP
			}
			ReleaseSkill(b, func(skillName string) {
				fmt.Println("从此我将立于天之上！", skillName)
			})
		case 4:
			fmt.Println("你使用的技能名为")
			fmt.Scan(&b)
			if charge(b) == true {
				goto LOOP
			}
			fmt.Println("请添加技能模板")
			fmt.Scan(&d)
			ReleaseSkill(b, func(skillName string) {
				fmt.Println(d, skillName)
			})
		}
	}
}
func charge(b string) bool {
	m := make(map[int]string)
	m[1] = "fuck"
	m[2] = "FUCK"
	m[3] = "你妈的"
	m[4] = "ikun"
	c := 1
	for i := 1; i <= 4; i++ {
		if strings.Contains(b, m[i]) == true {
			fmt.Println("敏感词汇")
			c = i
			break
		}
	}
	return strings.Contains(b, m[c])
}
func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
