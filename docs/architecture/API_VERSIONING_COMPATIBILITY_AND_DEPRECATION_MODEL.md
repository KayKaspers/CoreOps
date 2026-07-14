# CoreOps – API Versioning, Compatibility and Deprecation Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation API-versioning and compatibility model
> Implementation Status: Not implemented
> Versioning Notation: Not selected
> Schema Binding: Not selected
> Compatibility Validation: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-017 (docs-only / API architecture and security foundation)

## 1. Status

Technologieunabhängiges Modell für **API-Versionierung, Kompatibilität, Change Classes, Deprecation und Retirement**. Companion zu [API_GOVERNANCE_AND_OPERATION_MODEL.md](API_GOVERNANCE_AND_OPERATION_MODEL.md) und [API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md](../security/API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md). Keine Versionsnotation, kein Schema-Binding, keine Compatibility-Validation-Technologie.

## 2. Purpose

Eine „nur additive" Änderung kann semantisch breaking sein; `request-compatible ≠ response-compatible`; `same API version ≠ identical operation set`. Dieses Modell trennt die Versionsdimensionen und Kompatibilitätsklassen strikt und bindet jeden Compatibility Claim an Producer, Consumer, Version, Schema, Profil und Evidenz.

## 3. Scope

Version Dimensions · API/Operation Identity · Compatibility Classes · Change Classes · Request/Response/Error/Behavioural Compatibility · Producer/Consumer Compatibility · Deprecation · Retirement · Migration Guidance · Support/Validation · Domain-Pack-Relationship · Offline/Delayed Consumers · Audit/Evidence.

## 4. Non-Goals

- Keine Versionsnotation (SemVer o. ä.) erzwungen, kein Schema-Binding.
- Keine Compatibility-Validation-Runtime, kein Transport/API-Style.
- Keine Behauptung durchgeführter Compatibility-Validierung.

## 5. Version Dimensions

Getrennt: `API surface version · operation version · request contract version · response contract version · error contract version · data schema version · event/async-result version · Integration Contract version · Domain-Pack version · CoreOps product version · consumer version · producer version`. **Keine Dimension ersetzt still eine andere.** Keine Notation erzwungen (DEC-O-02 offen).

## 6. API and Operation Identity

Stabile API-/Operation-Identität (aus [Governance §7]) ist Bezugspunkt der Versionierung; `route/URL/transport ≠ identity`; `API version ≠ product version`. Retired IDs werden nicht wiederverwendet (§15).

## 7. Compatibility Classes

`not-assessed · backward-readable · forward-readable-with-notes · request-compatible · response-compatible · error-compatible · behaviourally-compatible · compatible-with-notes · validated-compatible · incompatible · deprecated-compatibility · unknown · conflicted`.
```text
request-compatible ≠ response-compatible
schema-compatible  ≠ behaviourally-compatible
same API version   ≠ identical operation set
newer version      ≠ automatically compatible
unknown            ≠ compatible
conflicted         ≠ safe for privileged automation
```
Claims gebunden an Producer · Consumer · API-/Operation-Version · Schema · Profil · Evidenz.

## 8. Change Classes

`additive optional operation · additive optional field · additive required field · field removal · operation removal · semantic change · default-behaviour change · authorization change · scope change · error-semantics change · side-effect change · idempotency change · pagination/ordering change · identifier change`. Je Klasse: Compatibility-Auswirkung · Consumer-Auswirkung · Migration Requirement · Deprecation Requirement · Security-Auswirkung · Validation Requirement. **Eine formal additive Änderung kann semantisch breaking sein.**

## 9. Request Compatibility

Betrifft, ob ältere/andere Consumer weiterhin gültige Requests stellen können. `request-compatible` sagt nichts über Response/Behaviour aus.

## 10. Response Compatibility

Betrifft, ob Consumer Responses weiterhin korrekt interpretieren. `response-compatible ≠ request-compatible`; additive Response-Felder können ältere Consumer brechen, wenn Pflichtsemantik geändert.

