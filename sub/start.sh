#/bin/sh
if [ ! -f /usr/share/nginx/html/conf/config.js ]; then
  cp /app/conf/config.js /usr/share/nginx/html/conf
fi

if [ $API_URL ]; then
  echo "当前 API 地址为: $API_URL"
  sed -i "s#http://127.0.0.1:25500#$API_URL#g" /usr/share/nginx/html/conf/config.js
else
  echo "当前为默认本地 API 地址: http://127.0.0.1:25500"
  echo "如需修改请在容器启动时使用 -e API_URL='https://sub.ops.ci' 传递环境变量"
fi

if [ $SHORT_URL ]; then
  echo "当前短链接地址为: $SHORT_URL"
  sed -i "s#https://s.ops.ci#$SHORT_URL#g" /usr/share/nginx/html/conf/config.js
fi

if [ $SITE_NAME ]; then
  sed -i "s#Subconverter Web#$SITE_NAME#g" /usr/share/nginx/html/conf/config.js
fi

nohup /base/subconverter & echo "启动成功"

init_nginx (){
  sed -i '$d' /etc/nginx/conf.d/default.conf
  sed -i '$d' /etc/nginx/conf.d/default.conf
  sed -i '$d' /etc/nginx/conf.d/default.conf
  cat >> /etc/nginx/conf.d/default.conf <<EOF
    location ~* /(sub|render|getruleset|surge2clash|getprofile) {
        proxy_redirect off;
        proxy_pass http://127.0.0.1:25500;
    }
  }
EOF
}

if [[ ! $(cat /etc/nginx/conf.d/default.conf | grep 25500) ]]; then
	init_nginx
fi

nginx -g "daemon off;"
