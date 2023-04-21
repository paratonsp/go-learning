# learning-go
Go Rest API:


Installing:
- docker build -t learning-go .
- docker run -d -p 8080:8080 --name=learning-go learning-go:latest

InstallinV2:
- docker compose -f nginx-compose.yaml up -d
- docker compose -f app-compose.yaml up -d

### Go Rest API

* Simple CRUD to MySQL
* JWT Cookie
* MUX Middleware 
* Store File to AWS S3


## Installing Docker

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. First, update your existing list of packages:
   ```sh
   sudo apt update
   ```

2. Clone the repo
   ```sh
   git clone https://github.com/your_username_/Project-Name.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```
