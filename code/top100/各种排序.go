package main

import "fmt"

//给定一个无序单链表，实现单链表的排序(按升序排序)。

//排序的算法比较常见的有 冒泡排序 快速排序 归并排序 插入排序 堆排序

//https://blog.csdn.net/bigbaochen/article/details/119611625

//冒泡排序，就是双重for循环，俩俩比较相邻的元素
func maopao(nums []int) []int {
	swap := func(nums []int,i,j int) {
		num := nums[i]
		nums[i] = nums[j]
		nums[j] = num
	}
	flag := true //优化点：增加标志位，如果在中途没有比较置换过，说明这个数组已经排好序了，这样退出就好了，减少时间复杂度
	for i := 0; i < len(nums) && flag; i++ {
		flag = false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums,j,j+1)
				flag = true
			}
		}
	}
	return nums
}

//快排的核心：基准数 从数组中选择一个基准数（一般是最开头的内个元素）让其他比基准数小的元素移动到一边，大的移动另一边。 在对左右两边的区间重复上一次操作，直到区间中剩下一个元素 （就是分治法的思想）
func kuaipaiDG(nums []int) []int {
	//比较置换
	swap := func(nums []int,i,j int) {
		num := nums[i]
		nums[i] = nums[j]
		nums[j] = num
	}
	//选出基准数，进行分区
	partition := func(nums []int,left,right int) int{
		start := left
		left++
		for left < right {
			for left < right && nums[left] < nums[start] {
				left++
			}
			for left < right && nums[right] > nums[start] {
				right--
			}
			if left >= right {
				break
			}
			swap(nums,left,right)
		}
		//如果当前值比基准元素还大，那将基本元素和当前值的前一位置换
		if nums[right] > nums[start] {
			right--
		}
		swap(nums,start,right)
		return right
	}
	var sort func(nums []int,left,right int)
	sort = func(nums []int,left,right int) {
		if left < right {
			par := partition(nums, left, right)
			sort(nums,left,par-1)
			sort(nums,par+1,right)
		}
	}
	sort(nums,0,len(nums)-1)
	return nums
}
//快排非递归
func kuaipaiFDG(nums []int) []int {
	//比较置换
	swap := func(nums []int,i,j int) {
		num := nums[i]
		nums[i] = nums[j]
		nums[j] = num
	}
	//选出基准数，进行分区
	partition := func(nums []int,left,right int) int{
		start := left
		left++
		for left < right {
			for left < right && nums[left] < nums[start] {
				left++
			}
			for left < right && nums[right] > nums[start] {
				right--
			}
			if left >= right {
				break
			}
			swap(nums,left,right)
		}
		//如果当前值比基准元素还大，那将基本元素和当前值的前一位置换
		if nums[right] > nums[start] {
			right--
		}
		swap(nums,start,right)
		return right
	}
	//使用栈，栈里存储这每一次基准数的分区区间 每一次弹栈出来一个区间，在继续寻找基准数，区分分区...以此类推，直到每一个分区都是一个元素
	var stack []int
	//先把起始和结束放进去
	stack = append(stack,0)
	stack = append(stack,len(nums)-1)

	for len(stack) != 0 {
		right := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if left < right {
			par := partition(nums,left,right)
			//压栈
			stack = append(stack,left)
			stack = append(stack,par-1)
			stack = append(stack,par+1)
			stack = append(stack,right)
		}
	}
	return nums
}

