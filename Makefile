

build: 
	(cd ../ && GOOS=js GOARCH=wasm go build -o wasm/main.wasm)

start-server: 
	 goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
