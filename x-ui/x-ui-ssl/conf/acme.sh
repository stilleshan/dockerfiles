PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

# fonts color
Green="\033[32m"
Red="\033[31m"
Yellow="\033[33m"
GreenBG="\033[42;37m"
RedBG="\033[41;37m"
Font="\033[0m"
# fonts color

DOMAIN=$(cat /conf/account.conf | awk -F= '{if($1~"DOMAIN")print $2}')
DNSAPI=$(cat /conf/account.conf | awk -F= '{if($1~"DNSAPI")print $2}')

acme (){
    cat /conf/account.conf >/acme.sh/account.conf
    /root/.acme.sh/acme.sh --upgrade
    /root/.acme.sh/acme.sh --register-account -m your@domain.com --server zerossl
    /root/.acme.sh/acme.sh --issue $* --keylength 2048 --dns ${DNSAPI} -d ${DOMAIN} -d \*.${DOMAIN}
    rm -rf /acme.sh/ca
    rm -rf /acme.sh/http.header
    if [ -f /acme.sh/${DOMAIN}/fullchain.cer ] && [ -f /acme.sh/${DOMAIN}/${DOMAIN}.key ]; then
        rm -rf /ssl/${DOMAIN}.cer
        rm -rf /ssl/${DOMAIN}.key
        cp /acme.sh/${DOMAIN}/fullchain.cer /ssl/${DOMAIN}.cer
        cp /acme.sh/${DOMAIN}/${DOMAIN}.key /ssl/${DOMAIN}.key
        mv /acme.sh/${DOMAIN} /acme.sh/${DOMAIN}-$(date +%Y)-$(date +%m)-$(date +%d)-${RANDOM}
    else
        echo "证书申请失败,已退出脚本,请重新尝试."
        rm -rf /acme.sh/${DOMAIN}
        exit 1
    fi
}

if [ ! -n "${DOMAIN}" ] ; then
    echo "请检查 account.conf 信息是否正确"
else
    acme $*
fi

cat >/var/spool/cron/crontabs/root<<EOF
0 0 10 * * /conf/acme.sh >/dev/null 2>&1
EOF
