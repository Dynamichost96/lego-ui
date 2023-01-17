# Group Certificates per DNS Provider

## Problem

As input forms differ between different DNS Providers, makeing a concrete difference between different DNS provides makes sense, as it is clearer to the user, where they can configure DNS provider specific settings. However grouping certificates by DNS provider makes it more difficult to find the certificates the user is looking for.

## Constraints

1. There are no concrete differences between certificates, regardless of the chosen DNS provider
2. There are config paramteres per certificate, per provider and globally

## Assumptions

1. A user must be able to identify config they need quickly by provider, but doesn't care about providers when manageing certificates

## Solutions

### Option A

The Certificates are grouped by provider. We only show the providers that are in use, unless we create a new certificate. The provider specific configuration is placed on the "main view" of the provider and a list of certificates 

### Option B

Certificates are all listed without grouping them per provider. The provider specific config is synced across all certificates of the provider and displayed inline to the certificate specific config

### Option C

Certificates are all listed without grouping them per provider. Provider configuration can be done through the settings page

## Decision

We decided on Option B

## Rationale

This option is the least implementation effort, and makes it simple for a user, when creating a new certificate.

## Implications

- Sync some fields of the certificate form to keep provider config consistent
- Update the specification to integrate additional DNS providers