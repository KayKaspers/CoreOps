# CoreOps – Evidence Reference, Validation and Lineage Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation evidence-reference and validation model
> Implementation Status: Not implemented
> Evidence Storage: Not selected
> Integrity Mechanism: Not selected
> Legal Admissibility: None claimed
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-018 (docs-only / event, audit, evidence and security-governance foundation)

## 1. Status

Technologieunabhängiges Modell für **Evidence References, Provenance, Validation und Sufficiency**. Companion zu [EVENT_AND_AUDIT_CORRELATION_MODEL.md](EVENT_AND_AUDIT_CORRELATION_MODEL.md) und [AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md). Keine Evidence-Storage-/Integrity-Technologie; **keine rechtliche Beweiskraft behauptet**.

## 2. Purpose

Evidenz wird oft überinterpretiert: verfügbar heißt nicht gültig, gültig heißt nicht ausreichend, ausreichend heißt nicht compliant. Dieses Modell trennt sechs Evidence-Dimensionen und bindet Sufficiency an Entscheidung/Scope/Zeit — damit `evidence capability ≠ available ≠ current`, `integrity stated ≠ verified`, `validated ≠ sufficient` und `sufficient ≠ compliance certified`. Es baut auf dem Three-State-Evidence-Modell (CO-WP-004E) auf (kein Parallelmodell).

## 3. Scope

Evidence-Begriffe/-Klassen · Requirements · Sources/Producers · References · Sets · Availability · Freshness · Integrity · Provenance · Handling History · Validation · Sufficiency · Conflicts · Supersession · Retention · Disclosure · Offline Evidence · Audit/Evidence.

## 4. Non-Goals

- Keine Evidence-Storage-/Integrity-/Hash-/Signatur-/WORM-Technologie.
- Keine rechtliche Chain-of-Custody-/Admissibility-Behauptung.
- Keine Behauptung durchgeführter Validierung/Compliance.

## 5. Evidence Concepts

`evidence capability · evidence requirement · evidence source · evidence producer · evidence item · evidence reference · evidence set · evidence collection · evidence availability · evidence integrity · evidence provenance · evidence freshness · evidence validation · evidence sufficiency · evidence disclosure · evidence retention`.
```text
evidence capability     ≠ evidence available
evidence available      ≠ evidence current
evidence integrity stated ≠ integrity verified
evidence validated      ≠ evidence sufficient for every decision
evidence sufficient     ≠ compliance certified
```

## 6. Evidence Classes

Evidenz je Domäne (identity/policy/approval/authorization/execution/verification/state/migration/integration/deployment/audit/offline). Jede Klasse hat Owner, Source, Disclosure- und Retention-Erwartung; Evidence relevance ist an Events (Companion 1) geknüpft.

## 7. Evidence Requirements

Eine Evidence Requirement beschreibt, welche Evidenz für eine Entscheidung/Prüfung nötig ist. `evidence requirement ≠ evidence available`; eine Requirement ohne verfügbare, gültige, ausreichende Evidenz ist nicht erfüllt.

## 8. Evidence Sources and Producers

Source (Ursprung) und Producer (Sammler/Erzeuger) getrennt; Source Trust explizit. `multiple evidence items ≠ independent evidence sources` (mehrere Items aus derselben Quelle sind nicht unabhängig).

## 9. Evidence References

Ein Evidence Reference benötigt mindestens: stable evidence reference · evidence class · source · producer · owner · subject · object · workspace/scope · collection context · event/operation references · schema/format reference · creation/observation boundary · freshness state · provenance state · integrity state · validation state · availability state · retention state · disclosure classification · access owner · known limitations · supersession reference · audit reference. **Ein Reference ist kein Ersatz für das Artefakt und kein Nachweis seiner Existenz.**

## 10. Evidence Sets

Ein Evidence Set bündelt Referenzen für eine Entscheidung; Set-Sufficiency (§17) ist eigenständig. Ein vollständiges Set ist nicht automatisch ausreichend oder unabhängig.

## 11. Availability

`available` heißt referenzierbar/zugreifbar im Scope; `evidence capability ≠ available`; Availability ≠ Freshness/Validity/Sufficiency.

## 12. Freshness

`fresh` bezieht sich auf Aktualität relativ zu Entscheidung/Beobachtung. `evidence available ≠ current`; `fresh ≠ authoritative`.

## 13. Integrity

Integrity beschreibt Unverändertheit; `integrity stated ≠ integrity verified` (Mechanismus deferred). Integritätsmetadaten belegen nicht automatisch Integrität (Bezug THR-016, THR-017).