//归并排序 核心也是分治的思想 不同于快排的是 它是先处理子问题，在合并成分区.
func guibingDG(nums []int) []int {
	merge := func(nums []int, left,right,mid int) {
		//首选，创建一个存储归并后的大集合数组
		var mergeNums []int
		//其次，用两个指针，分别在两个区间内比较，归并两个区间
		tmp,tmp1 := left,mid+1
		for tmp <= mid || tmp1 <= right {
			if tmp > mid {
				mergeNums = append(mergeNums,nums[tmp1])
				tmp1++
				continue
			}
			if tmp1 > right {
				mergeNums = append(mergeNums,nums[tmp])
				tmp++
				continue
			}
			//俩俩比较落入新的集合数组
			if nums[tmp] <= nums[tmp1] {
				mergeNums = append(mergeNums,nums[tmp])
				tmp++
			}else {
				mergeNums = append(mergeNums,nums[tmp1])
				tmp1++
			}
		}
		//最后，将大集合 还原到原先数组中
		index := 0
		for i := left; i <= right && len(mergeNums) > 0; i++ {
			nums[i] = mergeNums[index]
			index++
		}
	}
	var sort func(nums []int, left,right int)
	sort = func(nums []int, left,right int) {
		if left < right {
			mid := left + (right-left) / 2
			sort(nums,left,mid)
			sort(nums,mid+1,right)
			merge(nums,left,right,mid)
		}
	}
	sort(nums,0,len(nums)-1)
	return nums
}
//归并排序 非递归
func guibingFDG(nums []int) []int {
	merge := func(nums []int, left,right,mid int) {
		//首选，创建一个存储归并后的大集合数组
		var mergeNums []int
		//其次，用两个指针，分别在两个区间内比较，归并两个区间
		tmp,tmp1 := left,mid+1
		for tmp <= mid || tmp1 <= right {
			if tmp > mid {
				mergeNums = append(mergeNums,nums[tmp1])
				tmp1++
				continue
			}
			if tmp1 > right {
				mergeNums = append(mergeNums,nums[tmp])
				tmp++
				continue
			}
			//俩俩比较落入新的集合数组
			if nums[tmp] <= nums[tmp1] {
				mergeNums = append(mergeNums,nums[tmp])
				tmp++
			}else {
				mergeNums = append(mergeNums,nums[tmp1])
				tmp1++
			}
		}
		//最后，将大集合 还原到原先数组中
		index := 0
		for i := left; i <= right && len(mergeNums) > 0; i++ {
			nums[i] = mergeNums[index]
			index++
		}
	}
	k := 1  //子集合大小 1，2，4，8....
	for k < len(nums) {
		var i int
		for i = 0; i < len(nums)-2*k; i += 2*k {
			merge(nums,i,i+2*k-1,i+k-1)
		}
		if  i+k < len(nums) {
			merge(nums,i,len(nums)-1,i+k-1)
		}
		k = 2*k
	}
	return nums
}

//插入排序 思路：将一个元素插入到已经排好序的序列中，得到一个新的序列。 也就是说我们要将原序列分成两个区间，有序区间和无序区间，每次从无序区间中取一个，插入到有序区间
func charu(nums []int) []int {
	//把第一个元素当做是有序区间 , 所以i=1 从1开始
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		var j int
		//有序区间从后往前，和元素比较并移动位置，为元素腾地方
		for j = i - 1; j >= 0; j-- {
			if tmp < nums[j] {
				nums[j+1] = nums[j]
				continue
			}
			break
		}
		//归位
		nums[j+1] = tmp
	}
	return nums
}

//堆排序
func heap(nums []int) []int {
	//比较置换
	swap := func(nums []int,i,j int) {
		num := nums[i]
		nums[i] = nums[j]
		nums[j] = num
	}
	//下沉建堆
	sink := func(nums []int, root,end int) {
		//如果该root没有左子节点了，说明下沉到位了
		for 2*root <= end {
			left := 2*root //左子节点的坐标
			right := left+1 //右子节点的坐标
			index := left
			//找出root的左右两个子节点的最大值（大顶堆），之后再和当前root值比较
			if  right <= end && nums[right] > nums[left]{
				index = right
			}
			if nums[index] > nums[root] {
				swap(nums,root,index)
			}else {
				break
			}
			root = index
		}
	}
	//排序 先创建一个新数组，起始位是空余节点
	newNums := []int{0}
	for i := 0; i < len(nums); i++ {
		newNums = append(newNums,nums[i])
	}
	//先下沉建堆，将数组变成一个大顶堆，找出所有可能的顶点..最小的是 len(nums)/2
	end := len(newNums) -1
	for i := len(nums)/2; i >= 1; i-- {
		sink(newNums,i,end)
	}
	for end > 1 {
		swap(newNums,1,end)
		end--
		sink(newNums,1,end)
	}
	for i := 1; i < len(newNums); i++ {
		nums[i-1] = newNums[i]
	}
	return nums
}

func main() {
	fmt.Println(maopao([]int{1,3,2,1,4,3,7,5}))
	fmt.Println(kuaipaiDG([]int{1,3,2,1,4,3,7,5}))
	fmt.Println(kuaipaiFDG([]int{1,3,2,1,4,3,7,5}))
	fmt.Println(guibingDG([]int{1,3,2,1,4,3,7,5}))
	fmt.Println(guibingFDG([]int{1,3,2,1,4,3,7,5}))
	fmt.Println(charu([]int{1,3,2,1,4,3,7,5}))
	fmt.Println(heap([]int{1,3,2,1,4,3,7,5}))
}
