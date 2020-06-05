# Docker Swarm
The files in this repository will build out an ELK cluster using Docker Swarm

## Prerequsites 

- Have a Docker swarm set up.
- Label the swarm nodes as either `elastic == master` or `elastic == data`. This is required as the Elasticsearch nodes are ran in global mode, limiting one instance of Elasticsearch to one machine.
- Modified the Docker service memory limits on each machine to `LimitMEMLOCK=infinity`
- Added VM max map count (`vm.max_map_count=262144`) to each machine
- Modified the memory (`resources.limits.memory: 18G`) for each service in the docker-compose.yml file to a suitable amount i.e. if your machine has 16GB of ram you must set it to a value less than 16GB. Remember to allow enough memory for both Logstash and Elasticsearch to run on the data nodes** 

You can use the ansible script `/bare_metal/ansible/setup_swarm.yml` to set up a swarm to try it out. **The ansible scripts do not set up a secure production ready swarm so you should only use this for test purposes.**

See the Ansible Docer Swarm section at the bottom of this readme.

### Docker Setup
To deploy the stack load the docker-compose.yml file onto one of your master nodes and run the following command:

    docker stack deploy --compose-file ./docker-compose.yml open_reporting
    
To see if all containers have started correctly run the command below (NOTE: It may take a few minutes for the services to start)

    docker service ls
    

To remove the stack run:

    docker stack rm open_reporting

By default the following services should be exposed on the ports below:

    kibana - 5601
    logstash - 5514

### Ansible Docker Swarm
This play will set up your cluster with swarm managers and workers. **Nothing will be run on the manager nodes**.

This should be ran from the `/bare_metal/ansible` directory.

Add in IP addresses to the different blocks in the hosts file './ansible/inventory/production/swarm_hosts'.

| Group                          | Explaination                                                                                                                                     |
| -------------------------------| ------------------------------------------------------------------------------------------------------------------------------------------------ |
| swarm_manager                  | Single host. Creates and administers the swarm                                                                                                   |
| additional_swarm_managers      | Additional swarm managers are listed here. We recommend 2 (2 here plus the swarm-manager)                                                        |
| swarm_workers                  | Add remaining hosts here. They will run the Elastic cluster                                                                                      |
| master_node                    | Nodes listed here will be reserved for Elasticsearch nodes configured as master nodes. Check the Elastic website for the latest recommendations. |
| data_nodes                     | Nodes listed here will be reserved for Elasticsearch nodes configured as data nodes                                                              |

#### Running
Once you have divvied up your machines, run the following command to create your swarm:

    ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook ./playbooks/setup_swarm.yml -i ./inventory/production/swarm_hosts --private-key=./{{ KEY_NAME }}.pem -u ec2-user --fork 10 -e serial_number=10