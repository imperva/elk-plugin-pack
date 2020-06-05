# Ansible
To quickly create a bare metal cluster, use the Ansible plays we have created. 

Current Ansible scripts are tested with Amazon Linux machines images, and is expected to work as-is for any RHEL / Cent OS based linux. This may not work with Debian based distributions

## Prerequsites 

- Add an IP addresses to the different blocks in the hosts file './ansible/inventory/production/bare_metal_hosts'.
- You must define 1 node-master. You can define 1 or more of all other node types.

| Group                          | Explaination                                                                                                                                     |
| -------------------------------| ------------------------------------------------------------------------------------------------------------------------------------------------ |
| node-master                  | The initial master node that all other masters/data nodes will communicate with. This node will start without waiting for other nodes. This prevents an endless loop of master nodes waiting for an initial master node to start.                                                                                                   |
| additional_master_nodes      | Nodes listed here will be reserved for Elasticsearch nodes configured as master nodes. Check the Elastic website for the latest recommendations.                                                        |
| data_nodes                  | Add remaining hosts here. They will run the Elastic cluster                                                                                      |
| master_node                    | Nodes listed here will be reserved for Elasticsearch nodes configured as data nodes |
| coordinating_nodes                     | Nodes listed here will be reserved for Elasticsearch nodes that are neither configured as data or master nodes. Coordinating nodes are reserved for queries only                 |

## Running
Before running the playbook update the following sections of the playbook to match your infrastructure.

- In the playbook (ansibe /playbooks/setup_bare_metal_cluster.yml
- Update es_heap_size in  to a suitable value (No more than 32GB is recommended by Elastic)
- Update es_data_dirs to point to a suitable location on your machines disk
- If you are not using AWS set the mount_volume variable to false.
- Under `./ansible/config-files/logstash/jvm.options` change -Xms1g and -Xmx1g to suitable values (more is better but both values must be the same) based on your machines memory.
- Install Elasticsearch ansible module using command 
ansible-galaxy install elastic.elasticsearch,7.7.1

You can then use the folloying command to set up the cluster.

    ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook ./playbooks/setup_bare_metal_cluster.yml -i ./inventory/production/bare_metal_hosts --private-key=./{{ KEY_NAME }} -u {{ USERNAME }} --fork 10 -e serial_number=10

## Current limitations to the Ansible script playbook
- Re-running the playbook will not reinstall the elasticsearch cluster.
