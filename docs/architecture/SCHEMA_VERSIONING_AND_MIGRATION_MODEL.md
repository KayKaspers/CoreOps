# CoreOps – Schema Versioning and Migration Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation schema-versioning and migration model
> Implementation Status: Not implemented
> Schema Format: Not selected
> Migration Framework: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-016 (docs-only / data-architecture, persistence and migration foundation)

## 1. Status

Technologieunabhängiges Modell für **Schema-Identität, Schema-Versionierung, Kompatibilität, Change Classes und den Migrations-Lebenszyklus**. Companion zu [DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md](DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md) und [DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md](../security/DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md). Es wählt **kein** Schemaformat, kein Migrationsframework, keine Versionierungsnotation.

## 2. Purpose

Schema-Versionen und Datenversionen werden häufig verwechselt; „Migration ausgeführt" wird häufig mit „Daten korrekt" verwechselt. Dieses Modell trennt Schema-Identität von Version, Version von Datenversion und Migration Execution von Validation — damit `schema version ≠ data version`, `migration executed ≠ validated`, `read-compatible ≠ write-compatible` und `unknown compatibility ≠ safe`.

## 3. Scope

Schema Identity · Version Dimensions · Compatibility Classes · Change Classes · Producer/Consumer Compatibility · Migration Requirements · Migration Plan · Preflight · Migration Lifecycle · Mixed-Version Operation · Partial Migration · Rollback/Forward Recovery · Deprecation · Offline Migration · Audit/Evidence.

## 4. Non-Goals

- Keine Schema-Sprache/-Notation (JSON Schema/Protobuf/Avro/XML/SQL DDL), kein Schema Registry.
- Kein Migrationsframework, keine Versionierungsnotation (SemVer o. ä.) erzwungen.
- Keine Serialisierung, keine DB-/Storage-Auswahl, kein Runtime-Code.
- Keine Behauptung durchgeführter Migration/Validierung.

## 5. Concepts

Aufbauend auf [Ownership/Persistence §5]. Zentral: `schema` (strukturelle/Semantik-Definition einer Datenklasse), `schema identity` (stabil), `schema version`, `data version` (Version eines Datensatzes/Bestands), `migration` (kontrollierter Übergang zwischen Schema-/Datenversionen), `compatibility`, `validation`, `integrity verification`.
```text
schema version       ≠ data version
migration executed   ≠ migration successful
migration successful ≠ data integrity verified
migration available  ≠ migration validated
```

## 6. Schema Identity

Ein Schema benötigt mindestens: stable schema identity · canonical name · owning module · data domain · schema version · compatibility classification · producer scope · consumer scope · read support state · write support state · migration requirement · validation state · evidence reference · deprecation state.
```text
display name      ≠ schema identity
file name         ≠ schema identity
storage location  ≠ schema identity
new schema version ≠ automatic migration requirement for every consumer
```
**Keine vollständige Schema-ID-Liste.** Owning module folgt der [Ownership](DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md).

## 7. Version Dimensions

Getrennt: `schema version · data-record version · producer version · consumer version · migration version · contract version · CoreOps product version · Domain-Pack version`. Eine Version wird **nicht** still für eine andere Dimension verwendet. Je Version dokumentiert: Identität · Source · Owner · Scope · Compatibility · Read Support · Write Support · Migration Path · Deprecation · Evidence. **Keine Versionierungsnotation erzwungen** (DEC-O-02 offen).

## 8. Compatibility Classes

`not-assessed · backward-readable · forward-readable-with-notes · read-compatible · write-compatible · round-trip-compatible · migration-required · incompatible · deprecated-compatibility · unknown · conflicted`.
```text
read-compatible   ≠ write-compatible
write-compatible  ≠ round-trip-compatible
backward-readable ≠ lossless downgrade
migration available ≠ migration validated
unknown           ≠ compatible
conflicted        ≠ safe for automatic migration
```
Compatibility Claims sind an Schema-, Producer-, Consumer- und Datenversionen gebunden (konsistent mit [Domain Pack Compatibility](DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md)).

