# Short-url

Visit the website: [small.ml](https://small.ml)

## Check List

1. Implement Backend API

- [x] Return Short Url
- [x] Return Long Url
- [x] Redirect using Short Url
- [x] Document usage on home page

2. Host locally on docker

- [x] Run on docker
- [x] Use Environmental variables to supply port and database connection
- [x] Create Docker-compose file

3. Publish on AWS

- [x] Set up remote mongodb
- [x] Create seperate local and prod docker-compose files
- [x] Configure app to run on HTTPS
- [x] Host on AWS

4. Refactor 
- [x] Move database logic to repository
- [x] Improve error responses. Go's error messages are short and useless.
- [ ] Unit Tests
- [ ] Set up CircleCI
- [x] Use swagger.io on the home screen

Useful Resources

1. [Automating HTTPS with LetsEncrypt using autocert library](https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go-with-little-help-of-lets-encrypt.html)