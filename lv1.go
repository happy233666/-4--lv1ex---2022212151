package main

import "fmt"

var a = 0
var b string = "鸡你的太美"

func main() {
LOOP:
	fmt.Println("请选择你的技能类型")
	fmt.Println("1.正常型，2.玄幻小说型，3.装逼型")
	fmt.Scan(&a)
	if a > 3 || a < 0 {
		fmt.Println("超过你的能力啦")
	} else {
		switch a {
		case 0:
			fmt.Println("你啥也没干啊")
			goto LOOP
		case 1:
			fmt.Println("你使用的技能名为")
			fmt.Scan(&b)
			ReleaseSkill(b, func(skillName string) {
				fmt.Println("尝尝我的厉害吧！", skillName)
			})
		case 2:
			fmt.Println("你使用的技能名为")

			fmt.Scan(&b)
			ReleaseSkill(b, func(skillName string) {
				fmt.Println("蝼蚁！", skillName)
			})
		case 3:
			fmt.Println("你使用的技能名为")
			fmt.Scan(&b)
			ReleaseSkill(b, func(skillName string) {
				fmt.Println("从此我将立于天之上！", skillName)
			})
		}
	}
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