## 9. Change Classes

`additive optional · additive required · semantic · constraint · field rename · field split · field merge · type change · identifier change · ownership change · retention change · deletion or destructive`. Je Change Class: Compatibility-Auswirkung · Migration Requirement · Rollback-/Forward-Recovery-Auswirkung · Evidence Requirement · Security-/Privacy-Auswirkung.

| Change Class | Typ. Compatibility | Migration | Destruktiv |
|---|---|---|---|
| additive optional | read-/write-compatible | selten erforderlich | nein |
| additive required | migration-required (Consumer) | ja | nein |
| semantic | conflicted/unknown möglich | ja + Validierung | evtl. |
| constraint | migration-required | ja | evtl. |
| field rename | migration-required | ja | evtl. (mapping) |
| field split/merge | migration-required | ja | ja (lossy möglich) |
| type change | migration-required | ja | ja (lossy möglich) |
| identifier change | incompatible/migration-required | ja | ja (hohes Risiko) |
| ownership change | governance-relevant | ja (Governance) | ja |
| retention change | governance-relevant | evtl. | ja (Datenverlust möglich) |
| deletion/destructive | incompatible | ja | ja (hohes Risiko) |

## 10. Producer and Consumer Compatibility

Producer (schreibt) und Consumer (liest) haben getrennte Kompatibilität. `new writer ≠ old consumer compatibility`; ein neuer Producer darf nicht Daten erzeugen, die ältere Consumer nicht sicher lesen, ohne dokumentierte Migration/Grenze. Read Support ≠ Write Support je Version.

## 11. Migration Requirements

Eine neue Schema-Version erzeugt **nicht automatisch** eine Migration-Pflicht für jeden Consumer (§6). Migration ist erforderlich, wenn Change Class + Consumer/Producer-Scope es verlangen (§9/§10). Migration Requirement ist je Schema/Version dokumentiert.

## 12. Migration Plan

Ein Migration Plan benötigt mindestens: migration identity · owner · accountable human owner · source schema/version · target schema/version · affected data domains · affected modules · affected workspaces/environments · affected records/scope · change classes · preconditions · backup/recovery expectation · integrity checks · validation checks · downtime/availability expectation · offline considerations · rollback feasibility · forward-recovery path · approval requirement · execution authorization reference · audit reference.
```text
migration plan       ≠ migration authorization
approved migration   ≠ executed migration
executed migration   ≠ successful migration
successful migration ≠ verified data integrity
```

## 13. Preflight

Vor Migration mindestens geprüft: source/target schema known · ownership known · write authority known · migration authority valid · affected scope bounded · dependencies identified · compatibility assessed · backup/recovery path assessed · capacity (konzeptionell) · conflicting operations controlled · offline nodes identified · retention/privacy constraints considered · audit start possible · verification plan present · rollback/forward-recovery plan present. **Bei unbekannter Ownership, unklarem Scope oder fehlender Recovery-Grenze gilt fail-closed für destruktive Migrationen.**

## 14. Migration Lifecycle

Konzeptionelle Statuswerte:
```text
proposed → analysis-pending → plan-ready → approval-pending → approved → scheduled →
preflight-pending → preflight-failed → execution-pending → executing → partially-migrated →
executed → validation-pending → validated → validation-failed → rollback-pending → rolled-back →
forward-recovery-pending → recovered → outcome-unknown → suspended → cancelled → closed
```
```text
executed   ≠ validated
rolled-back ≠ original state restored unless verified
recovered  ≠ full service recovery unless verified
closed     ≠ successful
```
`closed` benötigt: Closure Reason · Final Outcome · Validation Status · Remaining Risk/Exception · Owner · Audit Reference.

## 15. Mixed-Version Operation