## 11. Error Compatibility

Betrifft Error-/Problem-Semantik (Companion 3). Eine geänderte Error-Semantik (z. B. andere Klasse/Retry-Bedeutung) ist eine eigene Kompatibilitätsdimension und kann breaking sein, auch wenn Request/Response gleich bleiben.

## 12. Behavioural Compatibility

Betrifft beobachtbares Verhalten (Side-Effects, Idempotency, Ordering, Default-Behaviour). `schema-compatible ≠ behaviourally-compatible`; Verhaltensänderungen benötigen eigene Bewertung/Deprecation.

## 13. Producer and Consumer Compatibility

Producer und Consumer haben getrennte Versionen/Kompatibilität; `new producer ≠ old consumer compatibility`. Mixed-Version-Consumer (inkl. Offline, §19) benötigen explizite Grenzen.

## 14. Deprecation

Benötigt: affected API/operations · reason · replacement · migration guidance · consumer impact · support impact · security impact · announcement reference · review trigger · retirement condition · owner.
```text
deprecated        ≠ unavailable
replacement exists ≠ consumers migrated
retired           ≠ historical evidence deleted
consumer inactivity ≠ safe removal
```

## 15. Retirement

`retired` erhält keine neuen Consumers; retired API-/Operation-IDs werden **nicht** still wiederverwendet; historische Evidenz bleibt erhalten. Consumer-Inaktivität ist kein Nachweis sicherer Entfernung.

## 16. Migration Guidance

Deprecation/Change liefert Migrationshinweise für Consumer (Replacement, Zeitfenster, Verhaltensunterschiede). Migration Guidance ist keine Garantie erfolgter Consumer-Migration.

## 17. Support and Validation

Support-/Validation-Status je API sind eigene Dimensionen (konsistent mit [Domain Pack Support Model](DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md)): `advertised/documented ≠ supported ≠ validated`; `validated-compatible` benötigt aktuelle scopegebundene Evidenz und wird nicht generalisiert.

## 18. Domain-Pack Relationship

Domain Packs (CO-WP-015) können API-/Operation-Profile referenzieren; Pack-Compatibility-Claims sind an API-/Operation-Version gebunden. Ein Pack-Status validiert eine API nicht automatisch.

## 19. Offline and Delayed Consumers

Offline/verzögerte Consumer können ältere API-/Schema-Versionen nutzen; Kompatibilität ist explizit version-/scope-gebunden; unbekannte Pflichtsemantik wird nicht still ignoriert (fail-safe). Reconciliation nach Reconnection.

## 20. Audit and Evidence

Erfasst: version assignment · compatibility classification · change class · deprecation/retirement · validation activity/evidence. `compatibility evidence ≠ compatibility for every consumer`; scopegebunden.

## 21. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. API version dimensions remain separate; no version silently replaces another.
2. API version is not the CoreOps product version.
3. Request-, response-, error- and behavioural-compatibility remain separate.
4. A formally additive change may be semantically breaking and requires evaluation.
5. Unknown or conflicted compatibility does not imply compatibility and is unsafe for privileged automation.
6. `validated-compatible` requires current, scope-bound evidence and is not generalized.
7. Deprecated or retired API/operation identities must not be reused silently.
8. Historical API identity and compatibility evidence must not be deleted.

## 22. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md): THR-011, THR-013, THR-023, THR-035, THR-038. Keine Duplikation, kein Parallelregister.

## 23. Technology Boundary

Nicht ausgewählt/implementiert: Versionsnotation, Schema-Binding, Compatibility-Validation-Runtime, Transport/API-Style.

## 24. Open Questions

- Verbindliche Versionsnotation (DEC-O-02).
- Detailregeln je Change Class.
- Standardisierte Behavioural-Compatibility-Prüfung.

## 25. Next Decision

Companion 3 trägt Error-/Idempotency-Semantik. Event/Audit Model (CO-WP-018) referenziert Versionierung. Notationswahl bleibt einer späteren ADR-Runde vorbehalten.
