$dist = "..\public\cgi-bin\index.cgi"

if($Args.Length -eq 0){
    $env:GOOS="linux";
    $env:GOARCH="amd64";
    go build -o $dist main.go;
}elseif($Args.Length -eq 1 -and $Args[0] -eq "clean"){
    Remove-Item $dist;
}
