# CloudFormationのサンプル（Lambda）

## ECRプッシュ

```sh
# ビルド
$ docker build . -t akthrms/go-lambda-sample

# タグつけ
$ docker tag akthrms/go-lambda-sample:latest <ID>.dkr.ecr.ap-northeast-1.amazonaws.com/akthrms/go-lambda-sample:latest

# プッシュ
$ docker push <ID>.dkr.ecr.ap-northeast-1.amazonaws.com/akthrms/go-lambda-sample:latest
```

## ローカル実行（RIE）

```sh
$ docker run -p 9000:8080 akthrms/go-lambda-sample

$ curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'
"Hello, Golang!!"
```

## スタック作成

```sh
$ aws cloudformation create-stack \
    --stack-name go-lambda-sample \
    --template-body file://template.yaml \
    --parameters ParameterKey=EcrImageUri,ParameterValue=<ECR URI> \
    --capabilities CAPABILITY_NAMED_IAM
```
