package main

import (
	"fmt"
)

type elementTp int

//冒泡排序，时间复杂度O(n2)
func bubble(data []elementTp) []elementTp {
	var tm elementTp
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				tm = data[i]
				data[i] = data[j]
				data[j] = tm
			}
		}
	}
	return data
}

//快速排序，最坏情况时间复杂度O(n2),eg:此时序列是已排过序的逆序
//最好情况和平均情况都算做O(nlogn)
//它的基本思想是：通过一趟排序将数据一分为二，其中一部分的所有数据都比另外一部分的所有数据都要小，然后对两部分递归，直至完成
//要避免最坏情况的出现，需要对枢纽元选择优化
func quickSort(data []elementTp) {
	if len(data) < 2 {
		return
	}
	var pivot elementTp
	//fmt.Println(data)
	pivot = data[0]
	var i, j = 0, len(data) - 1

	for i < j {
		for i < j && data[j] >= pivot {
			j--
		}
		data[i] = data[j]
		for i < j && data[i] <= pivot {
			i++
		}
		data[j] = data[i]
	}
	data[i] = pivot
	quickSort(data[:i])
	quickSort(data[j+1:])
}

//插入排序
//时间复杂度0(N2)
//由N-1趟排序组成，每趟排序都是选取值与有序序列的比较
func insertSort(data []elementTp) {
	var j int
	for i := 1; i < len(data); i++ {
		temp := data[i]
		for j = i; j > 0; j-- {
			if temp < data[j-1] {
				data[j] = data[j-1]
			} else {
				break
			}
		}
		data[j] = temp
	}
}

//希尔排序
//时间复杂度: 使用shell增量时最差为O(n2);使用hibbard增量最差为O(n3/2)
//通过比较相距一定间隔的元素来工作，各趟比较使用的间隔随着算法进行减小，直到为1(比较相邻元素)
//它实际上是对插入排序的改进，分组插入排序
//增量序列的选择对于算法复杂度有一定影响
func shellSort(data []elementTp) {
	var increment int
	var sz = len(data)
	var j int

	for increment = sz / 2; increment > 0; increment /= 2 { //排序趟数为增量序列长
		for i := increment; i < sz; i++ { //向每组有序区域插入
			var tmp = data[i]
			for j = i; j >= increment; j -= increment { //
				if tmp < data[j-increment] {
					data[j] = data[j-increment]
				} else {
					break
				}

			}
			data[j] = tmp
		}
	}
}

//归并排序
//时间复杂度(最坏情况):O(nlogn)
//使用的是经典的分治策略(divide and conquer),合并两个已排序的表
//|_1_|_8_|_19_|  +  |_7_|_13_|_29_|_33_|_44_|  ==> |__|__|__|__|__|__|__|__|
//虽然复杂度较好，但是线性的数据拷贝还是消耗很大的，会影响排序的速度
func mergeSort(data []elementTp) {
	buf := make([]elementTp, len(data))
	mergesort(data, buf, 0, len(data)-1)
}

func mergesort(data, buf []elementTp, first, last int) {
	if first < last {
		mid := (last + first) / 2
		mergesort(data, buf, first, mid)
		mergesort(data, buf, mid+1, last)
		merge(data, buf, first, mid, last)
	}
}

//传参buf引用，省去每次merge时临时分配buf空间，malloc既占用内存也占用cpu资源
func merge(data, buf []elementTp, first, mid, last int) {
	//buf := make([]elementTp, last-first+1)
	//fmt.Printf("--------%d--%d\n", len(buf), last)
	var ptrA, ptrB, ptrC int = first, mid + 1, first
	for ptrA <= mid && ptrB <= last {
		if data[ptrA] < data[ptrB] {
			buf[ptrC] = data[ptrA]
			ptrA++
		} else {
			buf[ptrC] = data[ptrB]
			ptrB++
		}
		ptrC++
	}
	for ptrA <= mid {
		buf[ptrC] = data[ptrA]
		ptrA++
		ptrC++
	}
	for ptrB <= last {
		buf[ptrC] = data[ptrB]
		ptrB++
		ptrC++
	}

	copy(data[first:], buf[first:last+1])
}

func main() {
	var src = []elementTp{10, 8, 22, 7, 10, 33, 9, 22, 10, 3}
	shellSort(src)
	fmt.Println(src)
}
