// @Author zhangjiaozhu 2023/8/15 10:53:00
package hard

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	nums1Size := len(nums1)
	if nums1Size%2 == 0 {
		return float64(nums1[nums1Size/2] + nums1[(nums1Size/2)+1])
	} else {
		return float64()
	}
}

// java
class Solution {
	public double findMedianSortedArrays(int[] nums1, int[] nums2) {
		int len = nums1.length + nums2.length;
		int[] num = new int[len];

		int a = 0, b = 0;

		for (int i = 0; i < len; i++) {
			if (b >= nums2.length) {
				num[i] = nums1[a];
				a++;
				continue;
			}
			if (a >= nums1.length) {
				num[i] = nums2[b];
				b++;
				continue;
			}
			if (nums1[a] < nums2[b]) {
				num[i] = nums1[a];
				a++;
			} else {
				num[i] = nums2[b];
				b++;
			}
		}

		double mid;
		if (len % 2 == 0) {
			mid = (double) (num[len / 2] + num[len / 2 - 1]) / 2;
		} else {
			mid = num[(len - 1) / 2];
		}

		return mid;
	}
}
