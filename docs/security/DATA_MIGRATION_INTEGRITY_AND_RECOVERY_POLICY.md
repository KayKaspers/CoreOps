# CoreOps – Data Migration Integrity and Recovery Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation migration-integrity and recovery policy
> Implementation Status: Not implemented
> Migration Technology: Not selected
> Backup Technology: Not selected
> Recovery Automation: Not implemented
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-016 (docs-only / data-architecture, persistence and migration foundation)

## 1. Status

Technologieunabhängige Policy für **Migrationsautorität, Klassifikation, Integrität, Backup/Recovery, destruktive Migration, Concurrency und Fail-Closed**. Companion zu [DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md](../architecture/DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md) und [SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md](../architecture/SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md). Kein Migration-/Backup-Mechanismus, keine Recovery-Automation.

## 2. Purpose

Migration ist der gefährlichste Datenvorgang: destruktiv, teilweise, offline, oder mit sensiblen Identity-/Audit-Daten. Diese Policy legt fest, wer migrieren darf, wie Integrität nachgewiesen wird und warum `executed ≠ validated`, `backup exists ≠ restorable`, `restore completed ≠ service recovered` und `migration must not reactivate revoked authority`.

## 3. Scope

Authority Model · Migration Classification · Planning · Approval · Preflight · Backup/Recovery · Execution Boundary · Integrity Verification · Validation · Partial Migration · Mixed Versions · Destructive Migration · Audit/Evidence-Daten · Identity/Authorization-Daten · Offline Migration · Concurrency/Change Freeze · Failure/Unknown Outcome · Rollback/Forward Recovery · Closure · Fail-Closed.

## 4. Non-Goals

- Keine Migration-/Backup-/Recovery-Technologie, keine Locking-/Snapshot-/Transaction-/Cluster-Auswahl.
- Keine tatsächliche Migration, kein Runtime-Code.
- Keine Behauptung implementierter/validierter/zertifizierter Kontrollen; kein SLA.

## 5. Authority Model

Getrennt (aus [Ownership §6]): Data Owner · Steward · Storage Responsibility · Write Authority · **Migration Authority** · Retention Owner · Recovery Owner · Evidence Owner. `storage responsibility ≠ authoritative ownership`; `write authority ≠ migration authority`; **Migration Authority ist keine uneingeschränkte Write-Autorität** und ändert Ownership nicht. Migration unterliegt Policy/Approval/Execution Authorization (CO-WP-013).

## 6. Migration Classification

Migrationen werden nach Risiko klassifiziert (non-destructive additive · migration-required · semantic · **destructive/irreversible** §16). Klasse bestimmt Approval-, Recovery-, Validierungs- und Window-Anforderungen (§7–§16).

## 7. Planning

Migration Plan wie [Schema/Migration §12], mit accountable human owner, Scope, Change Classes, Backup-/Recovery-Erwartung, Integrity-/Validation-Checks, Rollback-/Forward-Recovery-Pfad, Approval- und Execution-Authorization-Referenz. `migration plan ≠ migration authorization`.

## 8. Approval

Migration mit Write-/destruktiver Wirkung benötigt Policy-Decision und — abhängig von Scope/Profil — Approval (CO-WP-013). `approved migration ≠ executed migration`. Machine-Principals können Migration anfragen, aber privilegierte/destruktive Migration nicht selbst genehmigen; menschliche Verantwortlichkeit bleibt nachvollziehbar.

## 9. Preflight

Wie [Schema/Migration §13]. Bei unbekannter Ownership/unklarem Scope/fehlender Recovery-Grenze **fail-closed für destruktive Migrationen**. Preflight-Fehler → `preflight-failed`, keine Ausführung.

## 10. Backup and Recovery Expectation

Backup-Erwartung dokumentiert: reference · scope · creation time · source state · schema version · integrity/validation state · restore suitability · retention owner · access boundary · audit reference.
```text
backup exists          ≠ backup valid
backup valid           ≠ backup restorable
restore completed      ≠ data integrity verified
data integrity verified ≠ service recovered
rollback plan          ≠ guaranteed rollback
```
Keine Backupsoftware/Speicherarchitektur ausgewählt (Bezug THR-033).

