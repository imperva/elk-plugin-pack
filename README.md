### Files
The files located in the "files" directory are used to finish the configuration
of the system.  

  * lsar-export.json - contains the dashboards, visualizations, and other miscellaneous configurations for Elasticsearch and Kibana.  After the first few messages have arrived in Elasticsearch you will be able to create the index and then import the json file.  Do not try and import it until the index has been built.
  * syslog-message.txt - this file containes the properly formated syslog message ot be used in SecureSphere when creating the followed action to be used to send audit data to the LSAR system.

### Notes
Changing the listening ports in the docker-compose.yml file will require the configuration files for Kibana and FluentD to be changed.  Please make sure to check all configruation files if you make changes to the system names or their ports.