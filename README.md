# Install
1. ``git clone https://github.com/sepi0/shortener.git``

2. ``go build main.go``

## Run
1. ``go run main.go``

2. ``visit http://localhost:8080/``


### Usage
1. Visit http://localhost:8080/
2. Paste desired URL inside the input field
3. If you try to shorten invalid URLs, you'll get an 400 error
4. On success, you'll be redirected to http://localhost:8080/create where you'll find your shortened URL
5. Use your new URL


###*All shortened URLs are stored in a BoltDB database "URLs" bucket* 