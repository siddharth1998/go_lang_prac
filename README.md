

## To run and test your code

<details><summary>TCP Proxy </summary>
For server side

```console
go run $(pwd)/tcpProxy/tcpProxy.go
```

Create a simple http.server from python, where you have some files this will act like 

``` console
python -m http.server 10000
```

Curl command
```console
curl -i -X GET 127.0.0.1:80
```
</details>
 

 <details>
 <summary> Running the TCP Client to query the shodan</summary>
Run the following command to setup the shodan api key which you will get from the shodan account section
 ```console
 set SHODAND_API_KEY=APIKEY
 ```
 Create a mod.go if required ( i understand githublink should be added here wanted to test local driven development)<br>
if you see the mod file in the folder shodan is a virtual link /namespace which is not there.  
```console
go mod init shodan
```
Now just run the main.go in the cmd shodan( please remember the parent directory of main.go is shodan,but this is not refered in the go.mod)
```console
go run main.go
```
 </details>