# CoreOps – Context Pack: Foundation 0.1

Kompakter, wiederverwendbarer Kontext für die Foundation-Phase. Enthält keine Rohlogs, keine Chain-of-Thought, keine Secrets, keine privaten Daten und keine vollständige Kopie des Concept v3.0.

## Projekt und Phase

- **Projekt:** CoreOps — One Dashboard. Controlled Operations.
- **Phase:** `Foundation 0.1` (interner Phasenname; Release-Taxonomie: Foundation `v0.0.1-foundation`, Observe `v0.1.0-alpha.1`, beide `Proposed for acceptance`).

## NDF-Basis

- Nova-Development-Framework `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`).
- Branch `main` ist informativ, **nicht** normativ; Übernahme aus `main` nur via eigenes freigegebenes WP (CCR-04 geklärt).

## Akzeptierte Produktvision

Universelle, self-hosted und offline-fähige Operations Control Plane, die Source of Truth, Monitoring, Topologie, Desired State, Policy, vertrauenswürdige Automatisierung und nachvollziehbare Änderungen modular verbindet. Herkunft: CoreOps Concept v3.0 (bereitgestellt 12. Juli 2026). Technische Detailentscheidungen sind unbestätigte Foundation-Kandidaten.

## Aktueller Scope

- **In Scope:** Foundation-Dokumentation, Sicherheitsgrundlagen, Architekturmodelle, Capability-/Supportgrenzen, ADR-Vorbereitung, Teststrategie, Readiness Review.
- **Out of Scope:** Anwendungscode, produktive Integrationen, Agents, Deployments, Scans, Netzwerkänderungen, Druckerverwaltung, Workflow-Ausführung, Secrets-Verarbeitung, produktive Installation.

## Letztes Work Package

`CO-WP-024 – Secrets, Configuration Vault and Key Custody` (docs-only / security, configuration, secret-lifecycle and key-custody governance foundation) — **Nova Review `GO WITH NOTES`; Nova-Notes 1–4 vor dem Commit geschlossen; `completed-go-with-notes`; Human-Maintainer-Commit ausstehend**. Vorheriges WP: `CO-WP-023 – completed-go-with-notes` (Commit b324aad, gepusht).

## Secrets, Configuration Vault and Key Custody (CO-WP-024)

- Drei neue Dokumente: Secrets/Configuration Vault and Custody Governance, Key Material/Rotation/Revocation/Recovery Policy, Configuration Source of Truth and Secret Reference Model. Baut auf MOD-SEC-001/MOD-STA-001, CO-WP-010, CO-WP-009 auf; kein Parallelmodell.
- `secret ≠ ordinary configuration`; `secret reference ≠ secret value`; `credential possession ≠ authorization`; `retrieval ≠ use ≠ export`; `key identity ≠ key material ≠ custody`; `custody ≠ use authority`; `vault availability ≠ secret trust`. Vault = logische Governance-Grenze (kein Produkt); 15 Secret-/Config-Klassen; getrennte Autoritäten (owner ≠ custodian ≠ user; rotation ≠ reinstatement; backup operator ≠ recovery authority; machine principal ≠ human approval).
- Secret Lifecycle (proposed…destroyed/outcome-unknown); Secret Identity/Version/Value-Revision/Instances getrennt; Retrieval/Distribution/Injection/Use getrennt (`injected ≠ consumed`). Rotation vollständig erst nach Consumer-Assessment; Revocation mit Identity Binding + Freshness (`no local entry ≠ not revoked`); Bootstrap/Root/Recovery/Break-Glass bounded/one-time (keine dauerhafte Autorität); Recovery ≠ Reinstatement; Rollback reaktiviert kein revoked/expired/compromised Material.
- Configuration Source of Truth ≠ Runtime State; Secret References fail-closed (kein Fallback); Drift ≠ auto-remediated. CorePack Trust ≠ Secret Trust; Deployment Authorization ≠ Secret-Use Authorization; Offline erweitert keine Autorität (Offline-Authorization nicht wiederverwendbar); Audit/Evidence ohne Secret Values. Profile (Standard/Hardened/Government) = Kontrollstärke, keine Compliance/Zertifizierung.
- Decision Index +14 (DEC-S-318…331; nach Nova-GO-WITH-NOTES-Deduplizierung, ursprünglich 16), Risk Register +5 (RISK-295…299, gesamt 299; Wachstumsgrenze §36 eingehalten). Keine Vault-/KMS-/HSM-/TPM-/PKI-/Key-/Secret-/Config-/Rotation-/Backup-/Sync-Technologie; Rohsecret-/Rohschlüsselspeicherung nicht entschieden; keine ADR; bestehende Identity-/Authorization-/Trust-/Deployment-/Artifact-/Evidence-/Offline-/CorePack-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Restricted, Isolated, Air-Gapped Operation and CorePack (CO-WP-023)

- Drei neue Dokumente: Restricted/Isolated/Air-Gapped Operation Model, CorePack Identity/Content/Lifecycle Model, Offline Trust/Activation/Revocation/Transfer Policy.
- Connectivity Classes getrennt (`connected/restricted-connected/intermittently-connected/isolated/air-gapped/recovery-only`); `offline ≠ air-gapped`; `isolated ≠ trusted`; `network unavailable ≠ security controls optional`; `central unavailable ≠ local authority expands`; `local activation authority ≠ deployment execution authorization`; keine Klassifiziertnetz-Eignung.
- `CorePack ≠ Domain Pack/artifact/blueprint/backup/evidence package`; Identity/Version/Revision/Assembly-/Transfer-/Import-/Activation-Instance getrennt; `transferred ≠ imported ≠ trusted ≠ approved-for-activation ≠ activated ≠ deployment authorised`; Content Resolution bindet konkrete Revision (`mutable alias ≠ final binding`); Contents erben Trust/Compatibility nicht vom Container; Target Binding durch Transfer/Import/Activation erhalten.
- `quarantine release ≠ activation authorization`; Partial/Unknown Activation sichtbar (`no automatic retry → reconciliation`); Delta erfordert bestätigte exakte Baseline; Delayed Revocation sichtbar (`no local revocation entry ≠ not revoked centrally`); Offline Authorization action-/target-/scope-/version-/purpose-/time-bound und nicht wiederverwendbar; Rollback braucht aktuelle Trust-/Revocation-/Compatibility-Bewertung; Evidence Return (`returned ≠ complete ≠ sufficient`); Export-Autorität getrennt; keine Rohsecrets.
- Decision Index +14 (DEC-S-304…317), Risk Register +5 (RISK-290…294, gesamt 294; Wachstumsgrenze §4 eingehalten). Keine Package-/Manifest-/Installer-/Update-/Transfer-/Hash-/Signing-/Trust-Anchor-/PKI-/Encryption-/Reconciliation-Technologie und keine Offline-Runtime; keine ADR; Artifact-/Deployment-/Domain-Pack-/Integration-/Data-/API-/Event-/Evidence-/Telemetry-/Topology-/Identity-/Authorization-/State-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Artifact Trust, SBOM, Provenance and Revocation (CO-WP-022)

