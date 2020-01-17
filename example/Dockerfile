FROM swr.cn-north-4.myhuaweicloud.com/openeulergame2020/openeuler:v4.1
RUN yum -y install httpd --setopt=override_install_langs=en \
                         --setopt=tsflags=nodocs \
                         --setopt=install_weak_deps=false\
    && yum clean all \
    && touch /var/www/html/index.html \
    && echo "Hello World!" >> /var/www/html/index.html 
EXPOSE 80
CMD ["/usr/sbin/httpd","-DFOREGROUND"]
