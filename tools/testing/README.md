# Test Tools

## Build
To build run the following command from the **tools/testing** directory

    env TAG=${ENTER_TAG} SERVICE_NAME={{ log-sender | kibana-querier }} docker-compose -f testing/build.yml build

## Running
To run the test log sender tool

    log-sender {{ TARGET_IP_ADDRESS }} {{ NUMBER_OF_LOGS }} {{ THREADS_TO_USE }}

To run the test log sender tool using Docker

    docker run -d imperva/log-sender:{{ TAG }} {{ TARGET_IP_ADDRESS }} {{ NUMBER_OF_LOGS }} {{ THREADS_TO_USE }}

If you want to run multiple instances quickly you can use the helper script **'./scripts/run.sh'**. Create a file in the same directory called **hosts.txt** and populate with a list of IP addresses or DNS names.

    ./run.sh


## Useful commands

To stop and remove all running containers

    docker stop $(docker ps -a -q)
    docker rm $(docker ps -a -q)