$LogFilePath = "C:\Users\mhmta\OneDrive\Belgeler\Github\Dataguess\mid\scripts\LogFile.txt"
$EndpointUrl = "http://localhost:8080/v1/SmsQueue/TriggerWorker"
$Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjI3LCJuYW1lIjoiTXVyYXQiLCJzdXJuYW1lIjoiWWFsxLFuYXlhayIsInBob25lIjoiKzkwNTM0NTg4OTgwOCIsImVtYWlsIjoibXVyYXR5YWzEsW5heWFrQG1hcy5jb20iLCJzZXNzaW9uVXVpZCI6IjEyMTdjNDY3LTVmYzEtNGVhMy1iNWQ4LTc3ODUzNWViNzIwOCIsImV4cCI6MTcwMzc5NTY2OH0.BAXLKtqO7S29E54QPd1dfBIJi0sgWZu7JDPVpQVI6Wk"

# Fonksiyon: Log Dosyasına Yaz
function Write-ToLogFile {
    param (
        [string]$LogMessage
    )

    # Tarih ve saat bilgisini ekleyerek log mesajını oluştur
    $LogEntry = "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss') - $LogMessage"

    # Log dosyasına yaz
    Add-Content -Path $LogFilePath -Value $LogEntry
}

# Log dosyasını oluştur
New-Item -Path $LogFilePath -ItemType File -Force

try {
    # HTTP başlık eklemek için kullanılan özel bir hashmap
    $Headers = @{
        "Authorization" = "Bearer $Token"
    }

    # HTTP isteği gönder
    $Response = Invoke-RestMethod -Uri $EndpointUrl -Method Post -Headers $Headers

    # Yanıtın parçalarını al
    $StatusCode = $Response.StatusCode
    $ResponseContent = $Response.Content
    $HandledSmsCount = $Response.handledSmsCount

    # Log dosyasına yaz
    $LogMessage = "Status Code: $StatusCode, Response Content: $ResponseContent, Handled SMS Count: $HandledSmsCount"
    Write-ToLogFile -LogMessage $LogMessage

    Write-Host "Status Code: $StatusCode"
    Write-Host "Response Content: $ResponseContent"
    Write-Host "Handled SMS Count: $HandledSmsCount"

} catch {
    # Hata mesajını al
    $ErrorMessage = $_.Exception.Message

    # Log dosyasına hata mesajını yaz
    $LogMessage = "Hata Oluştu: $ErrorMessage"
    Write-ToLogFile -LogMessage $LogMessage

    Write-Host "Hata Oluştu: $ErrorMessage"
}