## 11. Execution Boundary

Migration erfolgt nur über autorisierte Execution-Pfade (CO-WP-013 §12); kein Modul/Adapter/Agent migriert außerhalb seiner Autorität; ein Migrationswerkzeug erhält keine Ownership/Audit-Manipulationsautorität.

## 12. Integrity Verification

Zu unterscheiden: format · schema · referential · semantic · ownership validation · record-count/scope reconciliation · integrity verification · application-level · operational verification. `migration tooling success ≠ data correctness`. Verifikation ist **scopegebunden** und wird nicht über die geprüften Daten generalisiert. Keine Prüfmethode ausgewählt.

## 13. Validation

`executed ≠ validated`. Validierung ist ein eigener Schritt (`validation-pending`/`validated`/`validation-failed`); ein `validation-failed` blockiert Closure als Erfolg.

## 14. Partial Migration

`partially-migrated` bleibt sichtbar; `partial ≠ complete failure ≠ complete success`. Betroffene/nicht betroffene Records referenziert; Teil-Rollback möglich, aber als solcher gekennzeichnet (Bezug THR-031).

## 15. Mixed Versions

Mixed-Version-Betrieb (§Schema-Model) benötigt explizite Compatibility- und Dauergrenzen; `new writer ≠ old consumer compatibility`; unbegrenzter Mixed-Version-Betrieb ist unzulässig.

## 16. Destructive Migration

Erhöhtes Risiko: field/record deletion · identifier replacement · ownership transfer · lossy type conversion · schema merge/split · retention reduction · evidence-history transformation · credential-reference transformation. Erfordert (scope-/profilabhängig): explizite Human Approval · aktualisierte Recovery-Bewertung · zusätzliche Validierung · Dry Run/Preview (sofern später technisch möglich) · Scope-Begrenzung · Maintenance/Change Window · Audit und Post-Migration Review. **Destruktive Migration erfordert explizite gebundene Autorität und anwendbare Approval.** Keine Dry-Run-Technologie ausgewählt.

## 17. Audit and Evidence Data

Normale Migration schreibt Historie **nicht** still um; Transformation erhält ursprüngliche Referenzen; Provenance bleibt sichtbar; Retention Owner bleibt verantwortlich; **Migration Authority erhält keine Audit-Manipulationsautorität**; gelöschte/zusammengeführte fachliche Datensätze orphanen erforderliche Evidenzreferenzen nicht. `append-only-governed` ist Governance-Anforderung, keine behauptete technische Unveränderlichkeit (Bezug THR-016, THR-017).

## 18. Identity and Authorization Data

Besonders sensibel: human identity references · memberships · role assignments · approval history · execution authorization · machine identity · credential references · break-glass records. Regeln:
- Migration erzeugt **keine** neue Autorität.
- Rollen/Memberships verlieren ihren Scope nicht.
- **Revoked Authority wird nicht reaktiviert; consumed Authorization wird nicht wiederverwendbar.**
- Credential Governance beansprucht weiterhin **keine** Roh-Secret-Speicherung (CO-WP-024 offen).
- Break-Glass-Historie bleibt zurechenbar (Bezug THR-002, THR-019, THR-020).

## 19. Offline Migration

Offline-Migrationspaket benötigt: source/target schema · target-environment binding · data-scope binding · migration-plan binding · authorization reference · provenance/integrity status · validity/usage boundary · import quarantine · local activation · partial execution · result package · reconciliation · delayed revocation · clock/sequence uncertainty. Nicht behauptet: implementiert, automatische konfliktfreie Migration, beliebige Air-Gap-Stufen, Klassifiziertnetz-Eignung, konkrete Signing-/Trust-Anchor-Technologie (Bezug THR-024). Fail-closed bei unklarer Provenance/Autorität.

