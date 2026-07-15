# CoreOps – Data Retention, Deletion and Preservation Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation data retention, deletion and preservation policy
> Implementation Status: Not implemented
> Retention Engine: Not selected
> Deletion/Sanitization Technology: Not selected
> Legal-Hold System: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Legal/Regulatory Mapping: None performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-025 (docs-only / data-governance, retention, minimization, redaction and controlled-disclosure foundation)

## 1. Status

Technologieunabhängige Policy für **Retention, Retention Start Events/Expiry, Preservation Holds, Kopien/Backups, Deletion/Purge/Destruction, Offline und Evidence/Unknown Outcomes**. Companion zu [DATA_CLASSIFICATION_AND_HANDLING_MODEL.md](DATA_CLASSIFICATION_AND_HANDLING_MODEL.md) und [REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md](../security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md). Konkretisiert die von [Audit/Evidence (CO-WP-018)](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md) an CO-WP-025 delegierten Retention-Fristen/Holds. Kein Retention-/Deletion-/Archiv-System implementiert.

## 2. Purpose

Retention und Löschung sind getrennte, autorisierte, nachweisbedürftige Vorgänge. Die Policy legt fail-closed-Regeln fest und warum `retention required ≠ indefinite retention`, `retention expired ≠ deletion completed`, `deletion requested ≠ deletion authorized`, `deletion executed ≠ every copy removed`, `logical deletion ≠ physical destruction` und `unknown deletion outcome ≠ deleted`.

## 3. Scope

Authority Model · Retention Policy · Start Events/Expiry · Preservation Holds · Copies/Caches/Replicas/Derived · Backup/Recovery Boundary · Deletion/Purge/Destruction · Deletion Evidence/Verification · Offline/CorePack/Evidence Return · Evidence Retention Boundary · Workspace Isolation · Failure/Fail-Closed.

## 4. Non-Goals

- Keine Retention-/Deletion-/Purge-/Archiv-/Backup-/DLP-/Discovery-/Legal-Hold-/Media-Sanitization-/Secure-Erase-Technologie; keine konkrete pauschale Frist.
- Keine gesetzliche/regulatorische Wirksamkeit; keine Legal-Compliance-/Zertifizierungs-Behauptung; kein Runtime-Code; keine Löschung realer Daten.

## 5. Concepts

`retention policy` · `retention period` · `retention start event` · `retention expiry` · `minimum/maximum retention boundary` · `preservation hold` · `hold release` · `deletion request/assessment/authorization` · `logical deletion` · `access withdrawal` · `purge` · `destruction` · `verification` · `reconciliation` · `primary data` · `replica` · `cache` · `snapshot` · `backup` · `archive` · `recovery copy` · `evidence copy` · `deletion evidence` · `outcome-unknown`.

## 6. Authority Model

Getrennt: retention authority · preservation-hold authority · hold-release authority · deletion authority · destruction authority · deletion operator. `retention authority ≠ hold-release authority`; `deletion operator ≠ deletion authority`; `backup operator ≠ recovery/deletion authority`; `machine principal ≠ human preservation approval`. Bindet an die Daten-Autoritäten des [Classification Model §8](DATA_CLASSIFICATION_AND_HANDLING_MODEL.md) und an [CO-WP-013](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md).

## 7. Retention Policy

Mindestattribute: policy identity · policy version · data class · data identity/scope · owner · purpose · retention start event · duration/bounded condition · minimum retention boundary · maximum retention boundary · review trigger · preservation-hold interaction · backup interaction · offline interaction · deletion/disposition action · evidence requirement · exceptions · known limitations.

```text
retention period starts only from a defined event
indefinite retention requires explicit authority and review
retention extension ≠ silent default
retention expiry    ≠ deletion authorization
retention policy present ≠ applied to every copy
```

Keine pauschale Dauer für alle Daten.

## 8. Retention Start Events and Expiry

