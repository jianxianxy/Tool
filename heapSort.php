<?php
/**
 *多叉堆排序
**/
header("Content-type=text/html;charset=utf-8");
date_default_timezone_set('Asia/Chongqing');
set_time_limit(0);

$arr = array();
for($i=1;$i<1000;$i++){
    $arr[] = rand(0,1000);
}


//多叉堆排序
$queen = 12;
$start = ctime();
$three  = heapSort($arr,$queen);
$mt = ctime($start);
echo $queen.' Fork:'.$mt;

//堆排序( M叉 )
function heapSort($arr,$queen){
    //第一步 构建大顶堆
    initMaxTopHeap($arr,$queen);
    //第二步 将堆顶和堆尾逐个交换,新的堆顶必将在其子节点产生,所以从堆尾向前成倒序排列
    $len = count($arr)-1;
    for($i=$len;$i>=0;$i--){
        $tmp = $arr[$i];
        $arr[$i] = $arr[0];
        $arr[0] = $tmp;
        maxForkOrder($arr,0,$i,$queen);
    }
    return $arr;
}

//构建大顶堆完全M叉树 子树节点值小于父节点值
function initMaxTopHeap(&$arr,$queen){
    //M叉树规则 假设有 N 个元素
    //那么就有 ceil(N/M) 棵树 即最后一棵树的索引是 ceil(N/M)-1
    //从最后一颗子树开始处理,把更大的值放在父节点上,直到树的顶端比较完,将最大的值放在树顶
    $start = ceil(count($arr)/$queen)-1;
    $end = count($arr)-1;
    for($start;$start>=0;$start--){
        maxForkOrder($arr,$start,$end,$queen);   
    }
    return $arr;
}

/*
    M叉树级联梳理 梳理成完全M叉树
    $start  开始处理的索引
    $end    结束处理的索引
    $queen  叉数
*/
function maxForkOrder(&$arr,$start,$end,$queen){ 
    $children = array();
    for($i=0;$i<$queen;$i++){
        $children[$i] = $queen*($start+1) - $i;   
    }
    //最大值的索引
    $maxValIndex = $start;      
    //数组最大索引
    $maxIndex = count($arr)-1;  
    foreach($children AS $key => $val){
        //如果有树
        if($val <= $maxIndex && $val < $end){
            if($arr[$val] > $arr[$maxValIndex]){
                $maxValIndex = $val;
            }
        }   
   
    }
    if($maxValIndex != $start){
        $temp = $arr[$start];
        $arr[$start] = $arr[$maxValIndex];
        $arr[$maxValIndex] = $temp;
        maxForkOrder($arr,$maxValIndex,$end,$queen);
    }
}

//计算耗时
function ctime($start='')
{
    $curr = microtime();
    $carr = explode(' ',$curr);
    $time = $carr[1].substr($carr[0],1,7);
    if($start)
    {
        $plus = sprintf("%.6f",$time-$start);
        return $plus;
    }
    else
    {
        return $time;
    }
}
?>