#!/bin/bash

curl https://acad.learn2earn.ng/assets/superhero/all.json | jq ' .[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'