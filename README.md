# Twitter Sentiment Analysis Backend

This backend service is designed for Twitter Sentiment Analysis, providing a robust and scalable solution for processing and analyzing tweets. Built using Golang, Kafka, gRPC, and Go-Fiber, it combines the efficiency of Go with the power of Kafka messaging and the flexibility of gRPC for seamless communication.

## Tech Stack

**Langauge:** GoLang

**Library:** GORM

**Framework:** gRPC, Apache Kafka

**Tools:** Docker

## System design

![App Screenshot](https://i.ibb.co/dt4WLVT/twtter-sentiments.png)

## Run Locally

Clone the project

```bash
  git clone https://github.com/samarthasthan/twitter-sentiments
```

Download Dataset csv file from kaggle, rename it to tweets.csv and copy it to twitter-sentiments/twitter-api

https://www.kaggle.com/datasets/kazanova/sentiment140

Copy tweets.csv file from local machine to VPS

```
scp -i "secret.pem" path/tweets.csv ubuntu@publicip:github/twitter-sentiments/twitter-api

```

Go to the project directory

```bash
  cd twitter-sentiments
```

Start the docker

```bash
  docker compose up -d
```

![App Screenshot](https://i.ibb.co/171krY5/Screenshot-2024-01-07-at-4-18-51-PM.png)

## API Reference

#### Welcome API

```http
  GET /
```

#### Get Tweets

```http
  GET /tweets
```

Return a JSON object containing the sentiment analysis results for the 10 most recent tweets. A score of 1 indicates a positive sentiment, while a null score indicates a negative sentiment.

#### Example

![App Screenshot](https://i.ibb.co/J2wCNqC/Screenshot-2024-01-07-at-4-20-34-PM.png)

## Authors

- [@samarthasthan](https://www.github.com/samarthasthan)
