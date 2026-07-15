# CoreOps – Work Package Queue

**NDF-Basis:** `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`) — `main` ist **nicht** normativ.

Diese Queue ist die **aktive, verbindliche** Foundation-Queue für CoreOps (autoritativ gegenüber der historischen Concept-§50-Queue; siehe [FOUNDATION_SCOPE_LOCK.md](../docs/governance/FOUNDATION_SCOPE_LOCK.md)). Änderungen an der Queue dürfen ausschließlich durch ein später freigegebenes Work Package erfolgen. Es erfolgt **keine** automatische Freigabe eines Folge-Work-Packages. Genau `CO-WP-013` ist als nächster geplanter Schritt markiert (`CO-WP-005`…`CO-WP-012` umgesetzt und in Nova Review; die 004er-Erweiterungsserie ist mit `CO-WP-004E` abgeschlossen).

Jedes Work Package hat genau **einen** primären Typ.

| ID        | Type              | Title                                                          | Status                        |
| --------- | ----------------- | -------------------------------------------------------------- | ----------------------------- |
| CO-WP-001 | docs-only         | NDF Project Bootstrap – Core Governance Skeleton               | completed-go-with-notes         |
| CO-WP-001A| docs-only         | Complete NDF v1.0.0 Skills Bootstrap                           | completed-go                    |
| CO-WP-002 | docs-only         | Concept v3.0 Registration and Decision Classification          | completed-go-with-notes         |
| CO-WP-003 | docs-only         | Project Brief, Scope Lock and Release Taxonomy                 | completed-go-with-notes         |
| CO-WP-004 | docs-only         | Foundation Capability Matrix and Initial Support Boundary      | completed-go-with-notes         |
| CO-WP-004A| gov-baseline      | Sovereignty, BSI Orientation and Concept Amendment Registration| completed-go-with-notes         |
| CO-WP-004B| gov-baseline      | Lessons Learned and NDF Feedback Governance                    | completed-go-with-notes         |
| CO-WP-004B1| transfer-preparation | First NDF Feedback Transfer Package                        | completed-go-with-notes        |
| CO-WP-004B2| gov-correction    | Finalize NDF Intake Approval Gates                             | completed-go-with-notes        |
| CO-WP-004B3| cross-project-traceability | Record Completed NDF Intake Transfer                      | completed-go-with-notes        |
| CO-WP-004B4| cross-project-adoption-traceability | Record Completed NDF Adoption and Close Transfer Package 001 | completed-go-with-notes        |
| CO-WP-004C| docs-only         | BSI and Public-Sector Readiness Baseline                       | completed-go-with-notes        |
| CO-WP-004D| gov-baseline      | ITIL and PRINCE2 Applicability and Tailoring Decision          | completed-go-with-notes        |
| CO-WP-004E| docs-only         | Capability Matrix Security and Governance Alignment            | completed-go-with-notes        |
| CO-WP-005 | docs-only         | Language Standard, Public Neutrality and Repository Governance | completed-go-with-notes        |
| CO-WP-006 | docs-only         | System Context, Plane Taxonomy and External Boundaries         | completed-go-with-notes        |
| CO-WP-007 | security-baseline | Threat Model and Trust Boundaries                              | completed-go-with-notes        |
| CO-WP-008 | docs-only         | Architecture and Module Boundaries                             | completed-go-with-notes        |
| CO-WP-009 | security-baseline | Human Identity, Workspaces, RBAC and Break Glass               | completed-go-with-notes        |
| CO-WP-010 | security-baseline | Machine Identity, Enrollment and Offline Credential Lifecycle  | completed-go-with-notes        |
| CO-WP-011 | docs-only         | Source of Truth and Field Provenance                           | completed-go-with-notes        |
| CO-WP-012 | docs-only         | Observed, Desired, Effective State and Drift                   | completed-go-with-notes        |
| CO-WP-013 | security-baseline | Policy, Approval and Execution Authorization                   | completed-go-with-notes        |
| CO-WP-014 | docs-only         | CoreOps Integration Contract v0.1                              | completed-go-with-notes        |
| CO-WP-015 | docs-only         | Domain Pack Governance, Support Levels and Compatibility       | completed-go-with-notes        |
| CO-WP-016 | docs-only         | Data Ownership, Persistence, Schema Versioning and Migration   | completed-go-with-notes        |
| CO-WP-017 | docs-only         | API Governance, Versioning, Errors and Idempotency             | completed-go-with-notes        |
| CO-WP-018 | docs-only         | Event, Audit Correlation and Evidence Model                    | completed-go-with-notes        |
| CO-WP-019 | docs-only         | Telemetry and Normalization Schema                             | completed-go-with-notes        |
| CO-WP-020 | docs-only         | Topology Graph, Evidence and Manual Authority                  | completed-go-with-notes        |
| CO-WP-021 | docs-only         | Deployment Control Plane and Blueprint Schema                  | completed-go-with-notes        |
| CO-WP-022 | security-baseline | Artifact Trust, SBOM, Provenance and Revocation                | completed-go-with-notes        |
| CO-WP-023 | docs-only         | Restricted, Isolated, Air-Gapped Operation and CorePack        | completed-go-with-notes        |
| CO-WP-024 | security-baseline | Secrets, Configuration Vault and Key Custody                   | completed-go-with-notes        |
| CO-WP-025 | docs-only         | Data Classification, Retention and Redaction                   | completed-go-with-notes        |
| CO-WP-026 | security-baseline | Self-Protection, Degraded Modes and Recovery Mode              | completed-go-with-notes        |
| CO-WP-027 | docs-only         | UX Information Architecture and Dashboard System               | planned (retained, not immediately executable) |
| CO-WP-028 | docs-only         | Test Strategy, Fixtures and Integration Lab                    | planned                       |
| CO-WP-029 | review-only       | Cross-Document Consistency and ADR Candidate Review            | planned                       |
| CO-WP-030 | review-only       | Foundation Readiness Review                                    | planned                       |
| CO-WP-031 | release-prep      | Foundation 0.1 Release Preparation                             | planned                       |

