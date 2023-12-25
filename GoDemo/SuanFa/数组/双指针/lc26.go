package main

/* 26 删除排序数组中的重复项*/
func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	println(removeDuplicates(arr))
}
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	// 初始化快慢指针的位置
	l, r := 0, 1
	//从第二个元素开始遍历
	for r = 1; r < len(nums); r++ {
		// 判断当前遍历的元素是否和num[l]不相等，如果两者不相等，就让左指针右移一位+1，交换左指针和当前遍历元素的值
		if nums[r] != nums[l] {
			// 左指针右移一位+1
			l++
			temp := nums[l]
			nums[l] = nums[r]
			nums[r] = temp
		}
	}
	return l + 1
}

// java
/*class Solution {
    public int removeDuplicates(int[] nums) {
        int j = 1;
        for(int i =1;i<nums.length;i++){
            if(nums[i]>nums[i-1]){
                nums[j++]=nums[i];
            }
        }
        return j;
    }
}*/
