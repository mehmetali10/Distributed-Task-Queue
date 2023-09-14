# tour log file path here
$LogFilePath = "...\Distributed-Task-Queue\scripts\LogFile.txt"
$EndpointUrl = "http://localhost:8080/v1/SmsQueue/TriggerWorker"
$Token = "Your Jwt token here"

function Write-ToLogFile {
    param (
        [string]$LogMessage
    )

    $LogEntry = "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss') - $LogMessage"

    Add-Content -Path $LogFilePath -Value $LogEntry
}

New-Item -Path $LogFilePath -ItemType File -Force

try {
    $Headers = @{
        "Authorization" = "Bearer $Token"
    }

    $Response = Invoke-RestMethod -Uri $EndpointUrl -Method Post -Headers $Headers

    $StatusCode = $Response.StatusCode
    $ResponseContent = $Response.Content
    $HandledSmsCount = $Response.handledSmsCount

    $LogMessage = "Status Code: $StatusCode, Response Content: $ResponseContent, Handled SMS Count: $HandledSmsCount"
    Write-ToLogFile -LogMessage $LogMessage

    Write-Host "Status Code: $StatusCode"
    Write-Host "Response Content: $ResponseContent"
    Write-Host "Handled SMS Count: $HandledSmsCount"

} catch {
    $ErrorMessage = $_.Exception.Message

    $LogMessage = "Error Occured: $ErrorMessage"
    Write-ToLogFile -LogMessage $LogMessage

    Write-Host "Error Occured: $ErrorMessage"
}
