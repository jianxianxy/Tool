package main
//"strings"
import(
    "fmt"
    "time"
    "strconv"
    "encoding/json"
    "io/ioutil"
    "net/http"  
)

//`json:"is_on_sale,string"`
type RetJson struct{
    Status int `json:"status"` 
    Info string `json:"info"`
    Data string `json:"data"`
}

func main(){
    b_Time := time.Now().Unix()
    page := 5
    count := 200
    master := make(chan int,page)
    for i := 0; i < page; i++ {
        start := i * count + 1;
        end := count + i * count 
        go CreateWork(start,end,master)
    }
    all := 0
    for i := 0; i < page; i++ {
        wno := <-master  
        all += wno  
    }
    e_Time := time.Now().Unix()
    u_Time := e_Time - b_Time
    fmt.Println("处理结束,共生成:",all,"用时:",u_Time)
}

//开启多个Work
func CreateWork(start int,end int,master chan int){
    succ := 0
    work := make(chan int)
    go ChanTimeOutGet(start,end,work)
    for{
        select{
            case back := <- work:
                if back == -2{
                    goto HERE
                }else if back == -1{
                    succ += 1
                    goto HERE    
                }else{
                    succ += back
                }
        }    
    }
    HERE:
    fmt.Println(start,"-",end," 成功:",succ)
    master <- succ
}

//多Channel处理
func ChanTimeOutGet(start int,end int,work chan int){
    url := "http://tpl.t.com/Product/MakeOne?productId=" + strconv.Itoa(start)
    ch := make(chan int)
    timeOut := make(chan int)
    go func(){
        time.Sleep(12e9)
        timeOut <- 0
    }()
    go ChanGet(url,ch)
    select{
        case back := <- ch:
            if start < end{   
                work <- back
                start += 1
                ChanTimeOutGet(start,end,work)            
            }else if back > 0{
                work <- -1 //成功并结束  
            }else{
                work <- -2 //失败并结束    
            }
        case back := <- timeOut:
            fmt.Println("[",start,":timeout]")
            if start < end{
                work <- back    
                start += 1
                ChanTimeOutGet(start,end,work)            
            }else{
                work <- -2 //失败并结束   
            }
    }
}

//Channel处理GET请求
func ChanGet(url string,ch chan int){
    body,err := GetUrl(url)
    if err != nil{
        fmt.Println(url + " >> " + body)
        ch <- 0
    }else{
        //fmt.Println(body)
        sta,_ := FormatOne(body)   
        ch <- sta
    }    
}

//发送GET请求
func GetUrl(url string)(ret string,err error){
    response,err := http.Get(url)
    if err != nil{
        return "Connecting failed!",err
    }
    defer response.Body.Close()
    body,_ := ioutil.ReadAll(response.Body)
    return string(body),nil
}

//处理返回结果:MakeOne
func FormatOne(body string)(sta int,ret *RetJson){
    ret = &RetJson{}
    errJson := json.Unmarshal([]byte(body),ret)
    if errJson == nil{
        return ret.Status,ret
    }else{
        return 0,ret
    }
}