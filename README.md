# GraphMicroService

this is a safe data transmission system made of 4 parts : (sender, reciever, broker, destination)

1_ Sender sends data with GRPC to the reciever
2_reciever recieves them and send it to borker with TCP Socket
3_broker recives them and log them in a file then send all of them to destination with TCP Socket

1.every buffer that have been used is preallocated 
2.all of functions are passing by refrence because we assume that data is huge and we dont want to kill GC (main data is on heap)
3.connection failures are getting handled by DATA_HANDLER(disk and memory implemention)
4.every failed request will be added to an array it will continue sending requests and handler on another goroutine will read them and send them


run the app with these following steps :
1. go build cmd/main/main.go -o ./app
2. ./app dest
3. ./app broker
4. ./app reciever
5. ./app sender

so many other optimization can be made :
1. using config.yaml
2. running them on 4 different machine (it will make the speed up to 4 times faster beacuse the limitaion is threads and their open sockets)
with creating more goroutine we are just making more context switchings not more cpu efficiency
3. using runtime.lockOSThread
4. implement the DATA_HANDLER better and remove memory and disk footprints there are ofc so much better tools these days (s3, seaweed, ceph)
5. use redis streams or temporal so data integrity and availablity will be easily achieved



