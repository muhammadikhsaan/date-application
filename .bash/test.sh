#!/bin/bash
if [[ $ENVIRONTMENT == "PRERELEASE" ]];
  then
    export $(echo $(cat .env.prerelease | xargs)) && go test -v pensiel.com/testing
elif [[ $ENVIRONTMENT == "DEVELOPMENT" ]];
  then
    export $(echo $(cat .env.development | xargs)) && go test -v pensiel.com/testing
elif [[ $ENVIRONTMENT == "PRODUCTION" ]];
  then
    export $(echo $(cat .env.production | xargs)) && go test -v pensiel.com/testing
elif [[ $ENVIRONTMENT == "LOCAL" ]];
  then
    export $(echo $(cat .env.local | xargs)) && go test -v pensiel.com/testing
else
  export $(echo $(cat .env.test | xargs)) && go test -v pensiel.com/testing
fi