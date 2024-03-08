## DataVizFlow

# Setup Instructions
1. Install redis
`$ brew install redis`  (for mac)

2. Run redis
`$ redis-server`

3. Run go application
```
go run main.go
Listening on port 8080...
```

Head over to the browser and open following:
`http://localhost:8080/datasets`


4. Go to vue-appn folder and run the following:
`npm run dev`

Open following link in the browser:
`http://localhost:5173/`


![alt text](<Screenshot 2024-03-08 at 5.20.37 PM.png>)


# Architecture Overview
This is simple application which has golang service which reads from static dataset.
Redis is used for caching. Data is stored in redis cache for fast retrieval and it is fetched from cache directly if present, otherwise it is first stored in the cache and then returned as response.

The vue application running listens on the apiUrl continuously and displays the data in a bar graph format.

# References:
https://vuejs.org/guide/quick-start.html
https://www.chartjs.org/docs/latest/getting-started/usage.html
https://vue-chartjs.org/guide/#installation