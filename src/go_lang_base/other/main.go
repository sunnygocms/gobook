package main

import (
	"fmt"

	"./singlechain/"
)

func main() {
	var snArray []singlechain.SunnyNavigation
	sn := singlechain.SunnyNavigation{1, 0, 0, "新闻管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{52, 0, 0, "推荐新闻管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{55, 0, 0, "频道管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{57, 0, 0, "广告&友链管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{49, 0, 0, "幻灯"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{44, 0, 0, "系统管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{2, 1, 1, "新闻添加"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{3, 1, 1, "新闻列表"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{45, 1, 44, "用户管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{46, 1, 44, "用户组管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{47, 1, 44, "权限管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{48, 1, 44, "菜单管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{51, 1, 49, "幻灯管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{53, 1, 52, "推荐新闻列表"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{54, 1, 52, "推荐分类列表"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{56, 1, 55, "频道管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{60, 1, 55, "城市信息管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{58, 1, 57, "广告列表"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{59, 1, 57, "友链的管理"}
	snArray = append(snArray, sn)
	sn = singlechain.SunnyNavigation{222, 1, 3, "新闻333"}
	snArray = append(snArray, sn)
	//	fmt.Println(snArray)
	for _, sn := range snArray {
		singlechain.Insert(sn)
	}
	fmt.Println("length---------", singlechain.GetLength())
	d := singlechain.GetFirst()
	fmt.Println(d.Data.Name)
	d = singlechain.GetLast()
	fmt.Println(d.Data.Name)

}
