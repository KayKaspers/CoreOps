# CoreOps – Redaction, Minimization and Controlled Disclosure Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation redaction, minimization and controlled-disclosure policy
> Implementation Status: Not implemented
> Redaction Engine: Not selected
> Masking/Anonymization Technology: Not selected
> Export Technology: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Legal/Regulatory Mapping: None performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-025 (docs-only / data-governance, retention, minimization, redaction and controlled-disclosure foundation)

## 1. Status

Technologieunabhängige Policy für **Data Minimization, Redaction/Masking/Suppression, abgeleitete Views, kontrollierte Disclosure/Export, Log-/Telemetry-/Audit-/Evidence-Grenzen, Secret-Grenze und Residual Risk**. Companion zu [DATA_CLASSIFICATION_AND_HANDLING_MODEL.md](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md) und [DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md](../governance/DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md). Konkretisiert [Audit Redaction/Disclosure (CO-WP-018)](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md) und [Telemetry Privacy (CO-WP-019)](TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md). Kein Redaction-/Masking-/Anonymisierungs-Mechanismus implementiert.

## 2. Purpose

Redaction erzeugt eine governte abgeleitete View und macht Daten nicht automatisch offenlegungssicher. Die Policy legt fest, warum `redacted view ≠ modified source record`, `redaction applied ≠ disclosure safe`, `masked ≠ anonymous`, `pseudonymized ≠ anonymous`, `hash or fingerprint ≠ anonymization`, `read authority ≠ export authority` und `data available ≠ disclosure authorized`.

## 3. Scope

Authority Model · Data Minimization · Redaction/Derived Views · Masking/Suppression/Pseudonymization/Anonymization Claims · Controlled Disclosure · Export · Logs/Events/Telemetry/Topology · Audit/Evidence · Secret-/Key-Grenze · Backup/Offline-Grenze · Workspace Isolation · Residual Disclosure/Unknown State · Fail-Closed.

## 4. Non-Goals

- Keine Redaction-/Masking-/Tokenisierungs-/Pseudonymisierungs-/Anonymisierungs-/DLP-/Scanner-/OCR-/ML-/Export-/Hash-/Encryption-Technologie.
- Keine Datenschutz-/Rechtskonformitäts-/Anonymitäts-Garantie; keine Compliance-/Zertifizierungs-Behauptung; kein Runtime-Code; keine Redaction realer Dateien.

## 5. Concepts

`data minimization` · `redaction` · `masking` · `suppression` · `field removal` · `aggregation` · `pseudonymization` · `anonymization claim` · `source identity/revision` · `redaction policy/purpose/authority` · `transformation category` · `derived-view identity/revision` · `reversibility classification` · `residual disclosure assessment` · `disclosure` · `export` · `derived disclosure view` · `recipient` · `redaction evidence` · `outcome-unknown`.

## 6. Authority Model

Getrennt: redaction authority · redaction-review authority · disclosure authority · export authority · publication authority. `redaction operator ≠ disclosure authority`; `redaction authority ≠ publication authority`; `read authority ≠ export authority`. Bindet an die Daten-Autoritäten des [Classification Model §8](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md) und an [CO-WP-013](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md).

## 7. Data Minimization

`technically collectible ≠ authorized to collect`; `useful ≠ necessary`; `collected once ≠ authorized for every later use`. Minimierung gilt für Collection, Derivation, Logging, Telemetry, Evidence und Export. Event Payloads/Logs erfassen nur Notwendiges; Identitätsreferenzen statt Vollprofile (Bezug [CO-WP-018 §20](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md)).

## 8. Redaction and Derived Views

Redaction verändert die autoritative Quelle/Source Evidence **nicht** still, sondern erzeugt eine gebundene abgeleitete View mit: source identity · source revision · redaction policy · redaction purpose · redaction authority · fields/regions affected · transformation category · derived-view identity · derived-view revision · reversibility classification · residual disclosure assessment · reviewer · target audience · validity · evidence · known limitations.

```text
redacted view    ≠ source record
redaction applied ≠ safe disclosure
visual hiding    ≠ source removal
```

