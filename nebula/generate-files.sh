#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <lighthouse_ip>"
  exit 1
fi

lighthouse_ip="$1"

machine_names=("manager" "worker1" "worker2")

TEMPLATE_CONFIG="template-config.yml"


# Insert machine_name and lighthouse_ip to relevant places.

for machine_name in "${machine_names[@]}"; do
	OUTPUT_CONFIG="config-${machine_name}.yml"

	envsubst '${machine_name} ${lighthouse_ip}' < "$TEMPLATE_CONFIG" > "$OUTPUT_CONFIG"

done
