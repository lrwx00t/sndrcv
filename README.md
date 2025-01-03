# sndrcv
Send Receive a file between two machines

```
❯ ./sndrcv -mode=receive 2333
Connect to: 172.16.227.92:2333
Waiting for incoming connections...
File 'hello' received successfully!

# send
❯ ./sndrcv -mode=send hello 172.16.227.92:2333              
File sent successfully!
```
