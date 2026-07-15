# CoreOps – Data Classification and Handling Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation data classification and handling model
> Implementation Status: Not implemented
> Classification Technology: Not selected
> Data Store: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Legal/Regulatory Mapping: None performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-025 (docs-only / data-governance, retention, minimization, redaction and controlled-disclosure foundation)

## 1. Status

Technologieunabhängiges Modell für **Datenklassifikation, Klassifikationsidentität/Freshness, Daten-Autoritäten, Collection/Minimization, Handling-Lebenszyklus und Isolation**. Companion zu [DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md](DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md) und [REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md](../security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md). Konkretisiert [DEC-P-08 Privacy by Design](../../project-system/DECISION_INDEX.md); baut auf [Data Ownership (CO-WP-016)](../architecture/DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [Audit/Evidence (CO-WP-018)](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md), [Secrets (CO-WP-024)](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md) auf; kein Parallelmodell.

## 2. Purpose

Klassifikation ist eine governte Aussage mit Scope und Evidenzgrenze, keine bewiesene Eigenschaft. Das Modell legt fest, warum `data classification ≠ deployment profile`, `data classification ≠ connectivity class`, `data classification ≠ national-security classification`, `classification label ≠ proven handling compliance`, `data owner ≠ unrestricted data user` und `collection permitted ≠ every later use permitted`.

## 3. Scope

Concepts · Klassifikationsklassen · Classification Identity/Freshness · Daten-Autoritäten · Collection/Minimization · Handling-Lebenszyklus · Workspace-/Environment-/Target-Isolation · Secret-/Key-Grenze · Profile · Failure/Unknown State. Retention/Deletion/Preservation im [Companion Retention Policy](DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md); Redaction/Disclosure im [Companion Disclosure Policy](../security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md).

## 4. Non-Goals

- Keine Datenbank/Object Store/Filesystem/Data Lake/Warehouse/Suchindex/Archiv/DLP/CASB/SIEM/Logging-/Telemetrie-/Klassifikations-/Scanner-/OCR-/ML-/KI-Klassifikationstechnologie.
- Keine gesetzliche/regulatorische/DSGVO-/BSI-Zuordnung; keine staatliche Geheimschutzklassifizierung; keine Compliance-/Rechtskonformitäts-/Zertifizierungs-/Einsatzreife-Behauptung; kein Runtime-Code.

## 5. Concepts

`data object` · `data record` · `data set` · `data field` · `metadata` · `content` · `source data` · `derived data` · `aggregated data` · `cached data` · `replicated data` · `exported data` · `evidence data` · `audit data` · `telemetry data` · `log data` · `topology data` · `configuration data` · `secret-bearing data` · `unknown-classification data` · `classification` · `classification label` · `classification authority` · `data owner` · `data steward` · `data custodian` · `data consumer` · `retention policy` · `retention period` · `retention start event` · `retention expiry` · `preservation hold` · `deletion` · `logical deletion` · `purge` · `destruction` · `redaction` · `masking` · `suppression` · `pseudonymization` · `anonymization claim` · `disclosure` · `export` · `derived disclosure view` · `deletion evidence` · `redaction evidence` · `outcome-unknown`.

```text
classification label = governte Aussage mit Scope und Evidenzgrenze
classification label ≠ automatisch bewiesener Inhalt oder Schutz
```

## 6. Data Classification Classes

Technologieunabhängige Klassen (keine staatliche Geheimschutzklassifizierung):

| Klasse | Kurzcharakter |
| ------ | ------------- |
| `public` | zur Veröffentlichung freigegeben (≠ unrestricted modification) |
| `internal` | intern (≠ safe for external disclosure) |
| `sensitive` | erhöhter Schutzbedarf (≠ secret) |
| `restricted` | eng begrenzter Kreis (≠ national-security classification) |
| `secret-bearing` | enthält/verweist Secret Values → an [CO-WP-024](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md) gebunden |
| `evidence-protected` | Audit-/Evidence-Bindung (≠ evidence accepted/sufficient) |
| `unknown-classification` | nicht bestimmt → fail-closed zur strengeren anwendbaren Grenze |

Je Klasse dokumentiert: Zweck · Definition · accountable owner · classification authority · zulässige Verbraucher · zulässige Zwecke · Storage Boundary · Processing Boundary · Export Boundary · Logging Boundary · Retention Expectation · Redaction Requirement · Offline Boundary · Backup Boundary · Review Trigger · Known Limitations.

```text
unknown-classification → fail-closed toward the stricter applicable handling boundary
public ≠ unrestricted modification
internal ≠ safe for external disclosure
sensitive ≠ secret
restricted ≠ national-security classification
secret-bearing data remains bound to CO-WP-024
evidence-protected ≠ evidence accepted or sufficient
```

## 7. Classification Identity and Freshness

Getrennt: data identity · content revision · classification decision · classification label · label version · classification scope · classification instance · reclassification event · declassification event · handling-policy version · assessment time · freshness boundary.

```text
same data identity ≠ same content revision
same bytes         ≠ same context or classification
classification inherited from container ≠ automatically valid
classification once assigned ≠ permanently current
reclassification   ≠ historical label erased
```

Materiale Änderungen an Inhalt, Herkunft, Zweck, Ziel, Workspace oder Disclosure Scope lösen **Reassessment** aus. Declassification ist eine explizite, autorisierte, auditierte Entscheidung.

## 8. Authority Model

Getrennt: data policy authority · classification authority · reclassification authority · data owner · data steward · data custodian · collection authority · processing authority · use authority · disclosure authority · export authority · retention authority · preservation-hold authority · hold-release authority · deletion authority · destruction authority · redaction authority · redaction-review authority · evidence-acceptance authority · audit authority.

```text
owner              ≠ unrestricted user
steward            ≠ export authority
custodian          ≠ classification authority
collection authority ≠ every later use authority
retention authority ≠ hold-release authority
deletion operator  ≠ deletion authority
redaction operator ≠ disclosure authority
local administrator ≠ unrestricted data authority
machine principal  ≠ human preservation or disclosure approval
```

Delegation bindet an: accountable human owner · action · data identity/class · workspace · environment · target · consumer · purpose · profile · policy version · time · maximum scope · review trigger · audit reference. Autorität bindet an [Policy/Approval/Execution (CO-WP-013)](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) — keine parallele Autorität.

## 9. Collection and Data Minimization

Definiert: collection purpose · collection authority · required data set · optional data set · prohibited data · minimization assessment · collection start · collection end · downstream purpose · reuse assessment · derived-data creation · unknown overcollection outcome.

```text
technically collectible ≠ authorized to collect
collected once          ≠ authorized for every later use
useful                  ≠ necessary
derived data            ≠ free of source restrictions
aggregation             ≠ anonymization
```

Keine Datenschutz-/Rechtskonformität behauptet. Detailliertes Minimization-/Redaction-Modell im [Companion Disclosure Policy](../security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md).

## 10. Handling Lifecycle

`proposed · collected/generated · received · registered · classified · quarantined · assessed · approved-for-bounded-use · active · copied · derived · exported · retention-pending · hold-active · deletion-pending · logically-deleted · purge-pending · destruction-pending · destroyed · superseded · archived · outcome-unknown`.

```text
received  ≠ trusted
registered ≠ classified
classified ≠ approved for every use
copied    ≠ independently authorized
exported  ≠ accepted by recipient
retention expired ≠ deleted
deletion initiated ≠ completed
destroyed ≠ historical evidence erased
unknown   ≠ absent or safe
```

Retention/Deletion-Details im [Companion Retention Policy](DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md).

## 11. Workspace, Environment and Target Isolation

Klassifikations- und Handling-Bindungen bleiben durch Collection, Import, Copy, Derivation, Export, Backup, Restore, Offline Transfer, Redaction und Evidence Return erhalten.

```text
data from workspace A ≠ usable in workspace B automatically
rename                ≠ target rebinding
same bytes            ≠ same disclosure authority
shared infrastructure ≠ shared data authority
derived data retains applicable source restrictions
cross-workspace disclosure requires explicit authorization
```

Bezug THR-035.

## 12. Secret and Key Boundary

Secret-bearing Data bleibt an [CO-WP-024](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md) gebunden: Secret Values sind keine gewöhnlichen Datenfelder; ein teilweise maskiertes Secret bleibt secret-bearing; Key Material wird nicht durch Hashing/Masking zu gewöhnlichen Daten umklassifiziert; Redaction ersetzt keine Secret Removal/Revocation/Rotation. Secret References dürfen klassifiziert/retained werden, ohne den Secret Value zu enthalten (Bezug THR-019, THR-020).

## 13. Profiles

`Standard`, `Hardened`, `Government` sind Governance-Profile (Dimensionen u. a.: Klassifikationstiefe · Review-Frequenz · Retention-Freshness · Hold-Freigabe · Rollentrennung · Disclosure-Freigabe · Redaction Review · Offline-Toleranz · Unknown-State-Verhalten · Evidence-Tiefe).

```text
Government profile ≠ government certification
Hardened profile   ≠ proven secure
Standard profile   ≠ unrestricted data handling
```

Keine gesetzlichen/regulatorischen Zuordnungen behauptet.

## 14. Failure and Unknown State

`classification unknown · classification conflict · stale classification · owner unknown · derived-data lineage incomplete · unauthorized disclosure suspected`.

```text
unknown          ≠ safe
unknown          ≠ absent
failure          ≠ no side effects
retry            ≠ automatically permitted
missing evidence ≠ operation did not occur
```

Fail-closed und Reassessment explizit.

## 15. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Data classification remains separate from deployment profiles and connectivity classes and claims no national-security classification.
2. Unknown-classification data is handled fail-closed toward the stricter applicable handling boundary.
3. A classification label is a governed statement with scope, version and freshness; it does not prove content or protection.
4. Classification inherited from a container is not automatically valid; material change triggers reassessment; reclassification preserves history.
5. Data ownership does not imply unrestricted access, use or export; owner, steward, custodian and the distinct authorities remain separate.
6. Collection authority does not imply unrestricted later use; useful ≠ necessary.
7. Derived, aggregated, cached and replicated data retain applicable source restrictions.
8. Secret-bearing data remains governed by CO-WP-024 and is not downgraded by masking or hashing.
9. Classification and handling bindings to workspace, environment and target are preserved and not replaced by rename.
10. Missing evidence does not prove that no operation occurred.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 16. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): Secret Exposure THR-019/THR-020; evidence to wrong recipient THR-018; offline export sensitive THR-025; stolen backup/evidence recipient THR-040; stale as current THR-013; manipulated telemetry THR-012; tenant/org boundary THR-035; insider THR-037; managed resource vs CoreOps THR-034; audit deletion/manipulation THR-016/THR-017. Keine erfundenen IDs; kein Parallelregister.

