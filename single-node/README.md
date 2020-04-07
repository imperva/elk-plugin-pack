# Large Scale Audit Report - Phase 1
## Create a single stack proof of concept for a large scale audit reporting platform for the Imperva Data Security platform.

## Based on [Elasticsearch](https://www.elastic.co/),[Kibana](https://www.elastic.co/products/kibana), and [FluentD](https://www.fluentd.org/).

The files in this repository will build out a single dockerized ELK or EFK stack which will consist of:

  * 1 - Elasticsearch v7.x node
  * 1 - Kibana v7.x node
  * 1 - Logstash node
  * 1 - FluentD node **uncomment docker-compose.yml to use**


### Logstash
The Dockerfile located in ./logstash  will build the logstash docker image based on the latest version the first time the docker-compose.yml file run.

The configuration file, ./logstash/pipeline/audit-pipeline.conf, is configured
to listen for JSON formatted syslog messages on TCP port 5514, and then write
those logs to the Elasticsearch data system in daily indexes titled
"audit-YYYY-MM-dd".  The index name can easily be changed by editing that file, and changing the name of the index.

### FluentD
The Dockerfile located in ./fluentd will build the fluentd docker image based on the latest version the first time the docker-compose.yml file is run and install the Elasticsearch output plugin.

The configuration file, ./fluentd/conf/fluent.conf, is configured to listen for JSON formatted syslog messages on TCP port 5514, and then write those logs to the Elasticsearch data system in daily indexes titled "lsar".  The index name can easily be changed by editing that file, and changing the name of the index.

### Kibana
The configuration file ./kibana/kibana.yml contains the basic Kibana
configuration.  Changing the elasticsearch.hosts configuration will allow you to
point Kibana at a differents host.  It is set to listen on the standard TCP port
of 5601 on the Docker host.  This can be changed in the docker-compose.yml file.

### Elasticsearch
This instance of elasticsearch is configured to listen on the standard TCP port
of 9200.  This can be changed in the docker-compose.yml file.  When you first
bring up the system, there is no index created yet.  Once SecureSPhere has been
configured to send audit data to the system, an index starting with "lsar-" will
be created.  A new index will be built daily.  Creating a master index of
"lsar-" will allow all of the daily indexes to be rolled up into a single index
for reporting and visualzation purposes.

### Docker Setup
From the repo's base directory, perform the following to get setup for the first time
```
docker-compose build # builds all containers
docker-compose up # Starts all containers
```
This should execute for a while and stop generating further messaegs ToDo: Add message to look for. Kill at this point (crtl+c) and run the process in background
```
docker-compose up -d
```

By default the following services should be running on the ports below:

| Service | Port |
| ------- | ---- |
| elastic | 9200 |
| kibana  | 5601 |
| fluentd | 5514 |

To check if they're running, execute:
```
netstat -na | grep 9200 # elastic
netstat -na | grep 5601 # kibana
netstat -na | grep 5514 # Logstash/FluentD syslog listener
```

each should return at least a line of output
