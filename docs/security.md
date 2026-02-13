# Security Notes

OPMLWatch should be deployed with these defaults:

- Request timeout and bounded retries for outbound requests
- Respect cache headers (`ETag`, `If-Modified-Since`) once fetcher is enabled
- Avoid internal network access when user-provided feed URLs are allowed (SSRF prevention)
- Sanitize HTML when rendering feed content in any web surface
- Keep secrets out of config files tracked by git

## Reporting

Please report security issues privately before public disclosure:
- Open a private security advisory in GitHub
- Or contact maintainers through repository security contact
