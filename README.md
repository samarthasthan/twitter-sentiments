Here's the revised version of your document with corrected English:

---

# Twitter Sentiment Analysis Backend

This backend service is designed for Twitter Sentiment Analysis, providing a robust and scalable solution for processing and analyzing tweets. Built using Golang, Kafka, gRPC, and Go-Fiber, it combines the efficiency of Go with the power of Kafka messaging and the flexibility of gRPC for seamless communication.

## Tech Stack

**Language:** GoLang

**Library:** GORM

**Framework:** gRPC, Apache Kafka

**Tools:** Docker

## System Design

![App Screenshot](https://i.ibb.co/dt4WLVT/twtter-sentiments.png)

## Live URL

[twitter-go.samarthasthan.com](https://twitter-go.samarthasthan.com)

## API Reference

#### Welcome API

```http
  GET /
```

#### Get Tweets

```http
  GET /tweets
```

Returns a JSON object containing the sentiment analysis results for the 10 most recent tweets. A score of 1 indicates a positive sentiment, while a null score indicates a negative sentiment.

#### Example

![App Screenshot](https://i.ibb.co/J2wCNqC/Screenshot-2024-01-07-at-4-20-34-PM.png)

## Run on VPS

1. Clone the project

```bash
git clone https://github.com/samarthasthan/twitter-sentiments
```

2. Download the dataset CSV file from Kaggle, rename it to `tweets.csv`, and copy it to `twitter-sentiments/twitter-api`.

   [Kaggle Dataset Link](https://www.kaggle.com/datasets/kazanova/sentiment140)

   Copy the `tweets.csv` file from your local machine to the VPS.

   ```bash
   scp -i "secret.pem" path/tweets.csv ubuntu@publicip:github/twitter-sentiments/twitter-api
   ```

   If you encounter a permission denied error, use the following command on the VPS:

   ```bash
   sudo chmod 777 (remote folder)
   ```

3. Navigate to the project directory

```bash
cd twitter-sentiments
```

4. Start Docker

```bash
docker compose up -d
```

![App Screenshot](https://i.ibb.co/171krY5/Screenshot-2024-01-07-at-4-18-51-PM.png)

5. Set up Nginx

```bash
apt update
apt install nginx
vim /etc/nginx/conf.d/samarthasthan.conf
```

Then copy your Nginx config to `{domain}.conf`

```nginx
server {
    listen 80;
    listen [::]:80;

    server_name twitter-go.samarthasthan.com;

    location / {
        proxy_pass http://localhost:9058;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

6. Install SSL certificate using Certbot

```bash
sudo apt update
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d twitter-go.samarthasthan.com
```

7. Automate SSL renewal with crontab

```bash
crontab -e
```

And paste the crontab job

```
0 6 * * 0 certbot renew -n -q --pre-hook "service nginx stop" --post-hook "service nginx start"
```

8. Finally, restart Nginx

```bash
nginx -t
sudo systemctl restart nginx
```

## Run Locally

1. Clone the project

```bash
git clone https://github.com/samarthasthan/twitter-sentiments
```

2. Download the dataset CSV file from Kaggle, rename it to `tweets.csv`, and copy it to `twitter-sentiments/twitter-api`.

   [Kaggle Dataset Link](https://www.kaggle.com/datasets/kazanova/sentiment140)

3. Navigate to the project directory

```bash
cd twitter-sentiments
```

4. Start Docker

```bash
docker compose up -d
```

![App Screenshot](https://i.ibb.co/171krY5/Screenshot-2024-01-07-at-4-18-51-PM.png)

Note: Now, you can access it using `localhost:9058/tweets`.

## Authors

- [@samarthasthan](https://www.github.com/samarthasthan)

---