Retention beginnt ausschließlich an einem definierten Start Event (z. B. creation, collection, last-use, case-close — konzeptionell, nicht fristfestlegend). Expiry markiert das Ende des Pflichtzeitraums, **nicht** eine Löschbestätigung: `retention expiry ≠ deletion authorization ≠ deletion completion`. Unbekannter Start/Expiry bleibt sichtbar und fail-closed.

## 9. Preservation Holds

Neutraler Begriff `preservation hold` (keine rechtliche Wirksamkeit/Legal-Compliance behauptet). Mindestens: hold identity · authority · reason category · affected data · scope · start · validity · review interval · permitted use · prohibited deletion · permitted redaction · export boundary · release authority · release condition · evidence · unknown-copy handling.

```text
hold active         ≠ unrestricted access
hold active         ≠ data trustworthy
hold release        ≠ immediate deletion
local hold          ≠ globally observed hold
missing hold entry  ≠ no central hold exists
```

Preservation Hold und Retention sind getrennte Autoritäten; ein Hold überschreibt Retention Expiry (keine Löschung während aktivem Hold).

## 10. Copies, Caches, Replicas and Derived Data

Retention, Holds und Deletion beziehen sich auf **alle** Instanzen: primary · replica · cache · snapshot · backup · archive · recovery copy · evidence copy · derived data. `retention policy present ≠ applied to every copy`; `derived data retains applicable source restrictions`. Unbekannte Copy-Inventarabdeckung bleibt explizit (§17).

## 11. Backup and Recovery Boundary

Getrennt: primary data · replica · cache · snapshot · backup · archive · recovery copy · evidence copy.

```text
primary deletion  ≠ backup deletion
backup retention  ≠ source retention
backup available  ≠ restore authorized
restore completed ≠ restored classification current
historical backup ≠ current disclosure authorization
```

Retention/Holds berücksichtigen Backup-/Recovery-Kopien ohne Auswahl einer Backup-Technologie ([CO-WP-024 Backup/Recovery §15](../security/KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md); Bezug THR-033, THR-040). Unknown Backup Coverage bleibt explizit.

## 12. Deletion, Purge and Destruction

Getrennt: deletion request · deletion assessment · deletion authorization · logical deletion · access withdrawal · purge request · purge execution · destruction request · destruction execution · verification · reconciliation · outcome-unknown.

```text
deletion requested   ≠ authorized
authorized           ≠ executed
primary record deleted ≠ all replicas deleted
cache invalidated    ≠ data absent
backup expired       ≠ backup destroyed
destruction reported ≠ destruction verified
unknown deletion outcome ≠ deleted
logical deletion     ≠ physical destruction
```

Keine konkrete Lösch-/Media-Sanitization-Technologie; keine „secure erase"-Behauptung ohne ausdrücklich definierte Evidenz.

## 13. Deletion Evidence and Verification

Deletion/Destruction benötigen Evidenz: affected data identity/scope · authority · reason · executed action · verification result · residual copies assessment · reconciliation state · audit reference. `destruction reported ≠ verified`; `deletion evidence available ≠ every copy covered`. Verifikation ist getrennt von der Ausführung.

## 14. Offline, CorePack and Evidence Return

Bindet an [CO-WP-023](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md): Offline-Kopien behalten Workspace-/Environment-/Target-/Classification-Bindings. `offline copy available ≠ current use authorized`; lokale Retention Expiry ≠ zentral autorisierte Löschung; lokale Holds können zentral unbekannt/stale sein; Offline Deletion/Redaction benötigt Evidence Return + Reconciliation; fehlende Rückmeldung ≠ keine Kopie/Verarbeitung. Offline-Betrieb erweitert **keine** Retention-/Deletion-/Disclosure-Autorität (Bezug DEC-S-305/314; THR-024). Keine Offline-/Sync-Technologie.

## 15. Evidence Retention Boundary

