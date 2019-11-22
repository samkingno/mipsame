package main

import (
    "fmt"
    "strconv"
    "strings"
    "flag"
    "os/exec"
)

/* IP地址转化为二进制数 */
func main() {

    var ip1 string
    var ip2 string
    var m int
    var hostname1 string
    var hostname2 string
    var t string
    //fmt.Println("请输入第一个ip：")
    //fmt.Scanln(&ip1)
    //fmt.Println("请输入第二个ip：")
    //fmt.Scanln(&ip2)
    //fmt.Println("请输入子网掩码的位数：")
    //fmt.Scanln(&m)

    flag.StringVar(&t,"t","ip","请输入判断类型")
    flag.StringVar(&ip1,"ip1","","请输入ip")
    flag.StringVar(&ip2,"ip2","","请输入ip")
    flag.StringVar(&hostname1,"h1","","请输入ip")
    flag.StringVar(&hostname2,"h2","","请输入ip")
    flag.IntVar(&m,"m",23,"请输入掩码位数")
    flag.Parse()

    if t == "ip"{
    addrInt1 := ipAddrToInt(ip1)
    /* 十进制转化为二进制 */
    c := strconv.FormatInt(addrInt1, 2) 
    k1 := ip32(c)
    fmt.Println("ip1:",ip1,k1)

    addrInt2 := ipAddrToInt(ip2)
    d := strconv.FormatInt(addrInt2, 2)
    k2 := ip32(d)
    fmt.Println("ip2:",ip2,k2)



    e := string([]byte(k1)[:m])
    f := string([]byte(k2)[:m])
    if e == f{
    fmt.Println("两个ip在同一网段")
    }else{
    fmt.Println("两个ip不在同一网段")
    }
    }else if t=="h"{
    ip1 = hosttoip(hostname1)
    ip2 = hosttoip(hostname2)
        addrInt1 := ipAddrToInt(ip1)
    /* 十进制转化为二进制 */
    c := strconv.FormatInt(addrInt1, 2)
    k1 := ip32(c)
    fmt.Println("hostname1:",hostname1,"ip:",ip1,k1)

    addrInt2 := ipAddrToInt(ip2)
    d := strconv.FormatInt(addrInt2, 2)
    k2 := ip32(d)
    fmt.Println("hostname2:",hostname2,"ip:",ip2,k2)



    e := string([]byte(k1)[:m])
    f := string([]byte(k2)[:m])
    if e == f{
    fmt.Println("两个ip在同一网段")
    }else{
    fmt.Println("两个ip不在同一网段")
    }

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

func hosttoip(hostname string) string{
cmd := exec.Command("host",hostname)
out, err := cmd.CombinedOutput()
if err != nil {
        fmt.Println(err)
    }
    outmap :=strings.Split(string(out), " ")
    ip:=outmap[3]
    return ip
}
