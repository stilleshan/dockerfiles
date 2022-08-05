#/bin/sh

sed -i "s#http://example.com#https://${MYURLS_DOMAIN}#g" /app/public/index.html

/app/myurls -domain ${MYURLS_DOMAIN} -conn redis:6379 -ttl ${MYURLS_TTL}
