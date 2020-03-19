# discord-push-bot  
push message to your discord by send http request  

## Getting Started  
### 1.download [release binary file](https://github.com/fengjijiao/discord-push-bot/releases).  
### 2.modify config.yml ,such as:  
```
bottoken: "your_discord_bot_token"
botapiurl: "https://your_discord_bot_domain"
botserverport: "127.0.0.1:8001"
```
### 3.start  
```
./dcpush -c config.yml
```
### 4.get url for send message  
open discord enter `getpushurl` to the bot  
you may be receive  
```
https://dcpush.tk/send?channelid=61222****9321&sign=90se8f6d555f7g***7d8f9
```
### 5.send message through http api  
for example:  
```
curl 'https://dcpush.tk/send?channelid=61222****9321&sign=90se8f6d555f7g***7d8f9' --data "this is a test message!!!"
```
### 6.finally  
your discord has received message
