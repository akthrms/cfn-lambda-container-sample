FROM --platform=linux/amd64 golang:1-alpine as builder

ENV GOARCH=amd64
ENV GOOS=linux

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o executable main.go


FROM --platform=linux/amd64 public.ecr.aws/lambda/go:1

COPY --from=builder /app/executable ${LAMBDA_TASK_ROOT}
CMD [ "executable" ]
