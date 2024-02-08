@echo off

rem Définit le nom de l'exécutable
set EXECUTABLE=leetspeakgo

rem Construit l'exécutable pour Linux
SET GOOS=linux
go build -o %EXECUTABLE%_linux

rem Construit l'exécutable pour macOS
SET GOOS=darwin
go build -o %EXECUTABLE%_macos

rem Construit l'exécutable pour Windows
SET GOOS=
go build -o %EXECUTABLE%.exe

