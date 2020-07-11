@ECHO OFF

SET flags=-trimpath -ldflags "-s -w"
SET output_name="pw.exe"
SET src_dir=./src

go test %src_dir% -test.v

IF %ERRORLEVEL% EQU 0 (
    go build %flags% -o %output_name% %src_dir%
)
