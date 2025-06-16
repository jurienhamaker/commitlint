!include "LogicLib.nsh"
;--------------------------------
;General information

;The name of the installer
Name "Commitlint installer"

;The output file path of the installer to be created
OutFile "commitlint_installer.exe"

;The default installation directory
InstallDir "$PROGRAMFILES\Commitlint"

;Request application privileges for user level privileges
RequestExecutionLevel admin


;--------------------------------
;Installer pages

;Show a page where the user can customize the install directory
Page directory
;Show a page where the progress of the install is listed
Page instfiles


;--------------------------------
;Installer Components

;A section for each component that should be installed
Section "Install commitlint"
	;Set output path to the installation directory
	SetOutPath $INSTDIR
	
	SetOverwrite on
	File /r "./unzipped/*"
	
	WriteUninstaller $INSTDIR\uninstaller.exe
SectionEnd

Section "Add to path"
	EnVar::SetHKCU
	EnVar::Check "Path" "$InstDir"
	Pop $0
	${If} $0 = 0
		DetailPrint "Already there"
	${Else}
		EnVar::AddValue "Path" "$InstDir"
		Pop $0 ; 0 on success
	${EndIf}
SectionEnd

Section "uninstall"
	EnVar::DeleteValue "Path" "$InstDir"
	Pop $0

	RMDir /r /REBOOTOK "$InstDir"
SectionEnd
