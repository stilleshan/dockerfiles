#/bin/sh

sed -i "s#https://s.ops.ci#https://${MYURLS_DOMAIN}#g" /app/public/index.html

/app/myurls -domain ${MYURLS_DOMAIN} -conn redis:6379
