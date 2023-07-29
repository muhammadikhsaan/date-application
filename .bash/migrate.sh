#!/bin/bash

if [[ $ENVIRONTMENT == "PRERELEASE" ]];
  then
    export $(echo $(cat .env.prerelease | xargs)) && go run pensiel.com/migrations
elif [[ $ENVIRONTMENT == "PRODUCTION" ]];
  then
    export $(echo $(cat .env.production | xargs)) && go run pensiel.com/migrations
elif [[ $ENVIRONTMENT == "DEVELOPMENT" ]];
  then
    export $(echo $(cat .env.development | xargs)) && go run pensiel.com/migrations
elif [[ $ENVIRONTMENT == "TEST" ]];
  then
    export $(echo $(cat .env.test | xargs)) && go run pensiel.com/migrations
else
    export $(echo $(cat .env.local | xargs)) && go run pensiel.com/migrations
fi