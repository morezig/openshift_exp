# Demo

Dependencies:
=============================================
go get -u github.com/kardianos/govendor  <br />
go get github.com/jteeuwen/go-bindata/... <br />
go get github.com/elazarl/go-bindata-assetfs/... <br />
govendor fetch +m

RUN:
=============================================
cd ui/backend <br />
npm install  <br />
npm run build <br />
<br />
會把ui/backend下的內容build到api/core/assets下<br /><br />

cd ../../api/core <br />
go-bindata-assetfs -pkg common -o common/bindata.go assets/... <br />
將api/core/assets下的內容打包成go file <br />
Local端直接執行: go run main.go  <br />
建立執行檔 : go build -o app <br />
./app <br />
