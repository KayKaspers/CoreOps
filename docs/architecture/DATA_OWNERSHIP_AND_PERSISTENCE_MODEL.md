# CoreOps – Data Ownership and Persistence Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation data-ownership and persistence model
> Implementation Status: Not implemented
> Storage Technology: Not selected
> Database Technology: Not selected
> Retention Validation: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-016 (docs-only / data-architecture, persistence and migration foundation)

## 1. Status

Technologieunabhängiges Modell für **Datenownership, Verantwortungsdimensionen, Datenklassen, Persistence-Klassen und Data Lifecycle** in CoreOps. Companion zu [SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md](SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md) und [DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md](../security/DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md). Es implementiert **keine** Datenbank, kein ORM, kein Schema, kein Storage-System.

## 2. Purpose

Wer Daten speichert, besitzt sie nicht. Dieses Modell trennt fachliche Ownership von Storage-Verantwortung und Write-/Migration-Autorität und ordnet jeder Datenklasse ein autoritatives Modul zu — damit `data owner ≠ storage operator`, `storage responsibility ≠ write authority` und `cached ≠ authoritative`. Es baut auf [Source of Truth](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) und dem [Module Catalog](COREOPS_MODULE_CATALOG.md) auf (kein Parallelmodell).

## 3. Scope

Ownership-Dimensionen · Datendomänen/-klassen · autoritative Module · Permitted Writers · Persistence-Klassen · Data Lifecycle · Retention/Archival · Recovery Ownership · Privacy/Disclosure · Offline-Betrachtungen · Audit/Evidence. Schema/Migration (Companion 2), Migration-Security (Companion 3).

## 4. Non-Goals

- Keine Datenbank-/SQL-/NoSQL-/ORM-Auswahl, kein konkretes Tabellen-/Dokumentenschema.
- Keine Dateiformate, kein Schema Registry, kein Migrationsframework, keine Backupsoftware.
- Keine Replikations-/Cluster-/Verschlüsselungstechnologie, kein Runtime-Code.
- Keine konkreten Retention-Fristen; keine Behauptung implementierter Persistenz.

## 5. Concepts

Begriffe (mindestens): data domain · data class · data owner · data steward · authoritative module · storage responsibility · write authority · migration authority · retention owner · schema · schema identity · schema version · data version · migration · migration plan/step/state · compatibility · validation · integrity verification · backup · restore · rollback · forward recovery · archive · purge.

**Grundregeln:**
```text
data owner            ≠ storage operator
storage responsibility ≠ write authority
write authority       ≠ migration authority
schema version        ≠ data version
migration executed    ≠ migration successful
migration successful  ≠ data integrity verified
backup created        ≠ backup restorable
restore completed     ≠ service recovered
archive               ≠ deletion
purge                 ≠ normal retention expiry
```

## 6. Ownership Dimensions

Getrennt dokumentiert: **Authoritative Data Owner · Data Steward · Storage Responsibility · Permitted Writers · Migration Authority · Retention Owner · Disclosure Owner · Recovery Owner · Evidence Owner**. Regeln:
- Pro kanonischem Datenkonzept existiert **eine** eindeutige autoritative Ownership-Regel.
- Storage übernimmt Ownership **nicht** still; Shared Storage erzeugt **keine** gemeinsame fachliche Autorität.
- Ein Adapter/Domain Pack wird **nicht** allein durch Speicherung zum Data Owner.
- Migration Authority ändert Datenownership **nicht**.
- Ownership-Wechsel benötigt explizite Governance-Entscheidung und Auditnachweis.

## 7. Data Domains and Classes

Je Klasse: authoritative module · owner · permitted writers · typical readers · persistence expectation · retention consideration · migration sensitivity · confidentiality · integrity requirement · offline behavior · recovery expectation · threat references. **Keine vollständige technische Feldliste.**

| Datenklasse | Autoritatives Modul (Konzept) | Persistence | Migration-Sensitivität | Threat refs |
|---|---|---|---|---|
| identity and membership data | MOD-IAM | persistent-authoritative | hoch | THR-002, THR-035 |
| policy and approval data | MOD-POL | append-only-governed / persistent-authoritative | hoch | THR-003, THR-002 |
| resource inventory | MOD-INV | persistent-authoritative | mittel | THR-014 |
| observed state | MOD-OBS | persistent-operational / cached | mittel | THR-012, THR-013 |
| desired state | MOD-STA | persistent-authoritative | hoch | THR-013 |
| effective and derived state | MOD-STA | derived-rebuildable | niedrig-mittel | THR-013 |
| topology data | MOD-TOP | derived-rebuildable / persistent-operational | mittel | THR-015 |
| workflow and job state | MOD-WFL | persistent-operational | mittel | THR-028 |
| execution plans and results | MOD-EXE | append-only-governed | hoch | THR-005, THR-029 |
| integration configuration | MOD-ADP | persistent-authoritative | mittel | THR-010 |
| machine identity and credential references | MOD-IAM/SEC | persistent-authoritative (Referenzen) | hoch | THR-019, THR-020 |
| Domain-Pack metadata | MOD-EXT | persistent-operational | mittel | THR-023 |
| audit records | MOD-EVD | append-only-governed | sehr hoch | THR-016, THR-017 |
| evidence references and metadata | MOD-EVD | append-only-governed | hoch | THR-016, THR-018 |
| notification state | MOD-NOT | persistent-operational / transient | niedrig | THR-018 |
| offline transfer and reconciliation metadata | MOD-OFF | persistent-operational | hoch | THR-024 |

