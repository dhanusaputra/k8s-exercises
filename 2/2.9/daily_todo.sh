#!/usr/bin/env bash

if [ "$URL" ]
then
curl -L -X POST "$URL" \
-H 'Content-Type: application/json' \
--data-raw '{"query":"mutation createTodo ($todo:String!,$done:Boolean!) {\n  createTodo(input:{text:$todo,done:$done}) {\n    text\n    done\n  }\n}","variables":{"todo":"test todo2","done":false}}'
fi