Während/nach einer Migration können Producer/Consumer/Module/Offline-Nodes unterschiedliche Versionen führen. `mixed-version operation ≠ automatically safe`; `schema migrated ≠ all data migrated`. Mixed-Version-Betrieb benötigt **explizite Compatibility- und Dauergrenzen** (§8) und wird auditiert.

## 16. Partial Migration

`partial migration ≠ complete failure ≠ complete success`. Eine teilweise migrierte Menge bleibt **sichtbar** (`partially-migrated`); betroffene/nicht betroffene Records werden referenziert; Rollback ist evtl. nur für einen Teil möglich (Detail Companion 3 §14).

## 17. Rollback and Forward Recovery

Rollback (zurück zur Quellversion) und Forward Recovery (vorwärts zu einem konsistenten Zielzustand) sind getrennt. `rolled-back ≠ restored unless verified`; `recovered ≠ verified`. Beide benötigen Verifikationsevidenz (Companion 3 §22, Bezug THR-032).

## 18. Deprecation

Deprecated Schema-Versionen erhalten Migrations-/End-of-Support-Informationen; `deprecated ≠ immediately removed`. Historische Schema-Identität/Evidenz bleibt erhalten (nicht still gelöscht).

## 19. Offline Migration

Offline-Migration folgt Companion 3 §19: target-environment/data-scope/migration-plan-binding, Provenance, Integrität, Import-Quarantäne, explizite lokale Aktivierung, Result-Reconciliation, delayed revocation, clock/sequence uncertainty. Nicht behauptet: implementiert, automatische konfliktfreie Migration, beliebige Air-Gap-Stufen, Klassifiziertnetz-Eignung, konkrete Signing-/Trust-Anchor-Technologie (Bezug THR-024).

## 20. Audit and Evidence

Erfasst: schema creation/version · compatibility classification · change class · migration plan/lifecycle transition · validation activity · evidence reference. `migration tooling success ≠ data correctness`; Validierung scopegebunden. Audit-/Evidence-Historie nicht still umgeschrieben (Bezug THR-016/017).

## 21. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Schema identity is stable and independent of display name, file name or storage location.
2. Schema version and data version remain separate; no version is silently reused for another dimension.
3. A new schema version does not automatically require migration for every consumer.
4. Read-, write- and round-trip-compatibility remain separate.
5. Unknown or conflicted compatibility blocks automatic (destructive) migration.
6. Executed migration does not imply validated data integrity.
7. Partial migration remains visible; mixed-version operation is explicitly bounded.
8. Rollback/forward-recovery claims require verification evidence.
9. Historical schema identity and migration evidence must not be deleted silently.

## 22. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md): THR-013, THR-014, THR-016, THR-017, THR-021, THR-022, THR-023, THR-024, THR-031, THR-032. Keine Duplikation, kein Parallelregister.

## 23. Technology Boundary

Nicht ausgewählt/implementiert: Schemaformat/-sprache, Schema Registry, Migrationsframework, Versionierungsnotation, Serialisierung, DB/Storage, Runtime-Code.

## 24. Compatibility

Konsistent mit [Ownership/Persistence](DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [SoT/Provenance](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Integration Contract Versioning](COREOPS_INTEGRATION_CONTRACT_V0_1.md), [Domain Pack Compatibility](DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md), [Migration Security Policy](../security/DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md). Konkretisiert DEC-O-02 (Versionierung offen).

## 25. Open Questions

- Verbindliche Versionierungsnotation (DEC-O-02, später).
- Detailregeln je Change Class (spätere Iteration).
- Standard-Verifikationsmethoden je Datenklasse.

## 26. Next Decision

Companion 3 (Migration-Security) trägt die Autoritäts-/Approval-/Recovery-Regeln. API Governance (CO-WP-017) und Event/Audit Model (CO-WP-018) referenzieren Schema-Versionierung. Format-/Framework-Wahl bleibt einer späteren ADR-Runde vorbehalten.
