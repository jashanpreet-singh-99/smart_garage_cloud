# Smart Garage Cloud Server (SGC)

This Code section contains the ASE cloud service Implementation. THe SGC is implemented using Golang for now. The golang file can be built into a binary file which is running on the Azure server instance. This server instance uses the port 80 to formulate this connection. Later on it can be switched to 443 for better security.

For now the Service contains 5 main endpoints that are required for testing the emulator and the mobile application. These endpoint are described as follow:
1. Get Light status
2. Set Light status
3. Get Door status
4. Open | Close Door
5. Lock | Unlock Door

## Endpoints

### 1. Get Light status

Request : GET  
End point : /Lights

Return : 0 (On) | 1 (Off)

Additional Headers
* Light : \<Position of the Light\> i.e Light_F_L  

![Alt text](misc/Endpoint_1.png?raw=true "Endpoint 1")

### 2. Set Light status

Request : PUT  
End point : /Lights

Additional Headers
* Light : \<Position of the Light\> i.e Light_F_L  
* Value : 0 (Off) | 1 (On) 

![Alt text](misc/Endpoint_2.png?raw=true "Endpoint 2")

### 3. Get Door status

Request : GET  
End point : /Door

Return : 0 (Closed) | 1 (Open)

Additional Headers = **None**

![Alt text](misc/Endpoint_3.png?raw=true "Endpoint 3")

### 4. Open | Close Door

Request : PUT  
End point : /Door

Additional Headers
* Command : OPEN | CLOSE  

![Alt text](misc/Endpoint_4.png?raw=true "Endpoint 4")

### 5. Lock | Unlock Door

Request : PUT  
End point : /DoorStop

Additional Headers
* Status : LOCK | UNLOCK 

![Alt text](misc/Endpoint_5.png?raw=true "Endpoint 5")

## To Do
- [ ] Login DB  
- [ ] JWT tokens (Security)  
- [ ] Test case 
- [ ] Translation  
- [ ] Modify return values  
