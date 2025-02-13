#!/bin/sh


KEY='' # https://console.cloud.google.com/google/maps-apis/credentials

while IFS= read -r line; do
  DATA=$(curl -s "https://maps.googleapis.com/maps/api/geocode/json?address=component:locality=${line},component:country=GB&sensor=false&key=${KEY}")
  #DATA='{ "results" : [ { "address_components" : [ { "long_name" : "Wrexham", "short_name" : "Wrexham", "types" : [ "locality", "political" ] }, { "long_name" : "Wrexham", "short_name" : "Wrexham", "types" : [ "postal_town" ] }, { "long_name" : "Wrexham Principal Area", "short_name" : "Wrexham Principal Area", "types" : [ "administrative_area_level_2", "political" ] }, { "long_name" : "Wales", "short_name" : "Wales", "types" : [ "administrative_area_level_1", "political" ] }, { "long_name" : "United Kingdom", "short_name" : "GB", "types" : [ "country", "political" ] } ], "formatted_address" : "Wrexham, UK", "geometry" : { "bounds" : { "northeast" : { "lat" : 53.0730607, "lng" : -2.9454931 }, "southwest" : { "lat" : 53.0291108, "lng" : -3.0308429 } }, "location" : { "lat" : 53.04304, "lng" : -2.992494 }, "location_type" : "APPROXIMATE", "viewport" : { "northeast" : { "lat" : 53.0730607, "lng" : -2.9454931 }, "southwest" : { "lat" : 53.0291108, "lng" : -3.0308429 } } }, "partial_match" : true, "place_id" : "ChIJXTc5V9C1ekgRjHfMNGlgddU", "types" : [ "locality", "political" ] } ], "status" : "OK" }'
  ONE=$(echo ${DATA} | jq -r '.results[0].address_components[] | select(.types[] | contains("administrative_area_level_1")) | .long_name')
  TWO=$(echo ${DATA} | jq -r '.results[0].address_components[] | select(.types[] | contains("administrative_area_level_2")) | .long_name')
  echo ${line},${TWO},${ONE}
done < address.txt