## 20. Concurrency and Change Freeze

Migration kann mit Writes/Jobs kollidieren. Konzeptionell: Write Freeze/kontrollierte Write-Grenze · Maintenance Window · konkurrierende Workflows · laufende Execution-Pläne · verspätete Offline Writes · Snapshot-/Consistency Boundary (spätere technische Entscheidung) · Konflikterkennung · Reconciliation nach Migration. Keine Locking-/Snapshot-/Transaction-/Cluster-Technologie ausgewählt.

## 21. Failure and Unknown Outcome

`outcome-unknown` bleibt explizit; weder Erfolg noch sichere Nichtausführung; erfordert Reconciliation; kein automatischer Retry ohne Governance (konsistent mit [Integration Trust/Failure](INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md)). Migrationswerkzeug-Fehler ≠ Datenkorrektheit.

## 22. Rollback and Forward Recovery

`rolled-back ≠ original state restored unless verified`; `recovered ≠ full service recovery unless verified`. Beide benötigen Verifikationsevidenz (Bezug THR-032). Recovery Owner (§5) verantwortet den Nachweis.

## 23. Closure

`closed ≠ successful`. Closure benötigt Closure Reason · Final Outcome · Validation Status · Remaining Risk/Exception · Owner · Audit Reference (wie [Schema/Migration §14]).

## 24. Fail-Closed Rules

Keine destruktive/privilegierte Migration bei: unbekannter Ownership · fehlender Migration Authority/Approval · unklarem/unbounded Scope · unbekannter/`conflicted` Compatibility · fehlender Recovery-Grenze · fehlgeschlagenem Preflight · fehlender Provenance (offline) · unbekanntem Outcome · unkontrollierten konkurrierenden Writes · sensiblen Identity-/Audit-Daten ohne erhaltene Provenance/Zurechenbarkeit.

## 25. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Storage responsibility must not imply authoritative ownership.
2. Migration authority must not imply unrestricted write authority.
3. Schema version and data version must remain separate.
4. Unknown compatibility must block automatic destructive migration.
5. Executed migration must not imply validated data integrity.
6. Partial migration must remain visible.
7. Rollback claims require verification evidence.
8. Migration must not reactivate revoked authority or reusable consumed authorization.
9. Migration must preserve audit and evidence provenance.
10. Offline migration requires provenance, integrity, target binding and explicit activation.
11. Destructive migration requires explicit bounded authority and applicable approval.
12. Historical schema identity and migration evidence must not be deleted silently.

## 26. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-002, THR-013, THR-014, THR-016, THR-017, THR-019, THR-020, THR-021, THR-022, THR-023, THR-024, THR-026, THR-031, THR-032, THR-033, THR-035. Keine Duplikation, kein Parallelregister.

## 27. Technology Boundary

Nicht ausgewählt/implementiert: Migration-/Backup-/Recovery-Technologie, Locking/Snapshot/Transaction/Cluster, Verschlüsselung, DB/Storage, Runtime-Code.

## 28. Compatibility

Konsistent mit [Ownership/Persistence](../architecture/DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [Schema/Migration](../architecture/SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md), [Policy/Approval/Execution](POLICY_DECISION_AND_EVALUATION_MODEL.md), [Machine Identity](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [Offline Reconciliation](OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md), [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md). Konkretisiert DEC-P-04, DEC-G-04 (Backup vor Änderung, Verifikation), DEC-P-08.

## 29. Open Questions

- Konkrete Verifikations-/Reconciliation-Methoden (spätere ADR).
- Snapshot-/Consistency-Boundary-Mechanismus (CO-WP-021/026).
- Roh-Secret-Migration bleibt offen (CO-WP-024).

## 30. Next Decision

Self-Protection/Degraded/Recovery Mode (CO-WP-026), Deployment Control Plane (CO-WP-021) und Secrets/Key Custody (CO-WP-024) konkretisieren Recovery-/Integritätsaspekte. Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
