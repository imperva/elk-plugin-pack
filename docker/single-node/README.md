# Single Node
The files in this repository will build out a single dockerized ELK stack which will consist of:

  * 1 - Elasticsearch node
  * 1 - Kibana node
  * 1 - Logstash node


### Logstash
The Dockerfile located in ./logstash  will build the logstash docker image based on the latest version the first time the docker-compose.yml file run.

The configuration file, ./logstash/pipeline/audit-pipeline.conf, is configured
to listen for JSON formatted syslog messages on TCP port 5514, and then write
those logs to the Elasticsearch data system in daily indexes titled
"audit-YYYY-MM-dd".  The index name can easily be changed by editing that file, and changing the name of the index.

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

## Docker Setup
From the repo's base directory, perform the following to get setup for the first time

### Start containers
This should execute for a while and stop generating further messaeges:

    docker-compose up 

To run in the background use:

    docker-compose up -d


By default the following services should be running on the ports below:

| Service  | Port  |
| -------- | ----- |
| elastic  | 9200  |
| kibana   | 5601  |
| logstash | 5514 |


To check if they're running, execute:
```
netstat -na | grep 9200 # elastic
netstat -na | grep 5601 # kibana
netstat -na | grep 5514 # Logstash syslog listener
```

each should return at least a line of output
