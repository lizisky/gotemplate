#!/bin/bash

if [ ! -f ./liziskyd ]; then
  echo "binary file not exist, building it ......"
  cd ..
  make
  cd -
fi

if [ ! -d ./logs ]; then
  mkdir logs
fi

if [ -z "$1" ];
then
    # ./liziskyd -logtostderr
    ./liziskyd -log_dir=logs -alsologtostderr
elif [ -d $1 ];
then
    ./liziskyd -logtostderr -data_dir=$1
    # echo "yes", $1
fi


