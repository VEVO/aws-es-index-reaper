# aws-es-index-reaper

[![GoDoc](https://godoc.org/github.com/VEVO/aws-es-index-reaper?status.svg)](https://godoc.org/github.com/VEVO/aws-es-index-reaper)
[![Test Status](https://github.com/VEVO/aws-es-index-reaper/workflows/tests/badge.svg)](https://github.com/VEVO/aws-es-index-reaper/actions?query=workflow%3Atests)
[![Go Report Card](https://goreportcard.com/badge/github.com/VEVO/aws-es-index-reaper)](https://goreportcard.com/report/github.com/VEVO/aws-es-index-reaper)

Remove AWS ES indexes that exceed a configurable threshold.

If you are using something like logstash or zipkin that makes daily indexes then `aws-es-index-reaper` can be used to keep the index count in check.

The index list is retreived from elasticsearch using the GET settings API (https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-settings.html#indices-get-settings).

For example if the env var INDEX_PREFIX has the value logstash the query would be.

```
GET /logstash*/_settings
```

The indexes are sorted using the CreationDate attribute on each index and the oldest indexes above the threshold MAX_COUNT are removed.

# Configuration

Control the configuration by setting the following environmental variables

```
INDEX_PREFIX = Index name prefix used to locate matching indexes.
ES_URL = Url to interact with AWS ES domain.  Do not include trailing slash.
MAX_COUNT = Any matching indexes over this count will be removed.
```

# Deploy
To deploy aws-es-index-reaper you can modify the [reaper-scheduledjob.yaml](reaper-scheduledjob.yaml) to your needs and then run

```
kubectl apply -f reaper-scheduledjob.yaml
```
# Docker Image
You can find the docker image [here](https://hub.docker.com/r/vevo/reaper/tags/).

