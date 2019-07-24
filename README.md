# Large Scale Audit Report - Phase 1
## Create a single stack proof of concept for a large scale audit reporting platform for the Imperva Data Security platform.

## Based on [Elasticsearch](https://www.elastic.co/),[Kibana](https://www.elastic.co/products/kibana), and [FluentD](https://www.fluentd.org/).

The files in this repository will build out a single dockerized EFK stack which will consist of:

  * 1 - Elasticsearch v7.x nod
  * 1 - Kibana v7.x nod
  * 1 - FluentD node

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

### Files
The files located in the "files" directory are used to finish the configuration
of the system.  

  * lsar-export.json - contains the dashboards, visualizations, and other miscellaneous configurations for Elasticsearch and Kibana.  After the first few messages have arrived in Elasticsearch you will be able to create the index and then import the json file.  Do not try and import it until the index has been built.
  * syslog-message.txt - this file containes the properly formated syslog message ot be used in SecureSphere when creating the followed action to be used to send audit data to the LSAR system.

### Notes
Changing the listening ports in the docker-compose.yml file will require the configuration files for Kibana dn FluentD to be changed.  Please make sure to check all configruation files if you make changes to the system names or their ports.
