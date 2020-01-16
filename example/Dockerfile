FROM huaxinlu/openeulergame2019:v3.1
ADD run.sh /run.sh
RUN yum -y install httpd \
    && yum clean all \
    && touch /var/www/html/index.html \
    && echo "Hello World!" >> /var/www/html/index.html \
    && chmod -v +x /run.sh
EXPOSE 80
CMD ["/run.sh"]
