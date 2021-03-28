@ECHO OFF

SET comp_flags=-trimpath -ldflags "-s -w"
SET test_flags=-test.v

SET src_dir=.\src
SET app_fname=build\pw.exe
SET test_fname=build\test.exe

go test -c -o %test_fname% %src_dir%
%test_fname% %test_flags%

IF %ERRORLEVEL% EQU 0 (
    go build %comp_flags% -o %app_fname% %src_dir%

    IF "%~1"=="clean" (
        del %test_fname%
    )
)
