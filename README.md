### Simple Message Queue
---

Simple implementation message queue with go and presistence data in redis running on port `localhost:8888` by default, change inside config file if the address already use for another service. this message queue use authentication by default please define user config inside ```.env``` or use default user at ```root/rahasia2020```  :D 

1. Requirement
    - Go installed and support go mod
    - Redis installed 
2. Setup
    - clone my git repoistory ```git clone https://github.com/cag000/simple-message-queue.git && cd ./simple-message-queue```
    - build app ```go build -o app```
    - run ```./app -h```
    ```
    > ./app -h
    Usage of ./app:
    -auth
            auth [./app -auth=true -user=$username -pwd=$password]
    -config string
            config file .env format
    -consumer
            consumer [./app -consumer=true -queue=$name]
    -create_queue
            create queue [./app -create_queue=true -queue=$name]
    -delete_queue
            delete queue [./app -delete_queue=true -queue=$name]
    -producer
            producer [./app -producer=true -queue=$name]
    -pwd string
            password [./app -pwd=$password]
    -queue string
            name queue [./app -queue=$name]
    -server_queue
            Make live server [./app -server_queue=true]
    -user string
            username [./app -user=$user]
    ```
3. Useage<br>
    Before use this app please define ```cfg.env``` config for your system
    - Start up server queue<br>
    ```./app -config=cfg.env -auth=true -user=$user -pwd=$pwd -server_queue=true```
    - Authentication<br>
    ```Define your match inside .env file```
    - Create Empty Queue<br>
    ```./app -config=cfg.env -auth=true -user=$user -pwd=$user -create_queue=true -queue=$queue_name```
    > or another method use nc in linux (nc localhost 8888), format: Delete=$queue_name
    - Delete Queue<br>
    ```./app -config=cfg.env -auth=true -user=$user -pwd=$user -create_queue=true -queue=$queue_name```
    > Note: Delete queue is delete any message inside of it
    - Producer<br>
    ```./app -config=cfg.env -auth=true -user=$user -pwd=$user -producer=true -queue=$queue_name```
    > Note: default command use data dummy.json, if you want to try another data use nc in linux, (nc localhost 8888) Format: $queue=$message. Produce message automatically create new queue if there is no any exists with that name
    - Consumer<br>
    ```./app -config=cfg.env -auth=true -user=$user -pwd=$user -consumer=true -queue=$queue_name```

4. Bugs<br>
    The useage of CPU high, because not properly implement the routine in go, still on research for better improvement :D

Feel free to modify and use it, or open new request