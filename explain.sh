#!/bin/bash

# Display the first and last name of the key witness
echo "FirstNameOfWitness LastNameOfWitness"

# Display the interview number of the witness
echo "123456"

# Display the color and make of the car of the main suspect
echo "Red Ferrari"

# Display the names of the 3 other main suspects in alphabetical order of their last name
echo "FirstNameOfSuspect1 LastNameOfSuspect1"
echo "FirstNameOfSuspect2 LastNameOfSuspect2"
echo "FirstNameOfSuspect3 LastNameOfSuspect3"

chmod +x explain.sh

./explain.sh | cat -e