- Drei neue Dokumente: Artifact Identity/Provenance/SBOM Model, Artifact Dependency/Compatibility/Distribution Model, Artifact Trust/Quarantine/Revocation Policy.
- `artifact available ≠ trusted`; `identity ≠ alias ≠ file name ≠ repo path`; `version ≠ revision`; `same version label ≠ same revision`. 19 Artifact-Klassen; 13 Rollen getrennt; Mutable Alias ≠ finale privilegierte Deployment-Bindung (Resolution bindet Identität/Version/Revision).
- Provenance/Integrity/Validation/Trust/Support/Compatibility getrennt (`integrity checked ≠ safe`; `validated ≠ trusted`); SBOM (`available ≠ complete ≠ accurate`; `missing ≠ absence`); Component Identity (`same name/version ≠ same component`); Vulnerability (`reference ≠ affected ≠ deployment exploitable`; `no finding ≠ no vulnerability`). Trust Decision use-/target-/scope-/version-/time-bound (≠ execution authorization); Quarantine release ≠ deployment authorization.
- Withdrawal/Revocation/Repository-Removal getrennt; `revocation → new deployment blocked unless governed exception`; `existing deployment ≠ permission for new`; Reinstatement erhält Historie. Deployment Binding an Resolution + Pre-Execution-Recheck; Offline mit Provenance/Integrität/Target-Binding/Trust/Import-Governance; Delayed Revocation sichtbar.
- Decision Index +15 (DEC-S-289…303), Risk Register +5 (RISK-285…289, gesamt 289; Wachstumsgrenze §4 eingehalten). Keine Registry-/SBOM-/Signing-/Scanner-Technologie; keine ADR; Deployment-/Integration-/Domain-Pack-/Data-/API-/Event-/Evidence-/Telemetry-/Topology-/Identity-/Authorization-/State-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Deployment Control Plane and Blueprint Schema (CO-WP-021)

- Drei neue Dokumente: Deployment Control Plane and Execution Model, Deployment Blueprint Versioning and Compatibility Model, Deployment Targeting/Execution/Recovery Policy.
- Deployment Control Plane = Governance-/Koordinationsgrenze (≠ Runtime). `blueprint ≠ plan`; `intent ≠ approved plan`; `approved plan ≠ execution authorization`; `executed ≠ successful ≠ verified`; `closed ≠ successful`. Write/Execution an CO-WP-013 gebunden (keine parallele Authorization Authority).
- Dynamische Topology-Selektion → materialisierter, begrenzter Target-Set-Snapshot (`topology query ≠ authorised target set`); Pre-Execution-Revalidation; conflicted Identität nicht still privilegiertes Ziel. Blueprint Identity stabil; Parameter/Overlays schwächen Security nicht still (`secret reference ≠ raw secret`); `artifact available ≠ trusted ≠ compatible`. Waves (`wave success ≠ remaining authorised`); Partial/Unknown sichtbar; Cancellation ≠ keine Nebenwirkung; Verification getrennt; Rollback braucht Verifikation; Machine ≠ Human Deployment Authority; Offline mit Provenance/Integrität/Target-Binding/gebundener Authorization/Aktivierung.
- Decision Index +15 (DEC-S-274…288), Risk Register +5 (RISK-280…284, gesamt 284; Wachstumsgrenze §4 eingehalten). Keine Deployment-Engine-/Blueprint-/Registry-Technologie; keine ADR; Policy-/Integration-/Domain-Pack-/Data-/API-/Event-/Evidence-/Telemetry-/Topology-/Identity-/Authorization-/State-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Milestone Lessons Review (CO-WP-013…020)

- Gebündelter docs-only Review der acht Foundation-WPs (24 Dokumente); Ergebnis **GO WITH NOTES FOR CO-WP-021**. Dokument: `project-brain/MILESTONE_REVIEW_CO_WP_013_TO_020.md`.
- Foundation-Kette (Policy/Approval/Execution → Integration → Domain Pack → Data/Migration → API → Event/Evidence → Telemetry → Topology) kohärent; alle Cross-Foundation-Invarianten gehalten; keine Technologie ausgewählt.
- Acht Lessons LL-023…030 (gesamt 30); drei NDF-Kandidaten NDF-FC-COREOPS-011…013 (`candidate-pending-nova-review`; gesamt 13, 008…013 pending).
- Read-only-Befunde (Follow-up): Risk Register 279 → Konsolidierung ~CO-WP-029/030; Decision Index DEC-S-273 (Alt-Kombistatus vs. getrennte Dimensionen) → Harmonisierung ~CO-WP-029; Capability-Count „74→94" ~CO-WP-029; Dokumentationsökonomie ~CO-WP-030. **Decision Index/Risk Register read-only unverändert; kein NDF-Transfer; CO-WP-021 nicht begonnen.**

## Topology Graph, Evidence and Manual Authority (CO-WP-020)

- Drei neue Dokumente: Topology Graph and Relationship Model, Topology Evidence/Confidence/Conflict Model, Topology Manual Authority and Disclosure Policy.
- `topology graph ≠ authoritative physical reality`; `node ≠ managed resource`; `edge ≠ verified connectivity`; `relationship assertion ≠ validated`; `missing edge ≠ no relationship`. 24 Node-/21 Relationship-Klassen; 10 Assertion Origins (`manual ≠ fact`; `inferred ≠ independently observed`). `same name/address/alias ≠ same canonical node`; Merge/Split erhalten Historie/Evidence; kein Timestamp-Last-Write-Wins.
- Evidence-Dimensionen getrennt (Trust/Authority/Confidence/Validation/Evidence/Sufficiency/Completeness); `complete local graph ≠ complete physical topology`; Konflikte sichtbar, ungelöste blockieren privilegierte Automation. Manual Authority human-attributable/scope-bound; Override löscht keine Observations; `suppressed from view ≠ absent`. `topology ≠ state/network/execution authority` (Write/Execution an CO-WP-013 gebunden); Offline mit Provenance/Integrität/Target-Binding/Governance.
- Decision Index +19 (DEC-S-255…273), Risk Register +10 (RISK-270…279, gesamt 279). Keine Graph-DB-/Discovery-/Visualization-Technologie; keine ADR; Telemetry-/Event-/Audit-/Evidence-/API-/Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Telemetry and Normalization Schema (CO-WP-019)

