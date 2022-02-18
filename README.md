# Simple API implementaion in Go

## To Get Started
#### Clone repo
Run the command to clone this github repository and change the directory to the project's folder:
```bash
git clone https://github.com/vsevdrob/api-go-gin-viper.git && cd api-go-gin-viper/src
```
#### Before run some commands, download some Go dependencies first.
This will download one of the coolest HTTP web framework written in Go (Golang).
```bash
go get github.com/gin-gonic/gin
```
Viper helps us to to operate with the predefined `config.yaml` file in order to export some required values from it.
```bash
go get github.com/spf13/viper
```
# Usage 
## Start server
Run the command that starts the server on host `127.0.0.1` and port `8080`.
```bash
go run main.go
```
After that open another terminal window and insure the path of current working directory is `*/api-go-gin-viper/src/`
___
## Add funder
Adds a funder to the DB. Assuming that *address* and *amount* keys/values in `single_funder.json` are predefined by Frontend.
```bash
curl localhost:8080/add-funder --header "Content-Type: application json" -d @single_funder.json --request "POST"
```
## Add funders
Adds a list of funders one by one to the DB. The same keys/values from above are also predefined in `numerous_funders.json`.
```bash
curl localhost:8080/add-funders --header "Content-Type: application json" -d @numerous_funders.json --request "POST"
```
## Fetch funder
Get a funder's data by assiging his/her **`id`** in the query.
```bash
curl localhost:8080/fetch-funder?id=1 --request "GET"
```
## Fetch funders
Get a list of each funder's data from the DB.
```bash
curl localhost:8080/fetch-funders --request "GET"
```
## Update funder
Update funder's funded total amount by increasing it. Assign the **`id`** in the query to update a specific funder. 
```bash
curl localhost:8080/update-funder?id=1 --request "PATCH"
```
## Delete funder
Delete a funder from DB. Assign funder's **`id`** in the query to delete the prefered one. 
```bash
curl localhost:8080/delete-funder?id=1 --request "DELETE"
```
# Licence
**MIT**
