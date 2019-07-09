go fmt  ./...
go test ./... 

@Choice /C YN /D Y /T 7 /M ">>>>>>>>>> Continue?"
@If ErrorLevel = 2 goto done

go vet -composites=false ./...

@Choice /C YN /D Y /T 7 /M ">>>>>>>>>> Continue?"
@If ErrorLevel = 2 goto done

@r.cmd

:done
