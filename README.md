# dataVizFlow
To build "DataVizFlow," a web application that reads datasets from a source, caches processed data in Redis, and displays the data on interactive graphs using Vue.js, you'll need to follow several steps. Below, I'll outline the general architecture and the steps to implement it.

step 1. start radis on local machine
    cd redis-dataset-gen
    python .\redis_data_gen.py

step 2. Run  Go server
    cd server
    go run .

step 3 run client server
    npm run serve 