- Drei neue Dokumente: Telemetry Signal and Normalization Model, Telemetry Mapping/Quality/Compatibility Model, Telemetry Trust/Privacy/Disclosure Policy.
- `telemetry ≠ event ≠ audit event ≠ evidence ≠ command`; `signal received ≠ source state currently true`. 22 Signal-Klassen; signal/series/sample/resource identity getrennt; `producer ≠ source`. Raw/Normalized/Derived/Aggregated getrennt. Metric/Log/Trace/Health-Semantik (`missing ≠ zero`; `self-reported ≠ externally verified`; `telemetry absent ≠ service down`).
- Canonical Fields (`same label ≠ same semantics`); Normalization Profiles; Units/Scale/Precision getrennt (`unknown/conflicting unit ≠ safe conversion`); Quality/Confidence/Freshness/Validation getrennt; `recently ingested ≠ recently observed`; Sampling/Aggregation sichtbar; Cardinality/Label-Grenze. `telemetry ≠ authoritative state`; `metric threshold ≠ execution authorization`; `anomaly ≠ incident`; Telemetry-to-Event/-Evidence mit Klassifikation/Provenance; Offline mit Provenance/Integrität/Target-Binding/Governance.
- Decision Index +18 (DEC-S-237…254), Risk Register +10 (RISK-260…269, gesamt 269). Keine Telemetry-/Protokoll-/Storage-/Mapping-Technologie; keine ADR; Event-/Audit-/Evidence-/API-/Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Event, Audit Correlation and Evidence Model (CO-WP-018)

- Drei neue Dokumente: Event and Audit Correlation Model, Evidence Reference/Validation/Lineage Model, Audit Integrity/Retention/Disclosure Policy.
- Konsolidiert Audit-/Evidence-Begriffe ohne Parallelmodell. `event ≠ command ≠ notification ≠ audit event ≠ evidence`; `event identity ≠ correlation/request/operation/attempt`. 20 Event-Klassen. Vier Zeitbegriffe getrennt; `timestamp/sequence ≠ globale Ordnung`; `correlation ≠ causation`; `recorded ≠ validated/complete`; `missing event ≠ Nichtausführung`.
- Sechs Evidence-Dimensionen getrennt (`available ≠ valid ≠ sufficient`); Sufficiency decision-/scope-/time-bound; `handling history ≠ legal admissibility`. `append-only-governed ≠ technisch unveränderlich`; `audit administrator ≠ unrestricted disclosure`; `read ≠ export permission`. Offline mit Provenance/Integrität/Target-Binding/Governance; Failure sichtbar; `closed ≠ complete/validated/sufficient/compliance`.
- Decision Index +17 (DEC-S-220…236), Risk Register +10 (RISK-250…259, gesamt 259). Keine Event-Bus-/Logging-/SIEM-/Storage-/Schema-/Hash-/Signatur-/WORM-/Redaction-Technologie; keine ADR; API-/Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## API Governance, Versioning, Errors and Idempotency (CO-WP-017)

- Drei neue Dokumente: API Governance and Operation Model, API Versioning/Compatibility/Deprecation Model, API Error/Idempotency/Replay Policy.
- Konkretisiert Integration Contract für programmatische Schnittstellen. 13 API-Klassen; API Identity stabil (≠ route/URL/transport/product version); 17 Operation Classes mit Side-Effect-/Privilege-Klassifikation. `API availability ≠ authorization`; `acceptance ≠ execution`; `successful response ≠ verified`; `error response ≠ proof of no side effect`.
- Zwölf Versionsdimensionen getrennt; Request/Response/Error/Behavioural-Compatibility getrennt (`unknown ≠ compatible`; formal additiv kann breaking sein). Idempotency Context ≠ authorization; `unknown outcome → kein Auto-Retry` + Reconciliation; Duplicate/Replay erhält Historie; Bulk erhält Per-Target-Autorität; Pagination ≠ Snapshot/aktuelle Authorization; Consumer-safe Error-Disclosure; Workspace-Isolation.
- Decision Index +16 (DEC-S-204…219), Risk Register +10 (RISK-240…249, gesamt 249). Keine Transport-/API-Style-/Schema-/Statuscode-/Gateway-/Idempotency-/Replay-Technologie; keine ADR; Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Data Ownership, Persistence, Schema Versioning and Migration (CO-WP-016)

- Drei neue Dokumente: Data Ownership and Persistence Model, Schema Versioning and Migration Model, Data Migration Integrity and Recovery Policy.
- Neun Ownership-Dimensionen getrennt; `data owner ≠ storage operator`; `storage responsibility ≠ write authority ≠ migration authority`. 16 Datenklassen mit autoritativem Modul; 10 Persistence-Klassen (`cached ≠ authoritative`; `unknown persistence ≠ safe for destructive migration`).
- Schema Identity stabil; `schema version ≠ data version`; 8 Versionsdimensionen; 11 Compatibility-Klassen (`read ≠ write ≠ round-trip`; `unknown ≠ compatible`); 12 Change Classes. `executed ≠ validated`; `backup exists ≠ restorable`; `restore ≠ service recovery`; `closed ≠ successful`. Partial sichtbar; Mixed-Version bounded; destruktive Migration mit gebundener Autorität + Approval, fail-closed.
- Migration reaktiviert keine widerrufene/konsumierte Authorization; Audit-/Evidence-Provenance erhalten; Offline mit Target-Binding/Provenance/Integrität/Aktivierung.
- Decision Index +18 (DEC-S-186…203), Risk Register +10 (RISK-230…239, gesamt 239). Keine Storage-/DB-/ORM-/Schema-/Migration-/Backup-/Cluster-Technologie; keine ADR; Integration-/Domain-Pack-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Domain Pack Governance, Support Levels and Compatibility (CO-WP-015)

- Drei neue Dokumente: Domain Pack Governance Model, Domain Pack Support and Compatibility Model, Domain Pack Trust/Provenance/Lifecycle Policy.
- Domain Pack = versionierte Governance-/Produktgrenze (≠ adapter/plugin/deployment unit/certification); 15 Pack-Klassen; stabile, nicht wiederverwendbare IDs. Neun Statusdimensionen getrennt; `active ≠ implemented ≠ maintained ≠ supported ≠ validated ≠ universally compatible`. Support Levels SUP-0/1/2/3/D; Support Level ≠ SLA; SUP-3 version-/target-/profil-/limitation-/evidence-bound; Compatibility Claims version-/target-/profil-/evidence-bound; expected ≠ validated; unknown ≠ compatible.
- Community/External ≠ trust/support/endorsement; Vendor ≠ Endorsement; Pack-Aktivierung ≠ Runtime Authority; Dependencies nicht still Core-Pflicht; Offline mit Target-Binding/Provenance/Integrität/Aktivierung; kompromittierter Pack/Maintainer suspendierbar; retired IDs nicht wiederverwendet, historische Evidenz erhalten.
- Decision Index +17 (DEC-S-169…185), Risk Register +10 (RISK-220…229, gesamt 229). Keine Packaging-/Marketplace-/Plugin-/Update-/Dependency-Resolution-/Signaturtechnologie; keine ADR; Integration-/Identity-/Authorization-/State-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## CoreOps Integration Contract v0.1 (CO-WP-014)