Bindet an [CO-WP-018](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md): `evidence retained ≠ all source data retainable`; `source deletion ≠ evidence deletion automatically`. Evidence-Retention und Source-Data-Retention sind getrennte Entscheidungen; eine Preservation-Hold- oder Retention-Pflicht auf Evidence begründet keine unbegrenzte Source-Retention und umgekehrt.

## 16. Workspace, Environment and Target Isolation

Retention-/Hold-/Deletion-Bindungen bleiben workspace-/environment-/target-gebunden; `data from workspace A ≠ deletable under workspace B authority`; geteilte Infrastruktur erzeugt keine gemeinsame Deletion Authority (Bezug THR-035).

## 17. Failure and Unknown State

`retention start unknown · retention expiry uncertain · hold state unknown · copy inventory incomplete · deletion outcome unknown · purge outcome unknown · destruction evidence missing · backup coverage unknown · offline reconciliation pending`.

```text
unknown          ≠ safe
unknown          ≠ absent
failure          ≠ no side effects
retry            ≠ automatically permitted
missing evidence ≠ operation did not occur
```

## 18. Fail-Closed Rules

Keine Löschung/Purge/Destruction bei: aktivem Preservation Hold · fehlender Deletion Authorization · unklarer Copy-Inventarabdeckung für destruktive Löschung · unbekannter Backup Coverage · unbekanntem Offline-Reconciliation-Stand · fehlendem Audit-Start. Unbekannter Deletion-/Redaction-Ausgang blockiert unsichere Wiederholung und erfordert Reconciliation.

## 19. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Retention starts only from a defined start event; indefinite retention requires explicit authority and review.
2. Retention expiry implies neither deletion authorization nor deletion completion.
3. Preservation hold and retention are separate authorities; an active hold blocks deletion.
4. Deletion request, assessment, authorization, execution and verification remain separate states.
5. Primary deletion does not imply deletion of replicas, caches, backups, archives or derived data.
6. Logical deletion is not physical destruction; destruction reported is not destruction verified.
7. Unknown deletion, purge or destruction outcomes remain visible and block unsafe retry.
8. Retention, holds and deletion apply to all copies; unknown copy/backup coverage remains explicit.
9. Evidence retention is separate from source-data retention.
10. Offline operation expands no retention, deletion or disclosure authority.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 20. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): audit deletion/manipulation THR-016/THR-017; backup/restore manipulated THR-033; stolen backup/evidence recipient THR-040; offline import THR-024; offline export sensitive THR-025; stale as current THR-013; tenant/org boundary THR-035; insider THR-037. Keine erfundenen IDs; kein Parallelregister.

## 21. Technology Boundary

Nicht ausgewählt: Retention-/Deletion-/Purge-/Archiv-/Backup-/DLP-/Discovery-/Legal-Hold-/Media-Sanitization-/Secure-Erase-/Workflow-/Policy-Engine · Datenformat · Cloud Provider. Alle bleiben `deferred`.

## 22. Compatibility

Konkretisiert DEC-P-08 und die Retention-/Hold-Offene-Fragen aus [CO-WP-018](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md); konsistent mit [Data Ownership (CO-WP-016)](../architecture/DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [Secrets/Key (CO-WP-024)](../security/KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md), [Offline/CorePack (CO-WP-023)](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md). Additiv; keine bestehende Invariante geschwächt.

## 23. Open Questions

- Konkrete Retention-Fristen je Datenklasse/Domäne (deferred, bewusst nicht festgelegt).
- Zuverlässige Copy-/Backup-Inventarabdeckung für Löschung (deferred).
- Verhältnis interner Preservation Holds zu späteren rechtlichen Anforderungen (bewusst nicht gemappt).

## 24. Next Decision

Nächstes lokal registriertes Work Package gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): `CO-WP-026 – Self-Protection, Degraded Modes and Recovery Mode`. Zuerst Nova Review von CO-WP-025, danach Human-Maintainer-Commit.
