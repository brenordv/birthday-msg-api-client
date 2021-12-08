# Birthday Quote API
## Client

This is the client for a silly project I've made.
The client will try to get a random birthday quote from the API and write it to the clipboard.

### Linux users
For this client to work properly, you must have 'xclip' or 'xsel' command to be installed.

## How to use
### Get random quote using default API language
Currently, the default API language is Brazilian Portuguese (pt-br)
```shell
bdayq.exe
```

### Get random quote in english 
```shell
bdayq.exe en-us
```


### Get random quote in Brazilian portuguese
```shell
bdayq.exe pt-br
```


## Ok. Got it...but why?
For fun. The api was made using Python (Flask) and is published on Heroku (@ https://raccoon-ninja-birthday-api.herokuapp.com/).
I work for a medium size company that's always hiring new employees, which means new birthday cards to write something nice and sign. 
I'm just adding some automation to the process. (I know, it's not the most empathic, personal or human thing to do, but was fun coding all this.)

## API Repo
https://github.com/brenordv/birthday-msg-api-app