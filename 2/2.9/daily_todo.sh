#!/usr/bin/env bash

random=$(curl -sLI https://en.wikipedia.org/wiki/Special:Random | grep "^location:" | cut -d " " -f2 | tr -d "\r")

if [ "$URL" ]
then
  curl -sL -X POST "$URL" \
-H 'Content-Type: application/json' \
--data-raw '{"query":"mutation createTodo ($todo:String!,$done:Boolean!) {\n  createTodo(input:{text:$todo,done:$done}) {\n    text\n    done\n  }\n}","variables":{"todo":"Read '"$random"'","done":false}}'
fi
