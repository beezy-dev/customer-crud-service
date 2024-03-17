# traveler-processed-service

The ```traveler-processed-service``` is an API server to record processed travelers for a period of 48-hour window to speed up next hop processing. Once the 48-hour window is over, the record will be offloaded to a separate data backend for archive and compliance purposes.

## requirements
The requirements are the followings:

- [ ] a CRUD API service providing a JSON formatted output when queried   
- [x] a decoupled backend database service to connect to    
- backend databases to support at v1 
    - [ ] MariaDB
    - [x] Postgres   
    - [ ] sqlite
- [ ] the service should be recovering for a partioning from its backend database  
- [ ] the data structure should include the followings:  
    * an UUID 
    * first name
    * last name
    * multi hop travel
    * was previously processed
    * when previously processed 
    * should be reprocessed
- [ ] after the 48-hour window the record should be offloaded to a historical database
- [ ] concurrent writing to database and in-memory caching using sqlite to leverage gorm


## service sequence

1. starting service
1. validating backend support
1. validating backend connectivity 
1. validating if database exists
1. if not create it
1. reconnecting with backend targeting the database
1. migrating schema in database (including potential update)

## contributing

Clone the repository: 
```
git clone https://github.com/beezy-dev/traveler-crud-service.git 
```

Run CompileDaemon to catch any code changes to trigger a rebuild and start of your binary:
```
~/go/bin/CompileDaemon -build="go build -o traveler-processed-service.out" -command="./traveler-processed-service.out" -color -graceful-kill -graceful-timeout 5
``` 


