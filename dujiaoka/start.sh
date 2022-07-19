#!/bin/sh

if [ -f "/dujiaoka/.env" ]; then
    if [ ! -d "./storage/app" ]; then
        mv -n storage_bak/* storage/
    fi
    if [ "$INSTALL" != "true" ]; then
        echo "ok" > install.lock
    fi
    if [ "$MODIFY" != "false" ]; then
        mv ./modify/StripeController.php /dujiaoka/app/Http/Controllers/Pay/StripeController.php
    fi

    bash /dujiaoka/start-hook.sh

    chmod -R 777 storage

    php artisan clear-compiled
    php artisan optimize
    php artisan migrate

    supervisord
else
    echo "配置文件不存在，请根据文档修改配置文件！"
fi
