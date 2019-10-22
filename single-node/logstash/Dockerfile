# Logstash 7.3.2

# This image re-bundles the Docker image from the upstream provider, Elastic.
FROM docker.elastic.co/logstash/logstash:7.3.2
USER root
# The upstream image was built by:
#   https://github.com/elastic/dockerfiles/tree/v7.3.1/logstash

# For a full list of supported images and tags visit https://www.docker.elastic.co

# For Logstash documentation visit https://www.elastic.co/guide/en/logstash/current/docker.html

# See https://github.com/docker-library/official-images/pull/5039 for more details.
RUN mkdir -p /opt/logstash/incoming && \
mkdir -p /opt/logstash/logs && \
mkdir -p /opt/logstash/sincedb && \
chown logstash:logstash /opt/logstash/incoming && \
chown logstash:logstash /opt/logstash/logs && \
chown logstash:logstash /opt/logstash/sincedb && \

USER logstash
WORKDIR /home/logstash


