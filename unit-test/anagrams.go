package main

import (
	"fmt"
	"strconv"
)

func solution(a []int) int64 {
    // the value and the value length
    var intMap = make(map[int]int, len(a))
    var result int64 = 0
    
    for i := range a {
        // take index convert to string then find len compare to other indexes and if 
        var str = strconv.Itoa(a[i])
        
        var intLen = len(str)
        
        intMap[i] = intLen
        
    }
    
    // make sure to find values that are the same length
    for key, _ := range intMap {
        //var current = a[i]
        
        if key != len(a) - 1 {
            if intMap[key] > intMap[key + 1] {
                continue
                //fmt.Println("didn't work")
            } else {
                //i := key + 1
                
                var str = strconv.Itoa(a[key])
                
                // turn the first str into keys to compare much easier 
                var strMap = strMap(str)
                
                //fmt.Println(strMap)
                
                var b []int = make([]int, len(a))
                copy(b, a)
                
                a[key] = 0
                
                var val int 
                var value int
                fmt.Println(str)
                for i := range a  {
                    val = 0
                    var str2 = strconv.Itoa(a[i])
                    fmt.Println("str is: ", str, "str2 is: ", str2, "i is: ", i)
                    //fmt.Println(i)
                    
                    if len(str) == len(str) {}
                    val += validDigits(str2, strMap)
                    
                    fmt.Println("val is: ", val)
                    fmt.Println("---------------")
                    /*fmt.Println("str map is: ", strMap)
                    fmt.Println("str2 is: ", str2)
                    fmt.Println("val is: ", val)
                    fmt.Println("---------------")*/
                    
                    //val -= 1
                    //fmt.Println(val)
                    //result += int64(val)
                    //fmt.Println(result)
                    //strLen := 0 
                    value += val                   
                }
                //val--
                result += int64(value)
                //val = 0
            }
        }    
    }
    
    //fmt.Println(intMap)
    return result
}

func strMap(str string) map[string]int {
    var strMap = make(map[string]int, len(str))
    
    //fmt.Println(str)
    for i := range str{ 
        letter := string(str[i])
        strMap[letter] = 0
    }
    return strMap
}

func validDigits(str string, strMap map[string]int) int {
    var count int = 0
    
    /*_, duplicate := strMap[str]
    
    if duplicate {
        fmt.Println("strMap is: ", strMap[str], "and str is: ", str)
        return 0    
    }*/
    
    for i := range str {
        letter := string(str[i])
        fmt.Println("Letter is: ", letter)
        _, valid := strMap[letter]
        
        fmt.Println("valid: ", valid)
        if valid {
            if strMap[letter] != 1 {
                count++
                strMap[letter] = 1
            }
        }
    }
    
    
    if count == len(str) {
        clearMap(strMap)
        return 1
    }
    
    clearMap(strMap)
    return 0
}

func clearMap(strMap map[string]int) {
    for i := range strMap {
            strMap[i] = 0
    }
}