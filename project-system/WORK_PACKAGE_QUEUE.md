# CoreOps – Work Package Queue

**NDF-Basis:** `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`) — `main` ist **nicht** normativ.

Diese Queue ist die **aktive, verbindliche** Foundation-Queue für CoreOps (autoritativ gegenüber der historischen Concept-§50-Queue; siehe [FOUNDATION_SCOPE_LOCK.md](../docs/governance/FOUNDATION_SCOPE_LOCK.md)). Änderungen an der Queue dürfen ausschließlich durch ein später freigegebenes Work Package erfolgen. Es erfolgt **keine** automatische Freigabe eines Folge-Work-Packages. Genau `CO-WP-004` ist als nächster geplanter Schritt markiert.

Jedes Work Package hat genau **einen** primären Typ.

| ID        | Type              | Title                                                          | Status                        |
| --------- | ----------------- | -------------------------------------------------------------- | ----------------------------- |
| CO-WP-001 | docs-only         | NDF Project Bootstrap – Core Governance Skeleton               | completed-go-with-notes         |
| CO-WP-001A| docs-only         | Complete NDF v1.0.0 Skills Bootstrap                           | completed-go                    |
| CO-WP-002 | docs-only         | Concept v3.0 Registration and Decision Classification          | completed-go-with-notes         |
| CO-WP-003 | docs-only         | Project Brief, Scope Lock and Release Taxonomy                 | implemented-awaiting-nova-review|
| CO-WP-004 | docs-only         | Foundation Capability Matrix and Initial Support Boundary      | planned-next                    |
| CO-WP-005 | docs-only         | Language Standard, Public Neutrality and Repository Governance | planned                       |
| CO-WP-006 | docs-only         | System Context, Plane Taxonomy and External Boundaries         | planned                       |
| CO-WP-007 | security-baseline | Threat Model and Trust Boundaries                              | planned                       |
| CO-WP-008 | docs-only         | Architecture and Module Boundaries                             | planned                       |
| CO-WP-009 | security-baseline | Human Identity, Workspaces, RBAC and Break Glass               | planned                       |
| CO-WP-010 | security-baseline | Machine Identity, Enrollment and Offline Credential Lifecycle  | planned                       |
| CO-WP-011 | docs-only         | Source of Truth and Field Provenance                           | planned                       |
| CO-WP-012 | docs-only         | Observed, Desired, Effective State and Drift                   | planned                       |
| CO-WP-013 | security-baseline | Policy, Approval and Execution Authorization                   | planned                       |
| CO-WP-014 | docs-only         | CoreOps Integration Contract v0.1                              | planned                       |
| CO-WP-015 | docs-only         | Domain Pack Governance, Support Levels and Compatibility       | planned                       |
| CO-WP-016 | docs-only         | Data Ownership, Persistence, Schema Versioning and Migration   | planned                       |
| CO-WP-017 | docs-only         | API Governance, Versioning, Errors and Idempotency             | planned                       |
| CO-WP-018 | docs-only         | Event, Audit Correlation and Evidence Model                    | planned                       |
| CO-WP-019 | docs-only         | Telemetry and Normalization Schema                             | planned                       |
| CO-WP-020 | docs-only         | Topology Graph, Evidence and Manual Authority                  | planned                       |
| CO-WP-021 | docs-only         | Deployment Control Plane and Blueprint Schema                  | planned                       |
| CO-WP-022 | security-baseline | Artifact Trust, SBOM, Provenance and Revocation                | planned                       |
| CO-WP-023 | docs-only         | Restricted, Isolated, Air-Gapped Operation and CorePack        | planned                       |
| CO-WP-024 | security-baseline | Secrets, Configuration Vault and Key Custody                   | planned                       |
| CO-WP-025 | docs-only         | Data Classification, Retention and Redaction                   | planned                       |
| CO-WP-026 | security-baseline | Self-Protection, Degraded Modes and Recovery Mode              | planned                       |
| CO-WP-027 | docs-only         | UX Information Architecture and Dashboard System               | planned                       |
| CO-WP-028 | docs-only         | Test Strategy, Fixtures and Integration Lab                    | planned                       |
| CO-WP-029 | review-only       | Cross-Document Consistency and ADR Candidate Review            | planned                       |
| CO-WP-030 | review-only       | Foundation Readiness Review                                    | planned                       |
| CO-WP-031 | release-prep      | Foundation 0.1 Release Preparation                             | planned                       |

## Regeln

- Genau ein primärer Typ pro Work Package.
- Nur `CO-WP-004` ist der nächste geplante Schritt (`planned-next`).
- Keine automatische Freigabe eines Folge-Work-Packages.
- Die Queue ist noch **nicht** final scope-locked.
- Änderungen an der Queue nur durch ein später freigegebenes Work Package.
