Learning Golang
================

## Videos
Please take a few minutes when you get a chance to go look through the branches. That's where the "past" code exists. Thanks!


## RSA512
```bash
ssh-keygen -t rsa -b 4096 -f keys/app.rsa
openssl rsa -in keys/app.key -pubout -outform PEM -out keys/app.key.pub
```
4096 for RA512
2048 for RA256