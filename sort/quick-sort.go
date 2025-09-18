package sort

import "fmt"

// quickSort 对数组 arr 的指定范围 [left, right] 进行快速排序
func quickSort(arr []int, left, right int) {
	if left >= right { // 递归终止条件：子数组只有一个或没有元素
		return
	}

	// 选择中间元素作为基准 pivot
	pivot := arr[(left+right)/2]

	i, j := left, right // 定义左右指针

	// 分区操作：将小于 pivot 的移到左边，大于 pivot 的移到右边
	for i <= j {
		// 找到左边第一个大于或等于 pivot 的元素
		for arr[i] < pivot {
			i++
		}
		// 找到右边第一个小于或等于 pivot 的元素
		for arr[j] > pivot {
			j--
		}

		// 如果 i <= j，说明找到了需要交换的元素
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i] // 交换元素
			i++                             // 移动左指针
			j--                             // 移动右指针
		}
	}

	// 递归地对左右两个子数组进行排序
	if left < j {
		quickSort(arr, left, j)
	}
	if i < right {
		quickSort(arr, i, right)
	}
}

func main() {
	array := []int{34, 7, 23, 32, 5, 62}

	if len(array) <= 1 { // 处理空或单元素数组
		fmt.Println(array)
		return
	}

	quickSort(array, 0, len(array)-1)

	fmt.Println(array) // 打印排序后的数组
}
