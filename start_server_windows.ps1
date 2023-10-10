#!/bin/pwsh
while ($true) {
    Write-Output ("`n`nRunning Containers :")
    docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}"
    Write-Output "`nMenu :`n
            1) Run Normal`n
            2) First Time Run (Will Reset Database if exist)`n
            3) Run with PhpMyAdmin`n
            4) First Time Run with PhpMyAdmin (Will Reset Database if exist)`n
            5) Run in Debug Mode (TBA)`n
            6) Kill all Open Containers (TBA)`n
            7) CURL Libraies (Endpoints with Curl) (TBA)`n
            0) Exit Script`n`n"

    $option = Read-Host -Prompt 'Enter your Option: '

    if ($option -eq "0") {
        Write-Host -NoNewLine 'Press any key to continue...';
        $null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
        Clear-Host
        break
    }elseif ($option -eq "1") {
        Copy-Item ./config/compose-enviroments/normal-run/docker-compose.yml ./
        docker-compose up -d
        Write-Output "`nThe Website will be avaible on http://local.techwall.xyz"
        Write-Output "Give the container about 30 seconds to get Ready"
        Write-Host -NoNewLine 'Press any key to continue...';
        $null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
        Clear-Host
        Write-Output "`nThe Website is avaible on http://local.techwall.xyz`n"
    }elseif ($option -eq "2") {
        Copy-Item ./config/compose-enviroments/first-run/docker-compose.yml ./
        docker-compose up -d
        Write-Output "`nThe Website will be avaible on http://local.techwall.xyz"
        Write-Output "Give the container about 30 seconds to get Ready"
        Write-Host -NoNewLine 'Press any key to continue...';
        $null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
        Clear-Host
        Write-Output "`nThe Website is avaible on http://local.techwall.xyz`n"
    }elseif ($option -eq "3") {
        Copy-Item ./config/compose-enviroments/normal-php/docker-compose.yml ./
        docker-compose up -d
        Write-Output "`nThe Website will be avaible on http://local.techwall.xyz"
        Write-Output "The PhpMyAdmin page will be avaible on http://localhost:8080"
        Write-Output "Give the container about 30 seconds to get Ready"
        Write-Host -NoNewLine 'Press any key to continue...';
        $null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
        Clear-Host
        Write-Output "`nThe Website is avaible on http://local.techwall.xyz"
        Write-Output "The PhpMyAdmin page is avaible on http://localhost:8080`n"

    }elseif ($option -eq "4") {
        Copy-Item ./config/compose-enviroments/first-php/docker-compose.yml ./
        docker-compose up -d
        Write-Output "`nThe Website will be avaible on http://local.techwall.xyz"
        Write-Output "The PhpMyAdmin page will be avaible on http://localhost:8080"
        Write-Output "Give the container about 30 seconds to get Ready"
        Write-Host -NoNewLine 'Press any key to continue...';
        $null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
        Clear-Host
        Write-Output "`nThe Website is avaible on http://local.techwall.xyz"
        Write-Output "The PhpMyAdmin page is avaible on http://localhost:8080`n"
    }elseif ($option -eq "5") {
        Write-Output "Feature Coming Soon"
        Write-Host -NoNewLine 'Press any key to continue...';
        $null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
        Clear-Host
    }elseif ($option -eq "6") {
        docker-compose down -v
        Remove-Item ./docker-compose.yml
        Write-Host -NoNewLine 'Press any key to continue...';
        $null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
        Clear-Host
    }elseif ($option -eq "7") {
        Write-Output "Feature Coming Soon"
        Write-Host -NoNewLine 'Press any key to continue...';
        $null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
        Clear-Host
    }

}