- Drei neue Dokumente: CoreOps Integration Contract v0.1, Integration Capability and Operation Model, Integration Trust/Failure/Recovery Policy.
- Contract Version 0.1 ≠ Produktversion. Sechs Capability-Dimensionen getrennt (advertised/detected/permitted/implemented/supported/validated); Adapter erteilt sich keine Permission per Declaration. request acceptance ≠ authorization ≠ execution; completion ≠ success ≠ verification; transport failure ≠ keine Nebenwirkung; unknown outcome sichtbar/blockiert Retry; partial failure sichtbar; rollback/recovery brauchen Verifikation.
- Read-only ohne stille Write Authority; Write/Execution brauchen explizite Policy+Authorization (CO-WP-013 Guards autoritativ); Adapter/Agent erweitern Scope nicht; Integrationsergebnisse erben keine Autorität (CO-WP-011 autoritativ); Offline mit Target-Binding/Provenance/Integrität/expliziter Aktivierung; Extensions überschreiben keine Invariante.
- Decision Index +16 (DEC-S-153…168), Risk Register +13 (RISK-207…219, gesamt 219). Keine Protokoll-/Schema-/Transport-/SDK-/Adapter-/Replay-/Queue-/Messaging-Technologie; keine ADR; Identity-/Authorization-/State-/Provenance-/Modul-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Policy, Approval and Execution Authorization (CO-WP-013)

- Drei neue Dokumente: Policy Decision and Evaluation Model, Approval and Authorization Lifecycle, Execution Authorization and Guard Policy.
- Policy Evaluation, Approval und Execution Authorization getrennt. Policy permit ≠ Approval ≠ Execution Authorization; kein Default-Permit; `indeterminate`/`conflicted` fail-closed; Policy-Konflikte bleiben sichtbar.
- Approval explizit/zurechenbar/scope-bound/widerrufbar; Machine ohne Self-Approval; Approver-Rolle ≠ unbegrenzter Scope; SoD (kleine Deployments berücksichtigt). Execution Authorization action-/target-/scope-/plan-/time-bound; Plan-/Target-Änderung → Re-Evaluation; expired/revoked/consumed nicht wiederverwendbar; Pre-Execution Guards + fail-closed; Adapter/Agent erweitern Scope nicht; Break Glass kein Parallelmodell; Offline Authorization mit Target-Binding/Provenance/Integrität/expliziter Aktivierung.
- Executed ≠ successful ≠ verified; closed ≠ successful; audit evidence ≠ execution authority.
- Decision Index +16 (DEC-S-137…152), Risk Register +17 (RISK-190…206, gesamt 206). Keine Policy-/Approval-/Execution-Engine, kein Token/Artefakt, kein Replay-Mechanismus, keine Queue/Workflow-Runtime; keine ADR; Identity-/Modul-/State-/Provenance-/Capability-/Threat-Dateien + Lessons-Learned + NDF-Kandidaten unverändert; keine NDF-Rückführung.

## Observed, Desired, Effective State and Drift (CO-WP-012)

- Drei neue Dokumente: Observed/Desired/Effective State Model, Drift Detection and Convergence Model, Safe Remediation and State Change Policy.
- Desired/Observed/Reported/Effective/Last-Known getrennt; Effective indeterminate/conflicted bei unklarer Autorität; keine Beobachtung ≠ keine Drift; kein Drift ≠ Compliance. Detection/Recommendation/Plan/Approval/Execution/Verification getrennt; Executed ≠ successful ≠ verified ≠ compliant; Rollback braucht Verifikation; Auto-Remediation deferred; fail-closed.
- Decision Index +15 (DEC-S-122…136), Risk Register +17 (RISK-173…189). Keine Engine-/Scheduler-/Queue-/Auto-Remediation-Auswahl; keine ADR; SoT-/Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung. Milestone-Lessons-Review-Eignung `yes` (CO-WP-005…012).

## Milestone Lessons Review (CO-WP-005…012)

- Gebündelter docs-only Review abgeschlossen; Nova Review pending. Dokument: `project-brain/MILESTONE_REVIEW_CO_WP_005_TO_012.md`. Ergebnis: **GO WITH NOTES FOR CO-WP-013**.
- Sechs konsolidierte Lessons LL-017…022 (gesamt 22); drei NDF-Feedback-Kandidaten NDF-FC-COREOPS-008…010 (`candidate-pending-nova-review`).
- Read-only-Befunde (Follow-up, nicht umgesetzt): Risk Register 189 → Konsolidierung ~CO-WP-029/030; Decision Index Alt-Kombistatus (DEC-S-01…37) vs. getrennte Dimensionen (DEC-S-38…136); Capability-Count „74→94" in Alt-Abschnitten.
- **Risk Register und Decision Index unverändert (read-only); kein NDF-Transfer; CO-WP-013 nicht begonnen.**

## Source of Truth and Field Provenance (CO-WP-011)

- Drei neue Dokumente: Source of Truth and State Authority Model, Field Provenance and Data Lineage Standard, Offline Data Reconciliation and Conflict Policy.
- SoT ≠ System of Record; 10 Authority-Klassen; ein Owner pro Feldkonzept; Desired/Observed/Effective/Derived/Cached getrennt; stabile canonical Feldidentität; Freshness/Trust/Validation getrennt; Lineage erhält Input-Autorität.
- Kein universelles Last-Write-Wins; Konflikte bleiben sichtbar; widerrufene Quellen nicht autoritativ; Import ≠ Authority-Transfer; Offline-Reconciliation fail-closed; Audit-History nicht durch Reconciliation umgeschrieben. Decision Index +16 (DEC-S-106…121), Risk Register +17 (RISK-156…172). Keine Storage-/Merge-/Sync-/Krypto-Auswahl; keine ADR; Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung.

## Machine Identity, Enrollment and Offline Credential Lifecycle (CO-WP-010)

- Drei neue Dokumente: Machine Identity and Principal Governance, Machine Enrollment and Trust Lifecycle, Offline Credential and Rotation Governance.
- Human ≠ Machine; Identity ≠ Credential; Discovery ≠ Enrollment; Registration ≠ Trust; Enrollment ≠ Write Authority. 10 Principal-Klassen; Lifecycle; Bootstrap ≠ permanente Identität; Rotation/Renewal ohne Scope-Erweiterung; Revocation nicht nur Doku; Compromise → explizite Re-Enrollment; Credential-Governance ≠ Rohsecret-Ownership.
- Decision Index +16 (DEC-S-90…105), Risk Register +19 (RISK-137…155). Keine PKI-/Krypto-/Protokoll-/Secret-Store-Auswahl; keine ADR; Human-Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung.

## Human Identity, Workspaces, RBAC and Break Glass (CO-WP-009)

- Drei neue Dokumente: Human Identity and Access Governance, Workspace/RBAC/Scope Model, Break-Glass and Emergency Access Policy.
- person/identity/account/principal getrennt; Repository- ≠ Runtime-Autorität; Auth ≠ Authz; Workspace ≠ Security-Tenant; deny-by-default RBAC; Permission-Taxonomie + 8 Scope-Typen; Cross-Workspace explizit/auditierbar.
- Break Glass benannt/temporär/scope-bound/auditiert mit Ablaufpflicht + verpflichtendem Post-Event Review; Offline-Emergency governed. 14 Security-Invarianten (Designanforderungen).
- Decision Index +14 (DEC-S-76…89), Risk Register +18 (RISK-119…136). Keine Auth-/IdP-/Session-/Policy-Engine-Auswahl; keine ADR; Modul-/Capability-/Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung.

