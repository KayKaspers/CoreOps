# CoreOps – Public Neutrality and Disclosure Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation public-neutrality policy
> Public Repository Scope: Applicable
> Certification or Endorsement Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-005` (docs-only / repository-governance foundation)

## 1. Status

Foundation policy for public neutrality and disclosure across the CoreOps public repository. It consolidates and extends the neutrality expectations already referenced in the NDF feedback and BSI documents without duplicating their process-specific rules.

## 2. Purpose

Keep CoreOps public artifacts organisation-neutral, vendor-neutral, and free of unnecessary private or sensitive data, while still allowing necessary mention of vendors, products, and frameworks for capability, integration, reference, example, risk, or roadmap purposes.

## 3. Scope

All public repository content: documentation, governance, examples, diagrams, and any exported artifact. Applies alongside the [Language Standard](COREOPS_LANGUAGE_STANDARD.md) and [Repository Governance Standard](REPOSITORY_GOVERNANCE_STANDARD.md).

## 4. Public Neutrality

CoreOps is vendor-neutral, deployment-neutral, organisation-neutral, publicly reusable, and not tied to a specific private environment. Neutrality does **not** forbid naming vendors, platforms, or frameworks where required for capability description, integration, compatibility, reference, example, risk, or roadmap.

## 5. Vendor Neutrality

```text
Vendor mention        ≠ endorsement
Integration target    ≠ mandatory dependency
Roadmap entry         ≠ implemented support
Implemented connector ≠ certified interoperability
Example product       ≠ preferred product
```

Multiple vendors should be treated in a balanced way where several are relevant; no artificial parity is forced when only one vendor is technically relevant.

## 6. Organisation Neutrality

Public documents must not contain unnecessary data about: private persons; private organisations; customers; authorities; associations; internal networks; private domains; private email addresses; personal usernames; local absolute paths; internal hosts; private repositories; credentials; non-public infrastructure. Project authorship and public maintainer information may be named where deliberate and necessary.

## 7. Institutional and Public-Sector Boundary

Public-sector or Government profiles must not create the impression of official authority approval, political support, institutional partnership, ministerial recommendation, agency certification, military clearance, or suitability for classified information. CoreOps may support public or regulated deployment scenarios without asserting institutional affiliation.

## 8. Personal Data

No personal data (names, personal emails, usernames, personal identifiers) is published unless deliberately required (e.g. public maintainer attribution). Print/job, user, and identity examples use synthetic data.

## 9. Infrastructure Data

No private infrastructure details: real internal IP addresses where unnecessary, private domains, internal hostnames, private topology diagrams, internal repository URLs, or non-public infrastructure layouts.

## 10. Secrets and Credentials

Never publish: passwords; tokens; API keys; private keys; session secrets; real production credentials; private support cases; unredacted logs. Secrets are never committed, even in examples. A secret name may be referenced; its value never.

## 11. Logs and Screenshots

Log excerpts, screenshots, config examples, support reports, incident documents, topology images, and export files must be redacted before publication. Redaction must not distort the technical meaning.

## 12. Test and Example Data

Prefer: `example.com` / `example.org` / `example.net`; `localhost`; `127.0.0.1`; RFC-reserved documentation addresses; generic usernames/hostnames; synthetic UUIDs; obviously fake tokens. Do not use realistic-looking secrets. Example:

```text
token_example_not_a_real_secret
```

rather than a realistically formatted production-like token, unless a format check must be demonstrated (then still clearly synthetic).

## 13. Redaction

Minimum redaction rules for log excerpts, screenshots, config examples, support bundles, incident documents, topology images, and export files: remove secrets, personal data, private paths, and private infrastructure; keep the technical sense intact; mark redacted regions clearly.

## 14. Claims and Endorsement

No endorsement, partnership, accreditation, certification, or preferred-treatment claim arises from a vendor mention, integration, roadmap entry, or profile name. Certification and compliance claim boundaries follow the BSI reference/claims register and the ITIL/PRINCE2 tailoring claim boundaries.

## 15. External Sources

External sources must be identifiable; version or retrieval state is documented for time-dependent sources; incomplete sources are not treated as complete; no extensive copyrighted content is reproduced; no proprietary framework text is reconstructed; missing source ingestion yields `verification-pending`, not invented detail.

## 16. Exceptions

Any deliberate disclosure (e.g. public maintainer contact) requires a documented rationale and Human-Maintainer approval. No silent exceptions; no exception may publish a secret.

## 17. Human-Maintainer Gate

Publication, disclosure decisions, and any exception require Human-Maintainer approval. AI systems prepare and review but never publish or commit.

## 18. Validation Checklist

```text
No secrets · No private paths · No personal data · No private infrastructure ·
No customer data · No unsupported claims · No certification/endorsement overclaim ·
No misleading translation-parity claim · No unredacted logs · No accidental private data
```

## 19. Compatibility

Additive; consistent with existing NDF feedback/BSI neutrality references and the `ndf-public-neutrality-guard` intent. No accepted decision changed; no technology selected; no ADR. Breaking-change potential: low.

## 20. Open Questions

- Which public maintainer attribution (if any) is deliberately included?
- Whether an automated public-hygiene check is added later (follow-up, not this WP).
- Which example-data conventions become mandatory vs. recommended.
