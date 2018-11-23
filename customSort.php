<?php

/*
 *快速排序
 *取第一个元素做标准，循环把小于标准值的放左，大于的放右，然后同理递归左右两边的元素。
 */ 
function quickSort($arr) {
    //先判断是否需要继续进行
    $length = count($arr);
    if ($length <= 1) {
        return $arr;
    }
    //如果没有返回，说明数组内的元素个数 多余1个，需要排序
    //选择一个标尺
    //选择第一个元素
    $base_num = $arr[0];
    //遍历 除了标尺外的所有元素，按照大小关系放入两个数组内
    //初始化两个数组
    $left_array = array(); //小于标尺的
    $right_array = array(); //大于标尺的
    for ($i = 1; $i < $length; $i++) {
        if ($base_num > $arr[$i]) {
            //放入左边数组
            $left_array[] = $arr[$i];
        } else {
            //放入右边
            $right_array[] = $arr[$i];
        }
    }
    //再分别对 左边 和 右边的数组进行相同的排序处理方式
    //递归调用这个函数,并记录结果
    $left_array = quickSort($left_array);
    $right_array = quickSort($right_array);
    //合并左边 标尺 右边
    return array_merge($left_array, array($base_num), $right_array);
}


/*
 * 归并排序 https://cuijiahua.com/blog/2018/01/algorithm_7.html
 * 归并排序是指将两个或两个以上有序的数列（或有序表），合并成一个仍然有序的数列（或有序表）。
 * 这样的排序方法经常用于多个有序的数据文件归并成一个有序的数据文件。
 * 
 */
function mergeSort(&$arr, $left, $right) {
    if($left < $right) {
        //说明子序列内存在多余1个的元素，那么需要拆分，分别排序，合并
        //计算拆分的位置，长度/2 去整
        $center = floor(($left+$right) / 2);
        //递归调用对左边进行再次排序：
        mergeSort($arr, $left, $center);
        //递归调用对右边进行再次排序
        mergeSort($arr, $center+1, $right);
        
        //合并排序结果
        echo $left,' : ',$center, ' : ',$right,'<br/>';
        //设置两个起始位置标记
        $a_i = $left;
        $b_i = $center+1;
        while($a_i<=$center && $b_i<=$right) {
            //当数组A和数组B都没有越界时
            if($arr[$a_i] < $arr[$b_i]) {
                $temp[] = $arr[$a_i++];
            } else {
                $temp[] = $arr[$b_i++];
            }
        }
        //判断 数组A内的元素是否都用完了，没有的话将其全部插入到C数组内：
        while($a_i <= $center) {
            $temp[] = $arr[$a_i++];
        }
        //判断 数组B内的元素是否都用完了，没有的话将其全部插入到C数组内：
        while($b_i <= $right) {
            $temp[] = $arr[$b_i++];
        }
        dump($temp);
        //将$arrC内排序好的部分，写入到$arr内
        $len=count($temp);
        for($i=0; $i<$len; $i++) {
            $arr[$left+$i] = $temp[$i];
        }
    }
}
?>