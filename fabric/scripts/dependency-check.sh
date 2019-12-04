#!/bin/bash

#cpu check
cpu_min=1
cpu_num=`cat /proc/cpuinfo|grep 'processor'|sort|uniq|wc -l`
echo "cpu num:$cpu_num"
 if [ $cpu_num -lt $cpu_min ]; then
  echo "CPU core less then required!"
  exit
fi

#memory check
memory_min=8
memory_free=`free -m|grep 'Mem'|awk '{print $4}'`
if [ $memory_free -lt $memory_min  ]; then
  echo "memory less then required!"
  exit
fi

#check disk
disk_size_min=1952550
disk_size=`df / | awk 'NR==2{print}' | awk '{print $2}' |grep '[0-9]'`
echo "disk size:$disk_size"
if [ $disk_size -lt $disk_size_min ]; then
   echo "disk size less then required!"
   exit
else echo "disk size is enough !"
fi

#check open files num
open_file_require=65535
open_file=`ulimit -a | grep "open files" | awk '{print $4}'`
echo "open files:" ${open_file}
if [[ $open_files -lt $open_file_require ]]; then
 echo "open files less then required!"
 exit
else echo "open-files is enough !"
fi
