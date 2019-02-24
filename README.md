Redis cluster REST API
======================
## Usage:
`REDIS_HOST=localhost:7000 ./main &`

Using in docker-compose.yml file:
```yaml
services:
  redis-api:
    image: kraynev/redis-api:latest
    logging:
      driver: json-file
    ports:
      - 8000:8000
    environment:
      REDIS_HOST: "redis:7000"
``` 

## Commands:
1. [Get whole cluster keys](#keys)
2. [Set any key value](#set)
3. [Add to queue](#lpush)
4. [Review queue records](#lrange)

<a name="keys"></a>
## 1. Get whole cluster keys
#### request:
    curl -XGET http://container.host/keys?pattern=key*
#### response:
    {"result":["keyNames"]}
<a name="set"></a>
## 2. Set any key value
#### request:
    curl -XPOST http://container.host/set/keyName \
        -d 'KEY VALUE STRING'
#### response:
    {"result":"OK"}
<a name="lpush"></a>
## 3. Add to queue
#### request:
    curl -XPOST http://container.host/lpush/queueName \
        -d '{"queue":"record"}'
#### response:
    {"result":{"in_queue":1}}
<a name="lrange"></a>
## 4. Review queue records
#### request:
    curl -XGET http://container.host/lrange/queueName
#### response:
    {"result":["{\"queue\":\"record\"}"]}
