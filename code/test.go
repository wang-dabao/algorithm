package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ints := make([]int, 2, 3)
	ints = append(ints, 1, 2, 3, 4, 5)
	fmt.Printf("jj:%d,%d,%+v", len(ints), cap(ints), ints)

}

type name struct {
	Age int `json:"age"`
}

type StationBusOrders []*name

// Len  实现 sort.Interface 接口的 Len 方法
func (s StationBusOrders) Len() int {
	return len(s)
}

// Less 实现 sort.Interface 接口的 Less 方法，用于定义排序规则
func (s StationBusOrders) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// Swap 实现 sort.Interface 接口的 Swap 方法
func (s StationBusOrders) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main2() {
	n1 := &name{
		Age: 10,
	}
	n2 := &name{
		Age: 16,
	}
	names := []*name{n1, n2}
	sort.Sort(StationBusOrders(names))
	marshal, _ := json.Marshal(names)
	fmt.Printf("jj:%+v", string(marshal))
}

func GetTableNameByPid(baseTable string, pid string) string {
	passengerId := strings.Trim(pid, " ")
	if strings.HasSuffix(passengerId, "000000") {
		passengerId = passengerId[0:6]
	}
	size := len(passengerId)
	zeroFill := "000000"
	if size < 6 {
		passengerId += zeroFill[0 : 6-size]
		size = 6
	}
	passengerId = passengerId[size-6 : size]
	passengerId = strings.TrimLeft(passengerId, "0")
	if len(passengerId) < 6 {
		fillSize := 6 - len(passengerId)
		passengerId += zeroFill[0:fillSize]
	}
	partNum, _ := strconv.Atoi(passengerId)
	tableSuffix := partNum % 11
	return baseTable + "_" + strconv.Itoa(tableSuffix)
}