## 14. Provenance

Provenance beschreibt Herkunft/Verarbeitungskette; bleibt durch Normalisierung/Redaction/Export erhalten (konsistent mit [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md), kein Parallelmodell). `parsed ≠ validated provenance`.

## 15. Handling History

Dokumentiert: original source · collector/producer · collection context · transformations · normalisation · redaction · export · import · offline transfer · validation activity · owner changes · supersession · retention transition · withdrawal/invalidation. **Keine rechtliche Chain-of-Custody-/Admissibility-Behauptung.**
```text
handling history ≠ legal admissibility
```

## 16. Validation

`not-assessed · available-unvalidated · validation-pending · validated-with-notes · validated · invalid · stale · superseded · unavailable · conflicted`.
```text
validated ≠ sufficient
```

## 17. Sufficiency

`not-assessed · insufficient · partially-sufficient · sufficient-for-decision · sufficient-with-exception · conflicted · unknown`.
```text
sufficient for one decision ≠ sufficient for another decision
fresh                       ≠ authoritative
multiple evidence items     ≠ independent evidence sources
```
**Sufficiency ist an Entscheidung, Scope und Zeitpunkt gebunden.**

## 18. Conflicts

Konfligierende Evidenz (`conflicted`) bleibt sichtbar und wird nicht still aufgelöst; sie ist nicht sicher für privilegierte automatische Entscheidungen. Konfliktauflösung ist auditierbar (Bezug SoT-Konfliktmodell CO-WP-011).

## 19. Supersession

Neuere Evidenz kann ältere ablösen (`superseded`); Supersession löscht keine historische Referenz/Evidenz (Companion 3 §16). Supersession-Historie bleibt nachvollziehbar.

## 20. Retention

Retention Owner/Class/Trigger je Evidence (Detail Companion 3 §14). `retention expiry ≠ automatic purge authority`; `archived ≠ safe to purge`.

## 21. Disclosure

Disclosure/Export benötigt eigene Autorität (Companion 3 §17-18); `read permission ≠ export permission`; `evidence reference ≠ permission to access referenced artifact`. Redaction/Datenminimierung gilt (Bezug THR-018, THR-019, THR-020, THR-040).

## 22. Offline Evidence

Offline-Evidence-Pakete folgen Companion 3 §21 (target-environment binding, provenance, integrity, import governance). Nicht behauptet: implementiert, beliebige Air-Gap-Stufen, konfliktfreie Reconciliation, Klassifiziertnetz-Eignung, konkrete Signing-Technologie (Bezug THR-024).

## 23. Audit and Evidence

Evidence-Aktivitäten (collection, validation, supersession, export) werden auditiert (Companion 1/3). Trennung `evidence available ≠ valid ≠ sufficient ≠ compliance`.

## 24. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Evidence capability does not imply evidence availability.
2. Availability does not imply freshness, validity or sufficiency.
3. Stated integrity is not verified integrity.
4. Validated evidence is not automatically sufficient for every decision.
5. Sufficiency remains decision-, scope- and time-bound.
6. Multiple evidence items are not necessarily independent evidence sources.
7. An evidence reference is not the artifact and not proof of its existence.
8. Handling history does not establish legal admissibility.
9. Conflicted or superseded evidence must remain visible; supersession must not delete history.

## 25. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md): THR-013, THR-016, THR-017, THR-018, THR-019, THR-020, THR-024, THR-029, THR-030, THR-040. Keine Duplikation, kein Parallelregister.

## 26. Technology Boundary

Nicht ausgewählt/implementiert: Evidence-Storage, Integrity-/Hash-/Signatur-/WORM-Mechanismus, Schema, Runtime-Code.

## 27. Compatibility

Konsistent mit [Event/Audit Correlation](EVENT_AND_AUDIT_CORRELATION_MODEL.md), [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md), [Capability Matrix Spec](../project-system/CAPABILITY_MATRIX_SPEC.md) (Evidence-Capability-Trennung), [Audit Policy](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md). Konkretisiert DEC-P-06 und die Evidence-Capability-Entscheidungen (DEC-S-31…37).

## 28. Open Questions

- Standardisierte Sufficiency-Kriterien je Entscheidungstyp.
- Integrity-Mechanismus (spätere ADR).
- Independent-Source-Kriterien im Detail.

## 29. Next Decision

Companion 3 trägt Integrity/Retention/Disclosure. Test Strategy (CO-WP-028) und Readiness Review (CO-WP-030) referenzieren Evidence-Sufficiency. Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