## 9. Masking, Suppression, Pseudonymization and Anonymization Claims

Getrennt: redaction · masking · suppression · field removal · aggregation · pseudonymization · claimed anonymization.

```text
masked         ≠ anonymous
pseudonymized  ≠ anonymous
aggregation    ≠ guaranteed anonymity
hash or fingerprint ≠ anonymization
```

Anonymisierung wird **nur** als Claim mit Evidenz- und Scope-Grenze beschrieben, niemals als automatisch erreicht. Reversibilität und Residual Disclosure bleiben explizit bewertet.

## 10. Controlled Disclosure

Getrennt: local read · operational use · support view · diagnostic view · audit view · evidence view · redacted disclosure · raw disclosure · export · cross-workspace transfer · public publication. Eine Disclosure erfasst: disclosure identity · source · derived view · recipient · purpose · scope · classification · redaction state · approval · authority · validity · transfer instance · receipt · downstream limitations · audit reference.

```text
data available          ≠ disclosure authorized
support access          ≠ raw-data access
audit access            ≠ source-data access
evidence access         ≠ source-data export
local availability      ≠ disclosure authorization
public classification   ≠ publication approval
```

## 11. Export

`read authority ≠ export authority`; `redaction authority ≠ publication authority`; `configuration export ≠ secret export` (CO-WP-024). Export benötigt eigene Autorität, Zweck-/Scope-Bindung, Klassifikations- und Redaction-State, Recipient und Audit; ausgehende Pakete enthalten nur den zulässigen Scope (Bezug THR-018, THR-020, THR-025, THR-040).

## 12. Logs, Events, Telemetry and Topology

Bindet an [Event/Audit (CO-WP-018)](../architecture/EVENT_AND_AUDIT_CORRELATION_MODEL.md), [Telemetry (CO-WP-019)](../architecture/TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md), [Topology (CO-WP-020)](../architecture/TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md): Logs erben nicht automatisch eine niedrigere Klassifikation; Telemetry übernimmt sensitive/secret-bearing Inhalte nicht ungeprüft; Topology-Daten können sensitiv sein ohne Secret Values; Event Payloads minimiert; Correlation IDs sind **nicht** als personenbezogen-frei oder secret-frei garantiert; Redaction zerstört Audit-Korrelation nicht unbemerkt; Sampling ist keine Retention-/Löschgarantie; `missing logs ≠ no operation occurred`. Keine Logging-/Telemetry-/Redaction-Technologie.

## 13. Audit and Evidence

Bindet an [CO-WP-018](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md). Getrennt: source evidence · evidence package · redacted evidence view · disclosure package · audit metadata · evidence acceptance · evidence sufficiency · retention of evidence · retention of source data.

```text
evidence retained ≠ all source data retainable
redacted evidence ≠ complete evidence
evidence package received ≠ evidence accepted
evidence accepted ≠ claim proven beyond its scope
source deletion   ≠ evidence deletion automatically
```

Redaction von Evidence trägt Source Binding · Revision Binding · Transformation Evidence · Reviewer · Disclosure Scope · Known Limitations.

## 14. Secret and Key Boundary

Bindet an [CO-WP-024](SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md): Secret Values sind keine gewöhnlichen Datenfelder; Redaction ist **kein** zulässiger Ersatz für Secret Removal/Governance; ein teilweise maskiertes Secret bleibt secret-bearing; Logs/Telemetry/Evidence/Backups/CorePacks werden durch Redaction **nicht** „sicher" für Raw Secrets erklärt; Key Material wird nicht durch Hashing/Masking zu gewöhnlichen Daten; Deletion eines Secret Values ersetzt keine Revocation-/Rotation-Entscheidung (Bezug THR-019, THR-020).

## 15. Backup and Offline Boundary

Redaction/Disclosure-Bindungen gelten auch für Kopien/Backups/Offline: `historical backup ≠ current disclosure authorization`; Offline-Kopien behalten Classification-/Redaction-/Disclosure-Bindings; Evidence-Return-CorePacks enthalten nur den zulässigen redigierten/klassifizierten Scope; Import/Receipt ≠ Evidence Acceptance (Bezug [CO-WP-023](../architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md); THR-025, THR-040). Offline-Betrieb erweitert keine Disclosure-Autorität.

