        fmt.Printf("%d ", copyData[i])
        srcData[i] = i
    // 修改原始数据的第一个元素
    // 复制原始数据从4到6(不包含)
    // 将切片赋值
    // 将数据复制到新的切片空间中
    // 引用切片数据
    // 打印复制切片的第一个和最后一个元素
    // 打印引用切片的第一个元素
    // 设置元素数量为1000
    // 预分配足够多的元素切片
    // 预分配足够多的元素切片
    const elementCount = 1000
    copy(copyData, srcData)
    copy(copyData, srcData[4:6])
    copyData := make([]int, elementCount)
    fmt.Println(copyData[0], copyData[elementCount-1])
    fmt.Println(refData[0])
    for i := 0; i < 5; i++ {
    for i := 0; i < elementCount; i++ {
    refData := srcData
    srcData := make([]int, elementCount)
    srcData[0] = 999
    }
    }
func main() {
import "fmt"
package main
}