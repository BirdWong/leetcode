/*
 * @lc app=leetcode.cn id=4 lang=c
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
#include<stdlib.h>
#include<stdio.h>

double array_value(int* nums1, int nums1Size, int* nums2, int index){
    if(index < nums1Size){
        return (double)nums1[index];
    }
    return (double)nums2[index - nums1Size];

}


double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int *index_num = (int*) calloc(nums1Size + nums2Size, sizeof(int));
    if(nums1Size == 0 && nums2Size == 0){
        return 0.0;
    }
    if(nums1Size == 0 && nums2Size == 1){
        return (double)nums2[0];
    }

    if(nums1Size == 1 && nums2Size == 0){
        return (double)nums1[0];
    }

    if(nums1Size == 1 && nums2Size == 1){
        return ((double)nums1[0] + (double)nums2[0]) / 2;
    }


    int median = (nums1Size + nums2Size) / 2;
    int median_second = 0;
    int has_second = 0;
    if ((nums1Size + nums2Size) % 2 == 0)
    {
        has_second = 1;
        median_second = median - 1;
    }

    int num1_index = nums1Size / 2;
    int num2_index = nums2Size / 2;

    while (1)
    {
        if(index_num[median] != 0){
            if(has_second){
                if(index_num[median_second] != 0){
                    return (array_value(nums1, nums1Size, nums2, index_num[median]) + array_value(nums1, nums1Size, nums2, index_num[median - 1])) / 2;
                }
            }else{
                return array_value(nums1, nums1Size, nums2, index_num[median]);
            }
        }

        if(nums1[num1_index] < nums2[num2_index]){
            int tmp_index = num1_index + num2_index;
            index_num[tmp_index] = num1_index;
            //printf("index: %d is num1_index: %d, value is: %d\n", tmp_index, num1_index, nums1[num1_index]);
            if(tmp_index > median)
            {
                num2_index -= 1;
            }
            if(tmp_index < median)
            {
                num1_index += 1;
            }
            if(tmp_index == median && has_second){
                num2_index -= 1;
            }
        }else if (nums1[num1_index] > nums2[num2_index])
        {
            int tmp_index = num1_index + num2_index;
            index_num[tmp_index] = nums1Size + num2_index;
            //printf("index: %d is num2_index: %d, value is: %d\n", tmp_index, (num2_index + nums1Size), nums2[num2_index]);
            if(tmp_index > median)
            {
                num1_index -= 1;
            }
            if(tmp_index < median)
            {
                num2_index += 1;
            }
            if(tmp_index == median && has_second){
                num1_index -= 1;
            }
        }
        
    }
    



}
// @lc code=end

int main()
{
    int nums1[] = {1,2};
    int nums2[] = {3,4};
    int nums1Size = 2;
    int nums2Size = 2;
    double result = findMedianSortedArrays(nums1, nums1Size, nums2, nums2Size);
    printf("%lf", result);
    return 0;
}