## 16. Workspace, Environment and Target Isolation

`same bytes ≠ same disclosure authority`; `derived data retains applicable source restrictions`; `cross-workspace disclosure requires explicit authorization`; Rename ≠ Target-Rebinding; geteilte Infrastruktur ≠ gemeinsame Disclosure Authority (Bezug THR-035).

## 17. Residual Disclosure and Unknown State

`redaction outcome unknown · residual disclosure unknown · export receipt unknown · derived-data lineage incomplete · unauthorized disclosure suspected`.

```text
unknown          ≠ safe
failure          ≠ no side effects
retry            ≠ automatically permitted
missing evidence ≠ operation did not occur
```

Residual Disclosure wird explizit bewertet; unklare Reversibilität/Residualrisiko blockiert privilegierte Disclosure.

## 18. Fail-Closed Rules

Keine Disclosure/Export/Publikation bei: unklarer Disclosure-/Export-Autorität · unbewertetem Residual Disclosure · unvollständiger Derived-Data-Lineage · Secret-bearing Inhalt ohne CO-WP-024-Governance · Target-/Workspace-Mismatch · fehlendem Audit-Start · unbekanntem Export-Receipt für privilegierte Offenlegung. Kein stiller Fallback auf eine weitergehende Disclosure.

## 19. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Redaction creates a governed derived view and does not silently alter source data or source evidence.
2. Redacted does not automatically mean safe for disclosure.
3. Masking, suppression, pseudonymization, hashing and aggregation do not automatically prove anonymity; anonymization is only a scoped, evidence-bounded claim.
4. Read authority does not imply disclosure or export authority; disclosure/export/publication authorities remain separate and classification-bound.
5. Data availability does not imply disclosure authorization.
6. Secret-bearing data remains governed by CO-WP-024; a partially masked secret stays secret-bearing; redaction is not secret removal.
7. Logs and telemetry do not inherit a lower classification; correlation IDs are not guaranteed personal-data-free or secret-free.
8. Evidence retention is separate from source-data retention; redacted evidence is not complete evidence.
9. Derived, cached, backed-up and offline copies retain applicable disclosure restrictions; offline operation expands no disclosure authority.
10. Partial and unknown redaction or disclosure outcomes remain visible and block unsafe retry; missing evidence does not prove no operation occurred.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 20. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): evidence to wrong recipient THR-018; secret leak in logs/exports THR-019/THR-020; offline export sensitive THR-025; stolen backup/evidence recipient THR-040; manipulated telemetry THR-012; stale as current THR-013; audit deletion/manipulation THR-016/THR-017; tenant/org boundary THR-035; insider THR-037. Keine erfundenen IDs; kein Parallelregister.

## 21. Technology Boundary

Nicht ausgewählt: Redaction-/Masking-/Tokenisierungs-/Pseudonymisierungs-/Anonymisierungs-/DLP-/Scanner-/OCR-/ML-/Export-/Hash-/Encryption-Engine · Datenformat/Label-Format · Cloud Provider. Alle bleiben `deferred`.

## 22. Compatibility

Konkretisiert DEC-P-08, [Audit Redaction/Disclosure (CO-WP-018)](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md), [Telemetry Privacy (CO-WP-019)](TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md), [Public Neutrality/Disclosure (CO-WP-005)](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md), [Secrets (CO-WP-024)](SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md). Additiv; keine bestehende Invariante geschwächt.

## 23. Open Questions

- Konkrete Reversibilitäts-/Residual-Disclosure-Bewertungsverfahren (deferred, keine Technologieauswahl).
- Grenzen automatisierter Redaction ohne Scanner-/ML-Auswahl (deferred).
- Verhältnis interner Anonymisierungs-Claims zu späteren rechtlichen Anforderungen (bewusst nicht gemappt).

## 24. Next Decision

Nächstes lokal registriertes Work Package gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): `CO-WP-026 – Self-Protection, Degraded Modes and Recovery Mode`. Zuerst Nova Review von CO-WP-025, danach Human-Maintainer-Commit.
