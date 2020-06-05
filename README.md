# Imperva Open Reporting
The product intends to address reporting and visualization needs of customers that are current users of Imperva Database Activity Monitoring (DAM) solution with capabilities to ingest and process massive data. Built using open source technologies (ELK stack), it is intended as a recipe for customers to manage their own reports and dashboards to suit individual reporting use cases. The offering will be open source and community supported in Imperva’s Github repository, allowing for users build on the platform.
Intended use cases
1. Self- Service Visualization and Dashboarding
Kibana provides rich options to view DAM activity data flowing into the system. Querying in Kibana is simple with prompts and suggestions as one builds on the data to query. This can then be used to build visualizations by aggregating and bucketing data based on one or more fields.
2. Security Analytics
The dashboards provide ability to drill down deeper into visualizations to understand patterns and variations in the data. For example, a sudden spike in DDL queries or failed logins is a red flag that can be dug deeper into.
3. Audit and Compliance
By setting up a policy in SecureSphere to send all events to LSAR, customers are able to perform audit and compliance on all database activities. Each of the report or visualization in the Kibana dashboard can be exported to a CSV raw data file (for free) or a PDF (requires Elastic  platinum license). This CSV can be formatted into compliance reports using other BI tools such as Microsoft Excel and Tableau.

## Repository Overview
The repository is broken down into several different sections as outline below

### Bare Metal
Contains ansible playbooks and terraform scripts to quickly get you set up with a bare metal cluster.

### Docker
Contains Docker compose files for setting up a single node or a cluster on a Docker swarm

### Kibana Setup
The files located in the "kibana_setup" directory are used to quickly set up Kibana with some already made dashboards and policies

  * dashboards-export.json - contains the dashboards, visualizations, and other miscellaneous configurations for Elasticsearch and Kibana.  After the first few messages have arrived in Elasticsearch you will be able to create the index and then import the json file.  Do not try and import it until the index has been built.
  * index-lifecycle-management-policy - contains a basic ILM policy to move data between nodes.
  * index-template - A basic index template. **Modify shards/replicas based on your own hardware setup**

### SecureSphere
Files for integrating SecureSphere with the open reporting solution

  * syslog-message.txt - this file containes the properly formated syslog message ot be used in SecureSphere when creating the followed action to be used to send audit data to the reporting system.

### Tools
Contains tools used in benchmarking the elasticsearch cluster


## Supported Hardware
This repo has been extensively tested using Amazon Linux 2. You may encounter issues with other distributions
