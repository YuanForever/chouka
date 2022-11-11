package main

import (
	"fmt"
	"math/rand"
	"time"
)

var zi [10]int     //随机数组
var n [10]int      //紫和金的数组
var choose1 string //抽卡模式
var chance int = 0 //抽卡次数
var choose2 string //是否充值
var choose3 string //抽卡模式
var choose4 string //抽卡是否继续
var zhiling string //次数获取指令
var line string    //是否继续指令
var up [90]int     //卡池
var z int = 0      //获取次数后结束循环
var more int = 0   //增加的抽卡次数

func kachi(map[int]string) map[int]string {
	m := map[int]string{}
	rand.Seed(time.Now().UnixNano())
	chu := chouka(zi)
	for i := 0; i < 9; i++ {
		m[chu[i]] = "出紫"
	}
	m[chu[9]] = "出金"
	return m
} //将数组里的数写进map中
func chouka(n [10]int) [10]int {
	for i := 0; i < 10; i++ {
		n[i] = rand.Intn(89) + 1
		for j := 0; j < i; j++ {
			if n[i] == n[j] {
				n[i] = rand.Intn(89) + 1
				j = 0
			}
		}
	}
	return n
} //随机的十个不重复的数写进一个数组
func cishu(m int) int {
	more = 0
	for i := 1; ; i++ {
		fmt.Println("请输入次数获取指令：鑫姐姐好漂亮/袁神最帅")
		fmt.Scan(&zhiling)
		if zhiling == "鑫姐姐好漂亮" {
			more += 1
			fmt.Println("恭喜你获取成功，抽卡次数+1")
		} else if zhiling == "袁神最帅" {
			more += 10
			fmt.Println("恭喜你获取成功，抽卡次数+10")
		}
		fmt.Println("是否继续获取次数？")
		fmt.Scan(&line)
		if line == "否" {
			return more
		}
	}
} //获取抽卡次数
func dan(n int, m [90]int) int {
	for g := 0; g < n; g++ {
		if m[g] == m[n] {
			m[n] = rand.Intn(89) + 1
			g = 0
		}
	}
	return m[n]
} //单抽抽卡
func main() {
	for q := 0; q < 90; q++ {
		up[q] = 0
	}
	for ti := 1; ; ti++ {
		rand.Seed(time.Now().UnixNano())
		zi = chouka(n)
		m := kachi(map[int]string{}) //卡池里所有的紫和金
		fmt.Println("请选择抽卡：" +
			"单抽/十连")
		for i := 0; ; {
			fmt.Scan(&choose1)
			if choose1 == "单抽" {
				if chance < 1 { //单抽次数不足
					fmt.Println("抽卡次数不足，请获取足够的抽卡次数后再次抽卡\n" +
						"是否获取抽卡次数？")
					fmt.Scan(&choose2)
					if choose2 == "是" {
						z = 1
						chance += cishu(more)
						break
					} else {
						fmt.Println("钱都不充，你抽个锤子")
						break
					}
				}
				if chance >= 1 { //单抽次数足够
					z = 0
					chance -= 1
					up[i] = dan(i, up)
				}
				for d := 0; d < 10; d++ {
					if d < 9 && zi[d] == up[i] { //出紫
						fmt.Println(m[zi[d]])
						break
					} else if d == 9 && zi[d] == up[i] { //出金
						fmt.Println(m[zi[d]])
						for w := 0; w < 90; w++ {
							up[w] = 0
						}
						if i > 45 {
							fmt.Println("过半了才出金，你个非酋，还老想着抽卡")
						} else if i <= 45 {
							fmt.Println("出金这么早，死神都说好")
						}
						i = 0
						break
					} else if d == 9 && zi[d] != up[i] {
						fmt.Println("蓝色")
					}
				}
				i++
				fmt.Println("是否继续单抽？")
				fmt.Scan(&choose3)
				if choose3 == "是" {
					continue //这里我不知道该怎么写了，就是让他继续执行循环，这样写的话要输入一次后再输入一次才会进入下一次循环
				} else {
					break
				}
			} else {
				if chance < 10 { //十连次数不足
					fmt.Println("抽卡次数不足，请获取足够的抽卡次数后再次抽卡\n" +
						"是否获取抽卡次数？")
					fmt.Scan(&choose2)
					if choose2 == "是" {
						z = 1
						chance += cishu(more)
						break
					} else {
						fmt.Println("钱都不充，你抽个锤子")
						break
					}
				} else { //十连次数足够
					z = 0
					chance -= 10
					for tim := 10; tim > 0; i++ {
						tim--
						up[i] = rand.Intn(89) + 1
						for j := 0; j < i; j++ {
							if up[j] == up[i] {
								up[i] = rand.Intn(89) + 1
							}
						}
						for d := 0; d < 10; d++ {
							if d < 9 && zi[d] == up[i] { //出紫
								fmt.Println(m[zi[d]])
								break
							} else if d == 9 && zi[d] == up[i] { //出金
								fmt.Println(m[zi[d]])
								for w := 0; w < 90; w++ {
									up[w] = 0
								}
								if i > 45 {
									fmt.Println("过半了才出金，你个非酋，还老想着抽卡")
								} else if i <= 45 {
									fmt.Println("出金这么早，阎王都说好")
								}
								i = 0
								break
							} else if d == 9 && zi[d] != up[i] {
								fmt.Println("蓝色")
								break
							}
						}
					}
					break
				}
			}
			if z == 1 {
				break
			}
		}
		if z == 0 {
			fmt.Println("是否继续抽卡？")
			fmt.Scan(&choose4)
		}
		if choose4 == "否" {
			break
		}
	}
}
