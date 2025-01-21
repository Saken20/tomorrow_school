
# curl https://01.tomorrow-school.ai/assets/superhero/all.json | jq -r ".[] | select(.id == $HERO_ID) | .connections.relatives"

#!/bin/bash

# Ensure HERO_ID is set
if [ -z "$HERO_ID" ]; then
  echo "Error: HERO_ID is not set."
  exit 1
fi

# Fetch the JSON data from the given URL using curl
json_data=$(curl -s https://01.tomorrow-school.ai/assets/superhero/all.json)

# Extract the subject's family using HERO_ID and jq
family=$(echo "$json_data" | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .relatives')

# Check if family data was found
if [ "$family" == "null" ]; then
  echo "No family information found for HERO_ID: $HERO_ID"
  exit 1
fi

# Print the family members without quotes
echo "$family" | sed 's/"//g'
