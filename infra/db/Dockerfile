FROM postgres:16

# install locales
RUN apt-get update && \
  apt-get install -y locales && \
  sed -i '/ja_JP.UTF-8/s/^# //g' /etc/locale.gen && \
  locale-gen ja_JP.UTF-8

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ Asia/Tokyo
