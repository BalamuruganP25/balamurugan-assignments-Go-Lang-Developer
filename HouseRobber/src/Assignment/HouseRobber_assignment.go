package Assignment
import (
    "fmt"
)

func Rob(nums []int)  {
    prevMax := 0
    currMax := 0

    for i:=0; i < len(nums); i++ {
        temp:=currMax
        if prevMax + nums[i] > currMax {
            currMax = prevMax + nums[i]
        }
        prevMax = temp
    }
   fmt.Println(currMax)

}



func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
