# CoreOps – Configuration Source of Truth and Secret Reference Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation configuration source-of-truth and secret-reference model
> Implementation Status: Not implemented
> Configuration Store: Not selected
> Configuration Format: Not selected
> Merge Engine: Not selected
> Secret Store: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-024 (docs-only / security, configuration, secret-lifecycle and key-custody governance foundation)

## 1. Status

Technologieunabhängiges Modell für **Configuration Source of Truth, Configuration-Versionen/Overlays, sensitive configuration, Secret References, Resolution, Drift, Target Binding sowie Deployment-/CorePack-Integration und Reconciliation**. Companion zu [SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md) und [KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md](../security/KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md). Baut auf [MOD-STA-001 Configuration and State](COREOPS_MODULE_CATALOG.md) und [Source of Truth and State Authority (CO-WP-011)](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md) auf; kein Parallelmodell. Kein Configuration Store/Format, keine Merge Engine, kein Secret Store implementiert.

## 2. Purpose

Konfiguration und Secret References verbinden Governance mit Laufzeit, ohne Secret Values zu tragen. Das Modell legt fest, warum `configuration source of truth ≠ runtime state`, `configuration validation ≠ successful application`, `configuration applied ≠ desired state verified`, `reference present ≠ secret exists`, `reference resolved ≠ use authorized` und `drift detected ≠ drift automatically remediated`.

## 3. Scope

Configuration Source of Truth · Versions/Revisions/Overlays · Sensitive Configuration · Secret Reference Model · Reference Resolution · Runtime-Effective Configuration/Drift · Reconciliation · Target/Workspace/Environment Binding · Deployment/Blueprint Integration · CorePack Integration · Offline Configuration/Reference Sets · Audit · Failure/Fail-Closed.

## 4. Non-Goals

- Kein Configuration-Format/-Store, keine Merge-/Reconciliation-Engine, keine Schema-Sprache, kein Secret Store, kein Manifestformat, kein Policy-Engine-Produkt.
- Keine Behauptung implementierter/validierter Konfigurations- oder Resolution-Kontrollen; kein Runtime-Code; keine Compliance-Behauptung; keine Secret Values im Modell.

## 5. Concepts

`authoritative configuration source` · `configuration schema reference` · `configuration version` · `configuration revision` · `default configuration` · `profile configuration` · `environment overlay` · `workspace overlay` · `target-specific override` · `runtime-effective configuration` · `locally cached snapshot` · `pending change` · `applied change` · `rejected change` · `drift` · `reconciliation state` · `unknown effective state` · `secret reference` · `sensitive configuration` · `resolution constraint`.

## 6. Configuration Source of Truth

Die autoritative Konfigurationsquelle ist explizit, owner-/module-gebunden (MOD-STA-001) und getrennt von der Laufzeitrealität.

```text
default            ≠ approved effective configuration
runtime state      ≠ authoritative configuration
local snapshot     ≠ centrally current
configuration valid ≠ secret references resolved
```

Konsistent mit [CO-WP-011](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md): `source of truth ≠ system of record`; kein stiller Last-Write-Wins; Konflikte bleiben sichtbar.

## 7. Configuration Versions, Revisions and Overlays

Getrennt: configuration version · configuration revision · default configuration · profile configuration · environment overlay · workspace overlay · target-specific override · runtime-effective configuration. `overlay ≠ unrestricted override`. Prioritäts-/Konfliktregeln sind konzeptionell definiert (Default < Profile < Environment < Workspace < Target-Override), erhalten Provenance und schwächen Security-Anforderungen nicht still. Kein Configuration-Format/keine Merge Engine ausgewählt.

## 8. Sensitive Configuration

`sensitive configuration` und `secret configuration` sind von ordinary configuration getrennt (Klassifikation im [Companion Governance §6](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md)). Sensitive/secret configuration verweist über Secret References; sie trägt **keine** Secret Values. `secret ≠ ordinary configuration`.

## 9. Secret Reference Model

Konfigurationen, Blueprints, Domain Packs, CorePacks, Logs, Events und Evidence tragen **keine** Secret Values, sondern References mit: logical secret identity · required version/resolution constraint · workspace · environment · target · consumer identity · purpose · expected classification · required trust state · validity boundary · fallback prohibition/policy · unresolved behavior · audit reference.

```text
reference present  ≠ secret exists
secret exists      ≠ consumer authorized
reference resolved ≠ use authorized
resolution success ≠ target compatible
fallback to another secret ≠ permitted
```

## 10. Secret Reference Resolution

Resolution bindet eine Reference an eine konkrete, bewertete Secret-/Credential-Revision (nicht an einen Mutable Alias). Sie prüft Workspace/Environment/Target/Consumer/Classification/Trust-State/Validity gegen die Reference. Unaufgelöste, mehrdeutige oder konfliktbehaftete Secret References werden **fail-closed** behandelt (kein Fallback auf ein anderes Secret). Resolution ≠ Use Authorization ([Companion Governance §11](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md)).

## 11. Runtime-Effective Configuration and Drift

Getrennt: pending change · applied change · rejected change · runtime-effective configuration · drift · unknown effective state.

```text
configuration applied ≠ workload healthy
runtime state         ≠ authoritative configuration
drift detected        ≠ drift automatically remediated
```

Konsistent mit [Drift Detection (CO-WP-012)](DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md): Detection ohne Write Authority; keine automatische Remediation im MVP. Drift, der Secret Resolution verändert, bleibt sichtbar und blockiert unsafe automatische Nutzung.

## 12. Reconciliation

Reconciliation gleicht runtime-effective mit autoritativer Konfiguration und mit Secret-/Revocation-Snapshots ab. Unknown Effective State und unaufgelöste References lösen explizite Reconciliation aus, nicht automatischen Retry. `local snapshot ≠ centrally current`. Keine Reconciliation-Engine ausgewählt.