> **Milestone-Hinweis 1 (nicht nummeriert):** `Milestone Lessons Review CO-WP-005 through CO-WP-012` — `completed-go-with-notes`. Gebündelter docs-only Review der acht Foundation-WPs; Ergebnis **GO WITH NOTES FOR CO-WP-013** (Commit 74f8e32). Keine `CO-WP-*`-ID; verändert die Queue-Reihenfolge nicht. Follow-ups (Capability-Count-Reconciliation, Risk-Register-Konsolidierung, Decision-Status-Harmonisierung) sind für ~`CO-WP-029`/`CO-WP-030` vorgemerkt.
> **Milestone-Hinweis 2 (nicht nummeriert):** `Milestone Lessons Review CO-WP-013 through CO-WP-020` — `completed-go-with-notes` (Commit d09d91b). Gebündelter docs-only Review der acht Foundation-WPs (24 Dokumente); Ergebnis **GO WITH NOTES FOR CO-WP-021**. Neues Dokument `project-brain/MILESTONE_REVIEW_CO_WP_013_TO_020.md`; Lessons LL-023…030; NDF-Kandidaten NDF-FC-COREOPS-011…013 (`candidate-pending-nova-review`). Keine `CO-WP-*`-ID; verändert die Queue-Reihenfolge nicht. `CO-WP-021` bleibt `planned-next`. Follow-ups (Risk-Konsolidierung, Decision-Status-Harmonisierung, Capability-Count, Dokumentationsökonomie) für ~`CO-WP-029`/`CO-WP-030`. Decision Index und Risk Register read-only unverändert.
> **Milestone-Vormerkung 3 (Planungshinweis, nicht nummeriert):** After completion of `CO-WP-026`, perform a bundled `CO-WP-021…026` Foundation Milestone Review — Status: `planned-next` · `not started`. Umfang: gebündelte Lessons Learned · NDF Candidates · Decision-/Risk-Wachstum · Foundation-Konsistenz · Reporting-/Vulnerability-Roadmap · nächste Phase. Keine `CO-WP-*`-ID; verändert die Queue-Reihenfolge nicht; Lessons-Learned-/NDF-Register bleiben jetzt unverändert.

