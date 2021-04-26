# Install
1. Clone repository
```
git clone https://github.com/sepi0/shortener.git
```

2. Change into directory
```
cd shortener
```
3. Initialize go.mod 
```
go mod init
```
4. Install gorilla mux
```
go get -u github.com/gorilla/mux
```
5. Install BoltDB
```
go get github.com/boltdb/bolt/
```
6. Build the module
```
go build .
```

# Run
```
go run .
```
   


### Usage
1. Visit http://localhost:8080/
2. Paste desired URL inside the input field
3. If you try to shorten invalid URLs, you'll get an 400 error
4. On success, you'll be redirected to http://localhost:8080/create where you'll find your shortened URL
5. Use your new URL


###*All shortened URLs are stored in a BoltDB database "URLs" bucket* 