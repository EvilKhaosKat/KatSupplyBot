# KatSupplyBot
Supply telegram bot - registers supply 'requests' 


### How to use:
- create `token` file with bot token (can be obtained from 'Bot father')
- optional - create `admins` file, with admins usernames (each one on different line)
- to use local app exemplar
    - call `./launch.sh`
- to use docker container
    - call `./create_docker_image.sh` to create image using `Dockerfile`, 
    and tag it with `katsupplybot`
    - call `./run_docker_container.sh` to create and run docker container named `KatSupplyBot`
        - within that script db file `KatSupplyBot.db` is mapped to host filesystem (in current directory)
            so that the results of the work are accessible outside the container
    - p.s.
        - if `KatSupplyBot` container exists, it can be started again by calling `./start_docker_container.sh`
        - you can watch logs by calling `docker logs KatSupplyBot` 

### Current commands:
#### For all users:
- `/add` - Add supply request, for instance "/add купить лимон"
- `/list` - Show all current requests
- `/close` - Close request

#### Admins only:
- /shutdown - Shutdown bot instance