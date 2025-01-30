#!/bin/bash

# fonts color
Green="\033[32m"
Red="\033[31m"
Yellow="\033[33m"
GreenBG="\033[42;37m"
RedBG="\033[41;37m"
Font="\033[0m"
# fonts color

WORK_PATH=$(dirname $(readlink -f $0))


v2_ui_ssl (){
clear
echo -e "${Green}请输入需要申请证书的根域名(例如:ioiox.com):${Font}"
read -p "请输入:" DOMAIN_INPUT
if [ ! -n "${DOMAIN_INPUT}" ]; then
    echo -e "${Red}输入错误,请重新运行脚本.${Font}"
    exit 0
fi
DOMAIN=$DOMAIN_INPUT
echo -e "${Green}请选择域名服务商:${Font}"
echo -e "1) 腾讯云 dnspod.cn"
echo -e "2) 阿里云 aliyun"
echo -e "3) Cloudflare"
read -p "请选择:" DNSAPI_INPUT
case "$DNSAPI_INPUT" in
    1)
    PLATFORM_NAME='dnspod.cn'
    DNSAPI='dns_dp'
    API_ID_HEADER='DP_Id'
    API_KEY_HEADER='DP_Key'
    ;;
    2)
    PLATFORM_NAME='aliyun'
    DNSAPI='dns_ali'
    API_ID_HEADER='Ali_Key'
    API_KEY_HEADER='Ali_Secret'
    ;;
    3)
    ;;
    *)
    echo -e "${Red}输入错误,请重新运行脚本.${Font}"
    exit 0
    esac

if [ "$DNSAPI_INPUT" == "3" ]; then
    echo -e "${Green}=========================================================================================${Font}"
    echo -e  "${Red}注意: Cloudflare API 有三种:${Font}"
    echo -e  "${Red}请参考 https://github.com/acmesh-official/acme.sh/wiki/dnsapi#1-cloudflare-option 选择.${Font}"
    echo "1) Using the global API key"
    echo "2) Using the new cloudflare api token"
    echo "3) Using the new cloudflare api token for Single Zone"
    read -p "请选择:" CHOICE_CLOUDFLARE_INPUT
    echo -e "${Green}=========================================================================================${Font}"
    case "$CHOICE_CLOUDFLARE_INPUT" in
        1)
        PLATFORM_NAME='Cloudflare'
        DNSAPI='dns_cf'
        API_ID_HEADER='CF_Key'
        API_KEY_HEADER='CF_Email'
        ;;
        2)
        PLATFORM_NAME='Cloudflare'
        DNSAPI='dns_cf'
        API_ID_HEADER='CF_Token'
        API_KEY_HEADER='CF_Account_ID'
        ;;
        3)
        PLATFORM_NAME='Cloudflare'
        DNSAPI='dns_cf'
        API_ID_HEADER='CF_Token'
        API_KEY_HEADER='CF_Account_ID'
        API_ZONE_HEADER='CF_Zone_ID'
        ;;
        *)
        echo -e "${Red}输入错误,请重新运行脚本.${Font}"
        exit 0
        esac
fi

read -p "请输入 $API_ID_HEADER :" API_ID_INPUT
read -p "请输入 $API_KEY_HEADER :" API_KEY_INPUT
if [ "$CHOICE_CLOUDFLARE_INPUT" == "3" ]; then
    read -p "请输入 $API_ZONE_HEADER :" API_ZONE_HEADER_INPUT
fi


echo -e "${Green}=========================================================================================${Font}"
echo -e "${Red}请确认以下信息正确无误!${Font}"
echo -e "${Green}域名: ${Font}${Red}${DOMAIN}${Font}"
echo -e "${Green}域名服务商: ${Font}${Red}${PLATFORM_NAME}${Font}"
echo -e "${Green}${API_ID_HEADER}:${Font} ${Red}${API_ID_INPUT}${Font}"
echo -e "${Green}${API_KEY_HEADER}:${Font} ${Red}${API_KEY_INPUT}${Font}"
if [ "$CHOICE_CLOUDFLARE_INPUT" == "3" ]; then
    echo -e "${Green}${API_ZONE_HEADER}:${Font} ${Red}${API_ZONE_HEADER_INPUT}${Font}"
