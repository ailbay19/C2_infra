#!/bin/bash

machine_names=("lighthouse" "manager" "worker1" "worker2")

for machine_name in "${machine_names[@]}"; do
	rm -rf ${machine_name}
	mkdir ${machine_name}

	cp bin/nebula ${machine_name}

	cp keys/${machine_name}.* ${machine_name}
	cp keys/ca.crt ${machine_name}

	mv config-${machine_name}.yml ${machine_name}/config.yml
	
	cp run.sh ${machine_name}
	cp setup-machine.sh ${machine_name}
done
