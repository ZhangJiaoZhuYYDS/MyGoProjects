// @Author zhangjiaozhu 2023/4/26 13:27:00
package main

/*26. 删除有序数组中的重复项*/
func main()  {

}
func removeDuplicates(nums []int) int {
	// 初始化左慢指针位置
	l := 0
	// 从第1下标开始循环遍历
	for i := 1; i< len(nums) ; i++ {
		// 遇到右快指针与左慢指针不同时，就把右快指针的数给左慢指针+1的数，实现把不同的数都往前挪
		if nums[i] != nums[l] {
			nums[l+1] = nums[i]
			// 移动左慢指针
			l++
		}
	}
	return l+1
}
/*class Solution {
    public int removeDuplicates(int[] nums) {
        int l= 0;
        for (int i = 1 ; i < nums.length;i++){
            if (nums[i]!= nums[l]){
                nums[++l] = nums[i];
            }
        }
        return l+1;
    }
}*/