fi
echo -e "${Red}请再次确认以上信息正确无误!${Font}"
echo -e "${Green}=========================================================================================${Font}"
echo -e "1) 开始部署"
echo -e "2) 退出脚本"
read -p "请输入:" START_INPUT
case "$START_INPUT" in
    1)
    echo -e "${Green}开始部署中......${Font}"
    accout_conf $*
    ;;
    2)
    exit 0
    ;;
    *)
    echo -e "${Red}输入有误,请重新运行脚本.${Font}"
    exit 0
    esac
}


accout_conf (){
WORK_PATH=$(dirname $(readlink -f $0))
wget https://raw.githubusercontent.com/stilleshan/dockerfiles/main/x-ui/x-ui-ssl/x-ui-ssl.tar
tar -xvf x-ui-ssl.tar
cat >${WORK_PATH}/x-ui-ssl/conf/account.conf<<EOF
export ${API_ID_HEADER}="${API_ID_INPUT}"
export ${API_KEY_HEADER}="${API_KEY_INPUT}"
export DOMAIN=${DOMAIN}
export DNSAPI=${DNSAPI}
EOF
if [ "$CHOICE_CLOUDFLARE_INPUT" == "3" ]; then
    sed -i "2a export ${API_ZONE_HEADER}=\"${API_ZONE_HEADER_INPUT}\"" x-ui-ssl/conf/account.conf
fi
docker_compose_ssl
}

docker_compose_ssl (){
    cd x-ui-ssl
    docker-compose up -d
    cd ${WORK_PATH}
    rm -rf ${WORK_PATH}/x-ui-ssl.tar
    echo -e "${Green}部署完毕,静等 1 分钟查看 ssl 目录下是否生成证书文件.${Font}"
    echo -e "${Green}访问 http://服务器IP:54321 使用帐号 admin 密码 admin 登陆,尽快修改帐号密码.${Font}"
    rm $0
}

docker_compose_nossl (){
    WORK_PATH=$(dirname $(readlink -f $0))
    wget https://raw.githubusercontent.com/stilleshan/dockerfiles/main/x-ui/x-ui-ssl/x-ui-ssl.tar
    tar -xvf x-ui-ssl.tar
    sed -i '8d' ${WORK_PATH}/x-ui-ssl/docker-compose.yml
    sed -i '11,20d' ${WORK_PATH}/x-ui-ssl/docker-compose.yml
    rm -rf ${WORK_PATH}/x-ui-ssl/conf
    mv ${WORK_PATH}/x-ui-ssl ${WORK_PATH}/x-ui
    cd x-ui
    docker-compose up -d
    cd ${WORK_PATH}
    rm -rf ${WORK_PATH}/x-ui-ssl.tar
    echo -e "${Green}部署完毕.访问 http://服务器IP:54321 使用帐号 admin 密码 admin 登陆,尽快修改帐号密码.${Font}"
    rm $0
}


if ! type docker-compose >/dev/null 2>&1 ; then
    echo -e "${Red}本机未安装 docker compose 已退出脚本.${Font}";
    exit 0
fi

clear
echo -e "${Green}=========================================================================================${Font}"
echo -e "${Green}开始安装 x-ui-ssl${Font}"
echo -e "${Red}注意:本脚本需要服务器有 docker 和 docker compose 环境${Font}"
echo -e "${Green}=========================================================================================${Font}"
echo "1) x-ui + acme 自动申请和续签证书"
echo "2) x-ui 独立版本"
read -p "请输入:" VS_INPUT
case "$VS_INPUT" in
    1)
    v2_ui_ssl
    ;;
    2)
    docker_compose_nossl
    ;;
    *)
    echo -e "${Red}输入有误,请重新运行脚本.${Font}"
    exit 0
    esac
