# ansible
To run a playbook and ask for sudo

    ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook ./playbooks/setup_swarm.yml -i ../inventory/production -K --private-key=./{{}}KEY_NAME.pem -u ec2-user --fork 10 -e serial_number=10

If not using .pem key

    ssh-copy-id username@host_ip
