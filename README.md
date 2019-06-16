# Run

```
cd app

cd web 
yarn install
cd ..

go run main.go
```

http://localhost:8080/


# Build 
```
env GOOS=linux GOARCH=arm GOARM=5 go build -tags rpi -o wasserspender-rpi

go build
```