## 17. Technology Boundary

Nicht ausgewählt: Datenbank · Object Store · Filesystem · Data Lake/Warehouse · Suchindex · Archiv · DLP · CASB · SIEM · Logging-/Telemetrie-Plattform · Klassifikationsschema-/Label-Format · Scanner · OCR · ML/KI-Klassifikation · Cloud Provider. Alle bleiben `deferred`.

## 18. Compatibility

Konkretisiert DEC-P-08; konsistent mit [Data Ownership (CO-WP-016)](../architecture/DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [Audit/Evidence (CO-WP-018)](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md), [Telemetry Privacy (CO-WP-019)](../security/TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md), [Secrets (CO-WP-024)](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md), [Public Neutrality/Disclosure (CO-WP-005)](PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md). Additiv; keine bestehende Invariante geschwächt; kein Parallelmodell.

## 19. Open Questions

- Konkrete Klassifikationskriterien je Datendomäne (deferred, domänenabhängig).
- Verhältnis interner Klassifikationsklassen zu späteren regulatorischen Mappings (bewusst nicht vorgenommen).
- Automatisierte Klassifikationsunterstützung (deferred; keine ML/KI-Auswahl).

## 20. Next Decision

Nächstes lokal registriertes Work Package gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): `CO-WP-026 – Self-Protection, Degraded Modes and Recovery Mode`. Zuerst Nova Review von CO-WP-025, danach Human-Maintainer-Commit.