## Architecture and Module Boundaries (CO-WP-008)

- Drei neue Dokumente: Logical Module Architecture, Module Catalog, Module Boundary and Dependency Standard.
- 17 logische Module (MOD-*) mit stabilen IDs; Policy/Control/Execution getrennt; Experience ≠ direkte Execution; Adapter umgehen Governance nicht; Agenten optional; Offline-Intake ohne direkte Ausführung; ein Owner pro autoritativem Konzept; 2 Mermaid-Diagramme (conceptual).
- module ≠ microservice/deployment; keine Technologie/Deployment-Auswahl; Invarianten = Designanforderungen. Decision Index +12 (DEC-S-64…75), Risk Register +15 (RISK-104…118). Keine ADR; Capability Matrix + Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung.

## Threat Model and Trust Boundaries (CO-WP-007)

- Zwei neue Dokumente: Foundation Threat Model + Threat Scenario Register; Trust-Boundary-Dokument additiv erweitert (TB-IDs + Threat-Verweise).
- 24 Assets, 16 Threat Actors, 18 Kategorien; Angriffsflächen alle 10 Planes; TB-01…11; 40 Threat Scenarios (THR-001…040, stabile IDs, qualitative Ratings); 17 Sicherheitsinvarianten; 5 Abuse Cases; 2 Mermaid-Diagramme (conceptual).
- Threat Model = Bedrohungen + Anforderungen, **keine** implementierten/validierten Kontrollen; kein `mitigated`/`closed` ohne Evidenz. Decision Index +10 (DEC-S-54…63), Risk Register +10 (RISK-94…103; einzelne Threats im Register, nicht dupliziert). Kein Pentest; keine Technologie/Krypto; keine ADR; Capability Matrix + Lessons-Learned unverändert; keine NDF-Rückführung.

## System Context, Plane Taxonomy and External Boundaries (CO-WP-006)

- Drei neue Dokumente: System Context and External Boundaries, Plane Taxonomy, Trust/Deployment/Execution Boundaries.
- 7 Akteure; Produkt- ≠ Deployment- ≠ Managed-Grenze; 20 externe Systemklassen (optional); 15 Interaktionsklassen (read ≠ write); Kontrollautorität; Connected/Restricted/Offline-Modi; Failure-Grundregeln; 13 Datenklassen; zwei Mermaid-Diagramme (conceptual).
- 10 logische Planes (plane ≠ Deployment-Einheit); Edge/Agent Plane optional; Managed Resource Plane außerhalb Produktgrenze. 11 Vertrauensgrenzen; Control-to-Execution/Policy-to-Action getrennt; fail-closed; Threat Model deferred (CO-WP-007).
- Decision Index +10 (DEC-S-44…53), Risk Register +14 (RISK-80…93). Keine Technologie/Architektur/Threat Model; keine ADR; Lessons-Learned unverändert; keine NDF-Rückführung.

## Language Standard, Public Neutrality and Repository Governance (CO-WP-005)

- Drei neue Governance-Dokumente: Language Standard, Public-Neutrality-and-Disclosure-Policy, Repository-Governance-Standard.
- Englisch kanonisch für maschinenbezogene Bezeichner; DE/EN primäre Produktsprachen; Translation-Status-Modell; semantische Parität Pflicht; keine automatische Parität.
- Hersteller-/Organisations-/institutionelle Neutralität; Nennung ≠ Endorsement; Public-Sector ≠ Behördenfreigabe; Secrets-/Datenschutzgrenzen; synthetische Beispieldaten; Redaction.
- Human-Maintainer-Gates erhalten; Source-of-Truth-Hierarchie (erweitert FOUNDATION_SCOPE_LOCK); Dokumentstatus/Supersession; Dateinamen/stabile IDs; UTF-8/Zeilenenden; PowerShell-Korrekturstandard.
- Decision Index +6 (DEC-S-38…43), Risk Register +13 (RISK-67…79). Keine ADR; keine automatisierte Durchsetzung; Lessons-Learned unverändert; keine NDF-Rückführung.

## Capability Matrix Security and Governance Alignment (CO-WP-004E)

- Zwei neue Dokumente: Capability-Security-and-Governance-Alignment + Capability-Matrix-Spezifikation (letztere neu erstellt).
- Foundation Capability Matrix additiv um fünf Statusdimensionen (Roadmap/Implementation/Support/Evidence/Security-Governance), Profile Relevance, PSR-01…18-Zuordnung, Responsibility-Codes erweitert; keine Capability gelöscht/umbenannt/hochgestuft.
- **Zählkorrektur: 94 Capabilities (grep-verifiziert); frühere „74"-Summe korrigiert; andere Referenzen später abzugleichen.**
- PSR-Mapping = Readiness-Relevanz, nicht BSI-Compliance; kein Control-Mapping; Evidence `not-assessed`; kein `compliant`. Decision Index +7 (DEC-S-31…37), Risk Register +6 (RISK-61…66). Keine ADR; Lessons-Learned unverändert; keine NDF-Rückführung.

## ITIL and PRINCE2 Applicability and Tailoring (CO-WP-004D)

- Zwei neue Dokumente: ITIL/PRINCE2-Applicability-and-Tailoring und interne Service-/Project-Governance-Profile.
- ITIL `adopted-with-tailoring`; PRINCE2 Version 7 `optional-profile`; vollständige Implementierung beider `rejected`. ITIL 4 & Version 5 mit Versionsgrenzen (Bridge anerkannt).
- NDF bleibt primäres Framework; Konflikthierarchie + Overload-Guards; keine parallele Governance. Drei interne, optionale Profile (keine Zertifizierungsstufen).
- Decision Index +8 (DEC-S-23…30), Risk Register +11 (RISK-50…60). Capability Matrix + Lessons-Learned-Register unverändert; keine NDF-Rückführung; keine Zertifizierung/Endorsement/Tool-Abhängigkeit; keine ADR.

## BSI and Public-Sector Readiness Baseline (CO-WP-004C)

- Drei neue Dokumente: Readiness-Baseline (18 PSR-Domänen, Verantwortungs-/Evidenzmodell), Reference-/Claims-Register (offizielles BSI-Referenzset; C5/C3A bedingt), Public-Sector-Profil (Standard ⊂ Hardened ⊂ Government, intern).
- Claim Boundaries: keine BSI-Zertifizierung/Vollkonformität/Behörden-/VS-NfD-Freigabe; Produktreife ≠ Deployment-Compliance; keine einzelnen BSI-Anforderungen erfunden.
- Evidenz-Drei-Zustands-Modell (capability ≠ available ≠ satisfied). Offline/Air-Gap und Souveränität/Cloud eingeordnet.
- Decision Index +7 (DEC-S-16…22), Risk Register +10 (RISK-40…49). BSI-Positionierung additiv aktualisiert.
- **Capability Matrix unverändert (→ CO-WP-004E); keine ADR; keine Technologie ausgewählt.**

