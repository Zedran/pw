@ECHO OFF

SET comp_flags=-trimpath -ldflags "-s -w"
SET test_flags=-test.v

SET src_dir=.\src
SET app_fname=build\pw.exe
SET test_fname=test.exe

SET res_origin="res"
SET res_destin="build/res"

go test -c -o %test_fname% %src_dir%
%test_fname% %test_flags%

IF %ERRORLEVEL% EQU 0 (
    go build %comp_flags% -o %app_fname% %src_dir%

    xcopy /e /i /s /y %res_origin% %res_destin%
    del %test_fname%
)
