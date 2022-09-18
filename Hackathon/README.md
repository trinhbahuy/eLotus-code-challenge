# How to run
## Setup environment
 - Download `eLotus-code-challenge.tar.gz` and uncompress
 - `cd /your-path/Hackathon`
 - Run `docker-compose up -d`
 - Run `docker ps` to check that we have 2 containers is running (`backend` is Go and `db` is Mysql)

## Try uploading file
- Access http://localhost:6969/upload
- upload load some files
- if success, a message `upload file success` will be shown, an image file will be added to `tmp` folder, and a record will be insert to database.

## Oauth2 Server
- I regretted for not completing this feature, although I try a lots.
- A good test, I will try to complete this Oauth2 later

### Note: If you want to access to database, please use username, password, and db name at `docker-compose.yml` 