## NDF-Transferpaket – Vollständiger Lebenszyklus (CO-WP-004B1 bis 004B4)

- **CO-WP-004B1:** Transfer Package 001 mit 3 Bundles erstellt; alle 7 Kandidaten auf `approved-for-transfer`.
- **CO-WP-004B2:** Package Status und 7 Human-Maintainer-Gates auf `approved`; Lesson LL-016 erfasst.
- **CO-WP-004B3:** Alle 7 Kandidaten auf `transferred-to-ndf` (NDF-INTAKE-COREOPS-001, Commit d08e35e).
- **CO-WP-004B4:** Alle 7 Kandidaten auf `adopted-in-ndf`, verteilt auf drei Adoption-WPs:
  - Adoption A (`NDF-ADOPT-COREOPS-001A`, Commit 1ebffa6): 001, 003, 004.
  - Adoption B (`NDF-ADOPT-COREOPS-001B`, Commit e894c6f): 002.
  - Adoption C (`NDF-ADOPT-COREOPS-001C`, Commit ebf716c): 005, 006, 007.
- **Transfer Package 001 Status:** `Closed – all candidates processed`.
- **Release-Grenze:** NDF-Release-Zuordnung bleibt `not yet assigned`; keine NDF-Version behauptet. `adopted-in-ndf` ≠ in veröffentlichter NDF-Version enthalten.
- **NDF-Repository unverändert durch CoreOps; keine weiteren Adoption-WPs gestartet.**

## Lessons Learned und NDF Feedback (CO-WP-004B)

- Lessons-Learned-Prozess und NDF-Feedback-Prozess etabliert (Proposed for acceptance).
- 15 Lessons erfasst (13 retrospektiv aus CO-WP-001…004A, 2 aus CO-WP-004B selbst) im Lessons-Learned-Register.
- 7 reservierte NDF-Feedback-Kandidaten (NDF-FC-COREOPS-001…007) bewertet, alle `ready-for-bundling`.
- **Kein Kandidat übertragen/adoptiert; kein NDF-Repository verändert; kein NDF-Work-Package erstellt.**
- Transferentscheidungen bleiben Nova Review + Human-Maintainer-Gate vorbehalten.

## Souveränität und BSI-Orientierung (CO-WP-004A)

- Concept-v3.1-Amendment, Sovereignty-and-Dependency-Policy und BSI-Alignment-Positioning erstellt (Accepted Product Direction).
- Produktsouveränität akzeptiert; verpflichtende externe Managementprodukte als Kernabhängigkeit ausgeschlossen; technische Basisabhängigkeiten offen (keine akzeptiert).
- BSI-orientiert; **keine** Zertifizierung/Konformität/Zulassung/VS-Eignung behauptet. Standard-/Hardened-/Government-Profile registriert.
- Lessons Learned (Governance-Richtung) und kontrollierter NDF-Rückfluss (Kandidatenprozess) registriert; ITIL/PRINCE2 nur Kandidaten.
- Capability Matrix unverändert; technische Architektur offen; keine ADR.

## Capability Matrix und Support Boundary (CO-WP-004)

- Foundation Capability Matrix (`docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md`): 74 Capabilities, 12 Domains; alle `not-implemented`/`not-supported` (Proposed for acceptance).
- Initiale Observe-Supportgrenze (`docs/integrations/INITIAL_SUPPORT_BOUNDARY.md`): Observe vollständig read-only, max. Integrationslevel 2; Map/Write ausgeschlossen.
- Drei getrennte Statusdimensionen (Roadmap/Implementation/Support), Herausgeberklassen und 21-Punkte-Support-Evidence-Satz definiert.
- `CCR-12` vorgeschlagen aufgelöst: Herstellernennung = Kandidat, kein Support.
- **Keine Runtime-Capability implementiert; keine Integration `supported`; keine technische Architektur Accepted; keine ADR.**

## Brief, Scope Lock und Release-Taxonomie (CO-WP-003)

- Project Brief, Foundation Scope Lock und Release-Taxonomie erstellt (`Proposed for acceptance`); Foundation `v0.0.1-foundation`, Observe `v0.1.0-alpha.1`.
- Docker-first eingeordnet; aktive Queue autoritativ; NDF-`main` und NDF-Level geklärt; Repository verifiziert.

## Concept-Registrierung

- Concept v3.0 vollständig registriert (57 Abschnitte); Decision Classification, Decision Index und Risk Register vorhanden.

## Lokaler NDF-Skills-Bestand

Vollständiger NDF-v1.0.0-Skills-Pack lokal unter `.claude/skills/` verfügbar (38 Skills, 39 Dateien, byte-identisch, Commit `9dcadc1`). Skills werden **pro Work Package selektiv** verwendet, nicht automatisch vollständig aktiviert. Provenance/Lock: `project-system/NDF_SKILLS_PROVENANCE.md`, `project-system/ndf-skills-lock.json`.

## Nächstes Work Package

`CO-WP-025 – Data Classification, Retention and Redaction` (docs-only; planned-next; **noch nicht begonnen**; erst nach Human-Maintainer-Commit von CO-WP-024). Die 004er-Erweiterungsserie (004A…004E), CO-WP-005…023, beide Milestone Lessons Reviews (005…012, 013…020) und CO-WP-024 sind bearbeitet; CO-WP-023 ist `completed-go-with-notes` (Commit b324aad), CO-WP-024 ist nach Nova Review `GO WITH NOTES` und geschlossener Korrekturrunde (Notes 1–4) `completed-go-with-notes` — Human-Maintainer-Commit ausstehend, Push nicht als erfolgt dargestellt. CO-WP-021/022/023 wurden committet und gepusht (Branch bei CO-WP-024-Preflight level mit origin/main). Milestone-Bündelentscheidung für NDF-Kandidaten 008…013 liegt bei Nova/HM. External: NDF-Release-Zuordnung für die drei Adoption-Commits (001…007) ausstehend.

## Aktuelle Blocker

- Keine harten Blocker. In CO-WP-003/004/004A/004B/004B1 adressiert (proposed/treatment-planned/accepted-direction/approved-for-transfer): Release-Taxonomie, Docker-first, Queue-Autorität, NDF-`main`, NDF-Level, CCR-12, Observe-Scope, Legacy-Protokolle, Capability-Maturity, Drucktelemetrie, Produktsouveränität, BSI-Orientierung, Basisabhängigkeiten, Lessons-Learned-/NDF-Feedback-Prozess, erstes NDF-Transferpaket.
- Weiter offen: Konflikte CCR-01, 05, 06, 07, 08, 09; technische ADR-Kandidaten; ITIL/PRINCE2-Tailoring (CO-WP-004D); konkretes IT-Grundschutz-Mapping (CO-WP-004C); NDF-seitiger Intake Review für Transfer Package 001.