## Zukünftige Reporting- und Vulnerability-Roadmap (roadmap-candidate, nicht nummeriert)

> **Status je Capability:** `roadmap-candidate` · `not scheduled` · `WP identifier pending queue review` · `not implemented`. **Keine neuen Work-Package-Nummern vergeben**, keine Technologie/API/Datenbank/Scanner/SBOM/PDF/Rendering ausgewählt, keine Decision/kein Risk registriert. Reihenfolge indikativ, nicht bindend; Einplanung erst durch spätere Queue-Review.

Geplante Capability-Chain (ohne WP-Nummern):

1. **Reporting Foundation**
2. **Asset and Component Inventory**
3. **Vulnerability Intelligence Ingestion**
4. **Vulnerability Correlation**
5. **Exposure and Remediation**
6. **Reporting Implementation**

**Reporting-Mindestscope (konzeptionell):** professionelle PDF-Berichte · CoreOps-Standarddesign · eigenes Logo · Corporate-Design- und Mandantenprofile · Template-/Branding-Version · Management-/Standard-/Detailberichte · Deutsch/Englisch · Inventar-/Log-/Update-/Deployment-/Topologie-/Audit-/Evidence-/Vulnerability-Berichte · Redaction vor Rendering ([CO-WP-025](../docs/security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md)) · getrennte Disclosure-/Export-Autorisierung ([CO-WP-025](../docs/governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md)) · keine Raw Secrets oder wiederverwendbaren Credentials ([CO-WP-024](../docs/security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md)).

**Vulnerability-Mindestscope (konzeptionell):** Inventarisierung von Hardware/Betriebssystem/Firmware/Software/Paketen/Diensten/Containern/Images/Bibliotheken/SBOM-Komponenten · Identitäten wie CPE/purl/Paketkoordinate/Firmware-ID/Image Digest/SBOM Component Identity · konzeptionelle Quellen CVE Records/NVD-CPE/Hersteller-Advisories/CISA KEV/EPSS/SBOM-VEX · Match Confidence · getrennte Zustände Candidate/Applicability/Confirmed/Not-Affected/Mitigated/Remediated/Verification-Pending/Outcome-Unknown · Produktnamensmatch ≠ bestätigte Betroffenheit · Offline Vulnerability-Intelligence-Snapshots/CorePacks ([CO-WP-023](../docs/architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md)) · Remediation über Deployment Governance ([CO-WP-021](../docs/architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md)) · Re-Inventarisierung und Verifikation · professionelle PDF-Vulnerability-Berichte. **Keine konkrete Technologie/API/Datenbank/Scanner-/SBOM-/PDF-/Rendering-Implementierung ausgewählt.**

## Regeln

- Genau ein primärer Typ pro Work Package.
- **Verbindlicher nächster Schritt ist der gebündelte `CO-WP-021…026` Foundation Milestone Review** (Milestone-Vormerkung 3; keine `CO-WP-*`-ID; `planned-next`; **noch nicht begonnen**). `CO-WP-026` ist nach **Nova Review `GO WITH NOTES`** und geschlossener Notes-Runde `completed-go-with-notes`; der Human-Maintainer-Commit steht noch aus (Push wird **nicht** als bereits erfolgt dargestellt). `CO-WP-025` (`3419664`), `CO-WP-024` (`916ba66`) und `CO-WP-023` (`b324aad`) sind committet, gepusht und `completed-go-with-notes`.
- Spätere Queue-Einträge (`CO-WP-027`…`CO-WP-031`) bleiben erhalten (**retained**), sind aber **nicht unmittelbar ausführbar**, solange der `CO-WP-021…026` Milestone Review nicht abgeschlossen ist. Kein reguläres Folge-WP (insbesondere nicht `CO-WP-027`) ist freigegeben; der Milestone Review selbst ist ebenfalls `not started`.
- Keine automatische Freigabe eines Folge-Work-Packages.
- Die Queue ist noch **nicht** final scope-locked.
- Änderungen an der Queue nur durch ein später freigegebenes Work Package.