## 13. Target, Workspace and Environment Binding

Configuration- und Secret-Reference-Bindungen bleiben durch Resolution, Distribution und Application erhalten. Environment Binding nicht durch Rename ersetzbar; ein Config-/Reference-Set für Environment A gilt nicht automatisch für B; geteilte Infrastruktur erzeugt keine gemeinsame Configuration/Secret Authority; Cross-Workspace-/Cross-Environment-Use ist explizit governed (Bezug THR-035).

## 14. Deployment and Blueprint Integration

Bindet an [Deployment (CO-WP-021)](DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md)/[Blueprint (CO-WP-021)](DEPLOYMENT_BLUEPRINT_VERSIONING_AND_COMPATIBILITY_MODEL.md): Deployment Blueprints tragen nur Secret Requirements · Secret References · Consumer Bindings · Validity Requirements · Failure Behavior — **nicht** die Secret Values.

```text
deployment approved ≠ secret use approved
deployment execution authorization ≠ secret export authorization
```

Rollback bewertet Secret-/Credential-/Key-Zustände neu; kein Reaktivieren widerrufener/abgelaufener Secret-Versionen ([Companion Key Policy §16](../security/KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md)).

## 15. CorePack Integration

Bindet an [CorePack (CO-WP-023)](COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md): CorePacks enthalten keine Raw Secrets/wiederverwendbaren Credentials — nur Secret References/Requirements/Governance-Metadaten.

```text
CorePack trust    ≠ secret trust
CorePack activated ≠ secret retrieval/use authorized
```

Secret Resolution innerhalb eines CorePack-Kontexts erweitert Target Binding/Workspace Isolation nicht; Evidence-Return-CorePacks exportieren keine Secret Values.

## 16. Offline Configuration and Secret Reference Sets

Bindet an [CO-WP-023](RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md): offline configuration snapshot · offline secret reference set · locally available secret subset · local policy snapshot · revocation snapshot · freshness boundary · clock uncertainty · evidence return · reconciliation.

```text
central unavailable    ≠ local authority expands
offline cache available ≠ use currently authorized
no local revocation entry ≠ centrally valid
local clock recent     ≠ central policy current
```

Offline Secret Authorizations sind action-/secret-/target-/consumer-/version-/zeitgebunden und nicht wiederverwendbar. Keine Offline-Sync-Technologie.

## 17. Audit and Evidence

Konfigurations- und Resolution-Ereignisse sind audit-relevant (config change/approval/apply/reject · reference resolution · drift · reconciliation · offline import), enthalten aber **keine** Secret Values — nur sichere Metadaten und Secret-Identity/Version-Referenzen (MOD-EVD-001, [Companion Governance §15](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md)). `missing audit event ≠ no change occurred`.

## 18. Failure and Unknown State

`secret reference unresolved · ambiguous resolution · configuration rejected · pending change stuck · runtime unknown effective state · drift undetermined · reconciliation conflicted · target/workspace mismatch · offline freshness unknown`.

```text
unknown          ≠ safe
failure          ≠ no side effects
retry            ≠ automatically permitted
missing evidence ≠ operation did not occur
```

## 19. Fail-Closed Rules

Keine Anwendung/privilegierte Nutzung bei: unresolved/ambiguous/conflicted Secret Reference · Configuration-Konflikt ohne Auflösung · Target-/Workspace-/Environment-Mismatch · unbekanntem runtime-effective state für privilegierte Aktion · unbekannter Offline-Freshness ohne Exception · fehlendem Audit-Start. Kein stiller Fallback auf ein anderes Secret oder eine andere Environment-Bindung.

## 20. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Configuration source of truth remains separate from runtime state; applied configuration does not imply verified desired state.
2. Secret references remain separate from secret values; governed objects carry references, not raw secrets.
3. Reference presence/resolution does not imply secret existence, consumer authorization, target compatibility or use authorization.
4. Overlays do not act as unrestricted overrides and do not silently weaken security.
5. Unresolved, ambiguous or conflicted secret references are handled fail-closed with no fallback to another secret.
6. Drift that changes secret resolution remains visible and blocks unsafe automatic use; drift detection implies no automatic remediation.
7. Configuration and secret-reference bindings to workspace, environment and target are preserved and not replaced by rename.
8. CorePack and deployment authorization do not imply secret retrieval or use authorization.
9. Offline operation does not expand authority; offline secret authorizations are bounded and non-reusable.
10. Audit and evidence contain no secret values.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 21. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): Secret Exposure THR-019/THR-020; stale as current THR-013; queue/executed misrepresentation THR-028/THR-029; offline import THR-024; replay THR-026; false time THR-027; tenant boundary THR-035; managed resource vs CoreOps THR-034; audit manipulation THR-016/THR-017. Keine erfundenen IDs; kein Parallelregister.

## 22. Technology Boundary

Nicht ausgewählt: Configuration-Format/-Store · Merge-/Reconciliation-Engine · Schema-Sprache · Secret Store · Manifestformat · Policy Engine. Alle bleiben `deferred`.

## 23. Open Questions

- Genaue Prioritäts-/Konfliktauflösung bei überlappenden Overlays ohne festgelegtes Config-Format (deferred).
- Zusammenspiel von Drift-Remediation und Secret-Rotation (Companion Key Policy, deferred).
- Offline-Reconciliation-Verfahren für Secret Reference Sets (deferred).

## 24. Next Decision

Nächstes lokal registriertes Work Package gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): `CO-WP-025 – Data Classification, Retention and Redaction`. Zuerst Nova Review von CO-WP-024, danach Human-Maintainer-Commit.
