$dist = "..\public\cgi-bin\index.cgi"
$target = "main.go"
if($Args.Length -eq 0){
    $env:GOOS="linux";
    $env:GOARCH="amd64";
    go generate;
    go build -o $dist $target;
}elseif($Args.Length -eq 1 -and $Args[0] -eq "clean"){
    Remove-Item $dist;
}
