#!/bin/sh

# fonts color
Green="\033[32m"
Red="\033[31m"
Yellow="\033[33m"
GreenBG="\033[42;37m"
RedBG="\033[41;37m"
Font="\033[0m"
# fonts color

if [ ! -d /pic ]; then
    echo -e "${Red}未挂载目录,请重新执行.${Font}"
    exit 0
fi

if [ -d /pic/Output ]; then
    echo -e "${Red}Output 目录已存在,当前已暂停执行.${Font}"
    echo -e "${Red}请将 Output 目录移除或备份至其他目录.${Font}"
    echo -e "${Red}否则将会导致重复压缩已输出图片,浪费 API 次数.${Font}"
    exit 0
fi

if [ -f /pic/api_key.txt ]; then
    sed -i '17d' /tinypng.py
    LINE=17
    for APIKEY in $(cat /pic/api_key.txt)
    do
        sed -i "${LINE}i\    \"${APIKEY}\"," /tinypng.py
        LINE=$(($LINE+1))
    done
fi

python /tinypng.py

if [ ! -f /pic/api_key.txt ]; then
    echo -e "${Red}未检测到 api_key.txt${Font}"
    echo -e "${Red}已使用内置公开的 key 压缩图片,由于额度有限,图片压缩有可能失败.${Font}"
    echo -e "${Green}建议自行免费申请 API key 配置使用更加稳定.${Font}"
fi
