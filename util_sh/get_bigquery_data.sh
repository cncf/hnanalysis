#!/bin/bash
if [ -z "$1" ]
then
  echo "$0: you need to provide start date as an argument"
  exit 1
fi
cat BigQuery/query.sql | sed "s/{{from}}/$1/g" | bq --format=csv --headless query --use_legacy_sql=false -n 1000000 --use_cache 1>data/hn.csv
