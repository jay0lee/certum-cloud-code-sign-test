# Certum Cloud Code Signing Test & Verification Repository

This repository tests and validates the reusable [`jay0lee/certum-cloud-code-sign`](https://github.com/jay0lee/certum-cloud-code-sign) GitHub Action across **Windows x86_64** (`windows-latest`) and **Windows Arm** (`windows-11-arm`) runners.

---

## Verification Pipeline

The automated workflow in [`.github/workflows/verify-signing.yml`](.github/workflows/verify-signing.yml) runs through the following 4 stages:

1. **Executable Generation**:
   - Generates a small Windows executable (`helloworld.exe`) using pre-existing GitHub Actions tools (Go with automatic fallback to pre-installed C# `csc.exe`).
   - Verifies the binary runs prior to signing.

2. **Certum Action Execution**:
   - Invokes `GAM-team/certum-cloud-code-sign@main`.
   - Downloads and silently installs SimplySign Desktop.
   - Automatically bypasses ARM64 OOBE privacy screens if on ARM64.
   - Calculates the dynamic RFC 6238 HMAC-SHA256 TOTP token at the last second with expiry protection.
   - Automates GUI typing into SimplySign Desktop without logging credentials or OTP.
   - Waits for the certificate and private key to load into `Cert:\CurrentUser\My`.
   - Locates the x64 `signtool.exe` and exports action outputs.

3. **Code Signing**:
   - Signs `helloworld.exe` using `signtool.exe` with the certificate thumbprint output from the action:
     ```powershell
     signtool sign /sha1 ${{ steps.certum.outputs.cert-thumbprint }} `
       /tr http://time.certum.pl /td SHA256 /fd SHA256 /v `
       helloworld.exe
     ```

4. **Comprehensive Signature Verification**:
   - Verifies the Authenticode signature using `signtool verify /pa /v helloworld.exe`.
   - Verifies the signature status, signer certificate subject, issuer, and timestamp using PowerShell's `Get-AuthenticodeSignature`.
   - Asserts that the signer certificate thumbprint matches the expected certificate output.
   - Executes the signed binary to ensure it was not corrupted by the signing process.

---

## Required Secrets

To enable end-to-end verification, set the following GitHub Secrets in this repository (**Settings > Secrets and variables > Actions**):

> [!IMPORTANT]
> GitHub Secret names **can only contain alphanumeric characters (`[A-Za-z0-9]`) and underscores (`_`)**. Hyphens/dashes (`-`) are not allowed by GitHub.

| Secret Name | Supported Fallback | Description |
| :--- | :--- | :--- |
| `CERTUM_USERNAME` | `USERNAME` | Your Certum SimplySign account email / username |
| `CERTUM_TOTP_SECRET` | `TOTP_SECRET` | Your Base32 TOTP secret key for 2FA one-time password generation |
| `CERTUM_CERT_SHA1` | `CERT_SHA1` | *(Optional)* Expected certificate SHA-1 thumbprint |

---

## License

Apache License 2.0. See [LICENSE](LICENSE).
