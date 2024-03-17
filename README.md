# traveler-processed-service

The ```traveler-processed-service``` is an API server to record processed travelers for a 24 hour window to speed up next hop processing. 

The requirements are the followings:

- [ ] a CRUD API service providing a JSON formatted output when queried   
- [x] a decouple backend database service to connect to    
- backend databases to support at v1 should be 
    - [ ] MariaDB
    - [x] Postgres   
- [ ] the service should be recovering for a partioning from its backend database  
- [ ] the data structure should include the followings:  
    * an UUID 
    * first name
    * last name
    * multi hop travel
    * was previously processed
    * when previously processed 
    * should be reprocessed
- [ ] after the 24 hours window the record should be offloaded to a historical database
