#!/bin/bash
curl https://01.tomorrow-school.ai/assets/superhero/all.json | jq ".[] | select(.id == 70).name"
