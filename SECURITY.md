# Security Policy

## Reporting a vulnerability

Please do not open a public issue for a security problem. Use GitHub's private
[security advisory](https://github.com/tiqhmibisimax/GroveWatch/security/advisories/new) form, or contact a
maintainer directly if the form is unavailable to you.

Include what an attacker gains in one sentence, the version or commit you are running, the smallest
reproduction you can manage, and whether the input is attacker-controlled in your setup. You should
get an acknowledgement within a few working days and a disclosure timeline before anything is
published.

## Scope

GroveWatch fingerprints a workspace into a snapshot other tools trust. In scope: a
change that verifies as clean (hash collision, canonicalisation gap, a path that escapes
the snapshot root), a malformed snapshot that crashes the verifier instead of being
rejected, and anything that makes 'what drifted' disagree with what actually changed.
