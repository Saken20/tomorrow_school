
curl https://01.tomorrow-school.ai/assets/superhero/all.json | jq ".[] | select(.id == $HERO_ID) | .connections.relatives" | tr -d '"'
