
curl https://01.tomorrow-school.ai/assets/superhero/all.json | jq -r ".[] | select(.id == $HERO_ID) | .connections.relatives" | tr -d '"'
