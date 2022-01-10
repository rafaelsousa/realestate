#!/usr/bin/env bash

#Create property
realestated tx realestate create-property --from owner1 "123 happy st" "chainville"  "BC"  "Canada" "V5K0A2" "57.371889" -- -125.337451

# Request the certification from the property
realestated tx realestate request-certification 1 100domus cosmos1ru57u8eh9asxmqg7h73csuu0v6clp5nth29wvn --from owner1 --gas auto


# The surveyour creates the certificate
realestated tx

