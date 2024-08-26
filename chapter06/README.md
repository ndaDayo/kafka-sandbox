## FileStreamSource

```
curl -X POST -H "Content-Type: application/json" --data '{
  "name": "local-file-source",
  "config": {
    "connector.class": "FileStreamSource",
    "tasks.max": "1",
    "file": "/data/input.txt",
    "topic": "connect-test"
  }
}' http://localhost:8083/connectors
```
## FileStreamSink
```
curl -X POST -H "Content-Type: application/json" --data '{
  "name": "local-file-sink",
  "config": {
    "connector.class": "FileStreamSink",
    "tasks.max": "1",
    "file": "/data/output.txt",
    "topics": "connect-test"
  }
}' http://localhost:8083/connectors
```
