# update.ps1
try {
    Write-Host "Iniciando atualização do TermOTP..."

    $installPath = "C:\Program Files\TermOTP\totp.exe"
    $backupPath = "C:\Program Files\TermOTP\totp.bak"
    $newBinary = "C:\Program Files\TermOTP\totp_new.exe"
    $downloadUrl = "https://github.com/JVitoroliv3ira/termotp/releases/latest/download/totp-windows-amd64.exe"

    if (Test-Path $installPath) {
        Copy-Item $installPath $backupPath -Force
        Write-Host "Backup criado em $backupPath"
    }

    Invoke-WebRequest -Uri $downloadUrl -OutFile $newBinary
    Write-Host "Novo binário baixado para $newBinary"

    Remove-Item $installPath -ErrorAction SilentlyContinue
    Move-Item $newBinary $installPath -Force
    Write-Host "Novo binário movido para $installPath"

    if (Test-Path $backupPath) {
        Remove-Item $backupPath -Force
        Write-Host "Backup removido."
    }

    Write-Host "Atualização concluída com sucesso!"
}
catch {
    Write-Host "Ocorreu um erro: $_"
    Write-Host "Restaurando backup..."
    if (Test-Path $backupPath) {
        Move-Item $backupPath $installPath -Force
        Write-Host "Rollback realizado: binário restaurado."
    }
}
