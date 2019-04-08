# play with golang on google cloud function

# Pre

- install gcloud in http://macappstore.org/google-cloud-sdk/

## Deploy

Deploy function on google cloud function

first setup project id
```bash
# fill with project id
PROJECT_ID=<project-id>

# set config to spesific projectid
gcloud config set project $PROJECT_ID
```

let's deploy
```bash
# hello is function name, F is namefunction on file function.go
gcloud functions deploy hello --runtime go111 --entry-point F --trigger-http
```

delete function
```bash
# hello is function name
gcloud functions delete hello
```

deploy with env.yaml
```bash
gcloud functions deploy hello --runtime go111 --env-vars-file env.yaml --entry-point F --trigger-http
```

deploy with region
```bash
gcloud functions deploy hello --runtime go111 --entry-point F --region asia-northeast1 --trigger-http  
```

