/*
 * @lc app=leetcode.cn id=3 lang=c
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start

#include<stdlib.h>
#include<string.h>
#include<stdio.h>
int lengthOfLongestSubstring(char * s){
    int i = 0;
    int max = 0;
    int64_t* sizes = calloc(sizeof(int64_t), strlen(s));
    for(; i < strlen(s); i++){
        int tmp = 0;
        if(i!=0)tmp=sizes[i-1];
        int j = i-tmp;
        for(; j < i; j++){
            if(s[j] == s[i]){
                sizes[i] = i-j;
                break;
            }
        }
        if(!sizes[i]){
            sizes[i] = tmp+1;
        }
        if(max < sizes[i])
            max = sizes[i];
        
    }
    return max;

}
// @lc code=end
// int main(){
//     lengthOfLongestSubstring("abdabcbb");
//     return 0;
// }
