#!/bin/bash
ID=70 
curl -s https://01.tomorrow-school.ai/assets/superhero/all.json | jq -r ".[] | select(.id == $ID) | .name"
