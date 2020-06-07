@ECHO OFF

SET output_name="pw.exe"
SET src_dir=./src

go test %src_dir% -test.v

IF %ERRORLEVEL% EQU 0 go build -o %output_name% -trimpath -ldflags "-s -w" %src_dir%
