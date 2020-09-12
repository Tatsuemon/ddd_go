FROM nginx:1.18.0
#デフォルトのnginxファイルを削除して作成したものコンテナないにコピー
RUN rm -f /etc/nginx/conf.d/*
RUN apt-get update && \
    apt-get install -y apt-utils \
    locales && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8

COPY ./docker/nginx/nginx.conf /etc/nginx/conf.d/app.conf

#-c以降の設定ファイルを指定して起動 daemon offでフォアグラウンドで起動
CMD /usr/sbin/nginx -g 'daemon off;' -c /etc/nginx/nginx.conf
