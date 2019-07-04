go fmt  ./...
go test ./... 
go vet -composites=false ./...
@pause

set flags=-t

@cd .\cmd\NQueens\
go build .
call NQueens	%flags%
@cd .\..\..\

@cd .\cmd\NQueensR\
go build .
call NQueensR	%flags%
@cd .\..\..\

@cd .\cmd\Sudoku\
go build .
call Sudoku	%flags%
@cd .\..\..\

@pause