(Modulzuordnungen konzeptionell, konsistent mit [Module Catalog](COREOPS_MODULE_CATALOG.md).)

## 8. Authoritative Modules

Jede Datenklasse hat **ein** autoritatives Modul (§7). Andere Module lesen über definierte Grenzen ([Module Boundary Standard](MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md)); kein Modul beansprucht Autorität über fremde Datenklassen durch Speicherung oder Caching.

## 9. Permitted Writers

Schreibrechte sind je Datenklasse explizit (Permitted Writers); Write ≠ Ownership. Ein Reader wird nicht durch Zugriff zum Writer; ein Writer wird nicht durch Schreiben zum Owner. Write-Operationen unterliegen Policy/Authorization (CO-WP-013).

## 10. Persistence Classes

`transient · session-bound · cached · persistent-operational · persistent-authoritative · append-only-governed · archived · derived-rebuildable · external-reference · unknown` (konzeptionelle Guidance, keine Enums).
```text
cached              ≠ authoritative
derived-rebuildable ≠ safe to delete without dependency review
append-only-governed ≠ technically immutable
archived            ≠ currently active
external-reference  ≠ externally available
unknown persistence ≠ safe for destructive migration
```
Keine Storage-Technologie ausgewählt.

## 11. Data Lifecycle

Konzeptionelle Statuswerte:
```text
created → active → updated → superseded → suspended → archival-pending → archived →
retention-review → purge-pending → purged → recovery-pending → restored → invalidated
```
Dokumentiert: Creation Authority · Mutation Authority · Supersession · Archival · Retention Review · Purge Approval · Recovery · Historical Evidence · Legal/organisatorische Holds (mögliche spätere Anforderung). **Keine konkreten Retention-Fristen.**

## 12. Retention and Archival

Retention Owner (§6) verantwortet Aufbewahrung/Review. `archive ≠ deletion`; `purge ≠ normal retention expiry`. Purge benötigt Approval; ein Hold kann Purge blockieren. Retention-Fristen sind spätere Governance-Entscheidungen (Bezug DEC-P-08, CO-WP-025).

## 13. Recovery Ownership

Recovery Owner (§6) verantwortet Wiederherstellbarkeit. `backup created ≠ restorable`; `restore completed ≠ service recovered` (Detail Companion 3 §10). Recovery-Erwartung ist je Datenklasse dokumentiert, nicht behauptet.

## 14. Privacy and Disclosure

Disclosure Owner (§6) verantwortet Offenlegungsgrenzen. Personenbezogene Daten unterliegen Datenminimierung/Redaction (Bezug DEC-P-08, [Public Neutrality](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md), THR-019/020/025). Detaillierte Klassifikation/Retention später (CO-WP-025).

## 15. Offline Considerations

Offline gespeicherte/replizierte Daten bleiben provenance-/scope-gebunden; `cached`/`external-reference` offline sind nicht autoritativ; Reconciliation nach Reconnection folgt der [Offline Reconciliation Policy](../security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md). Keine automatische konfliktfreie Synchronisation.

## 16. Audit and Evidence

Erfasst: ownership assignment/change · persistence classification · lifecycle transition · retention/purge decision · recovery event. `append-only-governed` ist Governance-Anforderung, **keine** behauptete technische Unveränderlichkeit; Audit-/Evidence-Historie wird nicht still umgeschrieben (Bezug THR-016/017).

## 17. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. Storage responsibility must not imply authoritative ownership.
2. Shared storage must not create shared authoritative authority.
3. Write authority is not migration authority; migration authority does not change ownership.
4. Cached or derived data is not authoritative.
5. `derived-rebuildable` is not safe to delete without dependency review.
6. `append-only-governed` is a governance requirement, not asserted technical immutability.
7. Unknown persistence is not safe for destructive migration.
8. Ownership changes require explicit governance and audit.
9. Audit and evidence provenance must not be silently rewritten.

## 18. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) (nur reale IDs, keine Duplikation, kein Parallelregister): THR-002, THR-003, THR-005, THR-010, THR-012, THR-013, THR-014, THR-015, THR-016, THR-017, THR-018, THR-019, THR-020, THR-023, THR-024, THR-028, THR-029, THR-035.

## 19. Technology Boundary

Nicht ausgewählt/implementiert: Datenbank/SQL/NoSQL, ORM, Tabellen-/Dokumentenschema, Dateiformate, Storage-/Replikations-/Cluster-Technologie, Verschlüsselung, Runtime-Code.

## 20. Compatibility

Konsistent mit [SoT/State Authority](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md), [State/Drift](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md), [Module Catalog](COREOPS_MODULE_CATALOG.md), [Integration Contract](COREOPS_INTEGRATION_CONTRACT_V0_1.md), [Domain Pack Governance](DOMAIN_PACK_GOVERNANCE_MODEL.md), [Identity/Authorization](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Konkretisiert DEC-P-04, DEC-S-107 (authoritative field ownership), DEC-P-08.

## 21. Open Questions

- Verbindliche Retention-Fristen/Holds (CO-WP-025).
- Genaue Persistence-Klassifikation je konkreter Datenklasse (spätere Iteration).
- Roh-Secret-Speicherung bleibt offen (CO-WP-024).

## 22. Next Decision

Companion 2 (Schema/Migration) und Companion 3 (Migration-Security) bauen auf dieser Ownership-Basis auf. Data Classification/Retention (CO-WP-025) und Secrets/Key Custody (CO-WP-024) konkretisieren einzelne Grenzen. Storage-/DB-Technologie bleibt einer späteren ADR-Runde vorbehalten.
