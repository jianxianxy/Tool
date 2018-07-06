package main

import "fmt"

//定义接口
type Inface interface{
    GetName() string
    GetAge()  int
}
//创建父类
type claParent struct{
    Name string
    Age  int
}
//创建子类一号
type claOne struct{
    claParent
    City string
}
//子类一号实现接口
func (cla claOne) GetName() string{
    return cla.Name
}
func (cla claOne) GetAge() int{
    return cla.Age
}
//创建类实现接口
type claTwo struct{
    Inface
    Leval int    
}

func main() {
    var a interface{}
    a = 3
    fmt.Println(a.(int))
    a = "abc"
    fmt.Println(a.(string))
    
    //cone := claOne{claParent:claParent{Name:"Leiluo",Age:33},City:"北京"}
    cone := claOne{City:"北京"}
    cone.claParent = claParent{Name:"Leiluo",Age:33}
    fmt.Println(cone.GetName())
    
    //claInface类实现必须是一个实现了Inface接口的类
    ctwo := claTwo{Inface:claOne{claParent:claParent{Name:"Kylin",Age:33},City:"北京"}}
    ctwo.Leval = 1
    fmt.Println(ctwo.GetName())
}