## Zentrale Sicherheitsgrenzen

- Human-Maintainer-only für Freigabe, Staging, Commit, Push, Merge, Tags, Releases, Deployments, irreversible/privilegierte Aktionen.
- Kein Code in Foundation.
- Keine autonomen Git-, Netzwerk- oder Infrastrukturaktionen.
- Fail-closed bei Unklarheit.
- Keine Übernahme privater Daten in Projektdateien.

## Akzeptierte ADRs

- **Keine.** Es existiert keine akzeptierte technische ADR.

## Bekannte Notes und Risiken

- NDF-Level-Semantik geklärt (CCR-03: Bootstrap-Status, Zahlenwert unverändert).
- Repository-URL im Manifest verifiziert und gesetzt: `https://github.com/KayKaspers/CoreOps`.
- Release-Taxonomie vorgeschlagen aufgelöst (CCR-02): `v0.0.1-foundation` / `v0.1.0-alpha.1`.
- Docker-first eingeordnet (CCR-10); aktive Queue-Autorität geklärt (CCR-11); NDF-`main` geklärt (CCR-04); Herstellersupport-Grenze vorgeschlagen aufgelöst (CCR-12).
- Capability Matrix (74 Capabilities) und Observe-Supportgrenze vorhanden; alle Capabilities `not-implemented`/`not-supported`.
- Souveränität und BSI-Orientierung als Accepted Product Direction registriert (CO-WP-004A); keine Zertifizierung/VS-Eignung behauptet; keine konkrete Dependency akzeptiert.
- Lessons-Learned- und NDF-Feedback-Prozess etabliert (CO-WP-004B); 15 Lessons, 7 NDF-Kandidaten (alle `ready-for-bundling`); kein Kandidat übertragen/adoptiert.
- Weiter offen: Konflikte CCR-01, 05, 06, 07, 08, 09; 39 Foundation-Risiken (RISK-01…39, davon 22 `treatment-planned`).
- Keine akzeptierten technischen ADRs.

## Relevante Quelldateien

- [project-system/project-manifest.yaml](../project-system/project-manifest.yaml)
- [project-system/PROJECT_PROFILE.md](../project-system/PROJECT_PROFILE.md)
- [project-system/WORK_PACKAGE_QUEUE.md](../project-system/WORK_PACKAGE_QUEUE.md)
- [project-system/NEXT_PHASE.md](../project-system/NEXT_PHASE.md)
- [project-system/DECISION_INDEX.md](../project-system/DECISION_INDEX.md)
- [project-system/RISK_REGISTER.md](../project-system/RISK_REGISTER.md)
- [project-system/NDF_SKILLS_PROVENANCE.md](../project-system/NDF_SKILLS_PROVENANCE.md)
- [project-system/ndf-skills-lock.json](../project-system/ndf-skills-lock.json)
- [docs/architecture/COREOPS_CONCEPT_V3.md](../docs/architecture/COREOPS_CONCEPT_V3.md)
- [docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md](../docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md)
- [docs/architecture/PROJECT_BRIEF.md](../docs/architecture/PROJECT_BRIEF.md)
- [docs/governance/FOUNDATION_SCOPE_LOCK.md](../docs/governance/FOUNDATION_SCOPE_LOCK.md)
- [docs/governance/RELEASE_TAXONOMY.md](../docs/governance/RELEASE_TAXONOMY.md)
- [docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md](../docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md)
- [docs/integrations/INITIAL_SUPPORT_BOUNDARY.md](../docs/integrations/INITIAL_SUPPORT_BOUNDARY.md)
- [docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md](../docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md)
- [docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md](../docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md)
- [docs/security/BSI_ALIGNMENT_POSITIONING.md](../docs/security/BSI_ALIGNMENT_POSITIONING.md)
- [docs/governance/LESSONS_LEARNED_PROCESS.md](../docs/governance/LESSONS_LEARNED_PROCESS.md)
- [docs/governance/NDF_FEEDBACK_PROCESS.md](../docs/governance/NDF_FEEDBACK_PROCESS.md)
- [project-system/LESSONS_LEARNED_REGISTER.md](../project-system/LESSONS_LEARNED_REGISTER.md)
- [project-system/NDF_FEEDBACK_CANDIDATES.md](../project-system/NDF_FEEDBACK_CANDIDATES.md)
- [docs/governance/NDF_FEEDBACK_TRANSFER_PACKAGE_001.md](../docs/governance/NDF_FEEDBACK_TRANSFER_PACKAGE_001.md)
- [project-brain/PROJECT_BRAIN.md](PROJECT_BRAIN.md)
- [ROADMAP.md](../ROADMAP.md)

## Kompakte Historie

