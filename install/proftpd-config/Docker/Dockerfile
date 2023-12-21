FROM debian:stable-slim


RUN apt-get update -qq && apt-get install -y proftpd-basic proftpd-mod-mysql proftpd-mod-crypto \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


COPY ./proftpd/sql.conf /etc/proftpd/sql.conf
COPY ./proftpd/modules.conf /etc/proftpd/modules.conf


EXPOSE 20 21

CMD ["proftpd", "--nodaemon"]
