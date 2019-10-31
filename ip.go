package main

import (
    "fmt"
    "strconv"
    "strings"
)

/* IP地址转化为二进制数 */
func main() {

    var a string
    var b string
    var g int
    fmt.Println("请输入第一个ip：")
    fmt.Scanln(&a)
    fmt.Println("请输入第二个ip：")
    fmt.Scanln(&b)
    fmt.Println("请输入子网掩码的位数：")
    fmt.Scanln(&g)

    addrInt1 := ipAddrToInt(a)
    /* 十进制转化为二进制 */
    c := strconv.FormatInt(addrInt1, 2) 
    k1 := ip32(c)
    fmt.Println("c:",a,k1)
    addrInt2 := ipAddrToInt(b)
    d := strconv.FormatInt(addrInt2, 2)
    k2 := ip32(d)
    fmt.Println("d:",b,k2)



    e := string([]byte(k1)[:g])
    f := string([]byte(k2)[:g])
    if e == f{
    fmt.Println("两个ip在同一网段")
    }else{
    fmt.Println("两个ip不在同一网段")
    }



    /* 二进制转化为十进制 */
    //d, err := strconv.ParseInt(c, 2, 64)
    //fmt.Println("d:", d, err)
}

func ipAddrToInt(ipAddr string) int64 {
    bits := strings.Split(ipAddr, ".")
    b0, _ := strconv.Atoi(bits[0])
    b1, _ := strconv.Atoi(bits[1])
    b2, _ := strconv.Atoi(bits[2])
    b3, _ := strconv.Atoi(bits[3])
    
    var sum int64
    sum += int64(b0) << 24
    sum += int64(b1) << 16
    sum += int64(b2) << 8
    sum += int64(b3)

    return sum
}

func ip32(ipint string) string{
    var k string
    h := len(ipint)
    for i:=32-h;i>0;i--{
       k += "0" 
    }
    k += ipint
    return k
}