- `CO-WP-001` (docs-only): Core Governance Skeleton erstellt (7 Dateien). Nova-Bewertung: `GO WITH NOTES` (read-only-Git-Note).
- `CO-WP-001A` (docs-only): Vollständiger NDF-v1.0.0-Skills-Pack lokal bereitgestellt (38 Skills, 39 Dateien, byte-identisch, Commit `9dcadc1`), Provenance + Lock erstellt. Nova-Bewertung: `GO`.
- `CO-WP-002` (docs-only): Nach fail-closed-Blocker und Source-Handoff Concept v3.0 vollständig registriert (57 Abschnitte), Decision Classification + Decision Index + Risk Register erstellt. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-003` (docs-only): Project Brief, Foundation Scope Lock und Release-Taxonomie erstellt (Proposed); CCR-02/04/10/11 aufgelöst, NDF-Level geklärt, Repository verifiziert. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004` (docs-only): Foundation Capability Matrix (74 Capabilities) und initiale Observe-Supportgrenze erstellt (Proposed); drei Statusdimensionen, Support-Evidence, CCR-12 vorgeschlagen aufgelöst. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004A` (docs-only / gov-baseline): Concept-v3.1-Amendment, Sovereignty-and-Dependency-Policy und BSI-Alignment-Positioning erstellt; Souveränität und BSI-Orientierung als Accepted Product Direction; keine Zertifizierung/VS-Eignung; Standard-/Hardened-/Government-Profile; Lessons-Learned-/NDF-Feedback-Richtung; ITIL/PRINCE2 Kandidaten. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004B` (docs-only / governance): Lessons-Learned-Prozess und NDF-Feedback-Prozess etabliert; 15 Lessons erfasst; 7 reservierte NDF-Kandidaten bewertet (alle `ready-for-bundling`). Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004B1` (docs-only / transfer-preparation): Transfer Package 001 mit 3 Bundles erstellt; alle 7 Kandidaten auf `approved-for-transfer`. Erster NDF Intake-Versuch fail-closed blockiert.
- `CO-WP-004B2` (docs-only / governance-correction): Package Status und alle 7 Human-Maintainer-Gates auf `approved`; neue Lesson LL-016 erfasst.
- `CO-WP-004B3` (docs-only / cross-project-traceability): Transfer Package 001 dem NDF-Intake übertragen (NDF-INTAKE-COREOPS-001, Commit d08e35e); alle 7 Kandidaten auf `transferred-to-ndf`.
- `CO-WP-004B4` (docs-only / cross-project adoption traceability): Alle 7 Kandidaten auf `adopted-in-ndf` über drei Adoption-WPs (Commits 1ebffa6, e894c6f, ebf716c); Transfer Package 001 geschlossen; NDF-Release-Zuordnung bleibt offen. Nova Review ausstehend.
- `CO-WP-004C` (docs-only / security-governance baseline): BSI- und Public-Sector-Readiness-Baseline (18 PSR-Domänen), Reference-/Claims-Register und Public-Sector-Profil erstellt; Decision Index +7, Risk Register +10; BSI-Positionierung additiv aktualisiert. Keine Zertifizierung/Compliance/Behördenfreigabe behauptet; Capability Matrix unverändert; keine ADR. Nova Review ausstehend.
- `CO-WP-004D` (docs-only / governance-framework applicability review): ITIL `adopted-with-tailoring`, PRINCE2 Version 7 `optional-profile`; ITIL 4 & Version 5 mit Versionsgrenzen; drei interne Governance-Profile; NDF bleibt primär; Decision Index +8 (DEC-S-23…30), Risk Register +11 (RISK-50…60). Keine Zertifizierung/Endorsement/Tool-Abhängigkeit; Capability Matrix + Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-004E` (docs-only / capability-governance alignment): Foundation Capability Matrix um fünf Statusdimensionen, Profile, PSR-Mapping und Responsibility-Codes erweitert (94 Capabilities; Zählkorrektur von „74"); Alignment-Dokument + Capability-Matrix-Spezifikation erstellt; Decision Index +7 (DEC-S-31…37), Risk Register +6 (RISK-61…66). PSR-Mapping ≠ BSI-Compliance; keine Capability implementiert; keine ADR; Lessons-Learned unverändert. Nova Review ausstehend.
- `CO-WP-005` (docs-only / repository-governance foundation): Language Standard, Public-Neutrality-and-Disclosure-Policy und Repository-Governance-Standard erstellt; Englisch kanonisch (maschinenbezogen), DE/EN Produkt; Neutralitäts-/Disclosure-Grenzen; Source-of-Truth-Hierarchie; Human-Maintainer-Gates; UTF-8/Zeilenenden; PowerShell-Korrekturstandard. Decision Index +6 (DEC-S-38…43), Risk Register +13 (RISK-67…79). Keine ADR; keine automatisierte Durchsetzung; Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-006` (docs-only / architecture-context foundation): System Context and External Boundaries, Plane Taxonomy (10 Planes) und Trust/Deployment/Execution Boundaries erstellt; Produkt- ≠ Deployment- ≠ Managed-Grenze; 20 externe Systemklassen; Kontrollautorität; Connected/Restricted/Offline-Modi; zwei Mermaid-Diagramme (conceptual). Decision Index +10 (DEC-S-44…53), Risk Register +14 (RISK-80…93). Keine Technologie/Architektur/Threat Model; keine ADR; Capability Matrix + Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-007` (docs-only / security-baseline): Foundation Threat Model (24 Assets, 16 Actors, 18 Kategorien, 10 Plane-Angriffsflächen, TB-01…11, 17 Invarianten, 5 Abuse Cases, 2 Mermaid) und Threat Scenario Register (40 Szenarien THR-001…040) erstellt; Trust-Boundary-Dokument additiv erweitert. Decision Index +10 (DEC-S-54…63), Risk Register +10 (RISK-94…103). Keine implementierten/validierten Kontrollen; kein Pentest; keine Technologie/Krypto; keine ADR; Capability Matrix + Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-008` (docs-only / logical-architecture foundation): Logical Module Architecture (17 Module MOD-*), Module Catalog (17-Modul-Register, Daten-/Zustandsownership) und Module Boundary and Dependency Standard erstellt; Policy/Control/Execution getrennt; module ≠ microservice/deployment; 2 Mermaid-Diagramme. Decision Index +12 (DEC-S-64…75), Risk Register +15 (RISK-104…118). Keine Technologie/Deployment; keine ADR; Capability Matrix + Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-009` (docs-only / security-baseline): Human Identity and Access Governance, Workspace/RBAC/Scope Model und Break-Glass and Emergency Access Policy erstellt; person/identity/account getrennt; Repository ≠ Runtime; Workspace ≠ Tenant; deny-by-default RBAC + Permission-Taxonomie; Break Glass benannt/temporär/auditiert mit Post-Event Review; 14 Invarianten. Decision Index +14 (DEC-S-76…89), Risk Register +18 (RISK-119…136). Keine Auth-/IdP-/Session-/Policy-Engine-Auswahl; keine ADR; Modul-/Capability-/Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-010` (docs-only / security-baseline): Machine Identity and Principal Governance, Machine Enrollment and Trust Lifecycle und Offline Credential and Rotation Governance erstellt; Human ≠ Machine; Discovery ≠ Enrollment; Registration ≠ Trust; 10 Principal-Klassen; Bootstrap ≠ permanente Identität; Credential-Governance ≠ Rohsecret-Ownership; Offline mit Provenance/Approval. Decision Index +16 (DEC-S-90…105), Risk Register +19 (RISK-137…155). Keine PKI-/Krypto-/Protokoll-/Secret-Store-Auswahl; keine ADR; Human-Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-011` (docs-only / data-governance and architecture foundation): Source of Truth and State Authority Model, Field Provenance and Data Lineage Standard und Offline Data Reconciliation and Conflict Policy erstellt; SoT ≠ System of Record; ein Owner pro Feldkonzept; Desired/Observed/Effective/Derived getrennt; kein Last-Write-Wins; Konflikte sichtbar; Import ≠ Authority; fail-closed Reconciliation. Decision Index +16 (DEC-S-106…121), Risk Register +17 (RISK-156…172). Keine Storage-/Merge-/Sync-/Krypto-Auswahl; keine ADR; Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-012` (docs-only / state-management and safe-remediation foundation): Observed/Desired/Effective State Model, Drift Detection and Convergence Model und Safe Remediation and State Change Policy erstellt; State-Semantik getrennt; keine Beobachtung ≠ keine Drift; Detection/Recommendation/Approval/Execution/Verification getrennt; Executed ≠ successful ≠ verified ≠ compliant; Auto-Remediation deferred; fail-closed. Decision Index +15 (DEC-S-122…136), Risk Register +17 (RISK-173…189). Keine Engine-/Scheduler-/Queue-/Auto-Remediation-Auswahl; keine ADR; SoT-/Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned unverändert; keine NDF-Rückführung. Milestone-Lessons-Review-Eignung `yes`. Nova Review ausstehend.
