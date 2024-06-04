# Gflights Notifier
This program send mails when it find some available tables in [Google Flights](https://www.google.com/travel/flights) website.

## Run locally

```bash
MAIL_SENDER_ADDRESS="sender@gmail.com" MAIL_SENDER_PASSWORD="REPLACEME" MAIL_RECEIVER_ADDRESS="receiver@gmail.com" go run *.go
```

## Use it with Tor

To prevent your IP to be exposed to the booking website, you can use tor.

First run Tor proxy Docker container:
```bash
docker run -d --restart=always --name tor-socks-proxy -p 127.0.0.1:9150:9150/tcp peterdavehello/tor-socks-proxy:latest
```

Then export TOR_PROXY_URL environment variable :

```bash
export TOR_PROXY_URL="socks5://127.0.0.1:9150"
```
