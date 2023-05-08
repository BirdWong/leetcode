/*
 * @lc app=leetcode.cn id=5 lang=c
 *
 * [5] 最长回文子串
 */

// @lc code=start
#include<string.h>
#include<stdlib.h>
#include<stdio.h>


int is_check(char * s,int len){
    
    
    //printf("join %s %d\n", s, len);
    int i = 0;
    for(; i < len/2 ; i++){
        //printf("check %d ", i);
        if(s[i] == s[len-i-1]){
            continue;
        }else{
            //printf("end join\n");
            return 0;
        }
    }
    //printf("end join\n");

    return 1;
}

int slen(char * s){
    int i = 0;
    for(;;i++){
        if(s[i] == '\0'){
            printf("%d", i);
            return i;
        }
    }
}

char * longestPalindrome(char * s){
    int len = slen(s);
    int i = len;
    for(; i > 0; i--){
        int pf = len-i;
        int j = 0;
        do{
            int tmpsize = len - pf;
            //printf("%d", tmpsize);
            char* tmps = calloc(sizeof(char), tmpsize);
            //printf("end\nbegin copy\n");

            memcpy(tmps, s+j, i);
            //printf("end copy %s \n", tmps);
            int result = is_check(tmps, tmpsize);
            if(result){
                return tmps;
            }
            //printf("%d", tmpsize);
            tmps = calloc(sizeof(char), tmpsize);
            //printf("end\nbegin copy\n");
            
            memcpy(tmps, s+j, len-j);
            //printf("end copy\n");
            result = is_check(tmps, tmpsize);
            if(result){
                //printf("%s", tmps);
                return tmps;
            }
            j++;
        }while(j < pf);
        
        
    }
    char *tmp = malloc(sizeof(char));
    tmp[0] = s[0];
    //printf("%s", tmp);
    return tmp;
}

// @lc code=end


int main(){
    char s[5] = "babad";
    printf("%s\n", longestPalindrome(s));
    return 0;
}


