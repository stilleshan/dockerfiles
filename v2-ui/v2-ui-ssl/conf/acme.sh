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
    cat /conf/account.conf > /acme.sh/account.conf
    /root/.acme.sh/acme.sh --upgrade
    /root/.acme.sh/acme.sh --register-account -m your@domain.com --server zerossl
    /root/.acme.sh/acme.sh --issue $* --dns ${DNSAPI} -d ${DOMAIN} -d \*.${DOMAIN}
    cp /acme.sh/${DOMAIN}/fullchain.cer /ssl/${DOMAIN}.cer
    cp /acme.sh/${DOMAIN}/${DOMAIN}.key /ssl/${DOMAIN}.key
    mv /acme.sh/${DOMAIN} /acme.sh/${DOMAIN}-$(date +%Y)-$(date +%m)-$(date +%d)-${RANDOM}
}

if [ ! -n "${DOMAIN}" ] ; then
    echo "请检查 account.conf 信息是否正确"
else
    acme $*
fi

cat >/var/spool/cron/crontabs/root<<EOF
0 0 1 * * /conf/acme.sh >/dev/null 2>&1
EOF
